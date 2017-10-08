package gtd_bot

import "time"

type state struct {
	user *User
	repo *repo
}

func newState(repo *repo, userID int) *state {
	user := repo.findUser(userID)
	return &state{user: user, repo: repo}
}

func (s *state) activeQuestion() string {
	return s.user.ActiveQuestion
}

func (s *state) setActiveQuestion(activeQuestion string) {
	s.user.ActiveQuestion = activeQuestion
	s.repo.update(s.user)
}

func (s *state) userID() int {
	return s.user.ID
}

func (s *state) goalCount() int {
	return s.repo.count(s.repo.userActiveGoalScope(s.user.ID))
}

func (s *state) inboxCount() int {
	return s.repo.count(s.repo.userInboxItemScope(s.user.ID))
}

func (s *state) actionCount() int {
	return s.repo.count(s.repo.userActionScope(s.user.ID))
}

func (s *state) actionToDoCount() int {
	return s.repo.count(s.repo.userActionToDoScope(s.user.ID))
}

func (s *state) actionInContextCount(context *Context) int {
	return s.repo.count(s.repo.userActionInContextScope(context.UserID, context.ID))
}

func (s *state) waitingForCount() int {
	return s.repo.count(s.repo.userWaitingForScope(s.user.ID))
}

func (s *state) goalToReviewCount() int {
	return s.repo.count(s.repo.userGoalToReviewScope(s.user.ID))
}

func (s *state) goalWithNoActionCount() int {
	return s.repo.count(s.repo.userGoalWithNoActionScope(s.user.ID))
}

func (s *state) inboxItemToProcess() *InboxItem {
	return s.repo.inboxItemToProcess(s.user.ID)
}

func (s *state) startProcessing(inboxItem *InboxItem) {
	s.user.CurrentInboxItem = inboxItem
	s.user.CurrentInboxItemID = inboxItem.ID
	s.user.CurrentGoal = nil
	s.user.CurrentGoalID = 0
	s.user.CurrentAction = nil
	s.user.CurrentActionID = 0
	s.repo.update(s.user)
}

func (s *state) setLastMessageNow() {
	s.user.LastMessageAt = time.Now()
	s.repo.update(s.user)
}

func (s *state) someWorkToBeDone() bool {
	return s.inboxCount() > 0 ||
		s.actionToDoCount() > 0 ||
		s.waitingForCount() > 0 ||
		s.goalToReviewCount() > 0 ||
		s.goalWithNoActionCount() > 0
}

func (s *state) addInboxItem(text string) {
	inboxItem := &InboxItem{
		UserID: s.user.ID,
		Text:   text,
	}
	s.repo.insert(inboxItem)
}

func (s *state) trashCurrentInboxItem() {
	s.user.CurrentInboxItem.ProcessedAt = time.Now()
	s.repo.update(s.user.CurrentInboxItem)
}

func (s *state) createGoalAndMakeCurrent(text string) {
	goal := &Goal{
		UserID: s.user.ID,
		Text:   text,
	}
	s.repo.insert(goal)
	s.user.CurrentGoal = goal
	s.user.CurrentGoalID = goal.ID
	s.repo.update(s.user)
}

func (s *state) setGoalDue(date time.Time) {
	s.user.CurrentGoal.DueAt = date
	s.repo.update(s.user.CurrentGoal)
}

func (s *state) setGoalStatement(text string) {
	s.user.CurrentGoal.Text = text
	s.repo.update(s.user.CurrentGoal)
}

func (s *state) createActionAndMakeCurrent(text string) {
	action := &Action{
		UserID: s.user.ID,
		GoalID: s.user.CurrentGoalID,
		Text:   text,
	}
	s.repo.insert(action)
	s.user.CurrentAction = action
	s.user.CurrentActionID = action.ID
	s.repo.update(s.user)
	if s.user.CurrentInboxItem != nil {
		// the inbox item is considered processed once there is an action
		s.user.CurrentInboxItem.ProcessedAt = time.Now()
		s.repo.update(s.user.CurrentInboxItem)
	}
}

func (s *state) createWaitingForAndMakeCurrent(text string) {
	waitingFor := &WaitingFor{
		UserID: s.user.ID,
		GoalID: s.user.CurrentGoalID,
		Text:   text,
	}
	s.repo.insert(waitingFor)
	s.user.CurrentWaitingFor = waitingFor
	s.user.CurrentWaitingForID = waitingFor.ID
	s.repo.update(s.user)
	if s.user.CurrentInboxItem != nil {
		// the inbox item is considered processed once there is an waitingFor
		s.user.CurrentInboxItem.ProcessedAt = time.Now()
		s.repo.update(s.user.CurrentInboxItem)
	}
}

