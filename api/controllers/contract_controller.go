package controllers

import (
	"encoding/json"
	"github.com/knuckerr/go_rest/api/models"
	"github.com/knuckerr/go_rest/api/responses"
	"github.com/knuckerr/go_rest/api/validators"
	"net/http"
)

func (server *Server) GetContracts(w http.ResponseWriter, r *http.Request) {
	contract := models.Contract{}
	contracts, err := contract.Contracts(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, &contracts)

}

func (server *Server) CreateContract(w http.ResponseWriter, r *http.Request) {
	contract := models.Contract{}
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	err = validators.New(contract)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	contract_created, err := contract.SaveContract(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, contract_created)

}
