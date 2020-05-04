package controllers

import (
	"crawlab-lite/constants"
	"crawlab-lite/forms"
	"crawlab-lite/services"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTaskList(c *gin.Context) {
	var page forms.PageForm

	if err := c.ShouldBindQuery(&page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if total, tasks, err := services.QueryTaskPage(page); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccessList(c, total, tasks)
	}
}

func GetTask(c *gin.Context) {
	id := c.Param("id")

	if task, err := services.QueryTaskById(id); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		if task == nil {
			HandleError(http.StatusNotFound, c, errors.New("task not found"))
			return
		}
		HandleSuccess(c, task)
	}
}

func CreateTask(c *gin.Context) {
	var form forms.TaskForm

	if err := c.ShouldBindJSON(&form); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if res, err := services.AddTask(form); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccess(c, res)
	}
}

func UpdateTaskCancel(c *gin.Context) {
	id := c.Param("id")

	if res, err := services.CancelTask(id, constants.TaskStatusCancelled); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	} else {
		HandleSuccess(c, res)
	}
}
