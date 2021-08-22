package repository

import "mfundo.com/printers/entity"

type PrinterRepository interface {
	Save(printer *entity.Printer) (*entity.Printer, error)
	FindAll() ([]entity.Printer, error)
	FindByID(id string) (*entity.Printer, error)
	//Delete(printer *entity.Printer) error
}
