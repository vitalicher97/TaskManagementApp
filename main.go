package main

import (
	"./controller"
	"./handleDb"
	"./service"
	"net/http"
)

/*type Projects struct {
	ProjectId          int    `json:"project_id"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
}*/

/*type ProjectColumns struct {
	columnId		int
	columnName		string
	columnStatus	string
	columnListPos	int
	projectId		int
}

type Tasks struct {
	taskId			int
	taskName		string
	taskDescription	string
	taskPriority	int
	columnId		int
}

type Comments struct {
	commentId		int
	commentText		string
	commentCreation	string
	taskId			int
}*/

//var db *sql.DB

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	/*project_id, project_name, project_description := handleDb.TestQuery(db)
	w.Header().Set("Content-Type", "application/json")
	res := Projects {
		Project_id:          project_id,
		Project_name:        project_name,
		Project_description: project_description,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatal(err)
	}*/

}

func handler(w http.ResponseWriter, r *http.Request){
	/*//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	w.Header().Set("Content-Type", "application/json")
	projectsSlice := handleDb.GetProjects( "projects", "")

	var projects model.Projects
	for key := range projectsSlice {
		projects = model.Projects{
			ProjectId:          projectsSlice[key].ProjectId,
			ProjectName:        projectsSlice[key].ProjectName,
			ProjectDescription: projectsSlice[key].ProjectDescription,
		}
		json.NewEncoder(w).Encode(projects)
	}*/

	service.CreateProject("Impressive project 2", "To be continued")

}

func insert(w http.ResponseWriter, r *http.Request){
	var columnValue map[string]string

	columnValue = make(map[string]string)

	columnValue["project_name"] = "Fourth project"
	columnValue["project_description"] = "Check it!"

	//columnTable := "projects"

	//handleDb.CreateData(db, columnTable, columnValue)

}

func main() {
	handleDb.CreateDbObject()

	controller.MainController()
	//http.HandleFunc("/project/create", handler)
	//http.HandleFunc("/insert", insert)
	//http.ListenAndServe(":8080", nil)


	//db := handleDb.CreateDbObject()
	//handleDb.TestQuery(db)




}