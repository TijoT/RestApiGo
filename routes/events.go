package routes

import (
	models "RestApi/events"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch events"})
	}

	context.JSON(http.StatusOK, events)
}

func getEventByID(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	// event.ID = 1
	event.UserID = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not create event. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"Message": "Event created!!", "event": event})

}

func updateEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event ID"})
		return
	}

	// Check event exists with the ID
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	// add event ID else the update fails
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse event ID"})
		return
	}

	// Check event exists with the ID
	existingEvent, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch event"})
		return
	}

	fmt.Println(existingEvent)

	err = existingEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event deleted successfully"})

}
