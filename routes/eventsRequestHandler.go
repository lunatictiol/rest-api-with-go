package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"lunatictiol.com/resApi/model"
)

func getHello(contex *gin.Context) {
	contex.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func getEvents(contex *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	contex.JSON(http.StatusOK, events)
}
func addEvent(contex *gin.Context) {
	var event model.Event
	err := contex.ShouldBindJSON(&event)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event.UserID = contex.GetInt64("userId")
	err = event.Save()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot upload data"})
		return
	}
	contex.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})

}

func getEventById(contex *gin.Context) {

	eventId, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	event, err := model.GetEventByID(eventId)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve data with id"})
		return
	}
	contex.JSON(http.StatusOK, event)
}

func updateEvent(contex *gin.Context) {

	eventId, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	event, err := model.GetEventByID(eventId)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve data with id"})
		return
	}
	if event.ID != contex.GetInt64("userId") {
		contex.JSON(http.StatusUnauthorized, gin.H{"message": "cannot change events you didnt create"})
		return
	}

	var updateEvent model.Event
	err = contex.ShouldBindJSON(&updateEvent)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot update event with id", "error": err})
		return
	}

	contex.JSON(http.StatusCreated, gin.H{"message": "event updated", "event": updateEvent})
}

func deleteEvent(contex *gin.Context) {

	eventId, err := strconv.ParseInt(contex.Param("id"), 10, 64)
	if err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}
	event, err := model.GetEventByID(eventId)
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve data with id"})
		return
	}
	if event.ID != contex.GetInt64("userId") {
		contex.JSON(http.StatusUnauthorized, gin.H{"message": "cannot delete events you didnt create"})
		return
	}
	err = event.Delete()
	if err != nil {
		contex.JSON(http.StatusInternalServerError, gin.H{"message": "cannot delete data with id"})
		return
	}

	contex.JSON(http.StatusOK, gin.H{"message": "event deleted"})

}
