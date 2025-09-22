🚗 Ride-Hail System
A distributed ride-hailing platform built with Go, implementing Service-Oriented Architecture (SOA) principles. Features real-time driver matching, location tracking, and seamless passenger-driver coordination through microservices communicating via RabbitMQ message queues and WebSocket connections.

✨ Features
Core Functionality

🚕 Real-time Ride Matching: Advanced geospatial algorithms match passengers with nearby drivers
📍 Live Location Tracking: Real-time GPS tracking with sub-second updates
💰 Dynamic Pricing: Vehicle type-based fare calculation (Economy/Premium/XL)
⚡ WebSocket Communication: Bi-directional real-time updates for passengers and drivers
📊 Admin Dashboard: Comprehensive system monitoring and analytics

## ✨ Technical Features

🏗️ Microservices Architecture: Loosely coupled services following SOA principles
📨 Message Queue Integration: Robust event-driven communication via RabbitMQ
🗺️ PostGIS Geospatial: Efficient location-based queries and calculations
🔐 JWT Authentication: Secure role-based access control
📈 High Concurrency: Handles thousands of simultaneous ride requests
🔄 Graceful Recovery: Circuit breakers and automatic reconnection

## 📚 Tech Stack  

- **Go** `1.23+`
- **PostgreSQL**
- **RabbitMQ**
- **WebSocket**
- **Docker**


💻 Development

## 🧬 Project Structure

```
ride-hail-system/
├── cmd/                         # Application entry points
│   ├── ride-service/            # Handles ride requests & lifecycle
│   ├── driver-service/          # Manages driver state & assignments
│   └── admin-service/           # Admin dashboard & management
│
├── internal/                    # Private application code
│   ├── config/                  # Configuration loading & management
│   ├── db/                      # Database connections & repositories
│   ├── message/                 # Messaging (Kafka, RabbitMQ, etc.)
│   ├── auth/                    # Authentication & authorization
│   ├── websocket/               # Real-time communication
│   └── microservices/           # Business microservices logic
│       ├── ride/                # Ride domain logic
│       ├── driver/              # Driver domain logic
│       └── admin/               # Admin domain logic
│
├── pkg/                         # Public reusable library code
│   ├── models/                  # Shared domain models
│   ├── logger/                  # Logging utilities
│   └── utils/                   # Helper functions
│
├── migrations/                  # Database migrations
│
├── docker-compose.yml           # Development environment setup
├── Dockerfile                   # Container image definition
└── README.md                    # Project documentation
```

## 👨🏻‍💻 Authors

- [![Status](https://img.shields.io/badge/alem-tishmal-success?logo=github)](https://platform.alem.school/git/tishmal) <a href="https://t.me/tim_shm" target="_blank"><img src="https://img.shields.io/badge/telegram-@tishmal-blue?logo=Telegram" alt="Status" /></a>

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

This project has been created by:

- Shmal T, ***"FullStack overflow"***

## 📜 License

Apache License Version 2.0
