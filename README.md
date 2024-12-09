# Project: Star Wars API - Backend and Frontend

This project is a full-stack application consisting of a backend built in Go and a frontend built using Vue.js 3. The application allows users to interact with Star Wars data such as planets and people, and supports sorting functionality.

-----

# Installation

### Backend (Go)

1. Ensure that you have Go installed on your machine.

You can verify by running:
```
    go version
```
2. Clone the repository:
```
    git clone https://github.com/raulisimo/bookish-guacamole.git
    cd bookish-guacamole
```
3. Install backend dependencies:
```
    cd backend
    go mod tidy
```
### Frontend (Vue 3)

1. Ensure Node.js is installed.

Verify it by running:
```
node -v
```

2. Navigate to the frontend folder:
```
    cd frontend
```
3. Install frontend dependencies:
```
    npm install
```
## Running Locally

### Backend

In the backend folder, run the backend server:

    go run cmd/main.go

The backend should be running at http://localhost:8080.

### Frontend

In the frontend folder, run the Vue.js app:

    npm run dev

- In the frontend folder, create a .env file and add VITE_API_BASE_URL=http://127.0.0.1:8080/api
- The frontend should be running at http://localhost:3000.

The frontend will automatically communicate with the backend running on http://localhost:8080.

## Building and Deploying with Docker Compose

To run both the backend and frontend using Docker Compose, follow these steps:

- Check that There is a Dockerfile in the backend directory.
- Check that there is a Dockerfile in the frontend directory.

At the root of the project, there is a compose.yaml file to manage the backend and frontend containers.

## Building and Running the Containers

From the root of your project, run the following command to build and start the containers:

    docker-compose up --build

This will:
- Build and run the backend Go API on port 8080
- Build and run the frontend Vue app on port 6969

You can access:
- Backend: http://localhost:8080
- Frontend: http://localhost/6969

Stopping the Containers

To stop the containers, run:

    docker-compose down

# API Endpoints

Here are the available endpoints for the backend API:
GET /api/planets

Fetches all planets from the Star Wars API, supports sorting.
GET /api/people

Fetches all people from the Star Wars API, supports sorting.

Both endpoints accept the following optional query parameters for sorting:

    sort: The field to sort by (only name and created).
    order: The order in which to sort (asc or desc).

### Backend Features

Supports a simple cache of the results from planets and people that last for 15 minutes.

### Frontend Features

The frontend is a simple Vue 3 app that interacts with the backend API. It includes:

- A table view of the planets and people from the Star Wars API. 
- Sorting options for both planets and people.
- A basic UI to display the data retrieved from the backend.
- A search bar to filter planets and people.
