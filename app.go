package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/saidalisamed/muxwebapp/config"
	"github.com/saidalisamed/muxwebapp/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

// App main structure
type App struct {
	Router   *mux.Router
	DB       *sql.DB
	Cfg      *config.Configuration
	Store    *sessions.CookieStore
	Template *template.Template
}

// Init initializes the application
func (a *App) Init(env string) {
	// Read environment configuration
	a.readConfig(env)

	a.initDB()

	a.initSession()

	a.Router = mux.NewRouter()
	a.configureMiddlewares()
	a.configureStatic("res/assets", "/static/")
	a.configureRoutes()
	a.parseTemplates("res/templates/*")
}

// Run starts the application
func (a App) Run() {
	// Enable logging
	loggedRouter := handlers.LoggingHandler(os.Stdout, a.Router)

	if a.Cfg.App.Conn == "sock" {
		socket, err := net.Listen("unix", a.Cfg.App.Sock)
		if err != nil {
			log.Printf("%s", err)
		} else {
			// Make socket file group writable
			if err := os.Chmod(a.Cfg.App.Sock, 0770); err != nil {
				log.Fatal(err)
			}

			// Graceful shutdown. Unlink unix socket on exit
			sigc := make(chan os.Signal, 1)
			signal.Notify(sigc, os.Interrupt, os.Kill, syscall.SIGTERM)
			go func(c chan os.Signal) {
				sig := <-c
				log.Printf("Caught signal %s: shutting down.", sig)
				socket.Close()
				os.Exit(0)
			}(sigc)

			log.Fatal(http.Serve(socket, handlers.RecoveryHandler()(loggedRouter)))
		}
	} else {
		log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", a.Cfg.App.IP, a.Cfg.App.Port), handlers.RecoveryHandler()(loggedRouter)))
	}
}

// Read configuration file
func (a *App) readConfig(env string) {
	var configName string
	if env == "prod" {
		configName = "prod"
	} else {
		env = "dev"
		configName = env
	}

	viper.SetConfigName(configName) // name of config file (without extension)
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error: %s", err))
	}

	err = viper.Unmarshal(&a.Cfg)
	if err != nil {
		log.Fatalf("unable to unmarshal into struct, %v", err)
	}

	log.Printf("Environment: %s", env)
	log.Printf("Application port: %d", a.Cfg.App.Port)
	log.Printf("Database host: %s", a.Cfg.DB.Host)
}

// Initialize database connection
func (a *App) initDB() {
	connType := fmt.Sprintf("@tcp(%s:%d)", a.Cfg.DB.Host, a.Cfg.DB.Port)
	if a.Cfg.DB.Conn == "sock" {
		connType = fmt.Sprintf("@unix(%s)", a.Cfg.DB.Sock)
	}

	var err error
	a.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s%s/%s?parseTime=true", a.Cfg.DB.User, a.Cfg.DB.Pass, connType, a.Cfg.DB.Name))
	if err != nil {
		panic(err.Error())
	}
}

// Initialize sessions
func (a *App) initSession() {
	if a.Cfg.App.Secret == "" {
		token, err := utils.GenerateRandomString(32)
		a.Cfg.App.Secret = token
		utils.CheckErr(err, "Secret:")
	}

	a.Store = sessions.NewCookieStore([]byte(a.Cfg.App.Secret))
}

// Plugs in various middlewares
func (a *App) configureMiddlewares() {
	// Use middlewares
	// a.Router.Use(someMiddleware)
}

func (a *App) configureStatic(dir string, pathPrefix string) {
	fs := http.FileServer(http.Dir(dir))
	// or use github.com/yosssi/go-fileserver for faster performance and caching
	// fs := fileserver.New(fileserver.Options{}).Serve(http.Dir(dir))

	a.Router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix, fs))
}

func (a *App) parseTemplates(path string) {
	template := template.New("")
	_, err := template.ParseGlob(path)
	if err != nil {
		log.Println(err)
	}
	a.Template = template
}
