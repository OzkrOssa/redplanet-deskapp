package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"

	"github.com/OzkrOssa/redplanet-deskapp/internal/mekano"
)

var CurrentTimeForMekanoInterfaceField string = time.Now().Format("02/01/2006 15:04")
var FormatedTimeSQL string = time.Now().Format("2006-01-02")
var PaymentFileDirPath string = "C:/Users/devre/OneDrive/pagos_mekano/"
var BillingFileDirPath string = "C:/Users/devre/OneDrive/facturacion_mekano/"
var MekanoInterfacesDirPath string = "C:/APOLOSOFT/MEKANO_REMOTO/INTERFACES/"

var Cashier = map[string]string{
	"CLAUDIA PATRICIA ACEVEDO MOTATO":     "11050501",
	"BANCOLOMBIA B":                       "11200501",
	"DAVIVIENDA D":                        "11200510",
	"cartera":                             "13050501",
	"PAY U":                               "13452501",
	"SUSUERTE S":                          "13452505",
	"PAYU":                                "13452501",
	"OFICINA SUPIA":                       "11050501",
	"GILDARDO DE JESUS ESPINOSA GUAPACHA": "11050501",
	"ANA MARIA GARCES RIOS":               "11050501",
	"JAVIER SANCHEZ RAMIREZ":              "11050501",
	"LIZETH VANESSA MONTOYA":              "11050501",
	"ANGIE XIMENA RESTREPO":               "11050501",
	"MARIA ISABEL BETANCUR ZULUAGA":       "11050501",
	"JULIO SICARD BOLIVAR TAMAYO":         "11050501",
}

var CostCenter = map[string]string{
	"RIOSUCIO.": "101",
	"RIOSUCIO":  "101",
	"QUINCHIA":  "100",
	"SUPIA":     "102",
	"MANIZALES": "103",
}

