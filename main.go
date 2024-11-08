package main

import (
	"file-share/assets"
	"file-share/fileshare"
	"file-share/global"
	"file-share/views"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"net"
	"net/http"
)

func main() {

	if !isPortOpen(global.Cfg.Port) {
		http.Get(fmt.Sprintf("http://127.0.0.1:%d/restart", global.Cfg.Port))
		return
	}

	v := views.NewFileShare()

	a := app.NewWithID(fmt.Sprintf("io.fyne.%s", global.AppName))
	a.SetIcon(assets.LogoDataSR)
	// 设置字体
	a.Settings().SetTheme(&assets.MyDefaultTheme{})
	w := a.NewWindow("分享工具")

	go func() {
		if err := fileshare.Init(w); err != nil {
			a.Quit()
			return
		}
	}()

	// 设置desk
	setDesk(a, w)

	// 创建主窗口
	fullSize := fyne.NewSize(480, 380)
	w.Resize(fullSize)

	v.Show(w)
	w.SetMaster()
	w.ShowAndRun()
}

func setDesk(a fyne.App, w fyne.Window) {
	if desk, ok := a.(desktop.App); ok {
		w.SetCloseIntercept(func() { w.Hide() })

		quitMenuItem := fyne.NewMenuItem("退出", a.Quit)
		quitMenuItem.IsQuit = true
		quitMenuItem.Icon = assets.MultiplyDataSR

		mainMenuItem := fyne.NewMenuItem("打开", func() {
			w.Show()
		})
		// 2024/11/8 15:40 用自己的图标会出现desktop重复显示问题原因不明
		mainMenuItem.Icon = assets.TwitterDataSR

		m := fyne.NewMenu("文件分享",
			mainMenuItem,
			fyne.NewMenuItemSeparator(),
			quitMenuItem)
		desk.SetSystemTrayMenu(m)
	}
}

func isPortOpen(port int) bool {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false // 端口被占用
	}
	defer listener.Close()
	return true // 端口未被占用
}
