package model
import(
    "database/sql"
    "fmt"
    _"github.com/mattn/go-adodb"
)
type MsSql struct{
    *sql.DB
    Conn string
}
func (m *MsSql) Open()(err error)  {
    m.DB,err=sql.Open("adodb","Provider=SQLOLEDB;Server=127.0.0.1,1433; Initial Catalog=UserService;User ID=sa;password=1234;")
    if err !=nil{
        fmt.Println("mssql connect error")
        return err
    }
    return nil
}