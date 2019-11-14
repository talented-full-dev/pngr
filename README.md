# PNGR
dockerized (postgres + nginx + golang + react) starter kit

I've only implemented basic user signup, session management, and a toy "post" type in terms of actual features, the idea is that this scaffolding can be forked and extended to serve a huge variety of purposes.

Feel free to create issues with suggestions, or pull requests for security or developer ergonomics improvements.

## Requirements
- Install docker && docker-compose

## Quick Start
1) `sudo docker-compose up`
2) Visit https://localhost (and approve the self-signed cert)
3) Make changes to either golang or react code, and watch the app rebuild itself!

## Production Builds
*Warning: this code is pre-alpha quality! Run in production at your own risk*

- Generate a production container with `sudo docker build .` 
- You will need to setup nginx in production for SSL termination and port forwarding to `:5000` look at `nginx/nginx.prod.conf` for ideas on how to do this.

--- 

## Postgres
Some tips for working with your postgres docker instance

### Creating and running migrations
Migrations are run using [go-migrate](https://github.com/golang-migrate/migrate).

I put together little bash scripts to help you get stuff done.
- `sudo postgres/new-migration.sh my_migration_name` will create a template for the next migration-
- `sudo postgres/run-migrations.sh` will execute any new migrations 

You can do more advanced migrate commands 

### Opening a psql client
`sudo docker exec -it pngr_postgres_1 psql --username postgres --dbname postgres`
Remember to use `\q` to exit.

### Rebuilding your database from scratch
If you want to clear out your postgres instance and start fresh, due to a bad migration or some other issue, normal container recreation isn't enough, because docker compose creates a volume for postgres data.

To *fully* reset your postgres instance, run:
`sudo docker-compose down -v && sudo docker-compose up --build --force-recreate`

--- 

## Nginx
Nginx is simply used to route requests to the front-end and back-end based on path.
It also terminates SSL so that we don't have to deal with certs in our app layer.

--- 

## Golang
Almost-vanilla golang api:
- Makes use of go modules for dependencies(!)
- [jwt-gp](github.com/dgrijalva/jwt-go) for JSON Web Tokens
- [sqlx](https://github.com/jmoiron/sqlx) for cleaner interactions with postgres

--- 

## React
The basic building blocks of the front-end are:
- [Create React App](https://github.com/facebookincubator/create-react-app) (unejected!)
- [React Router](https://github.com/ReactTraining/react-router)
- [Unstated](https://github.com/jamiebuilds/unstated)
- [Semantic UI React](https://react.semantic-ui.com/)
