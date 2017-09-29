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
			"CurrentAction.Context",
			"CurrentWaitingFor",
			"CurrentWaitingFor.Goal").
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

func (r *repo) insert(obj interface{}) error {
	_, err := r.tx.Model(obj).Insert()
	return err
}

func (r *repo) userInboxItemScope(userID int) *orm.Query {
	return r.tx.Model(&InboxItem{}).
		Where(`user_id = ?
			AND processed_at IS NULL`, userID)
}

func (r *repo) inboxItemsProcessedTodayScope(userID int) *orm.Query {
	return r.tx.Model(&InboxItem{}).
		Where(`user_id = ?
			AND processed_at > current_timestamp - interval '1 day'`, userID)
}

func (r *repo) actionsCompletedTodayScope(userID int) *orm.Query {
	return r.tx.Model(&Action{}).
		Where(`user_id = ?
			AND completed_at > current_timestamp - interval '1 day'`, userID)
}

func (r *repo) goalsCreatedTodayScope(userID int) *orm.Query {
	return r.userActiveGoalScope(userID).
		Where("created_at > current_timestamp - interval '1 day'")
}

func (r *repo) userActionScope(userID int) *orm.Query {
	return r.tx.Model(&Action{}).
		Where(`action.user_id = ?
			AND action.completed_at IS NULL
			AND action.dropped_at IS NULL
			AND action.reviewed_at < current_timestamp - interval '10 minutes'`, userID)
}

func (r *repo) userActionToDoScope(userID int) *orm.Query {
	return r.userActionScope(userID).
		Where("action.reviewed_at < current_timestamp - interval '10 minutes'").
		Join("LEFT JOIN contexts ON action.context_id = contexts.id").
		Where("contexts.id IS NULL OR contexts.active")
}

func (r *repo) userActiveGoalScope(userID int) *orm.Query {
	return r.tx.Model(&Goal{}).
		Where(`goal.user_id = ?
			AND goal.completed_at IS NULL
			AND goal.dropped_at IS NULL`, userID)
}

func (r *repo) userGoalToReviewScope(userID int) *orm.Query {
	return r.userActiveGoalScope(userID).
		Where(`goal.reviewed_at < current_timestamp - interval '1 week'`)
}

func (r *repo) userGoalWithNoActionScope(userID int) *orm.Query {
	return r.userActiveGoalScope(userID).
		Join("LEFT JOIN actions ON actions.goal_id = goal.id AND actions.completed_at IS NULL AND actions.dropped_at IS NULL").
		Where("actions.id IS NULL")
}

func (r *repo) userWaitingForScope(userID int) *orm.Query {
	// TODO
	// either goal is not due today or tomorrow, and not reviewed today
	// OR goal is due today or tomorrow, and not reviewed in the last hour

	return r.tx.Model(&WaitingFor{}).
		Where(`waiting_for.user_id = ?
			AND waiting_for.completed_at IS NULL
			AND waiting_for.dropped_at IS NULL
			AND waiting_for.reviewed_at < current_timestamp - interval '1 day'`, userID)
}

func (r *repo) userContextScope(userID int) *orm.Query {
	return r.tx.Model(&Context{}).Where("user_id = ?", userID)
}

func (r *repo) count(query *orm.Query) int {
	count, err := query.Count()
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
	err := r.userActionToDoScope(userID).
		Limit(1).
		Column("action.*", "Goal", "Context").
		// FIXME removed these categories of goals:
		// WHEN goal.due_at IS NULL AND goal.created_at < current_timestamp - interval '14 days' THEN 2
		// WHEN goal.due_at IS NOT NULL THEN 3
		OrderExpr(`CASE
				WHEN goal.due_at < current_timestamp + interval '2 days' THEN 1
				ELSE 4
			END ASC, action.reviewed_at ASC`).
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

func (r *repo) goalToReview(userID int) *Goal {
	var goalsToReview []Goal
	err := r.userGoalToReviewScope(userID).
		Order("reviewed_at ASC").
		Limit(1).
		Select(&goalsToReview)
	if err != nil {
		panic(err)
	}
	if len(goalsToReview) > 0 {
		goalToReview := goalsToReview[0]
		return &goalToReview
	}
	return nil
}

func (r *repo) goalWithNoAction(userID int) *Goal {
	var goalsWithoutActions []Goal
	err := r.userGoalWithNoActionScope(userID).
		Limit(1).
		Select(&goalsWithoutActions)
	if err != nil {
		panic(err)
	}
	if len(goalsWithoutActions) > 0 {
		goalWithNoAction := goalsWithoutActions[0]
		return &goalWithNoAction
	}
	return nil
}

func (r *repo) contexts(userID int) []*Context {
	var contexts []*Context
	err := r.userContextScope(userID).Select(&contexts)
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

func (r *repo) dirtyStateScope() *orm.Query {
	return r.tx.Model(&User{}).Where("active_question != ?", questionCollectingInbox)
}

func (r *repo) usersInDirtyState() []*User {
	var users []*User
	err := r.dirtyStateScope().Select(&users)
	if err != nil {
		panic(err)
	}
	return users
}

func (r *repo) usersForDailyUpdate() []*User {
	var users []*User
	err := r.tx.Model(&users).Select()
	if err != nil {
		panic(err)
	}
	return users
}

func (r *repo) earliestDirtyActivityTime() time.Time {
	var result time.Time
	r.dirtyStateScope().
		Order("last_message_at ASC").
		Limit(1).
		Column("last_message_at").
		Select(&result)
	return result
}

func (r *repo) usersDirtySince(time time.Time) []*User {
	var users []*User
	err := r.dirtyStateScope().Where("last_message_at <= ?", time).Select(&users)
	if err != nil {
		panic(err)
	}
	return users
}

func (r *repo) findContextByText(userID int, text string) *Context {
	var contexts []*Context
	err := r.userContextScope(userID).Where("text = ?", text).Select(&contexts)
	if err != nil {
		panic(err)
	}
	if len(contexts) > 0 {
		return contexts[0]
	} else {
		return nil
	}
}

func (r *repo) makeAllContextsActive(userID int) {
	_, err := r.tx.Model(&Context{Active: true}).
		Where("user_id = ?", userID).
		Column("active").
		Update()
	if err != nil {
		panic(err)
	}
}
