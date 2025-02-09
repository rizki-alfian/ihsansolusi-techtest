# ihsansolusi-techtest
Technical Test of Ihsan Solusi

# Project Setup Guide

## Prerequisites

Before you begin, ensure you have the following installed and set up:

1. **Docker**  
   - Please install Docker on your machine if you haven't already. You can download it from [Docker's official website](https://www.docker.com/).


## Project Setup

### Step 1: Configure the `service` Settings

- Open the `docker-compose.yml` file and modify the `container_name` and `environment` values as per folder services.

### Step 2: Copy Environment Files

- For following folders, copy the `.env.example` file to `.env`:
  - `ihsansolusi-account`


### Step 3: Update Environment Variables

In each `.env.local` file, update the values for the environment variables according to the following instructions:

- **DB_HOST**: Set this based on `container_name` found in `docker-compose.yml`.
- **DB_PORT**: Set this based on `ports` found in `docker-compose.yml`.
- **DB_USER**: Set this to the value of `POSTGRES_USER` from the Postgres service environment.
- **DB_PASS**: Set this to the value of `POSTGRES_PASSWORD` from the Postgres service environment.
- **DB_NAME**: Set this to the value of `POSTGRES_DB` from the Postgres service environment.


## Running the Project

Once you've set up everything, run the following command to start the project:

```bash
docker-compose up -d
```


## RUN MIGRATIONS
Once docker container is ready, you need to run migrations

### Step 1: Get in to ihsansolusi-account container
use this following command 
```
docker exec -it ihsansolusi-account sh
```

### Step 2: Run the migrations
Once you in ihsansolusi-account container use this following command to run the migrations (without brackets)
```
migrate -database "postgres://{{your_db_user}}:{{your_db_pass}}@{{you_container_name}}:{{your_db_port}}/{{your_db_name}}?sslmode=disable" -path migrations up
```



Here the Postman Documentation:
https://documenter.getpostman.com/view/40407342/2sAYJ6BeYw