package model

type Projects struct {
	ProjectId          int    `json:"project_id"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
}

type ProjectColumns struct {
	ColumnId      int    `json:"column_id"`
	ColumnName    string `json:"column_name"`
	ColumnStatus  string `json:"column_status"`
	ColumnListPos int    `json:"column_listPos"`
	ProjectId     int    `json:"project_id"`
}

type Tasks struct {
	TaskId          int    `json:"task_id"`
	TaskName        string `json:"task_name"`
	TaskDescription string `json:"task_description"`
	TaskPriority    int    `json:"task_priority"`
	ColumnId        int    `json:"column_id"`
}

type Comments struct {
	CommentId       int    `json:"comment_id"`
	CommentText     string `json:"comment_text"`
	CommentCreation string `json:"comment_creation"`
	TaskId          int    `json:"task_id"`
}
