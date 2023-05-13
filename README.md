# Bookstore REST API
Welcome to the Bookstore REST API! This project was created to enhance my Golang skills by building a RESTful API that interacts with a MySQL database. The API provides endpoints to perform CRUD operations on books.

# Description
The Bookstore REST API is a Golang application that uses Gorilla mux router to provide RESTful endpoints for performing CRUD operations on a MySQL database. The API supports the following endpoints:

GET /books: Retrieves all books in the database.
GET /books/:id: Retrieves a specific book by ID.
POST /books: Creates a new book in the database.
PUT /books/:id: Updates an existing book by ID.
DELETE /books/:id: Deletes a book from the database by ID.
The Bookstore REST API includes error handling for invalid input and database errors. It also includes validation for book fields to ensure that they meet certain requirements.
