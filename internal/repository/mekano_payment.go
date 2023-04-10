package repository

import (
	"log"
	"strconv"
	"sync"

	"github.com/OzkrOssa/redplanet-deskapp/internal/mekano"
	"github.com/OzkrOssa/redplanet-deskapp/internal/utils"
	"github.com/xuri/excelize/v2"
)

type mekanoPayment struct {
	xlsx *excelize.File
	db   string
}

func NewMekanoPayment(excelFile *excelize.File, db string) *mekanoPayment {
	return &mekanoPayment{excelFile, db}
}

func (mp *mekanoPayment) GenerateFile() error {
	var wg sync.WaitGroup
	var paymentResult []mekano.MekanoData

	paymentFile, err := mp.xlsx.GetRows(mp.xlsx.GetSheetName(0))
	if err != nil {
		log.Println(err)
	}

	for i, row := range paymentFile[1:] {
		rowCount := i + 1
		wg.Add(1)
		go func(row []string) {
			defer wg.Done()
			paymentData := mekano.MekanoData{
				Tipo:          "RC",
				Prefijo:       "_",
				Numero:        strconv.Itoa(rowCount), //FIXME: dababase consecutive + i
				Secuencia:     "",
				Fecha:         row[4],
				Cuenta:        "13050501",
				Terceros:      row[1],
				CentroCostos:  "C1",
				Nota:          "RECAUDO POR VENTA SERVICIOS",
				Debito:        "0",
				Credito:       row[5],
				Base:          "0",
				Aplica:        "",
				TipoAnexo:     "",
				PrefijoAnexo:  "",
				NumeroAnexo:   "",
				Usuario:       "SUPERVISOR",
				Signo:         "",
				CuentaCobrar:  "",
				CuentaPagar:   "",
				NombreTercero: row[2],
				NombreCentro:  "CENTRO DE COSTOS GENERAL",
				Interface:     utils.CurrentTimeForMekanoInterfaceField,
			}
			paymentData2 := mekano.MekanoData{
				Tipo:          "RC",
				Prefijo:       "_",
				Numero:        strconv.Itoa(rowCount), //FIXME: dababase consecutive + i
				Secuencia:     "",
				Fecha:         row[4],
				Cuenta:        utils.Cashier[row[9]],
				Terceros:      row[1],
				CentroCostos:  "C1",
				Nota:          "RECAUDO POR VENTA SERVICIOS",
				Debito:        row[5],
				Credito:       "0",
				Base:          "0",
				Aplica:        "",
				TipoAnexo:     "",
				PrefijoAnexo:  "",
				NumeroAnexo:   "",
				Usuario:       "SUPERVISOR",
				Signo:         "",
				CuentaCobrar:  "",
				CuentaPagar:   "",
				NombreTercero: row[2],
				NombreCentro:  "CENTRO DE COSTOS GENERAL",
				Interface:     utils.CurrentTimeForMekanoInterfaceField,
			}

			paymentResult = append(paymentResult, paymentData, paymentData2)

		}(row)

		wg.Wait()
	}

	utils.FileExporter(paymentResult)
	mekano.ShowStatistics(paymentResult)

	return nil
}
