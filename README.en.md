# Ride-Hail System 🚗

A real-time distributed ride-hailing platform built with Go, implementing Service-Oriented Architecture (SOA) principles. This system demonstrates advanced microservices orchestration, real-time communication via WebSockets, and sophisticated message queue patterns using RabbitMQ.

## 🎯 Overview

This project simulates the backend infrastructure of modern ride-hailing platforms like Uber, handling:

- Real-time ride requests and driver matching
- Live location tracking and updates
- Complex business workflows across distributed services
- High-concurrency scenarios with thousands of simultaneous rides
- Event-driven architecture with message queues

## 🏗️ Architecture

The system consists of three main microservices communicating through PostgreSQL and RabbitMQ:

```
┌─────────────┐        ┌───────────────────┐              ┌───────────┐
│  Passenger  │        │   Ride Service    │              │   Admin   │
│ (WebSocket) │◄──────►│  (Orchestrator)   │              │ Dashboard │
└─────────────┘        └───────────────────┘              └───────────┘
                                 ▲                               ▲
                                 │                               │
                                 ▼                               ▼
                    ┌──────────────────────────────────────────────┐
                    │         RabbitMQ Message Broker              │
                    │                                              │
                    │  • ride_topic exchange                       │
                    │  • driver_topic exchange                     │
                    │  • location_fanout exchange                  │
                    └──────────────────────────────────────────────┘
                                         ▲
                                         │
                                         ▼
┌─────────────┐             ┌───────────────────────┐
│   Driver    │             │  Driver & Location    │
│ (WebSocket) │◄───────────►│      Service          │
└─────────────┘             └───────────────────────┘
```

### Services

1. **Ride Service** - Orchestrates ride lifecycle and manages passenger interactions
2. **Driver & Location Service** - Handles driver operations, matching algorithms, and real-time tracking
3. **Admin Service** - Provides monitoring, analytics, and system oversight

## 🚀 Features

### Core Functionality
- ✅ Real-time ride request and matching
- ✅ Dynamic pricing based on distance, duration, and vehicle type
- ✅ Intelligent driver matching algorithm (distance + rating based)
- ✅ Live location tracking with WebSocket updates
- ✅ Ride status management (REQUESTED → MATCHED → EN_ROUTE → ARRIVED → IN_PROGRESS → COMPLETED)
- ✅ Cancellation handling with refund logic

### Technical Highlights
- 🔄 Event-driven architecture with RabbitMQ
- 🌐 Bidirectional real-time communication via WebSockets
- 🗺️ Geospatial queries using PostGIS
- 🔐 JWT-based authentication and authorization
- 📊 Event sourcing for complete audit trail
- 🎯 Circuit breakers and retry patterns
- 📈 System-wide monitoring and metrics

## 🛠️ Tech Stack

- **Language:** Go (with gofumpt formatting)
- **Message Queue:** RabbitMQ (AMQP)
- **Database:** PostgreSQL with PostGIS extension
- **Real-time:** WebSockets (Gorilla WebSocket)
- **Authentication:** JWT tokens
- **Allowed Libraries:**
  - `github.com/rabbitmq/amqp091-go` - AMQP client
  - `github.com/gorilla/websocket` - WebSocket implementation
  - `github.com/golang-jwt/jwt/v5` - JWT authentication
  - `pgx/v5` - PostgreSQL driver

## 📋 Prerequisites

- Go 1.21+
- PostgreSQL 14+ with PostGIS extension
- RabbitMQ 3.12+
- Docker (optional, for containerized deployment)

## 🔧 Installation

### 1. Clone the repository

```bash
git clone <repository-url>
cd ride-hail-system
```

### 2. Set up PostgreSQL

```bash
# Create database
createdb ridehail_db

# Enable PostGIS extension
psql ridehail_db -c "CREATE EXTENSION IF NOT EXISTS postgis;"

# Run migrations
psql ridehail_db < migrations/ride_service.sql
psql ridehail_db < migrations/driver_service.sql
```

### 3. Set up RabbitMQ

```bash
# Using Docker
docker run -d --name rabbitmq \
  -p 5672:5672 -p 15672:15672 \
  rabbitmq:3-management

# Or install locally
# Visit: https://www.rabbitmq.com/download.html
```

### 4. Configure environment

```bash
cp .env.example .env
# Edit .env with your database and RabbitMQ credentials
```

Example `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=ridehail_user
DB_PASSWORD=ridehail_pass
DB_NAME=ridehail_db

RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest

WS_PORT=8080
RIDE_SERVICE_PORT=3000
DRIVER_LOCATION_SERVICE_PORT=3001
ADMIN_SERVICE_PORT=3004
```

### 5. Build and run

```bash
# Build
go build -o ride-hail-system .

# Run
./ride-hail-system
```

## 📖 API Documentation

### Ride Service

#### Create a Ride
```http
POST /rides
Authorization: Bearer {passenger_token}
Content-Type: application/json

{
  "passenger_id": "uuid",
  "pickup_latitude": 43.238949,
  "pickup_longitude": 76.889709,
  "pickup_address": "Almaty Central Park",
  "destination_latitude": 43.222015,
  "destination_longitude": 76.851511,
  "destination_address": "Kok-Tobe Hill",
  "ride_type": "ECONOMY"
}
```

#### Cancel a Ride
```http
POST /rides/{ride_id}/cancel
Authorization: Bearer {passenger_token}
Content-Type: application/json

{
  "reason": "Changed my mind"
}
```

### Driver & Location Service

