// +build appengine

package soccer

import (

	"io"
	"log"

	//"encoding/json"

	"net/http"
	"text/template"

	"appengine"


 	"github.com/gorilla/schema"
)



func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "Not Found")
}




type Params struct {
	IsDevAppServer  bool
}







func handleMainPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "text/html")

	params := Params{
		
		IsDevAppServer : appengine.IsDevAppServer(),
	}
		

	var index = template.Must(template.ParseFiles("view/index.html"))

	 err := index.Execute(w, params)
     if err != nil {
     	log.Fatalf("template execution: %s", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
     }

}




type Location struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Altitude  string `json:"altitude"`
	Time      string `json:"time"`
	Accuracy  string `json:"accuracy"`

}


func handleLocation(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)

 	err := r.ParseForm()
	if err != nil {
        c.Errorf("Can not parse form")	
    }

	mLocation := new(Location)

	decoder := schema.NewDecoder()

    err = decoder.Decode(mLocation, r.PostForm)
    if err != nil {
        c.Errorf("Can not decode form")	
    }


	c.Infof("Location: %v", r.PostForm)
			
}



type Position struct {
	Azimut  string `json:"azimut"`
	Pitch   string `json:"pitch"`
	Roll    string `json:"roll"`
}


func handlePosition(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)

 	err := r.ParseForm()
	if err != nil {
        c.Errorf("Can not parse form")	
    }

	mPosition := new(Position)

	decoder := schema.NewDecoder()

    err = decoder.Decode(mPosition, r.PostForm)
    if err != nil {
        c.Errorf("Can not decode form")	
    }


	c.Infof("Position: %v", r.PostForm)
			
}


type Height struct {
	Altitude  string `json:"altitude"`
	Pressure  string `json:"pressure"`
	
}


func handleHeight(w http.ResponseWriter, r *http.Request){

	c := appengine.NewContext(r)

 	err := r.ParseForm()
	if err != nil {
        c.Errorf("Can not parse form")	
    }

	mHeight := new(Height)

	decoder := schema.NewDecoder()


    err = decoder.Decode(mHeight, r.PostForm)
    if err != nil {
        c.Errorf("Can not decode form")	
    }


	//c.Infof("mHeight: %v", r.PostForm)
			
}



func init() {

	http.HandleFunc("/location", handleLocation)
	http.HandleFunc("/position", handlePosition)
	http.HandleFunc("/height",   handleHeight)
	
	http.HandleFunc("/", handleMainPage)
	 

}



