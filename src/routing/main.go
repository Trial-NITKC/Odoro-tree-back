package routing

import (
    "fmt"
    "github.com/labstack/echo"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "src/databases"
)

type Tree struct {
    TreeID  uint `gorm:"primaryKey" json:"id"`
    Name    string `json:"name"`
}

type Branch struct {
    BranchID  uint `gorm:"primaryKey" json:"id"`
    ParentTreeId int `json:"parent_tree_id"`
}

type Leaf struct {
    LeafID  uint `gorm:"primaryKey" json:"id"`
    FrontContent string `json:"front_content"`
    BackContent string `json:"back_content"`
    Rating int `json:"rating"`
    ParentBranchId int `json:"parent_branch_id"`
}

func (t Tree) String() string {
    return fmt.Sprintf("TreeID:%d \n Name:%s \n ",
        t.TreeID,
        t.Name)
}

func (b Branch) String() string {
    return fmt.Sprintf("ParentTreeId:%d \n ",
        b.ParentTreeId)
}

func (l Leaf) String() string {
    return fmt.Sprintf("LeafID:%d \n FrontContent:%s \n BackContent:%s \n Rating:%d \n ParentBranchId:%d \n ",
        l.LeafID,
        l.FrontContent,
        l.BackContent,
        l.Rating,
        l.ParentBranchId)
}

func GetTree(c echo.Context) error {
    db := databases.GormConnect()
    defer db.Close()

    return c.String(200, "")
}

func searchTree(db *gorm.DB) []Tree {
    var trees []Tree
    db.Find(&trees)
    return trees
}
