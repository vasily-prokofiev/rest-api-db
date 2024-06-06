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

	// use specific version of interfaces
	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/country/list", listCountries)
		apiV1.POST("/country/create", createCountry)
		apiV1.PUT("/country/upd", updateCountry)
		apiV1.DELETE("/country/del", deleteCountry)

		apiV1.GET("/city/list", listCities)
		apiV1.POST("/city/create", createCity)
		apiV1.PUT("/city/upd", updateCity)
		apiV1.DELETE("/city/del", deleteCity)

		apiV1.GET("/query/country_by_continent", queryCountryByContinet)
		apiV1.GET("/query/city_by_country", queryCityByCountry)
		apiV1.GET("/query/city_by_continent", queryCityByContinet)
	}
	router.Run("localhost:8080")
}
