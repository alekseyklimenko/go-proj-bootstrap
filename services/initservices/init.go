package initservices

import (
	"github.com/alekseyklimenko/go-proj-bootstrap/config"
	"github.com/alekseyklimenko/go-proj-bootstrap/services"
	"github.com/alekseyklimenko/go-proj-bootstrap/services/processing"
	"github.com/alekseyklimenko/go-proj-bootstrap/services/some"
	"github.com/alekseyklimenko/go-proj-bootstrap/services/validation"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, conf *config.Config) {
	services.Some = some.NewService(db)
	services.Validation = validation.NewService()
	services.Processing = processing.NewService(db, conf)
}

func Shutdown() {
	services.Processing.Shutdown()
}
