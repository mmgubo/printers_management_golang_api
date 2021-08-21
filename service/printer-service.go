package service

import (
	"mfundo.com/printers/entity"
)






type PrinterService interface {
	Validate(printer *entity.Printer) error
	Create(printer *entity.Printer) (*entity.Printer, error)
	FindAll() ([]entity.Printer, error)
	FindByID(id string) (*entity.Printer, error)
	Delete(printer *entity.Printer) error
}


type service struct{}