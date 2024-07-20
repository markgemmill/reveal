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

const version = "0.1.6"

type Config struct {
	Timeout      int64  `toml:"timeout"`
	CopyTimeout  int64  `toml:"copy_timeout"`
	SegmentSize  int    `toml:"segment_size"`
	DockPosition string `toml:"dock_position"`
}

var positions = map[string]bool{"center": true, "ul": true, "ur": true, "ll": true, "lr": true}

// App struct
type App struct {
	ctx     context.Context
	cfg     *Config
	cfgPath pathlib.Path
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

	a.cfgPath = cfgPth

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

func (a *App) ClearClipboard() {
	clipboard.Write(clipboard.FmtText, []byte(""))
}

func (a *App) GetVersion() string {
	return version
}

func (a *App) GetTimeout() int64 {
	return a.cfg.Timeout
}

func (a *App) GetCopyTimeout() int64 {
	return a.cfg.CopyTimeout
}

func (a *App) GetSegmentSize() int {
	if a.cfg.SegmentSize < 3 {
		return 3
	}
	if a.cfg.SegmentSize > 5 {
		return 5
	}
	return a.cfg.SegmentSize
}

func (a *App) GetDockPosition() string {
	// values - center, UL, UR, LL, LR
	return a.cfg.DockPosition
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

func (a *App) WriteConfig(timeout, copyTimeout int64, segmentSize int, docPosition string) {
	config := Config{
		Timeout:      timeout,
		CopyTimeout:  copyTimeout,
		SegmentSize:  segmentSize,
		DockPosition: docPosition,
	}
	cfgUpdate, err := toml.Marshal(&config)
	if err != nil {
		panic(err)
	}
	a.cfgPath.Write(cfgUpdate)
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
	runtime.LogDebug(a.ctx, fmt.Sprintf("CONFIG COPY TIMEOUT: %d", config.CopyTimeout))
	runtime.LogDebug(a.ctx, fmt.Sprintf("CONFIG SEGMENT SIZE: %d", config.SegmentSize))
	runtime.LogDebug(a.ctx, fmt.Sprintf("CONFIG DOCK POSITION: %s", config.DockPosition))

	writeConfig := false

	if config.SegmentSize <= 2 || config.SegmentSize > 5 {
		config.SegmentSize = 4
		writeConfig = true
	}

	if config.Timeout <= 10 || config.Timeout > 60 {
		config.Timeout = 30
		writeConfig = true
	}

	if config.CopyTimeout <= 0 || config.CopyTimeout > 30 {
		config.CopyTimeout = 20
		writeConfig = true
	}

	_, validDockPosition := positions[config.DockPosition]

	if !validDockPosition {
		runtime.LogDebug(a.ctx, fmt.Sprintf("Given Dock Position to overwrite: %s", config.DockPosition))
		config.DockPosition = "center"
		writeConfig = true
	}

	if writeConfig {
		cfgUpdate, err := toml.Marshal(&config)
		if err != nil {
			panic(err)
		}
		pth.Write(cfgUpdate)
	}

	a.cfg = &config
}
