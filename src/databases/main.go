package databases

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() (database *gorm.DB) {
    DBMS := "mysql"
    PROTOCOL := "tcp(db)"
    DBNAME := "golang"
    USER := "odoro-tree"
    PASS := "password"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("DB connected")
    }
    return db
}
