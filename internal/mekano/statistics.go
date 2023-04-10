package mekano

import (
	"log"
	"strconv"
)

var (
	efectivo    int = 0
	bancolombia int = 0
	davivienda  int = 0
	susuerte    int = 0
	payU        int = 0
)

func ShowStatistics(data []MekanoData) {

	for _, d := range data {
		account, err := strconv.Atoi(d.Cuenta)
		if err != nil {
			log.Println(err)
		}
		switch d.Cuenta {
		case "11050501": //Efectivo
			efectivo += account
		case "11200501": //Bancolombia
			bancolombia += account
		case "11200510": //Davivienda
			davivienda += account
		case "13452501": //Pay U
			susuerte += account
		case "13452505": //Susuerte
			payU += account
		}
	}

	log.Println(efectivo, bancolombia, davivienda, susuerte, payU)

}
