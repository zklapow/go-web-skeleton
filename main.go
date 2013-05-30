package main

import (
    "go-web-skeleton/app"
)

func main() {
    a := app.Load("config.toml")
    a.Run()
}
