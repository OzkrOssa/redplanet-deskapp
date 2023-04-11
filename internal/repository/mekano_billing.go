package repository

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/OzkrOssa/redplanet-deskapp/internal/mekano"
	"github.com/OzkrOssa/redplanet-deskapp/internal/utils"
	"github.com/mozillazg/go-unidecode"
	"github.com/xuri/excelize/v2"
)

var (
	billingNormal     mekano.MekanoData
	billingNormalPlus mekano.MekanoData
	billingIva        mekano.MekanoData
	billingIvaPlus    mekano.MekanoData
	billingCxC        mekano.MekanoData
	billingCxCPlus    mekano.MekanoData
)

type mekanoBilling struct {
	xlsx   *excelize.File
	extras *excelize.File
}

func NewMekanoBilling(ctx context.Context, xlsx, extras *excelize.File) *mekanoBilling {
	return &mekanoBilling{xlsx, extras}
}

func (mb *mekanoBilling) GenerateFile() error {
	var montoBaseFinal float64
	var billingResult []mekano.MekanoData
	var wg sync.WaitGroup

	billingFile, err := mb.xlsx.GetRows(mb.xlsx.GetSheetName(0))
	if err != nil {
		return err
	}

	itemsIvaFile, err := mb.extras.GetRows(mb.extras.GetSheetName(0))
	if err != nil {
		return err
	}

	for _, bRow := range billingFile[1:] {
		wg.Add(1)
		go func(bRow []string) {
			defer wg.Done()
			montoBase, err := strconv.ParseFloat(bRow[12], 64)

			if err != nil {
				log.Println(err)
			}

			montoIva, err := strconv.ParseFloat(strings.TrimSpace(bRow[13]), 64)
			if err != nil {
				log.Println(err)
			}

			_, decimal := math.Modf(montoBase)

			if decimal >= 0.5 {
				montoBaseFinal = math.Ceil(montoBase)
			} else {
				montoBaseFinal = math.Round(montoBase)
			}

			if !strings.Contains(bRow[21], ",") {
				billingNormal = mekano.MekanoData{
					Tipo:          "FVE",
					Prefijo:       "_",
					Numero:        bRow[8],
					Secuencia:     "",
					Fecha:         bRow[9],
					Cuenta:        utils.Account[bRow[21]],
					Terceros:      bRow[1],
					CentroCostos:  utils.CostCenter[unidecode.Unidecode(bRow[17])],
					Nota:          "FACTURA ELECTRÓNICA DE VENTA",
					Debito:        "0",
					Credito:       fmt.Sprintf("%f", math.Ceil(montoBase)),
					Base:          "0",
					Aplica:        "",
					TipoAnexo:     "",
					PrefijoAnexo:  "",
					NumeroAnexo:   "",
					Usuario:       "SUPERVISOR",
					Signo:         "",
					CuentaCobrar:  "",
					CuentaPagar:   "",
					NombreTercero: bRow[2],
					NombreCentro:  bRow[17],
					Interface:     utils.CurrentTimeForMekanoInterfaceField,
				}

				billingIva = mekano.MekanoData{
					Tipo:          "FVE",
					Prefijo:       "_",
					Numero:        bRow[8],
					Secuencia:     "",
					Fecha:         bRow[9],
					Cuenta:        "24080505",
					Terceros:      bRow[1],
					CentroCostos:  utils.CostCenter[unidecode.Unidecode(bRow[17])],
					Nota:          "FACTURA ELECTRÓNICA DE VENTA",
					Debito:        "0",
					Credito:       fmt.Sprintf("%f", montoIva),
					Base:          fmt.Sprintf("%f", montoBaseFinal),
					Aplica:        "",
					TipoAnexo:     "",
					PrefijoAnexo:  "",
					NumeroAnexo:   "",
					Usuario:       "SUPERVISOR",
					Signo:         "",
					CuentaCobrar:  "",
					CuentaPagar:   "",
					NombreTercero: bRow[2],
					NombreCentro:  bRow[17],
					Interface:     utils.CurrentTimeForMekanoInterfaceField,
				}

				billingCxC = mekano.MekanoData{
					Tipo:          "FVE",
					Prefijo:       "_",
					Numero:        bRow[8],
					Secuencia:     "",
					Fecha:         bRow[9],
					Cuenta:        "13050501",
					Terceros:      bRow[1],
					CentroCostos:  utils.CostCenter[unidecode.Unidecode(bRow[17])],
					Nota:          "FACTURA ELECTRÓNICA DE VENTA",
					Debito:        bRow[14],
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
					NombreTercero: bRow[2],
					NombreCentro:  bRow[17],
					Interface:     utils.CurrentTimeForMekanoInterfaceField,
				}
				billingResult = append(billingResult, billingNormal, billingIva, billingCxC)
			} else {
				splitBillingItems := strings.Split(bRow[21], ",")
				for _, item := range splitBillingItems {
					for _, itemIva := range itemsIvaFile[1:] {
						if itemIva[1] == strings.TrimSpace(item) && itemIva[0] == bRow[0] {
							itemIvaBase, _ := strconv.ParseFloat(itemIva[2], 64)
							billingNormalPlus = mekano.MekanoData{
								Tipo:          "FVE",
								Prefijo:       "_",
								Numero:        bRow[8],
								Secuencia:     "",
								Fecha:         bRow[9],
								Cuenta:        utils.Account[unidecode.Unidecode(strings.TrimSpace(item))],
								Terceros:      bRow[1],
								CentroCostos:  utils.CostCenter[unidecode.Unidecode(bRow[17])],
								Nota:          "FACTURA ELECTRÓNICA DE VENTA",
								Debito:        "0",
								Credito:       fmt.Sprintf("%f", math.Ceil(itemIvaBase-1)),
								Base:          "0",
								Aplica:        "",
								TipoAnexo:     "",
								PrefijoAnexo:  "",
								NumeroAnexo:   "",
								Usuario:       "SUPERVISOR",
								Signo:         "",
								CuentaCobrar:  "",
								CuentaPagar:   "",
								NombreTercero: bRow[2],
								NombreCentro:  bRow[17],
								Interface:     utils.CurrentTimeForMekanoInterfaceField,
							}
							billingResult = append(billingResult, billingNormalPlus)
						}
					}
				}
				billingIvaPlus = mekano.MekanoData{
					Tipo:          "FVE",
					Prefijo:       "_",
					Numero:        bRow[8],
					Secuencia:     "",
					Fecha:         bRow[9],
					Cuenta:        "24080505",
					Terceros:      bRow[1],
					CentroCostos:  utils.CostCenter[unidecode.Unidecode(bRow[17])],
					Nota:          "FACTURA ELECTRÓNICA DE VENTA",
					Debito:        "0",
					Credito:       fmt.Sprintf("%f", montoIva),
					Base:          fmt.Sprintf("%f", montoBaseFinal),
					Aplica:        "",
					TipoAnexo:     "",
					PrefijoAnexo:  "",
					NumeroAnexo:   "",
					Usuario:       "SUPERVISOR",
					Signo:         "",
					CuentaCobrar:  "",
					CuentaPagar:   "",
					NombreTercero: bRow[2],
					NombreCentro:  bRow[17],
					Interface:     utils.CurrentTimeForMekanoInterfaceField,
				}

				billingCxCPlus = mekano.MekanoData{
					Tipo:          "FVE",
					Prefijo:       "_",
					Numero:        bRow[8],
					Secuencia:     "",
					Fecha:         bRow[9],
					Cuenta:        "13050501",
					Terceros:      bRow[1],
					CentroCostos:  utils.CostCenter[unidecode.Unidecode(bRow[17])],
					Nota:          "FACTURA ELECTRÓNICA DE VENTA",
					Debito:        bRow[14],
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
					NombreTercero: bRow[2],
					NombreCentro:  bRow[17],
					Interface:     utils.CurrentTimeForMekanoInterfaceField,
				}
				billingResult = append(billingResult, billingIvaPlus, billingCxCPlus)
			}

		}(bRow)
		wg.Wait()
	}

	utils.FileExporter(billingResult)
	mekano.BillingStatistics(billingResult)

	return nil
}
