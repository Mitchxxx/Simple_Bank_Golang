# Simple Bank

This repository contains the code for a Backend Master class written in Golang.

<center>
<img src=./backend-master.png alt="backend poster">
</center>

## Overview

A simple banking system API built with Go, featuring account management, money transfers, and user authentication. This project demonstrates best practices in backend development including database transactions, API design, authentication, and testing.

## Features

- **User Management**: Create users with secure password hashing
- **Account Management**: Create and manage bank accounts with different currencies
- **Money Transfers**: Transfer money between accounts with proper transaction handling
- **Authentication**: JWT and PASETO token-based authentication
- **Session Management**: Refresh token implementation for secure sessions
- **Authorization**: Middleware-based route protection
- **Database Migrations**: Version-controlled database schema changes
- **Testing**: Comprehensive unit and integration tests

## Tech Stack

- **Language**: Go 1.23.1
- **Database**: PostgreSQL 17
- **Framework**: Gin (HTTP web framework)
- **ORM**: SQLC (SQL compiler)
- **Authentication**: JWT & PASETO tokens
- **Testing**: Testify, Gomock
- **Migrations**: golang-migrate
- **Containerization**: Docker & Docker Compose
- **Deployment**: AWS EKS

## Database Schema

The application uses PostgreSQL with the following main tables:
- **users**: Store user credentials and profile information
- **accounts**: Bank accounts with owner, balance, and currency
- **entries**: Account transaction entries
- **transfers**: Money transfer records between accounts
- **sessions**: User session and refresh token management

## Prerequisites

- Go 1.23.1 or higher
- Docker and Docker Compose
- PostgreSQL 17
- golang-migrate CLI
- SQLC

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/Mitchxxx/Simple_Bank_Golang.git
cd Simple_Bank_Golang
```

### 2. Start PostgreSQL with Docker

```bash
make postgres
```

### 3. Create the database

```bash
make createdb
```

### 4. Run database migrations

```bash
make migrateup
```

### 5. Run the server

```bash
make server
```

The server will start on `http://localhost:8080` (or the address specified in your `app.env` file).

## Configuration

Create an `app.env` file in the root directory with the following variables:

```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable
SERVER_ADDRESS=0.0.0.0:8080
TOKEN_SYMMETRIC_KEY=your-32-character-secret-key
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=24h
```

## Available Make Commands

- `make postgres` - Start PostgreSQL container
- `make createdb` - Create database
- `make dropdb` - Drop database
- `make migrateup` - Run all migrations
- `make migrateup1` - Run one migration up
- `make migratedown` - Rollback all migrations
- `make migratedown1` - Rollback one migration
- `make sqlc` - Generate SQL code
- `make test` - Run tests
- `make server` - Run the server
- `make mock` - Generate mock database

## API Endpoints

### Public Endpoints

- `POST /users` - Create a new user
- `POST /users/login` - User login
- `POST /tokens/renew_access` - Renew access token

### Protected Endpoints (require authentication)

- `POST /accounts` - Create a new account
- `GET /accounts/:id` - Get account by ID
- `GET /accounts` - List user accounts (paginated)
- `POST /transfers` - Create a money transfer

## Running Tests

```bash
make test
```

This will run all unit and integration tests with coverage reporting.

## Docker Deployment

### Build and run with Docker Compose

```bash
docker-compose up
```

### Build Docker image

```bash
docker build -t simplebank:latest .
```

## Project Structure

```
.
├── api/              # HTTP API handlers and routing
├── db/
│   ├── migration/    # Database migration files
│   ├── mock/         # Mock database for testing
│   └── sqlc/         # Generated SQL code
├── doc/              # Database documentation
├── token/            # Token generation and validation
├── util/             # Utility functions and configuration
├── main.go           # Application entry point
├── Makefile          # Build and task automation
└── README.md         # This file
```

## License

This project is part of a Backend Master Class course.