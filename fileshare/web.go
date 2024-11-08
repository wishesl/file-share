package fileshare

import (
	"file-share/global"
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func Init(w fyne.Window) error {
	err := global.DB.AutoMigrate(&FileShare{})
	if err != nil {
		return err
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/download/:key", func(c *gin.Context) {
		key := c.Param("key")
		var fileshare FileShare
		global.DB.Where("key = ?", key).Find(&fileshare)
		if fileshare.Path == "" {
			c.String(http.StatusOK, "无效的Key")
			return
		}
		fileName := filepath.Base(fileshare.Path)
		c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.File(fileshare.Path) // 文件路径
	})
	r.GET("/restart", func(c *gin.Context) {
		w.Show()
		w.RequestFocus()
		c.String(http.StatusOK, "重启中")
	})

	return r.Run(fmt.Sprintf(":%d", global.Cfg.Port))
}

func GetDownloadUrl(key string) string {
	return fmt.Sprintf("http://%s:%d/download/%s", global.Cfg.Ip, global.Cfg.Port, key)
}
