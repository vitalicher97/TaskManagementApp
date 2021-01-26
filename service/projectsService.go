package service

import (
	"../handleDb"
	"../model"
	"strconv"
)

const tableNameProject = "projects"

func CreateProject(projectName string, projectDescription string) {

	columnValue := map[string]string{
		"project_name":        projectName,
		"project_description": projectDescription,
	}
	handleDb.CreateData(tableNameProject, columnValue)
	condition := "project_name = " + "'" + projectName + "'"
	projectsSlice := handleDb.GetProjects(condition)
	projectId := projectsSlice[0].ProjectId
	projectIdString := strconv.Itoa(projectId)
	CreateColumn("default"+projectIdString, "default"+projectIdString, "1", projectIdString)
}

func GetAllProjects() []model.Projects {
	projectsSlice := handleDb.GetProjects("")
	return projectsSlice
}

func GetProjectById(projectId string) []model.Projects {
	projectsSlice := handleDb.GetProjects("project_id = " + projectId)
	return projectsSlice
}

func UpdateProject(projectId string, projectName string, projectDescription string) {
	columnValue := map[string]string{
		"project_name":        projectName,
		"project_description": projectDescription,
	}
	condition := "project_id = " + projectId
	handleDb.UpdateData(tableNameProject, columnValue, condition)
}

func DeleteAllProjects() {
	handleDb.DeleteData(tableNameProject, "")
	handleDb.DeleteData(tableNameProjectColumns, "")
	handleDb.DeleteData(tableNameTasks, "")
	handleDb.DeleteData(tableNameComments, "")
}

func DeleteProjectById(projectId string) {
	condition := "project_id = " + projectId
	handleDb.DeleteData(tableNameProject, condition)
	handleDb.DeleteData(tableNameProjectColumns, condition)
	handleDb.DeleteData(tableNameTasks, condition)
	handleDb.DeleteData(tableNameComments, condition)
}
