<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-3.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center" style="text-align: center;">
  Challenge #4. Persisting the data
</h1>

In this 4th challenge we are going to persist our data into PostgreSQL.

## Instructions


* Create a new service that will receive the Ad created in challenge1 and will persist it.
* The easiest way for local development is to use a docker image for PostgreSQL.
* Since we already went thorough error handling and testing we can handle all possible scenarios
* You can use as orm : https://gorm.io/docs/index.html 


## Resources
1. Connecting to database with go package: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
2. Saving structures in database: https://raw.githubusercontent.com/richyen/toolbox/master/pg/orm/gorm_demo/gorm_demo.go
