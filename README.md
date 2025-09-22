🚗 Ride-Hail System
A distributed ride-hailing platform built with Go, implementing Service-Oriented Architecture (SOA) principles. Features real-time driver matching, location tracking, and seamless passenger-driver coordination through microservices communicating via RabbitMQ message queues and WebSocket connections.

✨ Features
Core Functionality

🚕 Real-time Ride Matching: Advanced geospatial algorithms match passengers with nearby drivers
📍 Live Location Tracking: Real-time GPS tracking with sub-second updates
💰 Dynamic Pricing: Vehicle type-based fare calculation (Economy/Premium/XL)
⚡ WebSocket Communication: Bi-directional real-time updates for passengers and drivers
📊 Admin Dashboard: Comprehensive system monitoring and analytics

Technical Features

🏗️ Microservices Architecture: Loosely coupled services following SOA principles
📨 Message Queue Integration: Robust event-driven communication via RabbitMQ
🗺️ PostGIS Geospatial: Efficient location-based queries and calculations
🔐 JWT Authentication: Secure role-based access control
📈 High Concurrency: Handles thousands of simultaneous ride requests
🔄 Graceful Recovery: Circuit breakers and automatic reconnection

