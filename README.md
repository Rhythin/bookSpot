# 📚 Bookspot

> A sophisticated microservices-based book management platform showcasing modern backend architecture and cloud-native technologies.

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Kafka](https://img.shields.io/badge/Apache%20Kafka-231F20?style=for-the-badge&logo=apache-kafka&logoColor=white)](https://kafka.apache.org/)
[![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=for-the-badge&logo=grpc&logoColor=white)](https://grpc.io/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-326CE5?style=for-the-badge&logo=kubernetes&logoColor=white)](https://kubernetes.io/)
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://docker.com/)

## 🎯 Project Overview

Bookspot is a comprehensive microservices-based web application inspired by modern manga reader applications. This project serves as a **portfolio showcase** demonstrating expertise in:

- **Microservices Architecture** with inter-service communication
- **Event-Driven Architecture** using Apache Kafka
- **Cloud-Native Development** with Kubernetes deployment
- **Modern Backend Technologies** (Go, gRPC, PostgreSQL)
- **DevOps Practices** (Docker, K8s, Infrastructure as Code)

## 🏗️ Architecture

### Microservices Structure
```
📦 bookspot/
├── 🔐 auth-service/          # User authentication & authorization
├── 📚 books-service/         # Book and chapter management
├── 🔔 notification-service/  # Event-driven notifications
└── 🛠️ shared/               # Common libraries and utilities
    ├── kafkaclient/         # Custom Kafka client implementation
    ├── customlogger/        # Structured logging system
    ├── middlewares/         # HTTP/gRPC middlewares
    └── connection/          # Database connection management
```

### Key Components
- **JWT-based Authentication** with role-based access control
- **Event-Driven Notifications** using Kafka message queues
- **gRPC Inter-Service Communication** for high-performance data exchange
- **PostgreSQL Database** with optimized connection pooling
- **Custom Logging System** with structured logging
- **Kubernetes-Ready Deployment** with containerized services

## 🚀 Features

### Core Functionality
- ✅ **User Management** - Registration, authentication, and authorization
- ✅ **Book Management** - CRUD operations for books and metadata
- ✅ **Chapter Management** - Chapter organization and content handling
- ✅ **Reading Lists** - Personal book collections and progress tracking
- ✅ **Real-time Notifications** - Event-driven notification system
- ✅ **Centralized Logging** - Structured logging across all services

### Technical Features
- 🔄 **Event-Driven Architecture** with Kafka message streaming
- 🔒 **Security Middleware** with JWT token validation
- 📊 **Database Optimization** with connection pooling and migrations
- 🐳 **Containerization** with Docker and Kubernetes deployment
- 📈 **Scalable Design** supporting horizontal scaling

## 🛠️ Technology Stack

### Backend Technologies
- **Language**: Go (Golang) 1.21+
- **Framework**: Custom HTTP handlers with middleware pattern
- **Database**: PostgreSQL with connection pooling
- **Message Queue**: Apache Kafka with custom client
- **Communication**: gRPC for inter-service communication
- **Authentication**: JWT tokens with custom middleware

### Infrastructure & DevOps
- **Containerization**: Docker with multi-stage builds
- **Orchestration**: Kubernetes with custom manifests
- **Service Discovery**: Kubernetes native service discovery
- **Monitoring**: Custom logging with structured output
- **Development**: Docker Compose for local development

## 📋 Project Status

### ✅ Completed Features
- [x] **Microservices Architecture** - Complete service separation and organization
- [x] **Custom Kafka Client** - Full implementation with consumer groups and error handling
- [x] **PostgreSQL Integration** - Database drivers and connection management
- [x] **Custom Logger** - Structured logging system across all services
- [x] **JWT Authentication** - Token-based authentication middleware
- [x] **Event-Driven Notifications** - Kafka-based notification system
- [x] **Database Models** - Complete entity definitions and relationships
- [x] **HTTP Handlers** - RESTful API endpoints for all services
- [x] **Error Handling** - Comprehensive error handling and logging
- [x] **Docker Configuration** - Dockerfiles and docker-compose.yml

### 🔄 In Progress
- [ ] **gRPC Communication** - Inter-service communication implementation

- [ ] **Kubernetes Manifests** - K8s deployment configurations
- [ ] **API Gateway** - Centralized API routing and rate limiting

### 📅 Future Enhancements
- [ ] **Frontend Application** - React/Vue.js user interface
- [ ] **OpenTelemetry** - Distributed tracing and metrics
- [ ] **CI/CD Pipeline** - Automated testing and deployment
- [ ] **API Documentation** - OpenAPI/Swagger documentation
- [ ] **Load Testing** - Performance testing and optimization
- [ ] **Monitoring Dashboard** - Grafana/Prometheus integration

## 🏃‍♂️ Getting Started

### Prerequisites
```bash
- Go 1.21+
- PostgreSQL 14+
- Apache Kafka 2.8+
- Docker & Docker Compose
- Kubernetes cluster (optional)
```

### Local Development
```bash
# Clone the repository
git clone https://github.com/yourusername/bookspot.git
cd bookspot

# Start dependencies
docker-compose up -d postgres kafka

# Run individual services
cd services/auth-service && go run cmd/server.go
cd services/books-service && go run cmd/server.go
cd services/notification-service && go run cmd/server.go
```

## 🎨 Key Technical Highlights

### Custom Kafka Client Implementation
- **Consumer Groups** with automatic topic creation
- **Error Handling** with retry mechanisms and dead letter queues
- **Graceful Shutdown** with proper resource cleanup
- **Parallel Processing** supporting multiple listeners per service

### Microservices Communication
- **gRPC** for high-performance inter-service communication
- **Event Sourcing** with Kafka for data consistency
- **Service Discovery** using Kubernetes native features
- **Load Balancing** with automatic failover

### Database Architecture
- **Connection Pooling** for optimal resource utilization
- **Transaction Management** with proper rollback mechanisms
- **Entity Relationships** with optimized queries
- **Migration System** for schema versioning

## 🤝 Contributing

This project is primarily a portfolio showcase, but contributions and feedback are welcome! Please feel free to:

- 🐛 Report bugs or issues
- 💡 Suggest new features or improvements
- 📖 Improve documentation
- 🔧 Submit pull requests

## 📧 Contact

This project demonstrates proficiency in modern backend development, microservices architecture, and cloud-native technologies. For questions or collaboration opportunities, please reach out!

---

*This project showcases expertise in Go, microservices architecture, Kafka, PostgreSQL, Kubernetes, and modern DevOps practices.*
