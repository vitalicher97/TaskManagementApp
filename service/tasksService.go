package service

import (
	"../handleDb"
	"../model"
)

const tableNameTasks = "tasks"

func CreateTask(taskName string, taskDescription string, taskPriority string, columnId string){
	columnValue := map[string]string {
		"task_name": taskName,
		"task_description": taskDescription,
		"task_priority": taskPriority,
		"column_id": columnId,
	}
	handleDb.CreateData(tableNameTasks, columnValue)
}

func GetTasks(columnId string) []model.Tasks{
	tasksSlice := handleDb.GetTasks("column_id = " + columnId)

	return tasksSlice
}

func UpdateTask(taskId string, taskName string, taskDescription string, taskPriority string, columnId string) {
	columnValue := map[string]string {
		"task_name": taskName,
		"task_description": taskDescription,
		"task_priority": taskPriority,
		"column_id": columnId,
	}
	condition := "task_id = " + taskId
	handleDb.UpdateData(tableNameTasks, columnValue, condition)
}

func DeleteTaskById(taskId string){
	handleDb.DeleteData(tableNameComments, "task_id = " + taskId)
	handleDb.DeleteData(tableNameTasks, "task_id = " + taskId)
}


