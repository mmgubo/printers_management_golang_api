package service

import (
	"errors"
	"math/rand"
	"strconv"

	"mfundo.com/printers/entity"
	"mfundo.com/printers/repository"
)


type PrinterService interface {
	Validate(printer *entity.Printer) error
	Create(printer *entity.Printer) (*entity.Printer, error)
	FindAll() ([]entity.Printer, error)
	FindByID(id string) (*entity.Printer, error)
	
}


type service struct{}

var (
	repo repository.PrinterRepository
)


func NewPrinterService(repository repository.PrinterRepository) PrinterService {
	repo = repository
	return &service{}
}

func (*service) Validate(printer *entity.Printer) error {
	if printer == nil {
		err := errors.New("The post is empty")
		return err
	}
	if printer.Name == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) Create(printer *entity.Printer) (*entity.Printer, error) {
	printer.ID = rand.Int63()
	return repo.Save(printer)
}

func (*service) FindAll() ([]entity.Printer, error) {
	return repo.FindAll()
}

func (*service) FindByID(id string) (*entity.Printer, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return repo.FindByID(id)
}








