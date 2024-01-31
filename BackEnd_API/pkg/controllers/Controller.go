package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"test/pkg/db/models"
	"test/pkg/db/repository/organizations"
	"test/pkg/db/repository/users"
	"test/pkg/utils"
)

// Users Controllers
func GetUsersController() ([]models.User, error) {
	return users.GetAllUsers()
}
func GetUserController(UserID string) (*models.User, error) {
	return users.GetUserById(UserID)
}
func DeleteUserController(userID string) error {
	return users.DeleteUserById(userID)
}
func AddUserController(newuser models.UserBind) error {
	return users.InsertUser(newuser)
}
func LoginUserController(email string, password string) (map[string]string, error) {
	user, err := users.LoginUser(email, password)
	if err != nil {
		return nil, err
	}
	claims := jwt.MapClaims{
		"userid": user.UserID,
		"name":   user.Name,
		"email":  user.Email,
	}
	AccessToken, err := utils.JWTGenerateToken(claims)
	if err != nil {
		return nil, err
	}
	RefreshToken, err := utils.GenerateRefreshToken(user.UserID)
	myMap := make(map[string]string)
	myMap["access_token"] = AccessToken
	myMap["refresh_token"] = RefreshToken
	return myMap, err
}
func UpdateUserController(UserID string, newuser models.User) error {
	return users.UpdateUser(UserID, newuser)
}

// Organizations Controllers
func GetOrganizationsController() ([]models.Organization, error) {
	return organizations.GetAllOrganizations()
}
func GetOrganizationController(OrgID string) (*models.Organization, error) {
	return organizations.GetOrganizationById(OrgID)
}
func AddOrganizationController(NewOrg models.OrganizationBind) (string, error) {
	AccessLevel := "fullaccess"
	UserID := "65b8a2e738f20fe0f0476148"
	return organizations.InsertOrganization(NewOrg, UserID, AccessLevel)

}
func InviteOrganizationController(OrgID string, email string) error {
	return organizations.InviteMemberToOrganization(OrgID, email)
}
func UpdateOrganizationController(OrgID string, updatedOrg models.OrganizationOnly) error {
	return organizations.UpdateOrganization(OrgID, updatedOrg)
}
func DeleteOrganizationController(OrgID string) error {
	return organizations.DeleteOrganizationById(OrgID)
}
func TestController() error {
	AccessLevel := "fullaccess"
	UserID := "65b8a2e738f20fe0f0476148"
	return organizations.InsertMemberIntoOrganization("65b9029a4c5f266f78039a9f", UserID, AccessLevel)
}
