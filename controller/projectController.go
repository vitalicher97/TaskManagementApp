package controller

import (
	"../model"
	"../service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func showProjects(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	projectsSlice := service.GetAllProjects()

	var project model.Projects
	for key := range projectsSlice {
		project = model.Projects{
			ProjectId:          projectsSlice[key].ProjectId,
			ProjectName:        projectsSlice[key].ProjectName,
			ProjectDescription: projectsSlice[key].ProjectDescription,
		}
		json.NewEncoder(w).Encode(project)
	}

}

func showProjectById(w http.ResponseWriter, r *http.Request) {

	projectId := strings.TrimPrefix(r.URL.Path, "/project/")
	w.Header().Set("Content-Type", "application/json")

	projectsSlice := service.GetProjectById(projectId)
	columnsSlice := service.GetProjectColumns(strconv.Itoa(projectsSlice[0].ProjectId))

	var project model.Projects
	project = model.Projects{
		ProjectId:          projectsSlice[0].ProjectId,
		ProjectName:        projectsSlice[0].ProjectName,
		ProjectDescription: projectsSlice[0].ProjectDescription,
	}
	json.NewEncoder(w).Encode(project)

	var column model.ProjectColumns
	var tasksSlice []model.Tasks
	var task model.Tasks
	var commentsSlice []model.Comments
	var comment model.Comments
	for colKey := range columnsSlice {
		column = model.ProjectColumns{
			ColumnId:      columnsSlice[colKey].ColumnId,
			ColumnName:    columnsSlice[colKey].ColumnName,
			ColumnStatus:  columnsSlice[colKey].ColumnStatus,
			ColumnListPos: columnsSlice[colKey].ColumnListPos,
			ProjectId:     columnsSlice[colKey].ProjectId,
		}
		json.NewEncoder(w).Encode(column)

		tasksSlice = service.GetTasks(strconv.Itoa(columnsSlice[colKey].ColumnId))
		for tasKey := range tasksSlice{
			task = model.Tasks{
				TaskId:          tasksSlice[tasKey].TaskId,
				TaskName:        tasksSlice[tasKey].TaskName,
				TaskDescription: tasksSlice[tasKey].TaskDescription,
				TaskPriority:    tasksSlice[tasKey].TaskPriority,
				ColumnId:        tasksSlice[tasKey].ColumnId,
			}
			json.NewEncoder(w).Encode(task)

			commentsSlice = service.GetComments(strconv.Itoa(tasksSlice[tasKey].TaskId))
			for comKey := range commentsSlice{
				comment = model.Comments{
					CommentId:       commentsSlice[comKey].CommentId,
					CommentText:     commentsSlice[comKey].CommentText,
					CommentCreation: commentsSlice[comKey].CommentCreation,
					TaskId:          commentsSlice[comKey].TaskId,
				}
				json.NewEncoder(w).Encode(comment)
			}

		}

	}

}

func createProject(w http.ResponseWriter, r *http.Request) {
	projectData := strings.Split(strings.TrimPrefix(r.URL.Path, "/create/project/"), "/")

	projectDataLen := len(projectData)
	if projectDataLen != 2 {
		log.Fatal(projectData, ": inappropriate data for createProject!")
		return
	}

	projectDataLen1 := len(projectData[0])
	projectDataLen2 := len(projectData[1])

	if projectDataLen1 < 1 || projectDataLen1 > 500 || projectDataLen2 > 1000 {
		log.Fatal(projectData[0], ", ", projectData[1], ": field restriction")
		return
	}

	service.CreateProject(projectData[0], projectData[1])
}

func updateProject(w http.ResponseWriter, r *http.Request){
	projectData := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/project/"), "/")

	projectDataLen := len(projectData)
	if projectDataLen != 3 {
		log.Println(projectData, ": inappropriate data for updateProject!")
		return
	}

	projectNameLen := len(projectData[1])
	projectDescriptinonLen := len(projectData[2])

	if projectNameLen < 1 || projectNameLen > 500 || projectDescriptinonLen > 1000 {
		log.Println(projectData[1], ", ", projectData[2], ": field restriction")
		return
	}

	service.UpdateProject(projectData[0], projectData[1], projectData[2])
}

func deleteProjectById(w http.ResponseWriter, r *http.Request){
	projectId := strings.TrimPrefix(r.URL.Path, "/delete/project/")
	service.DeleteProjectById(projectId)
}

func deleteAllProjects(w http.ResponseWriter, r *http.Request){
	service.DeleteAllProjects()
}
