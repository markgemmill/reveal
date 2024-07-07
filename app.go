package main

import (
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/markgemmill/appdirs"
	"github.com/markgemmill/pathlib"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/clipboard"
)

const version = "0.1.2"

type Config struct {
	Timeout int64
}

// App struct
type App struct {
	ctx context.Context
	cfg *Config
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

	cfgPth, err := a.configInit()
	if err != nil {
		panic(err)
	}

	a.readConfig(cfgPth)
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

func (a *App) GetTimeout() int64 {
	return a.cfg.Timeout
}

func (a *App) configInit() (pathlib.Path, error) {
	appDirs := appdirs.NewAppDirs("reveal", "")
	appDir := pathlib.NewPath(appDirs.UserConfigDir(), 0777)
	if !appDir.Exists() {
		err := appDir.MkDirs()
		if err != nil {
			return pathlib.Path{}, err
		}
	}
	configFile := appDir.Join("config.toml")
	if !configFile.Exists() {
		err := configFile.Write([]byte("timeout = 30000"))
		if err != nil {
			return pathlib.Path{}, err
		}
	}

	return configFile, nil
}

func (a *App) readConfig(pth pathlib.Path) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("CONFIG PATH: %s", pth.String()))

	cfgText, err := pth.Read()
	if err != nil {
		panic(err)
	}

	runtime.LogDebug(a.ctx, fmt.Sprintf("CONFIG TEXT: %s", cfgText))

	var config Config
	err = toml.Unmarshal(cfgText, &config)
	if err != nil {
		panic(err)
	}

	runtime.LogDebug(a.ctx, fmt.Sprintf("CONFIG TIMEOUT: %d", config.Timeout))
	a.cfg = &config
}
