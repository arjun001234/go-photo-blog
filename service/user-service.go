package service

import (
	"clean-architecture/entity"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type uservice struct {
	userRepo       entity.UserRepository
	sessionService entity.SessionService
}

func NewUserService(r entity.UserRepository, s entity.SessionService) entity.UserService {
	return &uservice{r, s}
}

func (uservice) ValidateUser(u *entity.User) error {
	validate := validator.New()
	err := validate.Struct(u)
	return err
}

func (us uservice) IsUserLoggedIn(s *entity.Session) error {
	err := us.sessionService.GetUser(s)
	return err
}

func (us uservice) ValidateCredential(u *entity.User) (*entity.Session, error) {
	var err error
	s := entity.Session{}
	enteredPassword := u.Password
	err = us.userRepo.GetByEmail(u)
	if err != nil {
		return nil, err
	}
	err = us.ComparePassword(enteredPassword, u.Password)
	s.User = u
	us.sessionService.SetSession(&s)
	return &s, err
}

func (uservice) HashPassword(s string) (string, error) {
	var hashedPassword []byte
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (uservice) ComparePassword(p string, hp string) error {
	return bcrypt.CompareHashAndPassword([]byte(hp), []byte(p))
}

func (us uservice) CreateUser(u *entity.User) (*entity.Session, error) {
	s := entity.Session{}
	var err error
	err = us.ValidateUser(u)
	if err == nil {
		u.Password, err = us.HashPassword(u.Password)
		if err == nil {
			err = us.userRepo.Save(u)
			s.User = u
			us.sessionService.SetSession(&s)
		}
	}
	return &s, err
}

func (us uservice) Logout(s *entity.Session) error {
	return us.sessionService.RemoveSession(s)
}

func (us uservice) UpdateUser(u *entity.User) error {
	err := us.ValidateUser(u)
	if err != nil {
		return err
	}
	return us.userRepo.Update(u)
}

func (us uservice) DeleteUser(id int64) error {
	return us.userRepo.Delete(id)
}

func (us uservice) FindUser(id int64) (*entity.User, error) {
	return us.userRepo.GetOne(id)
}

func (us uservice) FindUsers() (*[]entity.User, error) {
	return us.userRepo.GetAll()
}
