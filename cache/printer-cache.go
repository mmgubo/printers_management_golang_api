package cache

import "mfundo.com/printers/entity"


type PrinterCache interface {
	Set(key string, value *entity.Printer)
	Get(key string) *entity.Printer
}
