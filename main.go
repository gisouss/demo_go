package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Technology type
type technology struct {
    ID string `json:"id"`
    Tech string `json:"tech"`
    Rate string `json:"rate"`
}

// Technologies data
var technologies = []technology{
    {ID: "1", Tech: "Python", Rate: "5/5"},
    {ID: "2", Tech: "JavaScript", Rate: "4/5"},
    {ID: "3", Tech: "C#", Rate: "3/5"},
}

// getTechnologies responds with the list of all technologies as JSON.
func getTechnologies(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, technologies)
}

// postTechnologies adds a technology from JSON received in the request body.
func postTechnologies(c *gin.Context) {
    var newTechnology technology

    // Call BindJSON to bind the received JSON to
    // newTechnology
    if err := c.BindJSON(&newTechnology); err != nil {
        return
    }

    // Add the new technology to the slice.
    technologies = append(technologies, newTechnology)
    c.IndentedJSON(http.StatusCreated, newTechnology)
}

// getTechnologyByID locates the technology whose ID value matches the id
// parameter sent by the client, then returns that technology as a response.
func getTechnologyByID(c *gin.Context) {
    id := c.Param("id")

	for _, a := range technologies {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Technology not found"})
}

// deleteTechnologyByID deletes the technology whose ID value matches the id
func deleteTechnologyByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range technologies {
        if a.ID == id {
            copy(technologies[i:], technologies[i+1:])
            technologies[len(technologies)-1] = technology{}
            technologies = technologies[:len(technologies)-1]
            c.IndentedJSON(http.StatusOK, "Technology with ID: " + id + " has been successfully deleted !")
            return
        }
    }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Technology not found"})
}

// patchTechnologyByID updates the technology whose ID value matches the id
func patchTechnologyByID(c *gin.Context) {
	id := c.Param("id")
    var newTechnology technology

    // Call BindJSON to bind the received JSON to
    // newTechnology
    if err := c.BindJSON(&newTechnology); err != nil {
        return
    }

	for i, a := range technologies {
        if a.ID == id {
            technologies[i] = newTechnology
            c.IndentedJSON(http.StatusOK, "Technology with ID: " + id + " has been successfully updated !")
            return
        }
    }
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Technology not found"})
}

func main() {
	router := gin.Default()
    router.GET("/technologies", getTechnologies)
	router.GET("/technologies/:id", getTechnologyByID)
	router.POST("/technologies", postTechnologies)
	router.PATCH("/technologies/:id", patchTechnologyByID)
    router.DELETE("/technologies/:id", deleteTechnologyByID)
    router.Run("localhost:8081")
}