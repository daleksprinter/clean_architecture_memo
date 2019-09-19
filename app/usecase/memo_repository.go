package usecase

import (
	"../domain"
)

type MemoRepository interface {
	Store(u *domain.Memo) (int, error)
	FindAll() (domain.Memos, error)
}
