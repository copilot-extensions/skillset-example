package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func Weather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Weather Called")
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "https://api.weatherapi.com/v1/current.json?key=YOUR_API_KEY&q=London", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		fmt.Println("Something unrecoverable went wrong")
		return
	}
}
