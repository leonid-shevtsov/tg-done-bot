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

func (s *state) inboxCount() int {
	return s.repo.inboxCount(s.user.ID)
}

func (s *state) actionCount() int {
	return s.repo.actionCount(s.user.ID)
}

func (s *state) inboxItemToProcess() *InboxItem {
	return s.repo.inboxItemToProcess(s.user.ID)
}

func (s *state) startProcessing(inboxItem *InboxItem) {
	s.user.CurrentInboxItemID = inboxItem.ID
	s.user.CurrentGoalID = 0
	s.user.CurrentActionID = 0
	s.repo.update(s.user)
}

func (s *state) setLastMessageNow() {
	s.user.LastMessageAt = time.Now()
	s.repo.update(s.user)
}

func (s *state) someWorkToBeDone() bool {
	return s.inboxCount() > 0 || s.actionCount() > 0
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
	s.user.CurrentGoalID = goal.ID
	s.repo.update(s.user)
}

func (s *state) createActionAndMakeCurrent(text string) {
	action := &Action{
		UserID: s.user.ID,
		GoalID: s.user.CurrentGoalID,
		Text:   text,
	}
	s.repo.insert(action)
	s.user.CurrentActionID = action.ID
	s.repo.update(s.user)
	if s.user.CurrentInboxItem != nil {
		// the inbox item is considered processed once there is an action
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

func (s *state) setSuggestedAction(action *Action) {
	s.user.CurrentActionID = action.ID
	s.user.CurrentGoalID = action.GoalID
	s.repo.update(s.user)
}

func (s *state) skipCurrentAction() {
	s.user.CurrentAction.ReviewedAt = time.Now()
	s.repo.update(s.user.CurrentAction)
}

func (s *state) completeCurrentAction() {
	s.user.CurrentAction.CompletedAt = time.Now()
	s.repo.update(s.user.CurrentAction)
}

func (s *state) completeCurrentGoal() {
	s.user.CurrentGoal.CompletedAt = time.Now()
	s.repo.update(s.user.CurrentGoal)
}
