Golang CRUD Application with MySQL and Redis
This project is a CRUD API in Golang, using Gin for routing, MySQL as the primary data store, and Redis for caching. The application imports data from an Excel file, performs CRUD operations, and caches records in Redis for faster access.

Prerequisites
Docker - Used for running MySQL and Redis containers.
Go - To build and run the application.
cURL or Postman - To test the API.
Getting Started
Step 1: Clone the Repository
Clone the repository to your local environment:

bash
Copy code
git clone <repository-url>
cd <repository-folder>
Step 2: Set Up Environment Variables
Create a .env file in the root directory with the following contents:

plaintext
Copy code
DB_USER="root"
DB_PASSWORD="root"
DB_HOST="127.0.0.1:3306"
DB_NAME="testdb"
REDIS_ADDR="localhost:6379"
REDIS_PASSWORD=""
Step 3: Start MySQL and Redis Using Docker
Run the following Docker commands to start MySQL and Redis containers:

bash
Copy code
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:latest
docker run --name redis-container -p 6379:6379 -d redis:latest
If the containers are stopped, you can start them again with:

bash
Copy code
docker start mysql-container
docker start redis-container
Step 4: Install Dependencies
Install the necessary dependencies with:

bash
Copy code
go mod download
Step 5: Build and Run the Application
Build and run the Go application with:

bash
Copy code
go run main.go
Step 6: Test API Endpoints
You can use the following cURL commands to perform CRUD operations. Ensure the application is running on http://localhost:8080.

Import Data (POST)
To import data from an Excel file:

bash
Copy code
curl -X POST "http://localhost:8080/records" -F "id=123" -F "file=@Sample_Employee_data_xlsx.xlsx"
Get All Records (GET)
Retrieve all records:

bash
Copy code
curl -X GET "http://localhost:8080/records"
Update a Record (PUT)
Update a record by ID:

bash
Copy code
curl -X PUT http://localhost:8080/records -H "Content-Type: application/json" -d '{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "company_name": "Updated Co.",
  "address": "123 New St.",
  "city": "New City",
  "county": "New County",
  "postal": "12345",
  "phone": "123-456-7890",
  "email": "john.doe@example.com",
  "web": "https://newwebsite.com"
}'
Delete a Record (DELETE)
Delete a record by ID:

bash
Copy code
curl -X DELETE "http://localhost:8080/records/1"
Additional Notes
Database Initialization: When the application runs, it will automatically create a table named Employee if it doesn’t exist.
Error Handling: Ensure MySQL and Redis containers are running before starting the application to avoid connection errors.
