package controller

import (
	"../service"
	"log"
	"net/http"
	"strings"
)

func createComment(w http.ResponseWriter, r *http.Request){
	commentData := strings.Split(strings.TrimPrefix(r.URL.Path, "/create/comment/"), "/")

	commentDataLen := len(commentData)
	if commentDataLen != 2 {
		log.Println(commentData, ": inappropriate data!")
		return
	}

	commentTextLen := len(commentData[0])
	if commentTextLen < 1 || commentTextLen > 5000 {
		log.Println(commentData[0], ": field restriction!")
		return
	}

	service.CreateComment(commentData[0], commentData[1])
}

func updateComment(w http.ResponseWriter, r *http.Request){
	commentData := strings.Split(strings.TrimPrefix(r.URL.Path, "/update/comment/"), "/")
	service.UpdateComment(commentData[0], commentData[1], commentData[2])
}

func deleteCommentById(w http.ResponseWriter, r *http.Request){
	commentId := strings.TrimPrefix(r.URL.Path, "/delete/comment/")
	service.DeleteCommentById(commentId)
}
