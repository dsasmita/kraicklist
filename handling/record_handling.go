package handling

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"challenge.haraj.com.sa/kraicklist/service"
)

type handlingSearch struct {
	serviceRecord service.RecordServiceInterface
}

func HandleSearch(services service.RecordServiceInterface) http.HandlerFunc {
	handler := &handlingSearch{
		serviceRecord: services,
	}

	return handler.GetSearch
}

func (v handlingSearch) GetSearch(w http.ResponseWriter, r *http.Request) {
	// fetch query string from query params
	q := r.URL.Query().Get("q")
	log.Println(q)
	if len(q) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("missing search query in query params"))
		if err != nil {
			log.Println(err)
		}
		return
	}

	records, err := v.serviceRecord.Search(q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := w.Write([]byte(err.Error()))
		if err2 != nil {
			log.Fatal(err2)
		}
		log.Fatal(err)
	}

	// output success response
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	err = encoder.Encode(records)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
