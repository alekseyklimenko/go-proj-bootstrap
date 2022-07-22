package processing

import (
	"fmt"
	"github.com/alekseyklimenko/go-proj-bootstrap/config"
	"github.com/alekseyklimenko/go-proj-bootstrap/logger"
	"github.com/alekseyklimenko/go-proj-bootstrap/models"
	"github.com/alekseyklimenko/go-proj-bootstrap/services"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	db                *gorm.DB
	mu                sync.Mutex
	processItemsCount uint
	maxItemsProcess   uint
	items             []models.Some
}

var startChan = make(chan models.Some)

func NewService(db *gorm.DB, conf *config.Config) *Service {
	service := &Service{
		db:                db,
		processItemsCount: 0,
		maxItemsProcess:   2, //conf.Process.MaxItemsProcess,
	}
	service.items = make([]models.Some, 0, service.maxItemsProcess)
	go service.startChanListener()
	go service.initProcessing()
	return service
}

func (s *Service) QueueItem(some models.Some) {
	startChan <- some
}

func (s *Service) Shutdown() {
	close(startChan)
	for _, some := range s.items {
		services.Some.FreeItem(some.ID)
	}
}

func (s *Service) startChanListener() {
	for {
		item, ok := <-startChan
		if !ok {
			logger.NewEntry().Info("Closing Init processing chanel")
			break
		}
		logger.NewEntry().WithField("id", item.ID).Info("Start processing")
		err := services.Some.LockItem(item.ID)
		if err != nil {
			logger.NewEntry().WithField("id", item.ID).Error("Skip item, got error:")
			logger.NewEntry().Error(err.Error())
		}
		s.items = append(s.items, item)
		// some logic
	}
}

func (s *Service) initProcessing() {
	logger.NewEntry().Info(fmt.Sprintf("Getting to process %d items", s.maxItemsProcess))
	items := services.Some.GetSomeToProcess(s.maxItemsProcess)
	for _, item := range *items {
		s.mu.Lock()
		if s.processItemsCount == s.maxItemsProcess {
			s.mu.Unlock()
			break
		}
		s.processItemsCount++
		s.mu.Unlock()
		startChan <- item
	}
}
