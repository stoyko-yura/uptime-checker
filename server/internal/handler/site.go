package handler

import (
	"encoding/json"
	"net/http"
	"uptime-checker/internal/model"

	"gorm.io/gorm"
)

type SiteHandler struct {
	DB *gorm.DB
}

func (h *SiteHandler) GetSites(w http.ResponseWriter, r *http.Request) {
	var sites []model.Site

	if err := h.DB.Find(&sites).Error; err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(sites)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (h *SiteHandler) CreateSite(w http.ResponseWriter, r *http.Request) {
	var site model.Site

	if err := json.NewDecoder(r.Body).Decode(&site); err != nil {
		http.Error(w, "Invalid data format", http.StatusBadRequest)
		return
	}

	if site.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&site).Error; err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(site)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (h *SiteHandler) DeleteSite(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	result := h.DB.Delete(&model.Site{}, id)
	if result.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Site not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
