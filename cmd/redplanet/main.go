package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/OzkrOssa/redplanet-deskapp/internal/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

func main() {
	ctx := context.Background()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))

	if err != nil {
		log.Println("Error al conectar la base de datos ", err)
	}

	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println("Error al pinar la base de datos ", err)
	}

	paymentExcel, err := excelize.OpenFile("C:/Users/devre/OneDrive/pagos_mekano/PAGOS 3 ENERO 2023 PARA MEKANO.xlsx")
	if err != nil {
		log.Println(err)
	}

	payment := repository.NewMekanoPayment(ctx, paymentExcel, db)
	err = payment.GenerateFile()
	if err != nil {
		log.Println("Error al generar el archivo ", err)
	}

	// billingExcel, err := excelize.OpenFile("C:/Users/devre/OneDrive/facturacion_mekano/REPORTE FACTURACION ELECTRONICA MARZO 2023.xlsx")
	// if err != nil {
	// 	log.Println(err)
	// }

	// extras, err := excelize.OpenFile("C:/Users/devre/OneDrive/facturacion_mekano/extras.xlsx")
	// if err != nil {
	// 	log.Println(err)
	// }

	// billing := repository.NewMekanoBilling(ctx, billingExcel, extras)
	// billing.GenerateFile()
}
