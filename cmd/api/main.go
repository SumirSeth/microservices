package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/healthy", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "success": true,
        })
    })

    if err := r.Run(":8080"); err != nil {
        panic(err)
    }

    r.Run(":8080")
}