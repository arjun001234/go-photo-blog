package repository

import (
	"clean-architecture/entity"
	storage "clean-architecture/storage/mysql"
	"fmt"
	"testing"
)

type photoRepoTest struct {
	pr entity.PhotoRepository
}

func NewPhotoTestRepo(pr entity.PhotoRepository) *photoRepoTest {
	return &photoRepoTest{pr}
}

func TestPhotoRepository(t *testing.T) {
	p := &entity.Photo{
		Id:  1,
		Url: "http://localhost:8080/public/xyz.png",
		User: &entity.User{
			Id:       1,
			Name:     "messi",
			Email:    "messi34@example.com",
			Password: "messi1234",
		},
	}
	db := storage.NewMySQLTestDatabase()
	pr := NewPhotoRepository(db)
	tr := NewPhotoTestRepo(pr)
	tr.TestSave(t, p)
	tr.TestGetOne(t, p)
	tr.TestGetAll(t, p)
	tr.TestDelete(t, p)
}

func (tr *photoRepoTest) TestSave(t *testing.T, p *entity.Photo) {
	if err := tr.pr.Save(p); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("New Photo Saved")
}

func (tr *photoRepoTest) TestGetOne(t *testing.T, p *entity.Photo) {
	sp, err := tr.pr.GetOne(p.Id)
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if sp.Id != p.Id || sp.Url != p.Url || sp.User.Id != p.User.Id {
		t.Errorf("Error: %v", "Input Output didn't match")
	}
}

func (tr *photoRepoTest) TestGetAll(t *testing.T, p *entity.Photo) {
	mp, err := tr.pr.GetAll()
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	if len(*mp) == 0 {
		t.Errorf("Error: %v", "There must be at least one user")
	}
	isThere := false
	for _, v := range *mp {
		if v.Id == p.Id || v.Url == p.Url || v.User.Id == p.User.Id {
			isThere = true
		}
	}
	if isThere == false {
		t.Errorf("Error: User with id %v not found", p.Id)
	}
}

func (tr *photoRepoTest) TestDelete(t *testing.T, p *entity.Photo) {
	if err := tr.pr.Delete(p.Id); err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	fmt.Println("Photo Deleted")
}
