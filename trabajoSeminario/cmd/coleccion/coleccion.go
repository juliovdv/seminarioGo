package main

import (
	"flag"
	"fmt"
	"os"
	"seminarioGo/trabajoSeminario/internal/config"
	"seminarioGo/trabajoSeminario/internal/database"
	"seminarioGo/trabajoSeminario/internal/service/coleccion"

	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := loadConfig()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = createSchema(db)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := coleccion.New(db, cfg)

	for _, p := range service.GetPeliculas() {
		fmt.Println(p)
	}

}

func loadConfig() *config.Config {
	archConfig := flag.String("config", "./config/config.yaml", "...")
	flag.Parse()

	cfg, err := config.LoadConfig(*archConfig)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS coleccion (
		id integer primary key autoincrement,
		nombre varchar,
		director varchar,
		anio integer);`
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	insertarPelicula := `INSERT INTO coleccion (nombre, director, anio) VALUES (?, ?, ?)`
	db.MustExec(insertarPelicula, "Rambo", "Rambo", "1990")
	return nil
}
