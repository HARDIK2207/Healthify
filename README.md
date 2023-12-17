# Healthify

Healthify is a CRUD application built using the Go language and the Gofr framework. It allows users to track their health-related data, including name, age, and daily calories intake.

## Table of Contents

- [Technologies Used](#features)
- [Features](#features)
- [Database Structure](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Screenshots](#screenshots)
- [Testing](#testing)
- [Contributing](#contributing)

## Technologies Used

Make sure you have the following installed before running the application:

- [Go (Programming Language)](https://golang.org/dl/)
- [GoFr Framework](https://gofr.dev/)
- [MySQL Database](https://dev.mysql.com/downloads/) 
- [Git (Version Control System)](https://git-scm.com/downloads/)
## Features

- **Create:** Add new entries with user information such as name, age, and calories intake.
- **Read:** Retrieve and display user data, calculate BMI, and assess calories intake.
- **Update:** Modify existing entries to keep health information up-to-date.
- **Delete:** Remove entries for users who no longer need to be tracked.

## Database Structure

The application uses a database table with the following structure:

| Field    | Type         | Description                   |
|----------|--------------|-------------------------------|
| id       | INT          | Unique identifier             |
| name     | VARCHAR(255) | User's name                   |
| age      | INT          | User's age                    |
| calories | INT          | Daily calories intake         |


## Getting Started

1. **Clone the Repository:**
    ```bash
    git clone https://github.com/hardik2207/healthify.git
    cd healthify
    ```

2. **Run the Application:**
    ```bash
    go run route.go
    ```

5. **Testing with Postman:**
    Use Postman to test the CRUD operations. Import the provided Postman collection for quick and easy testing.

## API Endpoints

- **POST /user:** 
     1. Open Postman and set up a POST request to the "addCalories" endpoint.
     2. Provide the {healthify/name/age/calories} in the request body.
     3. Send the request.
     ![addCalories]
     (https://imgur.com/a/gYqiaGw)
- **GET /users:** Retrieve a list of all users.

- **PUT /user/{id}:** Update information for a specific user.
- **DELETE /user/{id}:** Delete a user from the database.

## Screenshots

Include screenshots here to showcase different aspects of your application.


## Testing

The project includes a Postman collection for testing API endpoints. Import the collection into Postman for quick and easy testing.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.




