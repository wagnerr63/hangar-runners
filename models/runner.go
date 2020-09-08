package models

import (
	"time"

	"../db"
)

type Runner struct {
	ID          int
	Nome        string
	Funcao      string
	Aniversario string
	HangarRun   string
}

func GetAllRunners() []Runner {

	db := db.ConectDatabase()

	selectAllRunners, err := db.Query("SELECT * FROM runners ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	r := Runner{}
	runners := []Runner{}

	for selectAllRunners.Next() {
		var id int
		var nome, funcao, hangarRun string
		var aniversario time.Time

		err = selectAllRunners.Scan(&id, &nome, &funcao, &aniversario, &hangarRun)
		if err != nil {
			panic(err.Error())
		}

		r.ID = id
		r.Nome = nome
		r.Aniversario = aniversario.Format("02/01/2006")
		r.Funcao = funcao
		r.HangarRun = hangarRun

		runners = append(runners, r)
	}
	defer db.Close()
	return runners
}

func AddRunner(nome, funcao string, aniversario time.Time, hangarRun string) {
	db := db.ConectDatabase()

	insertRunner, err := db.Prepare("INSERT INTO runners (nome, funcao, aniversario, hangar_run) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertRunner.Exec(nome, funcao, aniversario, hangarRun)

	defer db.Close()

}

func DeleteRunner(id string) {
	db := db.ConectDatabase()

	deleteRunner, err := db.Prepare("DELETE FROM runners WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteRunner.Exec(id)
	db.Close()
}

func GetRunner(id string) Runner {
	db := db.ConectDatabase()

	runnerResult, err := db.Query("SELECT * FROM runners WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	runner := Runner{}
	for runnerResult.Next() {
		var id int
		var nome, funcao, hangarRun string
		var aniversario time.Time

		err = runnerResult.Scan(&id, &nome, &funcao, &aniversario, &hangarRun)
		if err != nil {
			panic(err.Error())
		}

		runner.ID = id
		runner.Nome = nome
		runner.Funcao = funcao
		runner.Aniversario = aniversario.Format("2006-01-02")
		runner.HangarRun = hangarRun
	}

	defer db.Close()
	return runner

}

func UpdateRunner(id int, nome, funcao string, aniversario time.Time, hangarRun string) {
	db := db.ConectDatabase()

	updateRunner, err := db.Prepare("UPDATE runners SET nome=$1, funcao=$2, aniversario=$3, hangar_run=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateRunner.Exec(nome, funcao, aniversario, hangarRun, id)

	defer db.Close()

}
