package service

import (
	"../handleDb"
	"../model"
	"time"
)

const tableNameComments = "comments"

func CreateComment(commentText string, taskId string){
	now := time.Now().Format("02-Jan-2006 03:04:05.000")
	columnValue := map[string]string {
		"comment_text": commentText,
		"comment_creation": now,
		"task_id": taskId,
	}
	handleDb.CreateData(tableNameComments, columnValue)
}

func GetComments(taskId string) []model.Comments{
	commentsSlice := handleDb.GetComments("task_id = " + taskId)

	return commentsSlice
}

func UpdateComment(commentId string, commentText string, taskId string) {
	now := time.Now().Format("02-Jan-2006 03:04:05.000")
	columnValue := map[string]string {
		"comment_text": commentText,
		"comment_creation": now,
		"task_id": taskId,
	}
	condition := "comment_id = " + commentId
	handleDb.UpdateData(tableNameComments, columnValue, condition)
}

func DeleteCommentById(commentId string){
	handleDb.DeleteData(tableNameComments, "comment_id = " + commentId)
}