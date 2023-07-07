package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

//The purpose of this function is to decode the JSON body of an HTTP request and map it to a Go struct or any other type that matches the 
//structure of the JSON data. The x parameter is passed as a pointer to the target object, so the unmarshaling operation can modify its value and 
//populate it with the data from the JSON body.
func ParseBody(r *http.Request,x interface{})  {
	if body, err:=io.ReadAll(r.Body); err==nil{
		if err := json.Unmarshal([]byte(body), x); err!=nil{
			return
		}
	}
}