# REST-API-DB

This is a test application demonstrated REST-API GoLang + Postgres DB

## Compile & Run Locally

To checkout repository please use command:

git clone https://github.com/vasily-prokofiev/rest-api-db.git


To run, from the root of the repo (rest-api-db folder) use the command:

```
go run .
```

If normally started, follow outprints shall appears:

...
go run .
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /country/list             --> main.listCountries (3 handlers)
[GIN-debug] POST   /country/create           --> main.createCountry (3 handlers)
[GIN-debug] PUT    /country/upd              --> main.updateCountry (3 handlers)
[GIN-debug] DELETE /country/del              --> main.deleteCountry (3 handlers)
[GIN-debug] GET    /city/list                --> main.listCities (3 handlers)
[GIN-debug] POST   /city/create              --> main.createCity (3 handlers)
[GIN-debug] PUT    /city/upd                 --> main.updateCity (3 handlers)
[GIN-debug] DELETE /city/del                 --> main.deleteCity (3 handlers)
[GIN-debug] GET    /query/country_by_continent --> main.queryCountryByContinet (3 handlers)
[GIN-debug] GET    /query/city_by_country    --> main.queryCityByCountry (3 handlers)
[GIN-debug] GET    /query/city_by_continent  --> main.queryCityByContinet (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:8080/v1
...

## Access the app 

Application suppose to access Postgress DB server (cofigured locally) with credential 
User/Pwd: postgres/postgres
DB: restapi
COnfiguration Script: init_db.sql <TODO>

The App has a few Endpoints

All api endpoints are prefixed with API `/v1`

To reach any endpoint use Postman application  (https://www.postman.com/downloads/)
Data submited in JSON format (Body/Raw, JSON in Postman GUI)

Follow entity implemented

Data Handling:

The Continent table is permanent, read-only, created on the DB configuration phase. 
Continent DB contains 7 entities, 

restapi=# select * from continent;
 id |    name
----+------------
  1 | Africa  
  2 | Asia
  3 | Antarctica
  4 | Australia
  5 | Europa
  6 | N.America
  7 | S.America
(7 rows)

IT DO NOT has corresponding REST API

1. Countries Entity

Contains Follow IEs:
restapi=# select * from country;
 id | continent_id |    name     | population |  area
----+--------------+-------------+------------+--------

1.1 Get a List of Countries
localhost:8080/v1/country/list
GET("/country/list")

1.2 Create a New Country within a Continent
localhost:8080/v1/country/create
Example of IE 
POST
    {
        "continent_name": "Europa",
        "name": "Finland",
        "population": 50000000,
        "area": 338462
    }

1.3 Update an Existing Country
PUT("/country/upd")

1.4 Delete an Existing Country
DELETE("/country/del")

2. Cities Entity

Contains Follow IEs:
restapi=# select * from city;
 id | country_id |   name   | population |  area  | is_capital
----+------------+----------+------------+--------+------------

2.1 Get a List of Cities
localhost:8080/v1/city/list
GET("/city/list")

2.1 Create a City 
POST("/city/create")
PUT("/city/upd")
DELETE("/city/del")



3. Data Access Entities:

3.1 GET("/query/country_by_continent")
localhost:8080/v1/query/city_by_country
GET
    {
        "continent_name": "Europa"
    }

3.2 GET("/query/city_by_country")
localhost:8080/v1/query/city_by_country
GET
    {
        "country_name": "Finland"
    }

3.3 GET("/query/city_by_continent")
localhost:8080/v1/query/city_by_country
GET
    {
        "continent_name": "Europa"
    }



## Deploy as Docker Container

Solution uses two containers (Pstgress DB + REST API application) <TO DO>

