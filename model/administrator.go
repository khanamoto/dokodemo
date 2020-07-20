package model

type Administrator struct {
	ID      uint64 `db:"id"`
	UserID  uint64 `db:"user_id"`
	EventID uint64 `db:"event_id"`
}
