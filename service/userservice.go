package service

import(
    "fmt"
    "github.com/labstack/echo"
    "net/http"
    "helloGo/goMsSql/model"
)
func Query(c echo.Context) error {
    db :=model.MsSql{}
    if err  :=db.Open();err !=nil{
        return echo.NewHTTPError(http.StatusInternalServerError,err)
    }
    defer db.Close()
    fmt.Println(db)

    rows,err :=db.Query("select Name,Code from dbo.Users")
    if err !=nil{
        fmt.Println("query excute error")
        fmt.Println(err)
        return echo.NewHTTPError(http.StatusInternalServerError,err)
    }else{
        var user model.User
        var users []model.User
        for rows.Next(){
            rows.Scan(&user.Code,&user.Name)
            users=append(users,user)
        }
        fmt.Printf("%+v",users)
        return c.JSON(http.StatusOK,users)
    }
}
