package api

import (
	"github.com/sunwei/hugoverse/internal/application"
	"net/http"
)

func (s *Server) handleConfig(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	allConfigs, err := application.AllConfigurationInformation(s.ProjPath)
	if err != nil {
		s.writeJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.writeJSONResponse(w, allConfigs, http.StatusOK)
}
