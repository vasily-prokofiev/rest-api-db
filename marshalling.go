package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// the Continent table marhalling (fetching)
// Id is a Primary Key, Name is unique
type continent struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// the Country table marhalling (fetching)
// Id is a Primary key, Name is unique
type country struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ContinentId string `json:"continent_id"`
	Population  int64  `json:"population"`
	Area        int64  `json:"area"`
}

// the Country table marhalling (creating new Country)
// The ContinentName use to Lookup for ContinentId
type country_create struct {
	ContinentName string `json:"continent_name"`
	Id            string `json:"id"`
	Name          string `json:"name"`
	Population    int64  `json:"population"`
	Area          int64  `json:"area"`
}

// the City table marhalling (fetching)
// Id is a Primary key, Name is unique
type city struct {
	CountryId  string  `json:"country_id"`
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Population float64 `json:"population"`
	Area       float64 `json:"area"`
	IsCapital  bool    `json:"is_capital"`
}

// the City table marhalling (creating new City)
// The CountrytName use to Lookup for CountryId
type city_create struct {
	CountryName string  `json:"country_name"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Population  float64 `json:"population"`
	Area        float64 `json:"area"`
	IsCapital   bool    `json:"is_capital"`
}

// list all Countries from DB, ordered by Name
func listCountries(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name, population, area, continent_id FROM country ORDER BY name")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var countries []country
	for rows.Next() {
		var a country
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.ContinentId)
		if err != nil {
			log.Print(err)
		}
		countries = append(countries, a)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.IndentedJSON(http.StatusOK, countries)
}

// create new Country, Continent Name (unique string) uses as lookup entry to bind with Continent
func createCountry(c *gin.Context) {

	var newCountry country_create
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("INSERT INTO country (continent_id, name, population, area) VALUES ((SELECT id from continent WHERE name = $1), $2, $3, $4)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCountry.ContinentName, newCountry.Name, newCountry.Population, newCountry.Area); err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusCreated, newCountry)
}

// update exiting Country, Id (unique integer) uses as the Country key,  Continent relationship cannot be updated
func updateCountry(c *gin.Context) {

	var newCountry country
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("UPDATE country SET name = $2, population = $3, area = $4 WHERE id = $1")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCountry.Id, newCountry.Name, newCountry.Population, newCountry.Area); err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusCreated, newCountry)
}

// delete exiting Country, Id (unique integer) uses as the Country key
func deleteCountry(c *gin.Context) {

	var newCountry country
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("DELETE FROM country WHERE id = $1")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCountry.Id); err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusCreated, newCountry)
}

// list all Cities from DB, ordered by Name
func listCities(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, country_id, name, population, area, is_capital FROM city ORDER BY name")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var cities []city
	for rows.Next() {
		var a city
		err := rows.Scan(&a.Id, &a.CountryId, &a.Name, &a.Population, &a.Area, &a.IsCapital)
		if err != nil {
			log.Print(err)
		}
		cities = append(cities, a)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}

// create new City, Country Name (unique string) uses as lookup entry to bind with Countries
func createCity(c *gin.Context) {

	var newCity city_create
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("INSERT INTO city (country_id, name, population, area) VALUES ((SELECT id from continent WHERE name = $1), $1, $2, $3)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCity.CountryName, newCity.Name, newCity.Population, newCity.Area, newCity.IsCapital); err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusCreated, newCity)
}

// update exiting City, Id (unique integer) uses as the Country key, Country relationship cannot be updated
func updateCity(c *gin.Context) {

	var newCity city
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("UPDATE city SET name = $2, population = $3, area = $4, is_capital = $5 WHERE id = $1")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCity.Id, newCity.Name, newCity.Population, newCity.Area, newCity.IsCapital); err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusCreated, newCity)
}

// delete exiting City, Id (unique integer) uses as the City key
func deleteCity(c *gin.Context) {

	var newCity city
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("DELETE FROM city WHERE id = $1")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCity.Id); err != nil {
		log.Print(err)
	}

	c.JSON(http.StatusCreated, newCity)
}

// fech Countries by Continent Name, exposing all Countries belonging to Continent, ordered by Name
func queryCountryByContinet(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var newCountry country_create
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := db.Query("SELECT id, name, population, area FROM country WHERE continent_id IN (SELECT id FROM continent WHERE name = $1) ORDER BY name", newCountry.ContinentName)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var countries []country_create
	for rows.Next() {
		var a country_create
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area)
		if err != nil {
			log.Print(err)
		}
		countries = append(countries, a)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.IndentedJSON(http.StatusOK, countries)
}

// fech Cities by Continent Name, exposing all Cities belonging to Continent, ordered by Name
func queryCityByContinet(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var newCountry country_create
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := db.Query("SELECT id, name, population, area, is_capital FROM city WHERE country_id IN (SELECT id FROM country WHERE continent_id = (SELECT id FROM continent WHERE name = $1) ) ORDER BY name", newCountry.ContinentName)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var cities []city
	for rows.Next() {
		var a city
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.IsCapital)
		if err != nil {
			log.Print(err)
		}
		cities = append(cities, a)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}

// fech Cities by Country Name, exposing all Cities belonging to Country, ordered by Name
func queryCityByCountry(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var newCity city_create
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := db.Query("SELECT id, name, population, area, is_capital FROM city WHERE country_id IN (SELECT id FROM country WHERE name = $1) ORDER BY name", newCity.CountryName)
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	var cities []city
	for rows.Next() {
		var a city
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.IsCapital)
		if err != nil {
			log.Print(err)
		}
		cities = append(cities, a)
	}
	err = rows.Err()
	if err != nil {
		log.Print(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}
