package gtd_bot

import "github.com/go-pg/pg"

type repo struct {
	tx *pg.Tx
}

func newRepo(db *pg.DB) *repo {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	return &repo{tx: tx}
}

func (r *repo) finalizeTransaction() {
	if err := recover(); err != nil {
		r.tx.Rollback()
		panic(err)
	} else {
		err := r.tx.Commit()
		if err != nil {
			panic(err)
		}
	}
}

func (r *repo) findUser(userID int) *User {
	user := &User{ID: userID}
	_, err := r.tx.Model(user).
		Column("user.*", "CurrentInboxItem", "CurrentGoal").
		SelectOrInsert()
	if err != nil {
		panic(err)
	}
	return user
}

func (r *repo) update(obj interface{}) {
	_, err := r.tx.Model(obj).Update()
	if err != nil {
		panic(err)
	}
}

func (r *repo) insert(obj interface{}) {
	_, err := r.tx.Model(obj).Insert()
	if err != nil {
		panic(err)
	}
}

func (r *repo) inboxItemToProcess(userID int) *InboxItem {
	var inboxItemsToProcess []InboxItem
	err := r.tx.Model(&inboxItemsToProcess).
		Where("user_id = ? AND processed_at IS NULL", userID).
		Order("created_at ASC").
		Limit(1).
		Select()
	if err != nil {
		panic(err)
	}
	if len(inboxItemsToProcess) > 0 {
		inboxItemToProcess := inboxItemsToProcess[0]
		return &inboxItemToProcess
	}
	return nil
}

func (r *repo) inboxCount(userID int) int {
	count, err := r.tx.Model(&InboxItem{}).
		Where("user_id = ? AND processed_at IS NULL", userID).
		Count()
	if err != nil {
		panic(err)
	}
	return count
}
