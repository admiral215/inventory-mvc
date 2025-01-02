# Inventory MVC Application

This is a simple Inventory Management application built using Beego framework, GORM for ORM, and PostgreSQL as the database. The application follows the MVC (Model-View-Controller) architecture.

## Technologies Used
- **Beego**: A powerful web application framework for Go (Golang).
- **GORM**: A Golang ORM library to interact with the PostgreSQL database.
- **PostgreSQL**: A relational database management system.

## Features
- User authentication (on going)
- CRUD operations for inventory items
- Manage categories for inventory items
- View inventory report (on going)

## Setup and Installation

### Prerequisites
- **Go** (version 1.22 or higher)
- **PostgreSQL** (version 13 or higher)
- **Beego** (use `go get github.com/astaxie/beego` to install)


## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/inventory-mvc.git
   cd inventory-mvc
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the database:
   - Make sure you have PostgreSQL running and create a database (or use your existing one).
   - Update the `DATABASE_URL` in the `.env` file with your database credentials.

4. Run the application:
   ```bash
   go run main.go
   ```

   The app will start running on `http://localhost:8001`.
