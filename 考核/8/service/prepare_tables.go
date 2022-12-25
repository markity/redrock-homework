package service

import (
	"depo/dao"
	"log"
)

func MustPrepareTables() {
	_, err := dao.DB.Exec(dao.SentenceCreateCargo)
	if err != nil {
		log.Panicf("failed to drop table cargo: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateDepository)
	if err != nil {
		log.Panicf("failed to drop table depository: %v\n", err)
	}

	_, err = dao.DB.Exec(dao.SentenceCreateCargoDepository)
	if err != nil {
		log.Panicf("failed to drop table cargo_depository: %v\n", err)
	}
}
