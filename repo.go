package gtd_bot

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

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
		Column(
			"user.*",
			"CurrentInboxItem",
			"CurrentGoal",
			"CurrentAction",
			"CurrentAction.Goal",
			"CurrentWaitingFor").
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

func (r *repo) userInboxItemScope(userID int) *orm.Query {
	return r.tx.Model(&InboxItem{}).
		Where(`user_id = ?
			AND processed_at IS NULL`, userID)
}

func (r *repo) userActionScope(userID int) *orm.Query {
	return r.tx.Model(&Action{}).
		Where(`action.user_id = ?
			AND action.completed_at IS NULL
			AND action.dropped_at IS NULL
			AND action.reviewed_at < current_timestamp - interval '10 minutes'`, userID)
}

// func (r *repo) userGoalScope(userID int) *orm.Query {
// }

func (r *repo) userWaitingForScope(userID int) *orm.Query {
	return r.tx.Model(&WaitingFor{}).
		Where(`waiting_for.user_id = ?
			AND waiting_for.completed_at IS NULL
			AND waiting_for.dropped_at IS NULL
			AND waiting_for.reviewed_at < current_timestamp - interval '1 day'`, userID)
}

func (r *repo) inboxCount(userID int) int {
	count, err := r.userInboxItemScope(userID).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func (r *repo) actionCount(userID int) int {
	count, err := r.userActionScope(userID).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func (r *repo) waitingForCount(userID int) int {
	count, err := r.userWaitingForScope(userID).Count()
	if err != nil {
		panic(err)
	}
	return count
}

func (r *repo) inboxItemToProcess(userID int) *InboxItem {
	var inboxItemsToProcess []InboxItem
	err := r.userInboxItemScope(userID).
		Order("created_at ASC").
		Limit(1).
		Select(&inboxItemsToProcess)
	if err != nil {
		panic(err)
	}
	if len(inboxItemsToProcess) > 0 {
		inboxItemToProcess := inboxItemsToProcess[0]
		return &inboxItemToProcess
	}
	return nil
}

func (r *repo) actionToDo(userID int) *Action {
	var actionsToDo []Action
	err := r.userActionScope(userID).
		Order("reviewed_at ASC").
		Limit(1).
		Column("action.*", "Goal").
		Select(&actionsToDo)
	if err != nil {
		panic(err)
	}
	if len(actionsToDo) > 0 {
		actionToDo := actionsToDo[0]
		return &actionToDo
	}
	return nil
}

func (r *repo) waitingForToCheck(userID int) *WaitingFor {
	var waitingForsToCheck []WaitingFor
	err := r.userWaitingForScope(userID).
		Order("reviewed_at ASC").
		Limit(1).
		Column("waiting_for.*", "Goal").
		Select(&waitingForsToCheck)
	if err != nil {
		panic(err)
	}
	if len(waitingForsToCheck) > 0 {
		actionToDo := waitingForsToCheck[0]
		return &actionToDo
	}
	return nil
}

func (r *repo) contexts(userID int) []*Context {
	var contexts []*Context
	err := r.tx.Model(&contexts).
		Where("user_id=?", userID).
		Select()
	if err != nil {
		panic(err)
	}
	return contexts
}

func (r *repo) dropGoalActionsAndWaitingFors(goalID int) {
	_, err := r.tx.Model(&Action{}).
		Set("dropped_at = ?", time.Now()).
		Where("goal_id = ? AND completed_at IS NULL AND dropped_at IS NULL", goalID).
		Update()
	if err != nil {
		panic(err)
	}
	_, err = r.tx.Model(&WaitingFor{}).
		Set("dropped_at = ?", time.Now()).
		Where("goal_id = ? AND completed_at IS NULL AND dropped_at IS NULL", goalID).
		Update()
	if err != nil {
		panic(err)
	}
}
