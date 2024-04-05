package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/sensors", getSensors)
    router.GET("/sensors/:id", getsensorByID)
    router.POST("/sensors", postSensor)

    fmt.Println("Server will run on http://localhost:8000")
    router.Run(":8000")
}

type Sensor struct {
    Sensor string `json:"sensor"`
    Value  string `json:"value"`
}

func getSensors(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, GetSensors())
}


func postSensor(c *gin.Context) {
    var data Sensor

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, PostSensor(data.Sensor, data.Value))
}

func getsensorByID(c *gin.Context) {

    id := c.Param("id")

    sensors := GetSensors()

    for _, a := range sensors {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sensor not found"})
}