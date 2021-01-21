package singleton

import (
	"fmt"
	"sync"
)

type Database interface {
	getPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) getPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once init() -- thread safety
// laziness

var once sync.Once
var instance *singletonDatabase

func getSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		fmt.Println("Creating instance..")
		data := make(map[string]int)
		data["New Deli"] = 100000
		data["Moscow"] = 100
		db := singletonDatabase{data}
		instance = &db
	})

	return instance
}

func getTotalPopulation(db Database, cities []string) int {
	result := 0

	for _, city := range cities {
		result += db.getPopulation(city)
	}

	return result
}

func Run() {
	db := getSingletonDatabase()
	pop := db.getPopulation("New Deli")
	fmt.Println("Pop of New Deli = ", pop)
	db = getSingletonDatabase()
	pop = db.getPopulation("Moscow")
	fmt.Println("Pop of Moscow = ", pop)
	cities := []string{"New Deli", "Moscow"}
	tp := getTotalPopulation(db, cities)
	fmt.Println("Total population: ", tp)
}
