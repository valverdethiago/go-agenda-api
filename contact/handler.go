package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

const (
	contactsBasePath = "/contacts"
	contactsPathByID = "/:id"
)

type contactIdRequest struct {
	ID string `uri:"id" binding:"required"`
}
type contactSearchRequest struct {
	searchTerm string `uri:"term"`
}
type contactRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

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
	basicRoute := router.Group(contactsBasePath)
	basicRoute.GET("", controller.search)
	basicRoute.POST("", controller.createContact)
	basicRoute.GET("/:id", controller.getContactByID)
	basicRoute.PUT("/:id", controller.updateContact)
	basicRoute.DELETE("/:id", controller.deleteContact)

}

func (controller *Controller) createContact(ctx *gin.Context) {
	req, err := parseContactRequest(ctx)
	if err != nil {
		return
	}
	arg := Contact{
		Name:   req.Name,
		Email:  req.Email,
		Active: true,
	}
	contact, err := controller.store.Create(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, contact)
	return
}

func (controller *Controller) getContactByID(ctx *gin.Context) {
	contact, err := controller.assertContactExists(ctx)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, contact)
	return
}

func (controller *Controller) updateContact(ctx *gin.Context) {
	contact, err := controller.assertContactExists(ctx)
	if err != nil {
		return
	}
	req, err := parseContactRequest(ctx)
	if err != nil {
		return
	}
	arg := Contact{
		ID:    contact.ID,
		Name:  req.Name,
		Email: req.Email,
	}
	err = controller.store.Update(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, arg)
}

func (controller *Controller) deleteContact(ctx *gin.Context) {
	contact, err := controller.assertContactExists(ctx)
	if err != nil {
		return
	}
	err = controller.store.Delete(contact.ID.Hex())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}

func (controller *Controller) search(ctx *gin.Context) {
	req := parseContactSearchRequest(ctx)
	var contacts []Contact
	var err error
	if req == nil || req.searchTerm == "" {
		contacts, err = controller.store.GetAll()
	} else {
		contacts, err = controller.store.FindByName(req.searchTerm)
	}
	if err != nil {
		if err == mgo.ErrNotFound {
			ctx.JSON(http.StatusNoContent, gin.H{"message": "No contact found for the given search term"})
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}
	if len(contacts) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "No contacts found"})
	} else {
		ctx.JSON(http.StatusOK, contacts)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func parseContactSearchRequest(ctx *gin.Context) *contactSearchRequest {
	searchTerm := ctx.Query("term")
	if searchTerm != "" {
		return &contactSearchRequest{
			searchTerm: searchTerm,
		}
	}
	return nil
}

func parseContactIdRequest(ctx *gin.Context) (contactIdRequest, error) {
	var req contactIdRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	return req, err
}

func parseContactRequest(ctx *gin.Context) (contactRequest, error) {
	var req contactRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	return req, err
}

func (controller *Controller) assertContactExists(ctx *gin.Context) (*Contact, error) {
	req, err := parseContactIdRequest(ctx)
	if err != nil {
		return nil, err
	}
	contact, err := controller.store.FindByID(req.ID)
	if err != nil {
		if err == mgo.ErrNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "No contact found for the given id"})
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
	}
	return &contact, err
}
