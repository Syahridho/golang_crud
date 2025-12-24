# Go Web CRUD Application

A simple web application for managing employee data using Go, MySQL, and HTML templates.

## Setup Instructions

### 1. Database Setup

Make sure you have MySQL installed and running on localhost:3306 with a root user (no password).

Run the following SQL script to create the database and table:

```bash
mysql -u root < database/schema.sql
```

Or manually execute the SQL commands in `database/schema.sql`.

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Application

```bash
go run main.go
```

The application will be available at http://localhost:8080

## Endpoints

- `/` - Hello World message
- `/employee` - List all employees
- `/employee/create` - Create a new employee (GET for form, POST for submission)

## Troubleshooting

If data is not displaying on `/employee`:

1. Make sure MySQL is running
2. Verify the database and table exist using the schema.sql script
3. Check the console output for debug messages
4. Ensure you're accessing `/employee` (not `/employees`)
