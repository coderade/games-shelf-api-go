# games-shelf-api-go

Games Shelf REST API developed in GO.

This application is created using Golang, and you have the possibilities to:

- Login using JWT tokens (the user and password need to be managed on the backend project)
- CRUD of the games, platforms and genres
- Connect with a PostgreSQL Database
- Use GraphQL to search games
- Work with data received from RAWG api, the biggest video game database
  - Data as the games images and rating will be got from this API.

Some technologies used are:

- Go 1.18.1
- PostgreSQL
- JWT
- Crypto
- GraphQL

Client application developed in React available here:
[games-shelf-client-react](https://github.com/coderade/games-shelf-client-react).

## Running the Application 

#### Environment variables

The following environment variables are required to run this application 

- `PORT`:  Server port to list on (default: **4000**)
- `ENV`: Application environment (development|production) (default: **development**) -
- `DB_DATA_SOURCE`: Database data source  (default: **postgres://admin@localhost/games_shelf?sslmode=disable**)
- `APP_SECRET`: Application secret, used to crete JWT token (default: **games-shelf-api-secret**)
- `RAWG_API_KEY`: API key used to connect on the RAWG API. To get an API Key check [here](https://rawg.io/apidocs).

#### Database
To create the database, roles and grant the needed privileges run the following SQL commands

    CREATE DATABASE games_shelf
    CREATE ROLE admin;
    ALTER ROLE admin login;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to admin;
    GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public to admin;
    GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public to admin;

To create all the tables, sequences, constraints and content needed for the application we need to run the following
command:

     psql -h localhost -d games_shelf -U admin -p 5432 -a -q -f db/games_shelf.sql

### Running the application

After we set the required environment variables and create the database you can run the application 
in the following ways:

##### Build and execute via `./main`
    go build cmd/*.go
    ./main

##### Run with `go run`
     go run cmd/*.go



