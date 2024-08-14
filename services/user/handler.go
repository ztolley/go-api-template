package user

import (
	"net/http"
	"strconv"

	"github.com/ztolley/goapi/utils"
)

/**
*
* User Request Handler
*
* This is the main route handler for requests to /users
* It has a reference to the user store, so it can make database calls
* and sets the routes up to listen for requests. All the request handlers are defined
**/

type Handler struct {
	store UserStore
}

func NewHandler(store UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /users", h.GetUsers)
	router.HandleFunc("GET /users/{userID}", h.GetUserByID)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetUsers()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")

	// convert the userID to an int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user, err := h.store.GetUserByID(userIDInt)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
