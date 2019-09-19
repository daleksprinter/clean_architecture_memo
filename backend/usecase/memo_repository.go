package usecase

import (
	"github.com/daleksprinter/clean_todo/backend/domain"
)

type MemoRepository interface {
	Store(u *domain.Memo) (int, error)
	FindAll() (domain.Memos, error)
}
