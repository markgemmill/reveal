package main

import (
	"context"

	"golang.design/x/clipboard"
)

const version = "0.1.1"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
}

// Greet returns a greeting for the given name
func (a *App) GetClipboard() string {
	value := clipboard.Read(clipboard.FmtText)
	return string(value)
}

func (a *App) PutClipboard(value string) {
	clipboard.Write(clipboard.FmtText, []byte(value))
}

func (a *App) GetVersion() string {
	return version
}
