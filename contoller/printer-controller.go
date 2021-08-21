package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mmgubo/printers/service"
)

type controller struct{}

var (
	printerService service.PrinterService
)


type PrinterController interface {
	GetPrinterByID(response http.ResponseWriter, request *http.Request)
	GetPrinters(response http.ResponseWriter, request *http.Request)
	AddPrinter(response http.ResponseWriter, request *http.Request)
	UpdatePrinter(response http.ResponseWriter, request *http.Request)
	DeletePrinter(response http.ResponseWriter, request *http.Request)
}

func NewPostController(service service.PrinterService) PrinterController {
	printerService = service
	
	return &controller{}
}

func (*controller) GetPrinters(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	posts, err := printerService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

