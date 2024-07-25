package routes

import (
	"restapi/models"

	"strconv"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"net/http"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context, database *sql.DB) {
	events, err := models.GetallEvents(database)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventbyID(context *gin.Context, database *sql.DB) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	event, err := models.GetEventbyID(database, eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context, database *sql.DB) {

	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	userid := context.GetInt64("userid")
	event.UserID = userid
	err = event.Save(database)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	context.JSON(http.StatusCreated, gin.H{"message": "event created "})

}

func UpdateEvent(context *gin.Context, database *sql.DB) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	_, err = models.GetEventbyID(database, eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var updatedevent models.Event
	updatedevent.ID = eventID
	err = context.ShouldBindBodyWithJSON(&updatedevent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = updatedevent.Update(database)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": " Event updated successfully"})

}

func DeleteEvent(context *gin.Context, database *sql.DB) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	event, err := models.GetEventbyID(database, eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = event.Delete(database)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": " Event deleted successfully"})

}
