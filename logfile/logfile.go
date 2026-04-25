package logfile

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Init 将标准 log 与 Gin 输出定向到 dir/server.log（不写控制台）。
func Init(dir string) (close func(), err error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	p := filepath.Join(dir, "server.log")
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags)
	gin.DefaultWriter = io.Writer(f)
	gin.DefaultErrorWriter = f
	return func() { _ = f.Close() }, nil
}
