package config

// Configuration of the entire application
type Configuration struct {
	App  App
	DB   DB
	Mail Mail
}

// App specific configuration
type App struct {
	IP     string
	Port   int
	Sock   string
	Conn   string
	Upload string
	Secret string
	Title  string
	Tz     string
}

// DB specific configruation
type DB struct {
	User string
	Pass string
	Name string
	Host string
	Port int
	Sock string
	Conn string
}

// Mail configuration
type Mail struct {
	Host string
	Port int
	User string
	Pass string
}
