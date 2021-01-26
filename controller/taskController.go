package controller

import (
	"../service"
	"net/http"
	"strings"
)

func createTask(w http.ResponseWriter, r *http.Request){
	taskData := strings.Split(strings.TrimPrefix(r.URL.Path, "/create/task/"), "/")
	service.CreateTask(taskData[0], taskData[1], taskData[2], taskData[3])
}

func updateTask(w http.ResponseWriter, r *http.Request){
	taskData := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/task/"), "/")
	service.UpdateTask(taskData[0], taskData[1], taskData[2], taskData[3], taskData[4])
}

func deleteTaskById(w http.ResponseWriter, r *http.Request){
	taskId := strings.TrimPrefix(r.URL.Path, "/delete/task/")
	service.DeleteTaskById(taskId)
}