#### Go Online
```http
POST /drivers/{driver_id}/online
Authorization: Bearer {driver_token}
Content-Type: application/json

{
  "latitude": 43.238949,
  "longitude": 76.889709
}
```

#### Update Location
```http
POST /drivers/{driver_id}/location
Authorization: Bearer {driver_token}
Content-Type: application/json

{
  "latitude": 43.238949,
  "longitude": 76.889709,
  "accuracy_meters": 5.0,
  "speed_kmh": 45.0,
  "heading_degrees": 180.0
}
```

#### Start Ride
```http
POST /drivers/{driver_id}/start
Authorization: Bearer {driver_token}
Content-Type: application/json

{
  "ride_id": "uuid",
  "driver_location": {
    "latitude": 43.238949,
    "longitude": 76.889709
  }
}
```

### Admin Service

#### Get System Overview
```http
GET /admin/overview
Authorization: Bearer {admin_token}
```

#### Get Active Rides
```http
GET /admin/rides/active?page=1&page_size=20
Authorization: Bearer {admin_token}
```

## 🔌 WebSocket Connections

### Passenger Connection
```javascript
const ws = new WebSocket('ws://localhost:8080/ws/passengers/{passenger_id}');

// Authenticate
ws.send(JSON.stringify({
  type: 'auth',
  token: 'Bearer {passenger_token}'
}));

// Receive updates
ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  // Handle ride_status_update, driver_location_update
};
```

### Driver Connection
```javascript
const ws = new WebSocket('ws://localhost:8080/ws/drivers/{driver_id}');

// Authenticate
ws.send(JSON.stringify({
  type: 'auth',
  token: 'Bearer {driver_token}'
}));

// Receive ride offers
ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  if (data.type === 'ride_offer') {
    // Accept or reject ride
    ws.send(JSON.stringify({
      type: 'ride_response',
      offer_id: data.offer_id,
      ride_id: data.ride_id,
      accepted: true
    }));
  }
};
```

## 📊 Message Queue Architecture

### Exchanges

| Exchange | Type | Purpose |
|----------|------|---------|
| `ride_topic` | Topic | Ride-related messages |
| `driver_topic` | Topic | Driver-related messages |
| `location_fanout` | Fanout | Location broadcasts |

### Key Message Flows

1. **Ride Request Flow:**
   - Passenger → Ride Service → `ride_topic` → Driver Service
   - Driver Service matches drivers → sends offers via WebSocket
   - Driver accepts → `driver_topic` → Ride Service
   - Ride Service notifies passenger via WebSocket

2. **Location Update Flow:**
   - Driver sends location → Driver Service
   - Driver Service → `location_fanout` → All interested services
   - Services process location → update ETAs

## 🧪 Testing

```bash
# Run unit tests
go test ./...

# Run integration tests
go test -tags=integration ./...

# Load testing
go test -bench=. ./...
```

## 📝 Logging

All services implement structured JSON logging:

```json
{
  "timestamp": "2024-12-16T10:30:00Z",
  "level": "INFO",
  "service": "ride-service",
  "action": "ride_requested",
  "message": "New ride request created",
  "hostname": "ride-service-01",
  "request_id": "req_123456",
  "ride_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

## 🔒 Security

- JWT-based authentication for all API endpoints
- Role-based access control (Passenger, Driver, Admin)
- WebSocket authentication with 5-second timeout
- TLS encryption for all communications
- Input validation and sanitization
- Rate limiting on location updates

## 📈 Performance Considerations

- **Location Updates:** Rate-limited to 1 update per 3 seconds per driver
- **Driver Matching:** PostGIS spatial queries with 5km radius
- **WebSocket Connections:** Keep-alive pings every 30 seconds
- **Message Queue:** Acknowledgment-based delivery for reliability
- **Database:** Indexed queries for status and location lookups

## 🎓 Learning Objectives

This project teaches:

- Advanced message queue patterns (pub/sub, routing, fanout)
- Real-time bidirectional communication with WebSockets
- Geospatial data processing with PostGIS
- Microservices orchestration and choreography
- High-concurrency programming patterns
- Distributed state management
- Service-Oriented Architecture design

## 🐛 Troubleshooting

### RabbitMQ Connection Issues
```bash
# Check RabbitMQ status
rabbitmqctl status

# View management interface
http://localhost:15672 (guest/guest)
```

### Database Connection Issues
```bash
# Verify PostgreSQL is running
pg_isready

# Check PostGIS extension
psql ridehail_db -c "SELECT PostGIS_Version();"
```

### WebSocket Connection Failures
- Ensure authentication message is sent within 5 seconds
- Verify JWT token is valid and not expired
- Check CORS settings if connecting from browser

## 👥 Contributing

This is a learning project. Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Follow gofumpt formatting
4. Write tests for new features
5. Submit a pull request

## 📄 License

This project is created for educational purposes as part of the Alem School curriculum.

## 👨‍💻 Author

**Timur Shmal**

- Email: arturtimur201998@gmail.com
- GitHub: [@saboopher](https://github.com/saboopher/)
- LinkedIn: [Sabrina Bakirova](https://kz.linkedin.com/in/sabrina-bakirova-651b821b1)
- Discord: saboopher

## 🙏 Acknowledgments

- Alem School for the project specification
- The Go community for excellent tooling
- RabbitMQ and PostgreSQL teams for robust infrastructure

---

**Note:** This is a learning project demonstrating distributed systems concepts. For production use, additional considerations around security, monitoring, scaling, and disaster recovery would be necessary.