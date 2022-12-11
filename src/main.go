package main

import (
    "github.com/labstack/echo"
    "src/routing"
)

func main() {
    e := echo.New()

    e.GET("/trees", routing.GetTree)
    e.POST("/trees", routing.PostTree)
    e.PUT("/trees/:id", routing.PutTree)
    e.DELETE("/trees/:id", routing.DeleteTree)
    e.GET("/branches", routing.GetBranch)
    e.POST("/branches", routing.PostBranch)
    e.PUT("/branches/:id", routing.PutBranch)
    e.DELETE("/branches/:id", routing.DeleteBranch)
    e.GET("/leaves", routing.GetLeaf)
    e.POST("/leaves", routing.PostLeaf)
    e.PUT("/leaves/:id", routing.PutLeaf)
    e.DELETE("/leaves/:id", routing.DeleteLeaf)

    e.GET("/trees/:id/childbranches", routing.GetChildBranches)
    e.GET("/branches/:id/childleaves", routing.GetChildLeaves)

    e.Logger.Fatal(e.Start(":8080"))
}
