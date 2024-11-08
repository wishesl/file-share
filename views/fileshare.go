package views

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"file-share/assets"
	"file-share/fileshare"
	"file-share/global"
	"fmt"
	dialogw "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"github.com/atotto/clipboard"
	"github.com/sqweek/dialog"
	"image/png"
	"os"
	"path/filepath"

	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type FileShare struct {
	fileshares []fileshare.FileShare

	list *fyne.Container

	w fyne.Window
}

var openChen = make(chan func(), 3)

func init() {
	go func() {
		for {
			select {
			case f := <-openChen:
				f()
			}
		}
	}()
}

func NewFileShare() *FileShare {
	return &FileShare{}
}

func (fs *FileShare) Show(w fyne.Window) {
	if fs.w == nil {
		fs.w = w
		content := fs.CanvasObject()
		fs.w.SetContent(content)
		fs.w.SetOnDropped(func(p fyne.Position, url []fyne.URI) {
			fs.OnDropped(p, url)
			content.Refresh()
		})
		fs.w.SetOnClosed(func() {
			fs.w = nil
		})
	} else {
		fs.w.CenterOnScreen()
	}

}

func (fs *FileShare) CanvasObject() fyne.CanvasObject {
	//var fileshares []fileshare.FileShare
	global.DB.Find(&fs.fileshares)

	fs.list = container.NewVBox()

	for _, file := range fs.fileshares {
		fs.list.Add(fs.fileItemCanvasObject(file))
	}

	var open = false
	// "点击添加或拖拽到此界面"
	createButton := widget.NewButton("点击添加或拖拽到此界面", func() {
		if open {
			return
		}
		open = true
		go func() {
			defer func() { open = false }()
			filePath, err := dialog.File().Load()
			if err != nil {
				return
			}
			fs.createShare(filePath)
		}()
	})
	createButton.SetIcon(assets.ShareDataSR)

	deleteAll := widget.NewButton("全部删除", func() {
		for _, file := range fs.fileshares {
			global.DB.Where("path = ?", file.Path).Delete(&file)
		}
		fs.fileshares = nil
		fs.list.RemoveAll()
	})
	deleteAll.SetIcon(theme.DeleteIcon())

	contain := container.NewBorder(nil, nil, nil, nil, createButton)
	filesView := container.NewBorder(contain, deleteAll, nil, nil, container.NewScroll(fs.list))
	return filesView
}

func (fs *FileShare) createShare(filePath string) {
	var item fileshare.FileShare
	global.DB.Where("path = ?", filePath).Find(&item)
	if item.Path != "" {
		dialogw.NewInformation("文件已存在", "文件已存在", fs.w).Show()
		return
	}

	start := time.Now().UnixMilli()
	hash := md5.New()
	hash.Write([]byte(fmt.Sprintf("%s-%s", filePath, time.Now().Format("2006-01-02_15-04-05"))))
	hashInBytes := hash.Sum(nil)
	md5String := hex.EncodeToString(hashInBytes)
	var newFileshares = fileshare.FileShare{Path: filePath, Key: md5String}
	global.DB.Create(&newFileshares)
	fs.fileshares = append(fs.fileshares, newFileshares)
	fs.list.Add(fs.fileItemCanvasObject(newFileshares))
	fmt.Println("创建时间：", time.Now().UnixMilli()-start)
}

func (fs *FileShare) fileItemCanvasObject(item fileshare.FileShare) fyne.CanvasObject {
	var cav fyne.CanvasObject

	deleteButton := widget.NewButton("", func() {
		global.DB.Where("path = ?", item.Path).Delete(&item)
		global.DB.Find(&fs.fileshares)
		fs.list.Remove(cav)
	})
	deleteButton.SetIcon(theme.CancelIcon())

	cav = container.NewBorder(nil, nil, nil, deleteButton,
		widget.NewButton(fmt.Sprintf("%s", filepath.Base(item.Path)), share(item)))
	return cav
}

func (fs *FileShare) OnDropped(_ fyne.Position, urls []fyne.URI) {
	for _, url := range urls {
		fileinfo, err := os.Stat(url.Path())
		if err != nil {
			continue
		}
		if fileinfo.IsDir() {
			continue
		}
		fs.createShare(url.Path())
	}
}

func share(f fileshare.FileShare) func() {
	return func() {
		fmt.Println(fileshare.GetDownloadUrl(f.Key))
		url := fileshare.GetDownloadUrl(f.Key)
		showImage(f.Path, getImages(url), url)
	}
}

func getImages(str string) []byte {
	// 生成QR码
	qrCode, _ := qr.Encode(str, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 300, 330)

	// 创建输出文件
	buffer := new(bytes.Buffer)

	// 将QR码保存为PNG图片
	err := png.Encode(buffer, qrCode)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

func showImage(filePath string, bytes []byte, url string) {
	w := fyne.CurrentApp().NewWindow(filePath)
	w.Resize(fyne.NewSize(300, 300))
	w.CenterOnScreen()
	img := widget.NewIcon(&fyne.StaticResource{
		StaticName:    "a.img",
		StaticContent: bytes,
	})

	shareCopyButton := widget.NewButton("复制链接", func() {
		openChen <- func() {
			err := clipboard.WriteAll(fmt.Sprintf("文件分享: %s, 点击以下链接下载\n%s", filepath.Base(filePath), url))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			drv := fyne.CurrentApp().Driver()
			if drv, ok := drv.(desktop.Driver); ok {
				w := drv.CreateSplashWindow()
				w.SetContent(widget.NewLabelWithStyle("已将字符串复制到剪贴板",
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
				w.Show()

				go func() {
					time.Sleep(time.Second * 1)
					w.Close()
				}()
			} else {
				dialogw.ShowInformation("ok", "已将字符串复制到剪贴板", w)
			}
		}
	})
	w.SetContent(container.NewBorder(nil, shareCopyButton, nil, nil, img))
	w.Show()
}
