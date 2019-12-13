# PNGR Stack
Dockerized (postgres + nginx + golang + react) starter kit

Only implements basic user signup, session management, and a toy `post` type to demonstrate basic CRUD. PNGR is _not_ a CMS.

## Features
- Hot-reload, front and back, including a test-runner for golang changes
- JSON Web-Token cookies with automatic refresh: ready for horizontal scaling
- Uses multi-stage builds for small production images
- Feature development is up to you!

## Requirements
- Install docker && docker-compose

## Quick Start
1) `docker-compose up`
2) Visit `https://localhost` (*note **https***)
3) Approve the self-signed cert
4) Make changes to either golang or react code, and enjoy hot-reload goodness!

Preview of the app:

![Screenshot of the app](docs/demo.png?raw=true "Screenshot")

## Rebuilding your dev environment
Maybe your postgres went sideways from a wonky migration and you don't want to muck with fixing it.

`docker-compose down -v && docker-compose up --build --force-recreate`

## Deploying to Production
*Warning: Run in production at your own risk - this code is not security hardened!*

You should install postgresql on the host.
Then you can use `docker-compose.prod.yml` to build lean images to use in production.
Don't forget to copy `.env.example -> .env` and setup your secrets/passwords.

--- 

## Postgres
Some tips for working with your postgres docker instance

### Creating and running migrations
Migrations are run using [go-migrate](https://github.com/golang-migrate/migrate).

I put together little bash scripts to help you get stuff done.
- `postgres/new-migration.sh my_migration_name` will create a template for the next migration
- `postgres/run-migrations.sh` will execute any new migrations (this is used when the container is created)
- `postgres/migrate.sh` makes it easy to run arbitrary [go-migrate](https://github.com/golang-migrate/migrate) commands (e.g. `postgres/migrate.sh down 1`)

You can do more advanced migrate commands 

### Opening a psql client
`docker-compose exec postgres psql -U postgres`
Remember to use `\q` to exit.

--- 

## Nginx
Nginx is simply used to route requests to the front-end and back-end based on path.
It also terminates SSL so that we don't have to deal with certs in our app layer.

--- 

## Golang
Almost-vanilla golang api:
- Makes use of go modules for dependencies
- [jwt-go](https://github.com/dgrijalva/jwt-go) for JSON Web Tokens
- [sqlx](https://github.com/jmoiron/sqlx) for better postgres interface

--- 

## React
The basic building blocks of the front-end are:
- [Create React App](https://github.com/facebookincubator/create-react-app) (unejected!)
- [React Router](https://github.com/ReactTraining/react-router)
- [Unstated](https://github.com/jamiebuilds/unstated) for state management
- [Semantic UI React](https://react.semantic-ui.com/) for component library
