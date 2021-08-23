package main

import (

	//"mfundo.com/printers/cache"

	"os"

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



func main() {


     //myRouter := mux.NewRouter().StrictSlash(true)

     //myRouter.HandleFunc("/", homePage)

     httpRouter.GET("/printers", printerController.GetPrinters)
     httpRouter.POST("/printers", printerController.AddPrinter)
     httpRouter.GET("/printers/{id}", printerController.GetPrinterByID)

    httpRouter.SERVE(os.Getenv("PORT"))
    
    
    //log.Fatal(http.ListenAndServe(":4209", myRouter))

}


// func homePage(w http.ResponseWriter, r *http.Request){
//     fmt.Fprintf(w, "Welcome to the HomePage!")
    
//     fmt.Println("Endpoint Hit: homePage")
// }



