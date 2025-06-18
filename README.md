# Bookspot

Bookspot is a web application inspired by my favourite manga reader application. It is a work in progress, meant to showcase my skills in technologies such as Docker, Kubernetes, Golang, gRPC, Kafka and PostgreSQL. 

The project is designed to be deployed in a Kubernetes cluster, with a PostgreSQL database, multiple backend services, and a simple frontend application. The backend services will communicate with each other using gRPC. and kafka for event based notifications.

## Features

- User authentication and authorization with JWT Tokens
- Book management (CRUD operations)
- Chapter management (CRUD operations)
- Reading list management (CRUD operations)
- Notification system (using Kafka for event based notifications)
- Logging system (using custom logger)
- A simple frontend application

## Technologies Used

- Golang
- PostgreSQL
- gRPC
- Kubernetes
- Kafka
- Docker

## Tasks In Progress

- Implementing user based authentication and authorization
- Implementing a simple custom Kafka client
- Implementing gRPC communication between backend services
- Defining dockerfiles and docker-compose.yml for the application 
- Defining the required resources and setup for k8s deployment 

## Completed Tasks

- Setting up PostgreSQL database drivers
- Implementing a custom logger

## Future Plans

- Implementing the frontend application
- OpenTelemetry for tracing and metrics
