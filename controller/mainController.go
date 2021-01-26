package controller

import (
	"net/http"
)

func MainController() {
	http.HandleFunc("/projects", showProjects)
	http.HandleFunc("/project/", showProjectById)

	http.HandleFunc("/create/project/", createProject)
	http.HandleFunc("/update/project/", updateProject)
	http.HandleFunc("/delete/project/", deleteProjectById)
	http.HandleFunc("/delete/project/all", deleteAllProjects)

	http.HandleFunc("/create/column/", createColumn)
	http.HandleFunc("/update/column/", updateColumn)
	http.HandleFunc("/delete/column/", deleteColumnById)

	http.HandleFunc("/create/task/", createTask)
	http.HandleFunc("/update/task/", updateTask)
	http.HandleFunc("/delete/task/", deleteTaskById)

	http.HandleFunc("/create/comment/", createComment)
	http.HandleFunc("/update/comment/", updateComment)
	http.HandleFunc("/delete/comment/", deleteCommentById)

	http.ListenAndServe(":8080", nil)
}
