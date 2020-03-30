package models

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/testfixtures.v2"

	"github.com/jinzhu/gorm"

	// testfixtures need Sqlite drivers
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func fatalTestError(fmtStr string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmtStr, args...)
	os.Exit(1)
}

// MainTest a reusable TestMain(..) function for unit tests that need to use a
// test database. Creates the test database, and sets necessary settings.
func MainTest(m *testing.M, pathToRoot string) {
	var err error

	fixturesDir := filepath.Join(pathToRoot, "pkg", "fixtures")
	if err = createTestEngine(fixturesDir); err != nil {
		fatalTestError("Error creating test engine: %v\n", err)
	}
	os.Exit(m.Run())
}

func createTestEngine(fixturesDir string) (err error) {

	db, err = gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return
	}

	db.AutoMigrate(tables...)

	newSession()
	return InitFixtures(&testfixtures.SQLite{}, fixturesDir)
}

// PrepareTestDatabase load test fixtures into test database
func PrepareTestDatabase() error {
	return LoadFixtures()
}
