package main

import (
	"log"

	"github.com/OzkrOssa/redplanet-deskapp/internal/repository"
	"github.com/xuri/excelize/v2"
)

func main() {
	db := "SQL"

	paymentExcel, err := excelize.OpenFile("/home/oscar/Documentos/PAGOS 3 ENERO 2023 PARA MEKANO.xlsx")
	if err != nil {
		log.Println(err)
	}

	payment := repository.NewMekanoPayment(paymentExcel, db)
	payment.GenerateFile()

	// billingExcel, err := excelize.OpenFile("/home/oscar/Documentos/REPORTE FACTURACION ELECTRONICA FEBRERO 2023.xlsx")
	// if err != nil {
	// 	log.Println(err)
	// }

	// extras, err := excelize.OpenFile("/home/oscar/Documentos/CLIENTES CON 2 SERVICIOS MENSUALES MARZO.xlsx")
	// if err != nil {
	// 	log.Println(err)
	// }

	// billing := repository.NewMekanoBilling(billingExcel, extras, db)
	// billing.GenerateFile()
}
