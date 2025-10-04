# DTQ


## Overview
Simple microservices-based rideshare system that calculates routes, estimates fares, and manages trip bookings.

## Architecture
- 🌐 **API Gateway** – REST API for client requests  
- 🗺️ **Trip Service** – Core rideshare logic with OSRM integration  
- 👨‍✈️ **Driver Service** – Manage drivers  
- 🧮 **OSRM** – Open Source Routing Machine for route calculation  

## Features
- 🗺️ **Route Calculation** – Real-time routing via OSRM  
- 💰 **Fare Estimation** – Dynamic pricing based on distance and duration  
- 🚗 **Trip Management** – Book and track rides  

## Tech Stack

| Tech | Logo |
|------|------|
| Go | ![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) |
| gRPC | ![gRPC](https://img.shields.io/badge/gRPC-448CDA?logo=grpc&logoColor=white) |
| RabbitMQ | ![RabbitMQ](https://img.shields.io/badge/RabbitMQ-FF6600?logo=rabbitmq&logoColor=white) |
| OSRM | ![OSRM](https://img.shields.io/badge/OSRM-000000?logo=openstreetmap&logoColor=white) |
| Protocol Buffers | ![Protobuf](https://img.shields.io/badge/Protobuf-4285F4?logo=google&logoColor=white) |

