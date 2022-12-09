package routing

import (
    "fmt"
    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "src/databases"
)

type Tree struct {
    gorm.Model
    Name    string `json:"name"`
}

func (t Tree) String() string {
    return fmt.Sprintf("Name:%s \n ",
        t.Name)
}

func TreeGet() echo.HandlerFunc {
    return func(c echo.Context) error {
        db := databases.GormConnect()
        defer db.Close()

        text1 := searchTree(db)
        return c.JSON(200, text1)
    }
}

func searchTree(db *gorm.DB) []Tree {
    var trees []Tree
    db.Raw("SELECT * FROM trees").Scan(&trees)
    return trees
}
