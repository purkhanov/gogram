package types

// Describes a task in a checklist.
type ChecklistTask struct {
	// Unique identifier of the task
	ID int `json:"id"`

	// Text of the task
	Text string `json:"text"`

	// Optional. Special entities that appear in the task text
	TextEntities []MessageEntity `json:"text_entities,omitempty"`

	// Optional. User that completed the task; omitted if the task wasn't completed
	CompletedByUser *User `json:"completed_by_user,omitempty"`

	// Optional. Point in time (Unix timestamp) when the task
	// was completed; 0 if the task wasn't completed
	CompletionDate int `json:"completion_date,omitempty"`
}

// Describes a checklist.
type Checklist struct {
	// Title of the checklist
	Title string `json:"title"`

	// Optional. Special entities that appear in the checklist title
	TitleEntities []MessageEntity `json:"title_entities,omitempty"`

	// List of tasks in the checklist
	Tasks []Checklist `json:"tasks"`

	// Optional. True, if users other than the
	// creator of the list can add tasks to the list
	OthersCanAddTasks bool `json:"others_can_add_tasks,omitempty"`

	// Optional. True, if users other than the creator
	// of the list can mark tasks as done or not done
	OthersCanMarkTasksAsDone bool `json:"others_can_mark_tasks_as_done,omitempty"`
}
