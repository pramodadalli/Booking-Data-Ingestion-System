
# Booking Data Ingestion System

The Booking Data Ingestion System is a backend service built with Go (Golang) and the Beego framework, designed to ingest and process booking data from various travel vendors such as airlines and hotels. The system is responsible for validating, normalizing, and storing booking data for downstream processes like reporting, GST reconciliation, and analytics.

## Features

- **Booking Data Ingestion**: Accepts and validates booking data in JSON format.
- **CRUD Operations**: Create, Read, Update, and Delete booking data through API endpoints.
- **API Endpoints**:
  - **POST /bookings**: Ingest a new booking.
  - **GET /bookings**: Retrieve a list of bookings with optional filters.
  - **GET /bookings/{id}**: Retrieve details of a specific booking by ID.
  - **DELETE /bookings/{id}**: Delete a specific booking by ID.

## Getting Started

Follow these steps to get the project running locally.

### Prerequisites

1. **Go (Golang)** version 1.20 or higher
2. **Beego Framework**: This project uses Beego as the web framework.
3. **Go Modules**: Dependencies are managed using Go modules.

### Installation

1. Clone the repository:
   ```bash
   git clone <https://github.com/pramodadalli/Booking-Data-Ingestion-System.git>
   cd booking-data-ingestion-system
