package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/pkg/controllers"
	"test/pkg/db/models"
)

// Users Handlers

func GetUsersHandler(c *gin.Context) {
	users, err := controllers.GetUsersController()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
	return
}
func GetUserHandler(c *gin.Context) {
	userID := c.Param("id")
	user, err := controllers.GetUserController(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
func DeleteUserHandler(c *gin.Context) {
	userID := c.Param("id")
	err := controllers.DeleteUserController(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
func AddUserHandler(c *gin.Context) {

	var newUser models.UserBind

	// Bind the JSON request body to the User struct
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controllers.AddUserController(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, newUser)
}
func LoginUserHandler(c *gin.Context) {

	// Parse JSON request
	var jsonRequest map[string]interface{}
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Extract email from the request
	email, ok := jsonRequest["email"].(string)
	password, ok := jsonRequest["password"].(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Email not provided or invalid"})
		return
	}

	myMap, err := controllers.LoginUserController(email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, myMap)
}
func UpdateUserHandler(c *gin.Context) {
	userID := c.Param("id")
	var newUser models.User
	// Bind the JSON request body to the User struct
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controllers.UpdateUserController(userID, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, newUser)
}

// Organizations Handlers
func GetOrganizationsHandler(c *gin.Context) {
	orgs, err := controllers.GetOrganizationsController()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orgs)
}
func GetOrganizationHandler(c *gin.Context) {
	OrgID := c.Param("id")
	org, err := controllers.GetOrganizationController(OrgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, org)
}
func AddOrganizationHandler(c *gin.Context) {

	var NewOrg models.OrganizationBind

	// Bind the JSON request body to the User struct
	if err := c.BindJSON(&NewOrg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org_id, err := controllers.AddOrganizationController(NewOrg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, org_id)
}
func InviteOrganizationHandler(c *gin.Context) {

	organization_id := c.Param("organization_id")

	// Parse JSON request
	var jsonRequest map[string]interface{}
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Extract email from the request
	email, ok := jsonRequest["email"].(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Email not provided or invalid"})
		return
	}
	err := controllers.InviteOrganizationController(organization_id, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, "Invitation Done Successfully")
}
func UpdateOrganizationHandler(c *gin.Context) {
	OrgID := c.Param("organization_id")
	var updatedOrg models.OrganizationOnly
	updatedOrg.OrgID = OrgID
	// Bind the JSON request body to the User struct
	if err := c.BindJSON(&updatedOrg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := controllers.UpdateOrganizationController(OrgID, updatedOrg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, updatedOrg)
}
func DeleteOrganizationHandler(c *gin.Context) {
	OrgID := c.Param("organization_id")
	err := controllers.DeleteUserController(OrgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted successfully"})
}

func TestHandler(c *gin.Context) {

	err := controllers.TestController()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the created user
	c.JSON(http.StatusCreated, "done")
}
