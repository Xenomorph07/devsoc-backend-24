# Devsoc Backend '24

The official Backend API for DEVSOC'24 Hackathon Portal

## Features

- User Authentication
- Teams and Projects
- Admin panel where Admins can directly score teams

# How To Run

## Prerequisites:

- [Docker](https://www.docker.com): A container platform using which we can ensure standards.
- [Postman](https://www.postman.com)/[Apidog](https://apidog.com): A tool to test backend APIs without having to write frontends.
- [goose](https://github.com/pressly/goose): Goose is a database migration tool. Manage your database schema by creating incremental SQL changes or Go functions.

## Steps

1.  Clone the Repository: `git clone https://github.com/CodeChefVIT/devsoc-backend-24`
2.  Start the containers: `cd devsoc-backend-24 && cp .env.example .env && make build` Please ensure that you put the correct SMTP credentials to get email services
3.  Run the migrations: `make migrate-up`
4.  Use postman and test the api at the endpoint `http://localhost/api`

# Contributors
