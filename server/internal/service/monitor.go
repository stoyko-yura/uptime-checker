package service

import (
	"log"
	"net/http"
	"time"
	"uptime-checker/internal/model"

	"gorm.io/gorm"
)

type MonitorService struct {
	DB *gorm.DB
}

func (s *MonitorService) Start() {
	for {
		var sites []model.Site
		s.DB.Find(&sites)

		for _, site := range sites {
			go s.checkSite(site)
		}

		time.Sleep(1 * time.Minute)
	}
}

func (s *MonitorService) checkSite(site model.Site) {
	startTime := time.Now()

	resp, err := http.Get(site.URL)
	duration := time.Since(startTime).Milliseconds()

	statusCode := 0
	if err == nil {
		statusCode = resp.StatusCode
		err := resp.Body.Close()
		if err != nil {
			return
		}
	} else {
		log.Printf("check error %s: %v", site.URL, err)
	}

	check := model.Check{
		SiteID:         site.ID,
		StatusCode:     statusCode,
		ResponseTimeMs: duration,
		CheckedAt:      time.Now(),
	}
	s.DB.Create(&check)

	s.DB.Model(&site).Update("LastStatus", statusCode)

	log.Printf("Checked: %s | Status: %d | Time: %dms", site.URL, statusCode, duration)
}
