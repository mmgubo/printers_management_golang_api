package service

import "printer_api/entity"






type PrinterService interface {
	Validate(printer *entity.Printer) error
	Create(printer *entity.Printer) (*entity.Post, error)
	FindAll() ([]entity.Printer, error)
	FindByID(id string) (*entity.Printer, error)
}


type service struct{}