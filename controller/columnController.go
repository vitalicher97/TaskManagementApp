package controller

import (
	"../service"
	"log"
	"net/http"
	"strings"
)

func createColumn(w http.ResponseWriter, r *http.Request){
	columnData := strings.Split(strings.TrimPrefix(r.URL.Path, "/create/column/"), "/")

	columnDataLen := len(columnData)
	if columnDataLen != 4 {
		log.Println(columnData, ": inappropriate data for createProject!")
		return
	}

	columnNameLen := len(columnData[1])
	columnStatusLen := len(columnData[2])

	if columnNameLen < 1 || columnNameLen > 255 || columnStatusLen < 1 || columnStatusLen > 255 {
		log.Println(columnData[1], ", ", columnData[2], ": field restriction")
		return
	}

	service.CreateColumn(columnData[0], columnData[1], columnData[2], columnData[3])
}

func updateColumn(w http.ResponseWriter, r *http.Request){
	columnData := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/column/"), "/")
	service.UpdateProjectColumn(columnData[0], columnData[1], columnData[2], columnData[3], columnData[4])
}

func deleteColumnById(w http.ResponseWriter, r *http.Request){
	columnId := strings.TrimPrefix(r.URL.Path, "/delete/column/")
	service.DeleteProjectColumnById(columnId)
}
