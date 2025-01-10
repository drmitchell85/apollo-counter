package handlers

import (
	"apollo-counter/internal/controllers"
	"apollo-counter/internal/models"
	"apollo-counter/internal/utils"
	"encoding/json"
	"net/http"
)

type EventHandler struct {
	eventController controllers.EventController
}

func NewEventHandler(controller controllers.EventController) *EventHandler {
	return &EventHandler{
		eventController: controller,
	}
}

func (h *EventHandler) NewEvent(w http.ResponseWriter, r *http.Request) {
	req := models.NewEventRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondFailure(w, http.StatusBadRequest, utils.ErrInvalidJsonFormat)
		return
	}

	// validate if fields are empty
	if req.Title == "" || req.Description == "" || req.DateTime == "" {
		respondFailure(w, http.StatusBadRequest, utils.ErrMissingFields)
		return
	}

	err := h.eventController.CreateEvent(req)

	// TODO fix error handler
	if err != nil {
		respondFailure(w, http.StatusConflict, err)
		return
	}

	respondSuccess(w, http.StatusOK, "Event created successfully")
}
