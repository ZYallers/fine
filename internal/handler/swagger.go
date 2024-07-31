package handler

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ZYallers/fine/frame/f"
	"github.com/ZYallers/fine/os/ffile"
	"github.com/gin-gonic/gin"
)

func Swagger(docDir string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		swagFile := filepath.Join(docDir + "swagger.json")
		if !ffile.Exists(swagFile) {
			ctx.AbortWithStatusJSON(http.StatusOK, f.MapStrAny{"code": http.StatusNotFound, "msg": "swagger.json file not exist"})
			return
		}
		file, err := os.Open(swagFile)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, f.MapStrAny{"code": http.StatusInternalServerError, "msg": err.Error()})
			return
		}
		defer file.Close()
		fd, err := ioutil.ReadAll(file)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, f.MapStrAny{"code": http.StatusInternalServerError, "msg": err.Error()})
			return
		}
		ctx.Header("Content-Type", "application/json")
		ctx.String(http.StatusOK, string(fd))
	}
}
