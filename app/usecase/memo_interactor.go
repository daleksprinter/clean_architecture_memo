package usecase

import (
	"../domain"

	"fmt"
)

type MemoInteractor struct {
	MemoRepository MemoRepository
}

func NewMemoInteractor(memorepository MemoRepository) *MemoInteractor {
	return &MemoInteractor{
		MemoRepository: memorepository,
	}
}

func (interactor MemoInteractor) Add(m *domain.Memo) (err error) {
	id, err := interactor.MemoRepository.Store(m)
	fmt.Println("isnerted id :", id)
	return
}

func (interactor MemoInteractor) FindAll() (memos []domain.Memo, err error) {
	memos, err = interactor.MemoRepository.FindAll()
	return
}
