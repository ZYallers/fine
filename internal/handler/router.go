package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/ZYallers/fine/errors/ferror"
	"github.com/ZYallers/fine/frame/fapp"
	"github.com/ZYallers/fine/frame/frouter"
	"github.com/ZYallers/fine/internal/route"
	"github.com/ZYallers/fine/os/fcfg"
	"github.com/ZYallers/fine/os/fctx"
	"github.com/ZYallers/fine/text/fstr"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ParseRouter(app *fapp.App) {
	if len(app.Router) == 0 {
		return
	}
	dumpOf := fcfg.GetBool("app.router.dumpRoutes")
	engine := app.Server.HttpServer.Handler.(*gin.Engine)
	for path, r := range app.Router {
		if r.Method == "" {
			continue
		}
		handlersChain := gin.HandlersChain{}
		if r.Sign {
			if app.Sign.SignHandler != nil {
				handlersChain = append(handlersChain, app.Sign.SignHandler)
			} else {
				handlersChain = append(handlersChain, signCheck(app))
			}
		}
		if r.Login {
			if app.Session.LoginHandler != nil {
				handlersChain = append(handlersChain, app.Session.LoginHandler)
			} else {
				handlersChain = append(handlersChain, loginCheck(app))
			}
		}
		handlersChain = append(handlersChain, callRouteHandler(r))
		for _, method := range strings.Split(r.Method, ",") {
			switch fstr.ToUpper(method) {
			case http.MethodGet:
				engine.GET(path, handlersChain...)
			case http.MethodPost:
				engine.POST(path, handlersChain...)
			}
		}
		if dumpOf {
			route.Dumper.Append("/"+path, r.Method, getRouteHandlerFullName(r))
		}
	}
}

func getRouteHandlerFullName(route *frouter.Route) string {
	pkgPath := reflect.TypeOf(route.Provider()).Elem().PkgPath()
	return fmt.Sprintf("%s.%s", pkgPath, route.HandlerName())
}

func callRouteHandler(route *frouter.Route) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handlerReq := reflect.New(route.HandlerReq())
		switch ctx.Request.Method {
		case http.MethodPost:
			if err := ctx.ShouldBindWith(handlerReq.Interface(), binding.FormPost); err != nil {
				fctx.AbortJson(ctx, http.StatusInternalServerError, err)
				return
			}
		case http.MethodGet:
			if err := ctx.ShouldBindWith(handlerReq.Interface(), binding.Query); err != nil {
				fctx.AbortJson(ctx, http.StatusInternalServerError, err)
				return
			}
		}
		results := route.Handler().Call([]reflect.Value{reflect.ValueOf(ctx), handlerReq})
		handleRouteResponse(ctx, results)
	}
}

func handleRouteResponse(ctx *gin.Context, values []reflect.Value) {
	if len(values) != 2 {
		return
	}
	result, err := values[0], values[1]
	if !err.IsNil() {
		switch e := err.Interface().(type) {
		case *ferror.Error:
			fctx.AbortJson(ctx, e.Code(), e.Msg())
		default:
			fctx.AbortJson(ctx, http.StatusInternalServerError, e)
		}
		return
	}
	fctx.AbortJson(ctx, http.StatusOK, "OK", result.Interface())
}
