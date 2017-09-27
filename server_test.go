package main

import (
    "encoding/json"
    "github.com/omise/omise-go"
    "testing"
    "net/http"
    "net/http/httptest"
    "os"
)

func TestCustomHeader(t *testing.T) {
    // This is what will be called when the request arrives
    skey := os.Getenv("OMISE_SKEY")
  	pkey := os.Getenv("OMISE_PKEY")
    // Now let's instantiate a client and tell it to do its request to our fake server
    clientOmise, e := omise.NewClient(pkey, skey)
    if(e != nil){
      t.Error("Failed the test")
    }

    handler :=&TamboonHandler{clientOmise}
    req, err := http.NewRequest("GET", "http://localhost:8080/", nil)

    if(err != nil){
      t.Error("Failed the test in GET ")
    }

    recorder := httptest.NewRecorder()
    return_value := `[
     { "id": 0, "name": "Ban Khru Noi", "logo_url": "http://rkdretailiq.com/news/img-corporate-baankrunoi.jpg" },
     { "id": 1, "name": "Habitat for Humanity Thailand", "logo_url": "http://www.adamandlianne.com/uploads/2/2/1/6/2216267/3231127.gif" },
     { "id": 2, "name": "Paper Ranger", "logo_url": "https://myfreezer.files.wordpress.com/2007/06/paperranger.jpg" },
     { "id": 3, "name": "Makhampom", "logo_url": "http://www.makhampom.net/makhampom/ppcms/uploads/UserFiles/Image/Thai/T14Publice/2554/January/Newyear/logoweb.jpg" }
   ]`

    handler.GET(recorder, req )
    if e := json.NewEncoder(recorder).Encode(charities); e != nil {
      t.Errorf("handler returned unexpected body: got %v want %v",
          recorder.Body.String(), return_value)
  	}


}
