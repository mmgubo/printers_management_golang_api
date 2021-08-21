package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"mfundo.com/printers/cache"
	"mfundo.com/printers/entity"
	"mfundo.com/printers/errors"
	"mfundo.com/printers/service"
)

type controller struct{}

var (
	printerService service.PrinterService
	printerCache   cache.PrinterCache
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

func (*controller) GetPrinterByID(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	printerID := strings.Split(request.URL.Path, "/")[2]
	var printer *entity.Printer = printerCache.Get(printerID)
	if printer == nil {
		printer, err := printerService.FindByID(printerID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No printer found!"})
			return
		}
		printerCache.Set(printerID, printer)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(printer)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(printer)
	}

}

func (*controller) AddPrinter(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var printer entity.Printer
	err := json.NewDecoder(request.Body).Decode(&printer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := printerService.Validate(&printer)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := printerService.Create(&printer)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the printer"})
		return
	}
	printerCache.Set(strconv.FormatInt(printer.ID, 10), &printer)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*controller) DeletePrinter(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	printerID := strings.Split(request.URL.Path, "/")[2]
	var printer *entity.Printer = printerCache.Get(printerID)
	if printer == nil {
		printer, err := printerService.FindByID(printerID)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(errors.ServiceError{Message: "No printer found!"})
			return
		}
		printerCache.Set(printerID, printer)
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(printer)
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(printer)
	}

}

func (*controller) UpdatePrinter(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var printer entity.Printer
	err := json.NewDecoder(request.Body).Decode(&printer)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}
	err1 := printerService.Validate(&printer)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := printerService.Create(&printer)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the printer"})
		return
	}
	printerCache.Set(strconv.FormatInt(printer.ID, 10), &printer)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}



