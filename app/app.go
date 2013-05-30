package app

import (
    "log"
    "net/http"
    "strconv"
    "github.com/laurent22/toml-go/toml"
    "github.com/gorilla/mux"
)

type App struct {
    Host string
    Port int
}

func Load(path string) *App {
    app := new(App)
    var parser toml.Parser
    config := parser.ParseFile(path)

    app.Host = config.GetString("server.host")
    app.Port = config.GetInt("server.port")

    return app
}

func (app *App) Run() {
    r := mux.NewRouter()

    r.Handle("/", Handler(Index))

    http.Handle("/", r)
    err := http.ListenAndServe(app.Host + ":" + strconv.Itoa(app.Port), nil)
    if err != nil {
        log.Fatal("Listen and Server: ", err)
    }
}
