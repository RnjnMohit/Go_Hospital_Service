Hospital Service API
Overview
The Hospital Service API is a Golang-based backend application designed to manage hospital services. This API provides endpoints to handle patient information, medical staff details, appointments, and other hospital-related data. It is built to be scalable, maintainable, and easy to integrate with frontend applications or other services.

Prerequisites
Before you begin, ensure you have met the following requirements:

Golang (>= 1.16)
MongoDB (or any other preferred database)

Clone the repository:
git clone https://github.com/yourusername/hospital-service-api.git

Configuration
Create a .env file in the root directory of the project.
Add the following environment variables
PORT=3000
MONGODB_URI=mongodb://localhost:27017/hospitaldb
JWT_SECRET=your_jwt_secret_key

Running the Server
go run main.go
