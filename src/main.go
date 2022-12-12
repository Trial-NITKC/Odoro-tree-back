package main

import (
    "github.com/labstack/echo"
    "src/routing"
)

func main() {
    e := echo.New()

    e.GET("/trees", routing.GetTrees)
    e.POST("/trees", routing.PostTree)
    e.GET("/trees/:id", routing.GetTree)
    e.PUT("/trees/:id", routing.PutTree)
    e.DELETE("/trees/:id", routing.DeleteTree)
    e.GET("/branches", routing.GetBranches)
    e.POST("/branches", routing.PostBranch)
    e.GET("/branches/:id", routing.GetBranch)
    e.PUT("/branches/:id", routing.PutBranch)
    e.DELETE("/branches/:id", routing.DeleteBranch)
    e.GET("/leaves", routing.GetLeaves)
    e.POST("/leaves", routing.PostLeaf)
    e.GET("/leaves/:id", routing.GetLeaf)
    e.PUT("/leaves/:id", routing.PutLeaf)
    e.DELETE("/leaves/:id", routing.DeleteLeaf)

    e.GET("/trees/:id/childbranches", routing.GetChildBranches)
    e.GET("/branches/:id/childleaves", routing.GetChildLeaves)

    e.Logger.Fatal(e.Start(":8080"))
}
