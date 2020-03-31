package config

type database struct {
	Driver   string
	Username string
	Password string
	Name     string
	Host     string
}

type server struct {
	Addr  string
	Root  string
	Pprof bool
	Debug bool `default:"true"`
}

type session struct {
	Key string
}

type logs struct {
	Color  bool
	Debug  bool
	Pretty bool
	Level  string `default:"debug"`
}

// ContextKey for context package
type ContextKey string

func (c ContextKey) String() string {
	return "user context key " + string(c)
}

var (
	// Database represents the current database connection details.
	Database = &database{}

	// Server represents the informations about the server bindings.
	Server = &server{}

	// Session represents the informations about the session handling.
	Session = &session{}

	// ContextKeyUser for user
	ContextKeyUser = ContextKey("user")

	// Logs for zerolog
	Logs = &logs{}
)
