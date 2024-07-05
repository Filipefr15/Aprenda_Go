package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string
	Ano   int
	Km    int
	Price float64
}

var cars []Car

/*
	func generateCars() {
		cars = append(cars, Car{"Gol", 2010, 100000, 5000.00})
		cars = append(cars, Car{"Uno", 2010, 100000, 15000.0})
		cars = append(cars, Car{"Celta", 2010, 100000, 35000.0})
		cars = append(cars, Car{"Hb20", 2010, 100000, 70000.0})
		cars = append(cars, Car{"Fusca", 2010, 100000, 10000.0})
	}
*/
func main() {
	// Inicializa o banco de dados e cria a tabela se necessário
	if err := initDB(); err != nil {
		log.Fatal(err)
	}

	//generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

func initDB() error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS cars (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"ano" NUMBER,
		"km" NUMBER,
		"price" REAL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, ano, km, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(car.Name, car.Ano, car.Km, car.Price)
	if err != nil {
		return err
	}
	return nil
}
