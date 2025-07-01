package handlers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// CleanURI API for shortening URLs
// https://cleanuri.com/docs

// Structure to receive the JSON request with the URL to shorten
type ShortenRequest struct {
    URL string `json:"url"`
}

// HTTP handler to shorten URLs using the CleanURI API
func ShortenURL(w http.ResponseWriter, r *http.Request) {
    fmt.Println("ShortenURL Called") // Log to know the handler was called

    // Decode the request body expecting a JSON with the "url" field
    var reqData ShortenRequest
    if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    fmt.Printf("URL to shorten: %s\n", reqData.URL) // Log the received URL

    // Prepare the form body for the request to CleanURI
    form := []byte("url=" + reqData.URL)
    // Create the HTTP POST request to the CleanURI API
    req, err := http.NewRequestWithContext(r.Context(), http.MethodPost, "https://cleanuri.com/api/v1/shorten", bytes.NewBuffer(form))
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }
    // Set the header to indicate the body is a form
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // Make the request to CleanURI
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        http.Error(w, "Failed to contact CleanURI", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // If CleanURI responds with a code other than 200 OK, return the error
    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        http.Error(w, fmt.Sprintf("CleanURI error: %s", string(body)), http.StatusInternalServerError)
        return
    }

    // If everything goes well, copy CleanURI's response to the client
    w.Header().Set("Content-Type", "application/json")
    io.Copy(w, resp.Body)
}