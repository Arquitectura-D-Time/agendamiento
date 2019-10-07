package agendadas_controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	driver "agendamiento/common"
	repository "agendamiento/data"
	agendadas "agendamiento/data/agendadas_mysql"
	model "agendamiento/model"
)

func NewAgendadasHandler(db *driver.DB) *Agendadas {
	return &Agendadas{
		repo: agendadas.NewSQLAgendadas(db.SQL),
	}
}

type Agendadas struct {
	repo repository.Agendadas
}

func (p *Agendadas) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Agendadas) Create(w http.ResponseWriter, r *http.Request) {
	agendadas := model.Agendadas{}
	json.NewDecoder(r.Body).Decode(&agendadas)

	newID, err := p.repo.Create(r.Context(), &agendadas)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	fmt.Println(err)
}

func (p *Agendadas) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	data := model.Agendadas{IDtutoria: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Agendadas) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	fmt.Println(id)
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Agendadas) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "IDtutoria"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
