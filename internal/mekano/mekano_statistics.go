package mekano

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type paymentStatistics struct {
	RangoRC     string `json:"rango-rc"`
	Bancolombia int    `json:"bancolombia"`
	Davivienda  int    `json:"davivienda"`
	Susuerte    int    `json:"susuerte"`
	PayU        int    `json:"payu"`
	Efectivo    int    `json:"efectivo"`
	Total       int    `json:"total"`
}
type billingStatistics struct {
	Debito  float64 `json:"debito"`
	Credito float64 `json:"credito"`
	Base    float64 `json:"base"`
}

var (
	efectivo    int = 0
	bancolombia int = 0
	davivienda  int = 0
	susuerte    int = 0
	payU        int = 0
	total       int = 0
)

func PaymentStatistics(data []MekanoData, initialRC, lastRC int) {

	for _, d := range data {
		debito, err := strconv.Atoi(d.Debito)
		total += debito
		if err != nil {
			log.Println(err)
		}
		switch d.Cuenta {
		case "11050501": //Efectivo
			efectivo += debito
		case "11200501": //Bancolombia
			bancolombia += debito
		case "11200510": //Davivienda
			davivienda += debito
		case "13452501": //Pay U
			susuerte += debito
		case "13452505": //Susuerte
			payU += debito
		}
	}

	s := paymentStatistics{
		RangoRC:     fmt.Sprintf("%d-%d", initialRC+1, lastRC),
		Efectivo:    efectivo,
		Bancolombia: bancolombia,
		Davivienda:  davivienda,
		Susuerte:    susuerte,
		PayU:        payU,
		Total:       total,
	}

	result, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(result))

}

var (
	d, c, b float64 = 0, 0, 0
)

func BillingStatistics(data []MekanoData) {

	for _, row := range data {
		debito, _ := strconv.ParseFloat(row.Debito, 64)
		d += debito
		credito, _ := strconv.ParseFloat(row.Credito, 64)
		c += credito
		base, _ := strconv.ParseFloat(row.Base, 64)
		b += base
	}

	bs := billingStatistics{
		Debito:  d,
		Credito: c,
		Base:    b,
	}

	result, err := json.MarshalIndent(bs, "", "    ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(result))

}
