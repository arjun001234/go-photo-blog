package repository

import (
	"clean-architecture/entity"
	storage "clean-architecture/storage/mysql"
	"fmt"
	"testing"
)

type userRepoTest struct {
	ur entity.UserRepository
}

func NewUserTestRepo(ur entity.UserRepository) *userRepoTest {
	return &userRepoTest{ur}
}

func TestUserRepository(t *testing.T) {
	u := &entity.User{
		Id:       2,
		Name:     "Taylor",
		Email:    "taylor32@example.com",
		Password: "taylor1234",
	}
	db := storage.NewMySQLTestDatabase()
	ur := NewUserRepository(db)
	tr := NewUserTestRepo(ur)
	tr.TestSave(t, u)
	tr.TestGetOne(t, u)
	tr.TestGetByEmail(t, u)
	tr.TestGetAll(t, u)
	uu := *u
	tr.TestUpdate(t, &uu, u)
	tr.TestDelete(t, u)
}

func (tr *userRepoTest) TestSave(t *testing.T, u *entity.User) {
	if err := tr.ur.Save(u); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("New User Saved")
}

func (tr *userRepoTest) TestGetOne(t *testing.T, u *entity.User) {
	su, err := tr.ur.GetOne(u.Id)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if su.Name != u.Name || su.Email != u.Email || su.Password != u.Password || su.Id != u.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (tr *userRepoTest) TestGetAll(t *testing.T, u *entity.User) {
	mu, err := tr.ur.GetAll()
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if len(*mu) == 0 {
		t.Errorf("Error: %v", "There must be at least one user")
	}
	isThere := false
	for _, v := range *mu {
		if v.Name == u.Name || v.Email == u.Email || v.Password == u.Password || v.Id == u.Id {
			isThere = true
		}
	}
	if isThere == false {
		t.Errorf("Error: User with id %v not found", u.Id)
	}
}

func (tr *userRepoTest) TestGetByEmail(t *testing.T, u *entity.User) {
	cu, err := tr.ur.GetByEmail(u.Email)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if cu.Name != u.Name || cu.Email != u.Email || cu.Password != u.Password || cu.Id != u.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (tr *userRepoTest) TestUpdate(t *testing.T, uu *entity.User, u *entity.User) {
	uu.Name = "Taylor Swift"
	err := tr.ur.Update(uu)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	ou, err := tr.ur.GetOne(uu.Id)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if ou.Name != uu.Name || ou.Id != uu.Id || ou.Email != uu.Email || ou.Password != uu.Password {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (tr *userRepoTest) TestDelete(t *testing.T, u *entity.User) {
	err := tr.ur.Delete(u.Id)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("User Deleted")
}
