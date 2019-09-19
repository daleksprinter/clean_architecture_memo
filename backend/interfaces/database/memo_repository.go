package database

import (
	"github.com/daleksprinter/clean_todo/backend/domain"
	"github.com/jinzhu/gorm"
)

type MemoRepository struct {
	Conn *gorm.DB
}

func NewMemoRepository(c *gorm.DB) *MemoRepository {
	return &MemoRepository{
		Conn: c,
	}
}

type Memo struct {
	gorm.Model
	Text string `json:"text"`
}

func (db *MemoRepository) Store(memo *domain.Memo) (id int, err error) {
	m := &Memo{
		Text: memo.Text,
	}

	if err = db.Conn.Create(m).Error; err != nil {
		return
	}

	return int(m.ID), err
}

func (db *MemoRepository) FindAll() (memos domain.Memos, err error) {
	var m []Memo

	if err = db.Conn.Find(&m).Error; err != nil {
		return
	}

	for _, memo := range m {
		newmemo := &domain.Memo{
			ID:   int(memo.ID),
			Text: memo.Text,
		}
		memos = append(memos, *newmemo)
	}

	return
}
