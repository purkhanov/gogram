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

// Describes a task to add to a checklist.
type InputChecklistTask struct {
	// Unique identifier of the task; must be positive and unique
	// among all task identifiers currently present in the checklist
	ID int `json:"id"`

	// Text of the task; 1-100 characters after entities parsing
	Text string `json:"text"`

	// Optional. Mode for parsing entities in the text.
	// See formatting options for more details.
	// (https://core.telegram.org/bots/api#formatting-options)
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the text,
	// which can be specified instead of parse_mode. Currently,
	// only bold, italic, underline, strikethrough, spoiler, and
	// custom_emoji entities are allowed.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

// Describes a checklist to create.
type InputChecklist struct {
	// Title of the checklist; 1-255
	// characters after entities parsing
	Title string `json:"title"`

	// Optional. Mode for parsing entities in the title.
	// See formatting options for more details.
	// (https://core.telegram.org/bots/api#formatting-options)
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the
	// title, which can be specified instead of parse_mode.
	// Currently, only bold, italic, underline, strikethrough,
	// spoiler, and custom_emoji entities are allowed.
	TitleEntities []MessageEntity `json:"title_entities,omitempty"`

	// List of 1-30 tasks in the checklist
	Tasks []InputChecklistTask `json:"tasks"`

	// Optional. Pass True if other users can add tasks to the checklist
	OthersCanAddTasks bool `json:"others_can_add_tasks,omitempty"`

	// Optional. Pass True if other users can mark
	// tasks as done or not done in the checklist
	OthersCanMarkTasksAsDone bool `json:"others_can_mark_tasks_as_done,omitempty"`
}

// Describes a service message about checklist
// tasks marked as done or not done.
type ChecklistTasksDone struct {
	// Optional. Message containing the checklist whose
	// tasks were marked as done or not done. Note that
	// the Message object in this field will not contain
	// the reply_to_message field even if it itself is a reply.
	ChecklistMessage *Message `json:"checklist_message,omitempty"`

	// Optional. Identifiers of the tasks that were marked as done
	MarkedAsDonTaskIDs []int `json:"marked_as_done_task_ids,omitempty"`

	// Optional. Identifiers of the tasks that were marked as not done
	MarkedAsNotDoneTaskIDs []int `json:"marked_as_not_done_task_ids,omitempty"`
}

// Describes a service message about tasks added to a checklist.
type ChecklistTasksAdded struct {
	// Optional. Message containing the checklist to which
	// the tasks were added. Note that the Message object
	// in this field will not contain the reply_to_message
	// field even if it itself is a reply.
	ChecklistMessage *Message `json:"checklist_message,omitempty"`

	// List of tasks added to the checklist
	Tasks []Checklist `json:"tasks,omitempty"`
}
