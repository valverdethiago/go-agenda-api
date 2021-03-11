package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	contactsBasePath = "/contacts"
	contactsPathByID = "/contacts/:id"
)

// Controller type for handling contact API requests
type Controller struct {
	store Store
}

// NewController builds a new instance of contact controller
func NewController(store Store, router *gin.Engine) *Controller {
	controller := &Controller{
		store: store,
	}
	controller.setupRoutes(router)
	return controller
}

func (controller *Controller) setupRoutes(router *gin.Engine) {
	router.GET(contactsBasePath, controller.listContacts)
	router.GET(contactsPathByID, controller.getContactByID)
}

func (controller *Controller) listContacts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "funciona"})
	return
}

func (controller *Controller) getContactByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "funciona"})
	return
}
