# E-COMMERCE Restful API with Golang

## Project Overview
This project is a RESTful API for an E-Commerce platform, built using Golang. It's designed to provide a robust backend for e-commerce applications, offering a range of features from user management to product handling.

## Tech Stack
- **Golang**: For building the core application.
- **PostgreSQL**: As the database for storing all application data.
- **Docker**: For containerizing the application and ensuring consistent environments.
- **GCP CLI**: For managing cloud resources.
- **Postman**: For API testing and documentation.

## Features
- User Authentication and Authorization
- Product Management
- Order Processing
- Payment Integration
- Robust Error Handling
- Logging and Monitoring

## Database Schema
The database is designed to efficiently handle e-commerce data. Here's the schema design:
![Database Schema Design](https://i.ibb.co/Fwn0N58/bankyy-Shop.png)

## Getting Started
### Prerequisites
- Golang
- PostgreSQL
- Docker (optional)
- GCP CLI (optional)

### Installation
1. Clone the repository:
   ```
   git clone https://github.com/wiraphatys/E-Commerce-REST-API-with-Golang.git
   ```
2. Navigate to the project directory:
   ```
   cd E-Commerce-REST-API-with-Golang
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Set up your `.env` file with the necessary configurations.

### Running the Application
Run the application using:
```
go run main.go
```

## API Documentation
For detailed API endpoints and their usage, refer to the Postman documentation (link to Postman docs).

## Contributing
Contributions to the project are welcome! Please refer to the contributing guidelines for more information.

## License
This project is licensed under the [MIT License](LICENSE).
