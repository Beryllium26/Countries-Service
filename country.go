package main

import (
   "encoding/json"
   "fmt"
   "net/http"
//   "net/url"
//   "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the countries service, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.HandleFunc("/api/v1/countries", countriesservice)
	http.ListenAndServe("127.0.0.1:8083", nil)
}

type Message struct {
	Text string
}

func countriesservice (w http.ResponseWriter, r *http.Request) {

// type sCountries struct {
//	Name	string	`json:"name"`
//	ISO	string	`json:"isoCode"`
// }

 //var sco sCountries

 sco := map[string]interface{}{
	"1": map[string]interface{}{
	  "isoCode": "GB",
	  "name": "United Kingdom",
	},
	"2": map[string]interface{}{
          "isoCode": "FR",
	  "name": "France",
        },
 }

 dCountries := map[string]map[string]string{
  "1": map[string]string{
        "name":"United Kingdom",
        "isoCode":"GB",
  },
  "2": map[string]string{
        "name":"France",
        "isoCode":"FR",
  },
  "3": map[string]string{
	"name":"Ireland",
	"isoCode":"IE",
  },
  "4": map[string]string{
        "name":"Spain",
        "isoCode":"ES",
  },
 }

	fmt.Println("GET params were:", r.URL.Query());
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	param1 := r.URL.Query().Get("target")
	if param1 != "" {
	if param1 == "source" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
//		if scone, ok := sco["1"]; ok {
//		    fmt.Println(scone["name"], scone["isoCode"])
//		}
		json.NewEncoder(w).Encode(sco)
		return
	}
	if param1 == "destination" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(dCountries)
		return
	}
	}
	w.WriteHeader(404)

}

func about (w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to the CountriesServiceAPI, build v0.0.1, 07/06/2017"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
