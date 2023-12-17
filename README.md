# Healthify

Healthify is a CRUD application built using the Go language and the Gofr framework. It allows users to track their health-related data, including name, age, and daily calories intake.

## Table of Contents

- [Technologies Used](#features)
- [Features](#features)
- [Diagrams](#diagrams)
- [Database Structure](#features)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
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

## Diagrams

**Sequence Diagram** ![Screenshot (350)](https://github.com/HARDIK2207/Healthify/assets/84044856/84d46d16-50ef-4897-8e2d-a0a876956010)
**Use Case Diagram** ![Screenshot (353)](https://github.com/HARDIK2207/Healthify/assets/84044856/2637ac0b-768a-4544-b482-3461732ce885)


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
    ![Screenshot (339)](https://github.com/HARDIK2207/Healthify/assets/84044856/a1d5ceb7-7429-4eb8-a311-50ce7aa82851)
    ![Screenshot (340)](https://github.com/HARDIK2207/Healthify/assets/84044856/c3a1bf31-6c4c-4547-98ef-44267d4ca1af)

- **GET /users:**
    1. Set up a GET request to your "/healthify" endpoint.
    ![Screenshot (356)](https://github.com/HARDIK2207/Healthify/assets/84044856/3819d937-cc7b-4cc0-98a8-a3c833d2fa2c)

- **PUT /user/{id}:**
    1. Set up a GET request to your "/healthify/id" endpoint.
    2. Enter Json details new name or calories that is to be edited.
    ![Screenshot (347)](https://github.com/HARDIK2207/Healthify/assets/84044856/f756c782-5432-42b1-abae-837e5a711021)
    ![Screenshot (342)](https://github.com/HARDIK2207/Healthify/assets/84044856/8aa3ac4a-9f45-4035-afec-3f0afbccc22c)

- **DELETE /user/{id}:**
    1. Delete a user from the database.
    ![Screenshot (343)](https://github.com/HARDIK2207/Healthify/assets/84044856/51b48bd3-c22b-4a85-bb5e-2deb21ed1246)
    ![Screenshot (344)](https://github.com/HARDIK2207/Healthify/assets/84044856/103a0c1c-2e7c-4b3e-ba06-bcb50489ceb1)

## Testing

The project includes a Postman collection for testing API endpoints.

 - **GET**
    ![Screenshot (355)](https://github.com/HARDIK2207/Healthify/assets/84044856/6fab24e0-c4af-4040-ae06-810b428554c5)
- **POST**
    ![Screenshot (357)](https://github.com/HARDIK2207/Healthify/assets/84044856/279e9406-7ae5-4f4a-b27f-08e74335f6b4)
- **PUT**
    ![Screenshot (349)](https://github.com/HARDIK2207/Healthify/assets/84044856/5a8c3d91-15f2-4188-be8a-c7f11cfa62d8)
- **DELETE**
    ![Screenshot (358)](https://github.com/HARDIK2207/Healthify/assets/84044856/9a4b7994-035f-4b0a-b6da-212233528e5b)

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to improve the project.




