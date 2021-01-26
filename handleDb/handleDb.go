package handleDb

import (
	"../model"
	"database/sql"
	_ "github.com/godror/godror"
	"log"
)

var Db *sql.DB

func CreateDbObject() {
	var err error
	Db, err = sql.Open("godror",
		`user="SYSTEM" password="admin" connectString="localhost:1521/taskAppDB"`)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()
}

func CreateData(tableName string, columnValue map[string]string) {

	insertSql := "INSERT INTO " + tableName + " ("
	valuesSql := "VALUES ("
	keyQuantity := len(columnValue)
	count := 0
	for key := range columnValue {
		count++
		insertSql += key
		valuesSql += "'" + columnValue[key] + "'"
		if count != keyQuantity {
			insertSql += ", "
			valuesSql += ", "
		} else {
			insertSql += ") "
			valuesSql += ")"
		}
	}

	insertSql += valuesSql

	tx, err := Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(insertSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

func GetProjects(condition string) []model.Projects {

	var projects model.Projects
	selectSql := "SELECT * FROM projects"
	if condition != "" {
		selectSql += " WHERE " + condition
	}

	selectSql += " ORDER BY project_name"

	rows, err := Db.Query(selectSql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var projectsSlice []model.Projects
	for rows.Next() {
		err := rows.Scan(&projects.ProjectId, &projects.ProjectName, &projects.ProjectDescription)
		if err != nil {
			log.Println(err)
		}
		projectsSlice = append(projectsSlice, projects)
		log.Println(projectsSlice)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return projectsSlice

}

func GetProjectColumns(condition string) []model.ProjectColumns {

	var projectColumns model.ProjectColumns
	selectSql := "SELECT * FROM project_columns"
	if condition != "" {
		selectSql += " WHERE " + condition
	}

	selectSql += " ORDER BY column_listPos"

	rows, err := Db.Query(selectSql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var projectColumnsSlice []model.ProjectColumns
	for rows.Next() {
		err := rows.Scan(&projectColumns.ColumnId,
			&projectColumns.ColumnName, &projectColumns.ColumnStatus, &projectColumns.ColumnListPos,
			&projectColumns.ProjectId)
		if err != nil {
			log.Println(err)
		}
		projectColumnsSlice = append(projectColumnsSlice, projectColumns)
		log.Println(projectColumnsSlice)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return projectColumnsSlice

}

func GetTasks(condition string) []model.Tasks {

	var tasks model.Tasks
	selectSql := "SELECT * FROM tasks"
	if condition != "" {
		selectSql += " WHERE " + condition
	}

	rows, err := Db.Query(selectSql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var tasksSlice []model.Tasks
	for rows.Next() {
		err := rows.Scan(&tasks.TaskId, &tasks.TaskName, &tasks.TaskDescription, &tasks.TaskPriority, &tasks.ColumnId)
		if err != nil {
			log.Println(err)
		}
		tasksSlice = append(tasksSlice, tasks)
		log.Println(tasksSlice)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return tasksSlice

}

func GetComments(condition string) []model.Comments {

	var comments model.Comments
	selectSql := "SELECT * FROM comments"
	if condition != "" {
		selectSql += " WHERE " + condition
	}

	selectSql += " ORDER BY comment_creation DESC"

	rows, err := Db.Query(selectSql)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var commentsSlice []model.Comments
	for rows.Next() {
		err := rows.Scan(&comments.CommentId, &comments.CommentText, &comments.CommentCreation, &comments.TaskId)
		if err != nil {
			log.Println(err)
		}
		commentsSlice = append(commentsSlice, comments)
		log.Println(commentsSlice)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return commentsSlice

}

func UpdateData(tableName string, columnValue map[string]string, condition string) {

	updateSql := "UPDATE " + tableName + " SET "
	keyQuantity := len(columnValue)
	count := 0
	for key := range columnValue {
		count++
		updateSql += key + " = " + "'" + columnValue[key] + "'"
		if count != keyQuantity {
			updateSql += ", "
		} else {
			updateSql += " "
		}
	}

	updateSql += "WHERE " + condition

	tx, err := Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(updateSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

func DeleteData(tableName string, condition string) {

	deleteSql := "DELETE FROM " + tableName
	if condition != "" {
		deleteSql += " WHERE " + condition
	}

	tx, err := Db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(deleteSql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

}

func GetCountAll(tableName string, condition string) int {

	selectSql := "SELECT COUNT(*) FROM " + tableName
	if condition != "" {
		selectSql += " WHERE " + condition
	}

	var quantity int
	err := Db.QueryRow(selectSql).Scan(&quantity)
	if err != nil {
		log.Println(err)
	}

	return quantity

}