func (s *state) markActionCompleted() {
	s.user.CurrentAction.CompletedAt = time.Now()
	s.repo.update(s.user.CurrentAction)
}

func (s *state) actionToDo() *Action {
	return s.repo.actionToDo(s.user.ID)
}

func (s *state) waitingForToCheck() *WaitingFor {
	return s.repo.waitingForToCheck(s.user.ID)
}

func (s *state) goalToReview() *Goal {
	return s.repo.goalToReview(s.user.ID)
}

func (s *state) markGoalReviewed() {
	s.user.CurrentGoal.ReviewedAt = time.Now()
	s.repo.update(s.user.CurrentGoal)
}

func (s *state) setSuggestedAction(action *Action) {
	s.user.CurrentAction = action
	s.user.CurrentActionID = action.ID
	s.user.CurrentGoal = action.Goal
	s.user.CurrentGoalID = action.GoalID
	s.repo.update(s.user)
}

func (s *state) setCurrentWaitingFor(waitingFor *WaitingFor) {
	s.user.CurrentWaitingFor = waitingFor
	s.user.CurrentWaitingForID = waitingFor.ID
	s.user.CurrentGoal = waitingFor.Goal
	s.user.CurrentGoalID = waitingFor.GoalID
	s.repo.update(s.user)
}

func (s *state) setCurrentGoal(goal *Goal) {
	s.user.CurrentGoal = goal
	s.user.CurrentGoalID = goal.ID
	s.repo.update(s.user)
}

func (s *state) skipCurrentAction() {
	s.user.CurrentAction.ReviewedAt = time.Now()
	s.repo.update(s.user.CurrentAction)
}

func (s *state) continueToWait() {
	s.user.CurrentWaitingFor.ReviewedAt = time.Now()
	s.repo.update(s.user.CurrentWaitingFor)
}

func (s *state) completeCurrentAction() {
	s.user.CurrentAction.CompletedAt = time.Now()
	s.repo.update(s.user.CurrentAction)
}

func (s *state) completeCurrentGoal() {
	s.user.CurrentGoal.CompletedAt = time.Now()
	s.repo.update(s.user.CurrentGoal)
}

func (s *state) completeCurrentWaitingFor() {
	s.user.CurrentWaitingFor.CompletedAt = time.Now()
	s.repo.update(s.user.CurrentWaitingFor)
}

func (s *state) dropCurrentGoal() {
	s.user.CurrentGoal.DroppedAt = time.Now()
	s.repo.update(s.user.CurrentGoal)
	s.repo.dropGoalActionsAndWaitingFors(s.user.CurrentGoalID)
}

func (s *state) dropCurrentAction() {
	s.user.CurrentAction.DroppedAt = time.Now()
	s.repo.update(s.user.CurrentAction)
}

func (s *state) goalWithNoAction() *Goal {
	return s.repo.goalWithNoAction(s.user.ID)
}

func (s *state) inboxItemsProcessedTodayCount() int {
	return s.repo.count(s.repo.inboxItemsProcessedTodayScope(s.user.ID))
}

func (s *state) goalsCreatedTodayCount() int {
	return s.repo.count(s.repo.goalsCreatedTodayScope(s.user.ID))
}

func (s *state) actionsCompletedTodayCount() int {
	return s.repo.count(s.repo.actionsCompletedTodayScope(s.user.ID))
}

func (s *state) markCurrentContextInactive() {
	s.user.CurrentAction.Context.Active = false
	s.repo.update(s.user.CurrentAction.Context)
}

func (s *state) allContexts() []*Context {
	return s.repo.contexts(s.user.ID)
}

func (s *state) allGoals() []*Goal {
	return s.repo.goals(s.user.ID)
}

func (s *state) allActions() []*Action {
	return s.repo.actions(s.user.ID)
}

func (s *state) allWaitingFors() []*WaitingFor {
	return s.repo.waitingFors(s.user.ID)
}

func (s *state) allInboxItems() []*InboxItem {
	return s.repo.inboxItems(s.user.ID)
}

func (s *state) setCurrentActionContext(context *Context) {
	s.user.CurrentAction.Context = context
	if context != nil {
		s.user.CurrentAction.ContextID = context.ID
	} else {
		s.user.CurrentAction.ContextID = 0
	}
	s.repo.update(s.user.CurrentAction)
}

func (s *state) findContextByText(text string) *Context {
	return s.repo.findContextByText(s.user.ID, text)
}

func (s *state) createContext(text string) (*Context, error) {
	context := Context{
		UserID: s.user.ID,
		Text:   text,
	}
	err := s.repo.insert(&context)
	if err != nil {
		return nil, err
	}
	return &context, nil
}

func (s *state) makeAllContextsActive() {
	s.repo.makeAllContextsActive(s.user.ID)
}
