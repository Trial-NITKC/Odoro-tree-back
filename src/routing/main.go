package routing

import (
    "fmt"
    "github.com/labstack/echo"

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

    var trees []Tree
    db.Find(&trees)

    return c.JSON(200, trees)
}

func PostTree(c echo.Context) error {
    db:= databases.GormConnect()
    defer db.Close()

    tree := new(Tree)
    if err := c.Bind(tree); err != nil {
        return err
    }

    db.Create(&tree)
    db.Where(&tree).Find(&tree)

    return c.JSON(201, tree)
}

func PutTree(c echo.Context) error {
    db := databases.GormConnect()
    defer db.Close()

    tree := new(Tree)
    id := c.Param("id")
    db.Where("tree_id = ?", id).Find(&tree)

    updateValues := new(Tree)
    if err := c.Bind(updateValues); err != nil {
        return err
    }

    db.Model(&tree).Updates(updateValues)
    db.Where(&tree).Find(&tree)

    return c.JSON(200, tree)
}

func DeleteTree(c echo.Context) error {
    db:= databases.GormConnect()
    defer db.Close()
    
    id := c.Param("id")
    db.Delete(&Tree{}, "tree_id = ?", id)

    return c.NoContent(204)
}

func GetBranch(c echo.Context) error {
    db := databases.GormConnect()
    defer db.Close()

    var branches []Branch
    db.Find(&branches)

    return c.JSON(200, branches)
}

func PostBranch(c echo.Context) error {
    db:= databases.GormConnect()
    defer db.Close()

    branch := new(Branch)
    if err := c.Bind(branch); err != nil {
        return err
    }

    db.Create(&branch)
    db.Where(&branch).Find(&branch)

    return c.JSON(201, branch)
}

func PutBranch(c echo.Context) error {
    db := databases.GormConnect()
    defer db.Close()

    branch := new(Branch)
    id := c.Param("id")
    db.Where("branch_id = ?", id).Find(&branch)

    updateValues := new(Branch)
    if err := c.Bind(updateValues); err != nil {
        return err
    }

    db.Model(&branch).Updates(updateValues)
    db.Where(&branch).Find(&branch)

    return c.JSON(200, branch)
}

func DeleteBranch(c echo.Context) error {
    db:= databases.GormConnect()
    defer db.Close()
    
    id := c.Param("id")
    db.Debug().Delete(&Branch{}, "branch_id = ?", id)

    return c.NoContent(204)
}

func GetLeaf(c echo.Context) error {
    db := databases.GormConnect()
    defer db.Close()

    var leaves []Leaf
    db.Find(&leaves)

    return c.JSON(200, leaves)
}

func PostLeaf(c echo.Context) error {
    db:= databases.GormConnect()
    defer db.Close()

    leaf := new(Leaf)
    if err := c.Bind(leaf); err != nil {
        return err
    }

    db.Create(&leaf)
    db.Where(&leaf).Find(&leaf)

    return c.JSON(201, leaf)
}

func PutLeaf(c echo.Context) error {
    db := databases.GormConnect()
    defer db.Close()

    leaf := new(Leaf)
    id := c.Param("id")
    db.Where("leaf_id = ?", id).Find(&leaf)

    updateValues := new(Leaf)
    if err := c.Bind(updateValues); err != nil {
        return err
    }

    db.Model(&leaf).Updates(updateValues)
    db.Where(&leaf).Find(&leaf)

    return c.JSON(200, leaf)
}

func DeleteLeaf(c echo.Context) error {
    db:= databases.GormConnect()
    defer db.Close()
    
    id := c.Param("id")
    db.Debug().Delete(&Leaf{}, "leaf_id = ?", id)

    return c.NoContent(204)
}
