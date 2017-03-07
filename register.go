package main
import(
    "net/http"
    "github.com/labstack/echo"
    "helloGo/goMsSql/service"
)

func RegisterApi()  {
    e.GET("/",func (c echo.Context) error {
        return c.String(http.StatusOK,"hello mssql")
    })
    api :=e.Group("/api")
    v1 :=api.Group("/v1")

    user :=v1.Group("/user")
    user.GET("",service.Query)
}