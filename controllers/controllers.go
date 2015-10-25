package controllers

import (  
    "encoding/json"
    "fmt"
    "strings"
    "net/http"
    "io/ioutil"
    "gopkg.in/mgo.v2/bson"
    "gopkg.in/mgo.v2"
    "github.com/julienschmidt/httprouter"
    "assignment2/models"
)

type (  
    // UserController represents the controller for operating on the User resource
    UserController struct{
        session *mgo.Session
    }
)

func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}

// GetUser retrieves an individual user resource
func (uc UserController) GetLocations(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Stub an example user
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Stub user
    l := models.Location{}

    // Fetch user
    if err := uc.session.DB("cmpe273").C("assignment2").FindId(oid).One(&l); err != nil {
        w.WriteHeader(404)
        return
    }

    // Marshal provided interface into JSON structure
    lj, _ := json.Marshal(l)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", lj)
}

// CreateUser creates a new user resource
func (uc UserController) CreateLocations(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Stub an user to be populated from the body
    l := models.Location{}

    // Populate the user data
    json.NewDecoder(r.Body).Decode(&l)
    str := getURL(l.Address, l.City, l.State)
    getLocation(&l, str)

    // Add an Id
    l.Id = bson.NewObjectId()
    // Write the user to mongo
    uc.session.DB("cmpe273").C("assignment2").Insert(l)

    // Marshal provided interface into JSON structure
    lj, _ := json.Marshal(l)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", lj)
}

// Update removes an existing user resource
func (uc UserController) UpdateLocations(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Stub an example user
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)
    l := models.Location{}
    // Populate the user data
    json.NewDecoder(r.Body).Decode(&l)
    str := getURL(l.Address, l.City, l.State)
    getLocation(&l, str)
    l.Id = oid
    // Write the user to mongo
    if err := uc.session.DB("cmpe273").C("assignment2").Update(bson.M{"_id": l.Id}, bson.M{"$set": bson.M{"address": l.Address, "city": l.City, "state": l.State, "zip": l.Zip, "coordinate.lat": l.Coordinate.Lat, "coordinate.lng": l.Coordinate.Lng}}); err != nil {
        w.WriteHeader(404)
        return
    }
    if err := uc.session.DB("cmpe273").C("assignment2").FindId(oid).One(&l); err != nil {
        w.WriteHeader(404)
        return
    }
    // Marshal provided interface into JSON structure
    lj, _ := json.Marshal(l)

    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", lj)
}

// RemoveUser removes an existing user resource
func (uc UserController) RemoveLocations(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Remove user
    if err := uc.session.DB("cmpe273").C("assignment2").RemoveId(oid); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)

}

// construct queryURL
func getURL(address string, city string, state string) string{
    addStr := address
    cityStr := city
    stateStr := state
    add := strings.Split(addStr, " ")
    var res string
    for i:=0;i< len(add);i++ {  
        if i==len(add)-1 {
            res = res + add[i] + ","
        }else{
            res = res + add[i] +"+"
        }
    }
    c := strings.Split(cityStr, " ")
    for i:=0;i<len(c);i++ {
        if i==len(c)-1{
            res = res + "+" + c[i] + ","
        }else{
            res = res + "+" + c[i]
        }
    }
    res = res + "+" + stateStr
    return res
}

// get data from googleMapAPI
func getLocation(l *models.Location, str string){
    // "http://maps.google.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&sensor=false"
    urlPath :=  "http://maps.google.com/maps/api/geocode/json?address="
    urlPath += str
    urlPath += "&sensor=false"
    res, err := http.Get(urlPath)
    if err!=nil {
        fmt.Println("GetLocation: http.Get",err)
        panic(err)
    }
    defer res.Body.Close()
    body,err := ioutil.ReadAll(res.Body)
    if err!=nil {
        fmt.Println("GetLocation: ioutil.ReadAll",err)
        panic(err)
    }

    mp := make(map[string]interface{})
    err = json.Unmarshal(body, &mp)
    if err!=nil {
        fmt.Println("GetLocation: json.Unmarshal",err)
        panic(err)
    }
    temp := mp["results"].(interface{})
    next := temp.([]interface{})
    geometry := next[0].(map[string]interface{})
    geometry = geometry["geometry"].(map[string]interface{})
    location := geometry["location"].(map[string]interface{})
    lat := location["lat"].(float64)
    lng := location["lng"].(float64)
    l.Coordinate.Lat = lat
    l.Coordinate.Lng = lng
}