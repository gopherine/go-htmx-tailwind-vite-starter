# Golang HTMX Tailwind Alpine Starter

## Overview

This project is a starter template for a web application built with Go. It includes Echo for web routing, Tailwind CSS for styling, Templ for HTML templating, SQLC for SQL code generation, and Golang-Migrate for database migrations. It also uses Air for live reloading in development.

## Project Structure

```bash
.
├── Makefile                # Makefile to manage build tasks and development commands
├── README.md               # Project documentation
├── air.toml                # Configuration file for Air live reloading
├── bun.lockb               # Bun lock file for dependency management (if using Bun)
├── go.mod                  # Go module file, listing project dependencies
├── go.sum                  # Go checksum file, ensuring the integrity of dependencies
├── handlers
│   ├── home.go             # HTTP handler for the home route
│   └── shared.go           # Shared HTTP handlers or middleware
├── infrastructure
│   └── postgres
│       └── migrations      # Database migration files for Golang-Migrate
├── internals               # Internal application logic and domain-specific code
├── main.go                 # Main application entry point
├── package.json            # Node.js package file for managing JavaScript dependencies
├── public
│   └── global.css          # Global CSS file for styling
├── sqlc.yaml               # Configuration file for SQLC code generation
├── static_dev.go           # Development-specific static file handling
├── static_prod.go          # Production-specific static file handling
├── tailwind.config.js      # Configuration file for Tailwind CSS
├── tmp
│   ├── build-errors.log    # Log file for build errors
│   └── main                # Temporary build artifacts directory
├── tools.go                # Go file to track development tools as dependencies
└── views
    ├── components
    │   ├── navigation.templ     # Templ template for navigation component
    │   └── navigation_templ.go  # Go file for navigation component logic
    ├── css
    │   └── global.css           # CSS file for styling the application
    ├── home
    │   ├── index.templ          # Templ template for the home page
    │   └── index_templ.go       # Go file for home page logic
    └── layouts
        ├── base.templ           # Templ template for the base layout
        └── base_templ.go        # Go file for base layout logic
```

## Getting Started

### Prerequisites

    •	Go 1.16 or later
    •	Node.js and npm
    •	Air (for live reloading in development)
    •	Templ (for HTML templating)
    •	Tailwind CSS (for styling)
    •	SQLC (for SQL code generation)
    •	Golang-Migrate (for database migrations)

### Installation

1. Clone the repository

```bash
git clone https://github.com/yourusername/yourprojectname.git
cd yourprojectname
```

2. Install go dependencies

```bash
go mod tidy
go install
# Note: if go install does not install every thing from tools.go you will have to install templ, sqlc and golang-migrate manually
```

```bash
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

To get started clone this template and run

```bash
yarn install
# or
npm install
# or
bun install
```

### Configuration

1. Create a .env file for your environment variables.

```bash
cp .env.example .env
```

2. Update the .env file with your configuration settings.

### Usage

Development

1. Run the application with Air for live reloading

```bash
make air
```

This will run the application in development mode with live reloading using Air, Templ, and Tailwind CSS.

2. If you only generate templates with Templ and Proxy

```bash
make templ
```

3. Watch and compile CSS with Tailwind CSS

```bash
make css
```

### Production

1. Build the application for production

```bash
make build
```

2. Run the production build

```bash
make run
```

### Running Migrations

First you need to write sql native queries in query folder inside `infrastructure/postgres/query`. To run database migrations, use Golang-Migrate with the following command:

```bash
migrate -database postgres://user:password@localhost:5432/dbname?sslmode=disable -path ./infrastructure/postgres/migrations up
```

you will need to compile css for the first time to include it in the binary

```
make css
```

# Migrations

For migrations make sure you have golang-migrate installed

```
migrate -path repo/migrations -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" -verbose up
```

# To start development server

You can access client at port :7331 and backend at :1323

```
make air
```

# To run prod

```
make run
```

credits:

The base structure is inspired by Anthony's video so big shoutout, please do visit his channel https://www.youtube.com/watch?v=1dqp1s72Z8E
