package cache

import "printer_api/entity"


type PrinterCache interface {
	Set(key string, value *entity.Printer)
	Get(key string) *entity.Printer
}
