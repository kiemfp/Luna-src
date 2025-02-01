package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mitchellh/panicwrap"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

var AcceptedVersion = "version-ecc9c250281b4c14"

//go:embed all:frontend/build
var assets embed.FS
var app *App

func GetLogName() string {
	return "debug-" +
		fmt.Sprintf("%v", time.Now().Unix()) + "-luna.txt"
}

var Logger = logger.NewFileLogger("./logs/" + GetLogName())

func main() {

	if _, err := os.Stat("scripts"); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("scripts", os.ModePerm)
	}
	if _, err := os.Stat("autoexec"); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("autoexec", os.ModePerm)
		os.Create("./autoexec/main.luau")
		os.WriteFile("./autoexec/main.luau", []byte(`
-- this is auto generated by lunas init system.
-- this is auto generated by lunas init system.
-- this is auto generated by lunas init system.
-- this is auto generated by lunas init system.

print([[


88     88   88 88b 88    db    
88     88   88 88Yb88   dPYb   
88  .o Y8   8P 88 Y88  dP__Yb  
88ood8 'YbodP' 88  Y8 dP""""Yb 
- http2
]])

print("welcome to luna!")
wait(2)
printidentity()`), 0644)
	}
	if _, err := os.Stat("workspace"); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("workspace", os.ModePerm)
	}
	if _, err := os.Stat("logs"); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("logs", os.ModePerm)
	}

	files, _ := os.ReadDir("./logs")
	if len(files) > 10 {
		for _, entry := range files {
			os.RemoveAll(filepath.Join("./logs", entry.Name()))
		}
	}

	exitStatus, err := panicwrap.BasicWrap(func(output string) {
		Logger.Fatal(output)
		os.Exit(1)
	})
	if err != nil {
		os.Exit(1)
	}
	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	app = NewApp()

	err = wails.Run(&options.App{

		LogLevelProduction: logger.ERROR,
		Logger:             Logger,
		Title:              "Luna",
		Width:              836,
		Height:             504,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		WindowStartState: options.Normal,
		OnDomReady: func(ctx context.Context) {
			go GetRobloxProccesses()
		},
		Windows: &windows.Options{
			WindowIsTranslucent: true,
			BackdropType:        windows.Acrylic,
		},
	})

	if err != nil {
		Logger.Error(err.Error())
	}
}

/*
	_, pid := memory.IsProcessRunning()

	fmt.Println(pid)

	if len(pid) > 0 {
		mem, err := memory.NewLuna(pid[0].Pid)
		if err != nil {
			return
		}

		instance.PatchRoblox(mem)

		renderView := utils.GetRenderVDM(pid[0].Pid, mem, utils.OffsetsDataPlayer, false)

		var Start time.Time = time.Now()
		datamodel, _ := mem.AOBSCANALL("52 65 6e 64 65 72 4a 6f 62 28 45 61 72 6c 79 52 65 6e 64 65 72 69 6e 67 3b", false, 1)
		fmt.Println(time.Since(Start))

		var New *instance.RobloxInstances = &instance.RobloxInstances{
			Pid:     int64(pid[0].Pid),
			ExeName: pid[0].Name,
			Mem:     mem,
			Instances: instance.Instances{
				RenderView: renderView,
				RobloxBase: uint64(mem.RobloxBase),
			},
			Offsets: utils.OffsetsDataPlayer,
		}

		var DM instance.Instance

		var (
			RenderView = 0x1E8
		)

		rv, _ := mem.ReadPointer(datamodel[0] + uintptr(RenderView))

		var (
			Fake = uintptr(0x118)
			Real = uintptr(0x1A8)
		)

		for {
			fakedm, _ := mem.ReadPointer(rv + Fake)
			realdm, _ := mem.ReadPointer(fakedm + Real)
			DM = instance.NewInstance(realdm, New)

			fmt.Println(DM.String())
			time.Sleep(time.Millisecond * 150)
		}

		os.Exit(0)

	}
*/
