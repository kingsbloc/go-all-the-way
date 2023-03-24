package repo

import (
	"time"

	"github.com/altschool/go-app/internal/models"
	"github.com/rs/xid"
)

var chefs []models.Chef

func NewChefRepo() {
	chefs = make([]models.Chef, 0)
}

func NewChef(chef models.Chef) {

	chef.Id = xid.New().String()
	chef.CreatedAt = time.Now()
	chefs = append(chefs, chef)

}

func ListChefs() []models.Chef {
	return chefs
}

func DeleteChef(id string) bool {
	index := -1

	for i := 0; i < len(chefs); i++ {
		if chefs[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		return false
	}

	chefs = append(chefs[:index], chefs[index+1:]...)

	return true
}

func UpdateChef(body models.Chef, id string) (bool, *models.Chef) {

	index := -1

	for i := 0; i < len(chefs); i++ {
		if chefs[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		return false, nil
	}

	body.Id = id
	chefs[index] = body

	return true, &body
}

func ChefExists(id string) bool {
	for _, chef := range chefs {
		if chef.Id == id {
			return true
		}
	}
	return false
}
