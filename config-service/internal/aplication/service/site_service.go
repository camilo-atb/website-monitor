package service

import (
	"config-service/internal/domain/model"
	"config-service/internal/domain/ports"
	"fmt"
	"time"
)

// USE CASE = SERVICE

const (
	ErrURLRequired       = "URL no puede estar vacía"
	ErrReviewTimeInvalid = "ReviewTime debe ser mayor a 0"
	ErrIdInvalid         = "ID debe ser mayor a 0"
)

type SiteService struct {
	repo ports.OutputPort // inyección de dependencias
}

func NewSiteService(repo ports.OutputPort) ports.InputPort {
	return &SiteService{repo: repo}
}

func (s *SiteService) Create(input ports.CreateSiteInput) error {
	if input.URL == "" {
		return fmt.Errorf(ErrURLRequired)
	}

	if input.ReviewTime <= 0 {
		return fmt.Errorf(ErrReviewTimeInvalid)
	}

	err := s.repo.Save(model.MonitoredURL{
		URL:          input.URL,
		ReviewTime:   input.ReviewTime,
		CreationDate: time.Now(),
		ModifyDate:   time.Now(),
	})

	if err != nil {
		return fmt.Errorf("error al guardar el sitio: %w", err)
	}

	return nil
}

func (s *SiteService) Update(ID int, input ports.UpdateSiteInput) error {

	if ID <= 0 {
		return fmt.Errorf(ErrIdInvalid)
	}

	site, err := s.repo.FindByID(ID)
	if err != nil {
		return fmt.Errorf("site no encontrado: %w", err)
	}

	if input.URL != nil {
		if *input.URL == "" {
			return fmt.Errorf(ErrURLRequired)
		}
		site.URL = *input.URL
	}

	if input.ReviewTime != nil {
		if *input.ReviewTime <= 0 {
			return fmt.Errorf(ErrReviewTimeInvalid)
		}
		site.ReviewTime = *input.ReviewTime
	}

	site.ModifyDate = time.Now()

	err = s.repo.Update(site)
	if err != nil {
		return fmt.Errorf("error al actualizar el sitio: %w", err)
	}

	return nil
}

func (s *SiteService) Delete(ID int) error {

	if ID <= 0 {
		return fmt.Errorf(ErrIdInvalid)
	}

	err := s.repo.Delete(ID)

	if err != nil {
		return fmt.Errorf("error al eliminar el sitio: %w", err)
	}

	return nil
}

func (s *SiteService) List() ([]model.MonitoredURL, error) {

	URLs, err := s.repo.FindAll()

	if err != nil {
		return nil, fmt.Errorf("error al listar los sitios: %w", err)
	}

	return URLs, nil
}

/*
¿Qué es el site_service?

Es la implementación del InputPort.

👉 Traducción simple:

El usuario dice: “quiero crear un sitio”
El site_service decide cómo se hace eso
*/

// El service IMPLEMENTA el InputPort
