package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	callbackData := make([]byte, r.ContentLength)
	r.Body.Read(callbackData)
	log.Printf("Callback: %s\n", callbackData)

	response := map[string]string{"status": "success"}
	jsonRespone, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRespone)
}

func main() {
	http.HandleFunc("/api/v1/Checkout/Callback", callbackHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
