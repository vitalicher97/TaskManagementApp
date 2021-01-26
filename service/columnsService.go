package service

import (
	"../handleDb"
	"../model"
	"strconv"
)

const tableNameProjectColumns = "project_columns"

func CreateColumn(columnName string, columnStatus string, columnListPos string, projectId string){
	columnValue := map[string]string {
		"column_name": columnName,
		"column_status": columnStatus,
		"column_listPos": columnListPos,
		"project_id": projectId,
	}
	handleDb.CreateData(tableNameProjectColumns, columnValue)
}

func GetProjectColumns(projectId string) []model.ProjectColumns{
	columnsSlice := handleDb.GetProjectColumns("project_id = " + projectId)

	return columnsSlice
}

func UpdateProjectColumn(columnId string, columnName string, columnStatus string, columnListPos string, projectId string) {
	columnValue := map[string]string {
		"column_name": columnName,
		"column_status": columnStatus,
		"column_listPos": columnListPos,
		"project_id": projectId,
	}
	condition := "column_id = " + columnId

	handleDb.UpdateData(tableNameProjectColumns, columnValue, condition)
}

func DeleteProjectColumnById(columnId string){
	columnsSlice := handleDb.GetProjectColumns("column_id = " + columnId)
	projectId := strconv.Itoa(columnsSlice[0].ProjectId)
	columnListPos := columnsSlice[0].ColumnListPos
	columnsQuantity := handleDb.GetCountAll(tableNameProjectColumns, "project_id = " + projectId)
	if columnsQuantity > 1 {
		condition := "column_id = " + columnId
		tasksSlice := GetTasks(columnId)
		if len(tasksSlice) != 0 {
			columnsSlice = GetProjectColumns(projectId)
			var leftColId int
			for colKey := range columnsSlice {
				if columnListPos == columnsSlice[colKey].ColumnListPos {
					if colKey != 0 {
						leftColId = columnsSlice[colKey-1].ColumnId
					} else {
						leftColId = columnsSlice[colKey+1].ColumnId
					}
					break
				}
			}
			for tasKey := range tasksSlice {
				UpdateTask(strconv.Itoa(tasksSlice[tasKey].TaskId),
					tasksSlice[tasKey].TaskName, tasksSlice[tasKey].TaskDescription,
					strconv.Itoa(tasksSlice[tasKey].TaskPriority), strconv.Itoa(leftColId))
			}
		}
		handleDb.DeleteData(tableNameProjectColumns, condition)
	}


}


