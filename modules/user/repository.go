package user

import (
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type userRepository struct {
	schema.User
}

func createNewUser(firstName, lastName, email string) (*schema.User,[]error) {
	user := schema.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	errs := database.DB().Create(&user).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &user, nil
}


func (userRepo *userRepository) findUserById(id uint) ( *schema.User, []error ) {
	var user schema.User

	errs := database.DB().Where("id = ?",id).First(&user).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return &user, nil
}