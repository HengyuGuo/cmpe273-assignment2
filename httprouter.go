package main

import (  
    "fmt"
    "net/http"
    "os"
    "assignment2/controllers"
    // Third party packages
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"
)



func getSession() *mgo.Session {  
    // Connect to our local mongo
    // s, err := mgo.Dial("mongodb://localhost")
    // Connect to our remote mongo
    s, err := mgo.Dial("mongodb://admin:123@ds041404.mongolab.com:41404/cmpe273")
    // Check if connection error, is mongo running?
    if err != nil {
        fmt.Println("Can't connect to mongo, go error %v\n", err)
        os.Exit(1)
    }
    return s
}

func main() {
    mux := httprouter.New()
    // Get a UserController instance
    uc := controllers.NewUserController(getSession())
    mux.GET("/locations/:id", uc.GetLocations)
    mux.POST("/locations/", uc.CreateLocations)
    mux.DELETE("/locations/:id", uc.RemoveLocations)
    mux.PUT("/locations/:id", uc.UpdateLocations)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
