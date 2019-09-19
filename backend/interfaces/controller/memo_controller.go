package controller

import (
	"net/http"

	"encoding/json"

	"github.com/daleksprinter/clean_todo/backend/domain"
	"github.com/daleksprinter/clean_todo/backend/usecase"
	"github.com/daleksprinter/clean_todo/backend/interfaces/database"
	"github.com/jinzhu/gorm"
)

type MemoController struct {
	MemoInteractor usecase.MemoInteractor
}

func NewMemoController(conn *gorm.DB) *MemoController {
	return &MemoController{
		MemoInteractor: *usecase.NewMemoInteractor(
			database.NewMemoRepository(conn),
		),
	}

}

func (m MemoController) Create(w http.ResponseWriter, r *http.Request) {
	var memo domain.Memo

	if err := json.NewDecoder(r.Body).Decode(&memo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured"))
		return
	}

	if err := m.MemoInteractor.Add(&memo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured"))
		return
	}
}

func (m MemoController) FindAll(w http.ResponseWriter, r *http.Request) {
	memos, err := m.MemoInteractor.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured"))
		return
	}

	if err = json.NewEncoder(w).Encode(memos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured"))
		return
	}
}
