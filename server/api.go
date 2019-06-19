package main

import (
	"encoding/json"
	"net/http"
)

type APIErrorResponse struct {
	ID         string `json:"id"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (p *Plugin) writeAPIError(w http.ResponseWriter, err *APIErrorResponse) {
	b, _ := json.Marshal(err)
	w.WriteHeader(err.StatusCode)
	if _, err := w.Write(b); err != nil {
		p.API.LogError("can't write api error http response", "err", err.Error())
	}
}

func (p *Plugin) writeAPIResponse(w http.ResponseWriter, resp interface{}) {
	b, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		p.API.LogError("Error encoding JSON response", "err", jsonErr.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Encountered an unexpected error. Please try again.", StatusCode: http.StatusInternalServerError})
	}
	if _, err := w.Write(b); err != nil {
		p.API.LogError("can't write response user to http", "err", err.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Encountered an unexpected error. Please try again.", StatusCode: http.StatusInternalServerError})
	}
}