var Account = map[string]string{
	"RESIDENCIAL BASICO RURAL":                     "41457001",
	"RESIDENCIAL BASICO GRAVADO":                   "41457002",
	"BANDA ANCHA PLAN BASICO HOGAR GRAVADO":        "41457051",
	"RESIDENCIAL BASICO":                           "41457090",
	"RESIDENCIAL AVANZADO":                         "41457003",
	"COMERCIAL AVANZADO":                           "41457011",
	"COMERCIAL BASICO":                             "41457010",
	"BANDA ANCHA FIBRA 50 MEGAS RESIDENCIAL":       "41457053",
	"BANDA ANCHA FIBRA COMERCIAL 20 MEGAS":         "41457071",
	"BANDA ANCHA FIBRA COMERCIAL BASICO":           "41457070",
	"PYME 1 GRAVADO":                               "41457015",
	"BANDA ANCHA FIBRA 20 MEGAS RESIDENCIAL":       "41457052",
	"BANDA ANCHA PLAN BASICO HOGAR":                "41457050",
	"BANDA ANCHA FIBRA COMERCIAL 40 MEGAS GRAVADO": "41457095",
	"RECONEXION GRAVADO":                           "41459010",
	"PYME FIBRA 20 MEGAS GRAVADO":                  "41457096",
	"SERVICIO DE INSTALACION":                      "41459014",
	"IP FIJA":                                      "41459018",
	"SERVICIO DE MANTENIMIENTO":                    "41459016",
	"SERVICIO PUNTO A PUNTO":                       "41459012",
	"CABLE":                                        "41459502",
	"CANALETA":                                     "41459501",
	"EQUIPOS":                                      "41459503",
	"CARGADOR":                                     "41459504",
	"BANDA ANCHA FIBRA COMERCIAL 50 MEGAS":         "41457072",
	"COMERCIAL PLUS":                               "41457092",
	"BANDA ANCHA FIBRA COMERCIAL 30 MEGAS GRAVADO": "41457094",
	"PYME 2 GRAVADO":                               "41457093",
	"RESIDENCIAL PLUS GRAVADO":                     "41457091",
	"TRASLADO":                                     "41459020",
	"CONECTIVIDAD LAN TO LAN":                      "41459024",
	"PUNTO DE ACCESO (AP)":                         "41459022",
	"MODIFICACION":                                 "41459026",
	"IP PUBLICA":                                   "41459030",
	"SERVICIO ROUTER ADICIONAL":                    "41459028",
	"PYME 2":                                       "41457093",
	"PYME FIBRA 20 MEGAS":                          "41457096",
	"PLAN BASICO HOGAR GRAVADO 2":                  "41457051",
	"COMERCIAL BASICO 10MBPS":                      "41457010",
	"F.O. COMERCIAL BASICO., IP PUBLICA":           "41457070",
	"RESIDENCIAL BASICO 10MBPS (ESPECIAL)":         "41457090",
	"FIBRA OPTICA 20 MEGAS":                        "41457052",
	"PLAN BASICO HOGAR GRAVADO (ESPECIAL)":         "41457051",
	"IP PUBLICA, RESIDENCIAL BASICO _C16":          "41457090",
	"RESIDENCIAL BASICO 2":                         "41457090",
	"RESIDENCIAL BASICO 10MBPS":                    "41457090",
	"FIBRA OPTICA 50 MEGAS 2":                      "41457053",
	"RESIDENCIAL BASICO 2 (ESPECIAL)":              "41457090",
	"F.O. COMERCIAL 50MBPS":                        "41457072",
	"PLAN BASICO HOGAR (ESPECIAL)":                 "41457050",
	"RESIDENCIAL BASICO _C16 (ESPECIAL)":           "41457090",
	"F.O. COMERCIAL 20MBPS":                        "41457071",
	"F.O. COMERCIAL 20MBPS, IP PUBLICA":            "41457071",
	"RESIDENCIAL BASICO _C16":                      "41457090",
	"RESIDENCIAL BASICO_RURAL":                     "41457001",
	"RESIDENCIAL AVANZADO 2":                       "41457003",
	"RESIDENCIAL BASICO 10MBPS 2":                  "41457090",
	"RESIDENCIAL BASICO 3":                         "41457090",
	"F.O. COMERCIAL BASICO.":                       "41457070",
	"RESIDENCIAL BASICO 5":                         "41457090",
	"PLAN BASICO HOGAR 3":                          "41457050",
	"IP PUBLICA, PLAN BASICO HOGAR 2":              "41457050",
	"RESIDENCIAL BASICO RURAL 3":                   "41457001",
	"RESIDENCIAL BASICO RURAL 2":                   "41457001",
	"FIBRA OPTICA 20 MEGAS 2":                      "41457052",
	"EQUIPO QUEMADO INT":                           "41459503",
	"PLAN BASICO HOGAR 2 (ESPECIAL)":               "41457050",
	"METRO CABLE UTP":                              "41459502",
	"CABLE FIBRA":                                  "41459502",
	"RESIDENCIAL BASICO (ESPECIAL)":                "41457090",
	"TRASLADOS INT":                                "41459020",
	"RECONEXION":                                   "41459010",
	"INSTALACION SERVICIO":                         "41459014",
	"INSTALACION INT":                              "41459014",
	"PLAN BASICO HOGAR.":                           "41457050",
	"RESIDENCIAL BASICO 6":                         "41457090",
	"PLAN BASICO HOGAR 2":                          "41457050",
	"PYME 1":                                       "41457015",
	"SERVICIO INSTALACION":                         "41459014",
	"IP PUBLICA, PLAN BASICO HOGAR 2 (ESPECIAL)":   "41457050",
	"PLAN BASICO HOGAR 3, ROUTER ADICIONAL":        "41457050",
	"PLAN SENIOR":                                  "41457054",
	"PLAN FINCA BÁSICO":                            "41457004",
	"PLAN FINCA AVANZADO":                          "41457005",
	"PLAN MÁSTER":                                  "41457055",
	"PYME 1 (ESPECIAL)":                            "41457015",
	"PLAN SENIOR COMERCIAL":                        "41457074",
	"UNIFI":                                        "41459038",
	"SERVICIO A LA MEDIDA 100 MBPS":                "41457076",
	"5 IPV4 /29":                                   "41459032",
	"INTERCONEXION ENTRE SEDES":                    "41459036",
	"INTERNET DEDICADO 100 MBPS":                   "41457020",
	"PLAN MASTER COMERCIAL":                        "41457074",
	"ROUTER ADICIONAL":                             "41459028",
	"F.O. COMERCIAL 30MBPS":                        "41457094",
	"TV DIGITAL":                                   "1",
	"SUSCRIPCION":                                  "41459034",
	"PLAN FINCA BASICO":                            "41457004",
	"COMERCIAL BASICO RURAL":                       "41457012",
	"PLAN MASTER":                                  "41457055",
	"PLAN SENIOR GRAVADO":                          "41457056",
}

var Municipalities = map[string]string{
	"RIOSUCIO": "101",
	"QUINCHIA": "100",
	"SUPIA":    "102",
}

// Create a file in C:/APOLOSOFT/MEKANO_REMOTO/INTERFACES/
func FileExporter(d []mekano.MekanoData) error {
	text, err := os.Create(filepath.Join("/home/oscar/Documentos/", "CONTABLE.txt"))
	if err != nil {
		return err
	}

	defer text.Close()

	csvWriter := csv.NewWriter(text)
	csvWriter.Comma = ','

	for _, data := range d {
		row := []string{
			data.Tipo,
			data.Prefijo,
			data.Numero,
			data.Secuencia,
			data.Fecha,
			data.Cuenta,
			data.Terceros,
			data.CentroCostos,
			data.Nota,
			data.Debito,
			data.Credito,
			data.Base,
			data.Aplica,
			data.TipoAnexo,
			data.PrefijoAnexo,
			data.NumeroAnexo,
			data.Usuario,
			data.Signo,
			data.CuentaCobrar,
			data.CuentaPagar,
			data.NombreTercero,
			data.NombreCentro,
			data.Interface,
		}
		csvWriter.Write(row)
	}
	csvWriter.Flush()

	return nil
}
