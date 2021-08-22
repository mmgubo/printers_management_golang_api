package main

import (

	//"mfundo.com/printers/cache"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	//"os"

	"mfundo.com/printers/cache"

	"mfundo.com/printers/controller"
	router "mfundo.com/printers/http"
	"mfundo.com/printers/repository"
	"mfundo.com/printers/service"
)
var (
	printerRepository repository.PrinterRepository = repository.NewDynamoDBRepository()
	printerService    service.PrinterService       = service.NewPrinterService(printerRepository)
	printerCache      cache.PrinterCache           = cache.NewRedisCache("localhost:6379", 1, 10)
	printerController controller.PrinterController = controller.NewPrinterController(printerService, printerCache)
	httpRouter     router.Router             = router.NewMuxRouter()
)

// type Printer struct {
//     Id string `json:"Id"`
//     Name string `json:"Name"`
//     Number string `json:"Number"`
//     Status string `json:"Status"`
// }

// let's declare a global Printers array
// that we can then populate in our main function
// to simulate a database
// var Printers []Printer


func main() {


    // httpRouter.GET("/printers", printerController.GetPrinters)
	// httpRouter.GET("/printers/{id}", printerController.GetPrinterByID)
	// httpRouter.POST("/printers", printerController.AddPrinter)

     myRouter := mux.NewRouter().StrictSlash(true)

     myRouter.HandleFunc("/", homePage)

     myRouter.HandleFunc("/printers", printerController.GetPrinters).Methods("GET")
     myRouter.HandleFunc("/printers", printerController.AddPrinter).Methods("POST")
     //myRouter.HandleFunc("/printers/{id}", deletePrinter).Methods("DELETE")
     myRouter.HandleFunc("/printers/{id}", printerController.GetPrinterByID).Methods("GET")


    log.Fatal(http.ListenAndServe(":4208", myRouter))

    // Printers = []Printer{
    // Printer{Id: "1", Name: "Xerox printer", Number: "192.168.280", Status: "Active"},
    // Printer{Id: "2", Name: "Samsung printer", Number: "190.168.177", Status: "Active"},
    // }

	 
    // handleRequests()
}



// func handleRequests() {

// 	 // creates a new instance of a mux router
//     myRouter := mux.NewRouter().StrictSlash(true)

//     myRouter.HandleFunc("/", homePage)

// 	// add our printers route and map it to our 
//     // returnAllPrinters function like so
//     myRouter.HandleFunc("/printers", returnAllPrinters)
//     myRouter.HandleFunc("/printer", createNewPrinter).Methods("POST")
//     myRouter.HandleFunc("/printer/{id}", deletePrinter).Methods("DELETE")
//     myRouter.HandleFunc("/printer/{id}", returnSinglePrinter)
    
//     log.Fatal(http.ListenAndServe(":7004", myRouter))
// }

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    
    fmt.Println("Endpoint Hit: homePage")
}


// func returnAllPrinters(w http.ResponseWriter, r *http.Request){
		

// 	json.NewEncoder(w).Encode(Printers)
//     fmt.Println("Endpoint Hit: returnAllPrinters")
    
// }

// func returnSinglePrinter(w http.ResponseWriter, r *http.Request){
//     vars := mux.Vars(r)
//     key := vars["id"]

//     // Loop over all of our Printers
//     // if the printer.Id equals the key we pass in
//     // return the article encoded as JSON
//     for _, printer := range Printers {
//         if printer.Id == key {
//             json.NewEncoder(w).Encode(printer)
//         }
//     }
// }

// func createNewPrinter(w http.ResponseWriter, r *http.Request) {
//     // get the body of our POST request
//     // return the string response containing the request body    
//     reqBody, _ := ioutil.ReadAll(r.Body)

//     var printer Printer 
//     json.Unmarshal(reqBody, &printer)
    
//     Printers = append(Printers, printer)

//     json.NewEncoder(w).Encode(printer)
// }

// func deletePrinter(w http.ResponseWriter, r *http.Request) {
//     // once again, we will need to parse the path parameters
//     vars := mux.Vars(r)
//     // we will need to extract the `id` of the article we
//     // wish to delete
//     id := vars["id"]

//     // we then need to loop through all our articles
//     for index, printer := range Printers {
//         // if our id path parameter matches one of our
//         // articles
//         if printer.Id == id {
//             // updates our Articles array to remove the 
//             // article
//             Printers = append(Printers[:index], Printers[index+1:]...)
//         }
//     }}

   

