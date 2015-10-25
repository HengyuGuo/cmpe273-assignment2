package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
)

// type (
// 	Results struct{
// 		Geometry struct{
// 			Location struct{
// 				Lat float32 		`json:"lat"`
// 				Lng float32 		`json:"lng"`
// 			}						`json:"location"`
// 		}							`json:"geometry"`	
// 	}								`json:"results"`
// )

func GetLocation(Symbol string) string{
	// "http://maps.google.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&sensor=false"
	urlPath :=  "http://maps.google.com/maps/api/geocode/json?address="
	urlPath += "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
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
	// u := results{}
	// fmt.Println(string(body[:]))
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
	fmt.Println(lat)
	fmt.Println(lng)
	return ""
}

func getURL() {
	str := "1600 Amphitheatre Parkway"
	city := "Mountain View"
	state := "CA"
	add := strings.Split(str, " ")
	var res string
	for i:=0;i< len(add);i++ {	
		if i==len(add)-1 {
			res = res + add[i] + ","
		}else{
			res = res + add[i] +"+"
		}
	}
	c := strings.Split(city, " ")
	for i:=0;i<len(c);i++ {
		if i==len(c)-1{
			res = res + "+" + c[i] + ","
		}else{
			res = res + "+" + c[i]
		}
	}
	res = res + "+" + state
	fmt.Println(res)
}

func main(){
	// "http://maps.google.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&sensor=false"
	urlPath :=  "http://maps.google.com/maps/api/geocode/json?address="
	urlPath += "1600+Amphitheatre+Parkway,+Mountain+View,+CA"
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
	// u := results{}
	// fmt.Println(string(body[:]))
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
	fmt.Println(lat)
	fmt.Println(lng)
}