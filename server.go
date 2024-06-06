package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/restapi?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/country/list", listCountries)
	router.POST("/country/create", createCountry)
	router.PUT("/country/upd", updateCountry)
	router.DELETE("/country/del", deleteCountry)

	router.GET("/city/list", listCities)
	router.POST("/city/create", createCity)
	router.PUT("/city/upd", updateCity)
	router.DELETE("/city/del", deleteCity)

	router.GET("/query/country_by_continent", queryCountryByContinet)
	router.GET("/query/city_by_country", queryCityByCountry)
	router.GET("/query/city_by_continent", queryCityByContinet)

	router.Run("localhost:8080")
}
