package some

import (
	"fmt"
	"github.com/alekseyklimenko/go-proj-bootstrap/models"
	"github.com/alekseyklimenko/go-proj-bootstrap/models/requests"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
	m  models.Some
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
		m:  models.Some{},
	}
}

func (s *Service) CreateNew(item *models.Some, formData requests.Some) error {
	item.Name = formData.Name
	item.Url = formData.Url
	item.ClientId = 1
	result := s.db.Create(&item)
	return result.Error
}

func (s *Service) GetSomeToProcess(count uint) *[]models.Some {
	var something []models.Some
	s.db.Where("status = ?", models.SomeStatusIdle).Limit(int(count)).Find(&something)
	return &something
}

func (s *Service) LockItem(id uint) error {
	someModel := models.NewSome()
	table := someModel.TableName()
	err := s.db.Transaction(func(tx *gorm.DB) error {
		tx.Exec(fmt.Sprintf("LOCK TABLE %s IN SHARE ROW EXCLUSIVE MODE", table))
		result := tx.First(&someModel, "id = ?", id)
		if result.RowsAffected == 0 {
			return fmt.Errorf("can't set some id %d to processing: something not found", someModel.ID)
		}
		if someModel.Status != models.SomeStatusIdle {
			return fmt.Errorf("can't set some id %d to processing: something status is not idle", someModel.ID)
		}
		tx.Exec(fmt.Sprintf("UPDATE %s SET status = ? WHERE id = ?", table), models.SomeStatusProcessing, someModel.ID)
		return nil
	})
	return err
}

func (s *Service) FreeItem(id uint) {
	s.db.Exec(fmt.Sprintf("UPDATE %s SET status = ? WHERE id = ?", s.m.TableName()), models.SomeStatusIdle, id)
}
