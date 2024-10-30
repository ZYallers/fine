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
		chain := gin.HandlersChain{}
		if r.Sign {
			if app.Sign.SignHandler != nil {
				chain = append(chain, app.Sign.SignHandler)
			} else {
				chain = append(chain, signCheck(app))
			}
		}
		if r.Login {
			if app.Session.LoginHandler != nil {
				chain = append(chain, app.Session.LoginHandler)
			} else {
				chain = append(chain, loginCheck(app))
			}
		}
		chain = append(chain, callRouteHandler(r))
		for _, method := range strings.Split(r.Method, ",") {
			switch fstr.ToUpper(method) {
			case http.MethodGet:
				engine.GET(path, chain...)
			case http.MethodPost:
				engine.POST(path, chain...)
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
		if err := ctx.ShouldBind(handlerReq.Interface()); err != nil {
			fctx.AbortJson(ctx, http.StatusInternalServerError, err)
			return
		}
		results := route.Handler().Call([]reflect.Value{reflect.ValueOf(ctx), handlerReq})
		handleRouteResponse(ctx, results)
	}
}

func handleRouteResponse(ctx *gin.Context, resp []reflect.Value) {
	if len(resp) != 2 {
		return
	}
	res, err := resp[0], resp[1]
	if !err.IsNil() {
		switch e := err.Interface().(type) {
		case *ferror.Error:
			fctx.AbortJson(ctx, e.Code(), e.Msg())
		default:
			fctx.AbortJson(ctx, http.StatusInternalServerError, e)
		}
		return
	}
	fctx.AbortJson(ctx, http.StatusOK, "OK", res.Interface())
}
