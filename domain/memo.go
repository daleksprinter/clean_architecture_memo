package domain

type Memo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Memos []Memo
