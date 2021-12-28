package service

import (
	"clean-architecture/entity"
	repo "clean-architecture/repo/mysql"
	storage "clean-architecture/storage/mysql"
	"fmt"
	"testing"
)

type userTestService struct {
	us entity.UserService
}

func NewUserServiceTest(us entity.UserService) *userTestService {
	return &userTestService{us}
}

func TestUserService(t *testing.T) {
	u := &entity.User{
		Id:       4,
		Name:     "Erling",
		Email:    "erling20@example.com",
		Password: "erling1234",
	}
	db := storage.NewMySQLTestDatabase()
	ur := repo.NewUserRepository(db)
	sr := repo.NewSessionRepository(db)
	ss := NewSessionService(sr)
	us := NewUserService(ur, ss)
	ts := NewUserServiceTest(us)
	ts.TestValidateUser(t, u)
	ts.TestCreateUser(t, u)
	u.Password = "erling1234"
	ts.TestValidateCredential(t, u)
	ts.TestFindUser(t, u)
	ts.TestFindUsers(t, u)
	uu := *u
	ts.TestUpdateUser(t, &uu, u)
	ts.TestHashAndComparePassword(t, u)
	ts.TestDeleteUser(t, u)
}

func (ts *userTestService) TestCreateUser(t *testing.T, u *entity.User) {
	_, err := ts.us.CreateUser(u)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("Service: User Added")
}

func (ts *userTestService) TestFindUser(t *testing.T, u *entity.User) {
	su, err := ts.us.FindUser(u.Id)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if su.Name != u.Name || su.Email != u.Email || su.Id != u.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (ts *userTestService) TestFindUsers(t *testing.T, u *entity.User) {
	mu, err := ts.us.FindUsers()
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

func (ts *userTestService) TestUpdateUser(t *testing.T, uu *entity.User, u *entity.User) {
	uu.Name = "Erling Haaland"
	err := ts.us.UpdateUser(uu)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	ou, err := ts.us.FindUser(uu.Id)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if ou.Name != uu.Name || ou.Id != uu.Id || ou.Email != uu.Email || ou.Password != uu.Password {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (ts *userTestService) TestDeleteUser(t *testing.T, u *entity.User) {
	if err := ts.us.DeleteUser(u.Id); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("Service: User Deleted")
}

func (ts *userTestService) TestValidateUser(t *testing.T, u *entity.User) {
	if err := ts.us.ValidateUser(u); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func (ts *userTestService) TestHashAndComparePassword(t *testing.T, u *entity.User) {
	hp, err := ts.us.HashPassword(u.Password)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	err = ts.us.ComparePassword(u.Password, hp)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
}

func (ts *userTestService) TestValidateCredential(t *testing.T, u *entity.User) {
	s, err := ts.us.ValidateCredential(u)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if s.User.Name != u.Name || s.User.Email != u.Email || s.User.Id != u.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}
