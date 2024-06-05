package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type continent struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type country struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	ContinentId string  `json:"continent_id"`
	Population  float64 `json:population`
	Area        float64 `json:area`
}

type country_create struct {
	ContinentName string  `json:"continent_name"`
	Id            string  `json:"id"`
	Name          string  `json:"name"`
	Population    float64 `json:population`
	Area          float64 `json:area`
}

type city struct {
	CountryId  string  `json:"country_id"`
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Population float64 `json:"population"`
	Area       float64 `json:"area"`
	IsCapital  bool    `json:"is_capital"`
}

type city_create struct {
	CountryName string  `json:"country_name"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Population  float64 `json:"population"`
	Area        float64 `json:"area"`
	IsCapital   bool    `json:"is_capital"`
}

func listCountries(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name, population, area, continent_id FROM country")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var countries []country
	for rows.Next() {
		var a country
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.ContinentId)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, countries)
}

func createCountry(c *gin.Context) {

	var newCountry country_create
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("INSERT INTO country (continent_id, name, population, area) VALUES ((SELECT id from continent WHERE name = $1), $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCountry.ContinentName, newCountry.Name, newCountry.Population, newCountry.Area); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newCountry)
}

func updateCountry(c *gin.Context) {

	var newCountry country
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("UPDATE country SET name = $2, population = $3, area = $4 WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCountry.Id, newCountry.Name, newCountry.Population, newCountry.Area); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newCountry)
}

func deleteCountry(c *gin.Context) {

	var newCountry country
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("DELETE FROM country WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCountry.Id); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newCountry)
}

func listCities(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, name, population, area, is_capital FROM city")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cities []city
	for rows.Next() {
		var a city
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.IsCapital)
		if err != nil {
			log.Fatal(err)
		}
		cities = append(cities, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}

func createCity(c *gin.Context) {

	var newCity city_create
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("INSERT INTO city (country_id, name, population, area) VALUES ((SELECT id from continent WHERE name = $1), $1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCity.CountryName, newCity.Name, newCity.Population, newCity.Area, newCity.IsCapital); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newCity)
}

func updateCity(c *gin.Context) {

	var newCity city
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("UPDATE city SET name = $2, population = $3, area = $4, is_capital = $5 WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCity.Id, newCity.Name, newCity.Population, newCity.Area, newCity.IsCapital); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newCity)
}

func deleteCity(c *gin.Context) {

	var newCity city
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	stmt, err := db.Prepare("DELETE FROM city WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newCity.Id); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newCity)
}

func queryCountryByContinet(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var newCountry country_create
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := db.Query("SELECT id, name, population, area FROM country WHERE continent_id IN (SELECT id FROM continent WHERE name = $1) ORDER BY name", newCountry.ContinentName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var countries []country_create
	for rows.Next() {
		var a country_create
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area)
		if err != nil {
			log.Fatal(err)
		}
		countries = append(countries, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, countries)
}

func queryCityByContinet(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var newCountry country_create
	if err := c.BindJSON(&newCountry); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := db.Query("SELECT id, name, population, area, is_capital FROM city WHERE country_id IN (SELECT id FROM country WHERE continent_id = (SELECT id FROM continent WHERE name = $1) ) ORDER BY name", newCountry.ContinentName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cities []city
	for rows.Next() {
		var a city
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.IsCapital)
		if err != nil {
			log.Fatal(err)
		}
		cities = append(cities, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}

func queryCityByCountry(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var newCity city_create
	if err := c.BindJSON(&newCity); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := db.Query("SELECT id, name, population, area, is_capital FROM city WHERE country_id IN (SELECT id FROM country WHERE name = $1) ORDER BY name", newCity.CountryName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var cities []city
	for rows.Next() {
		var a city
		err := rows.Scan(&a.Id, &a.Name, &a.Population, &a.Area, &a.IsCapital)
		if err != nil {
			log.Fatal(err)
		}
		cities = append(cities, a)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, cities)
}
