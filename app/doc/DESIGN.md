# Courier Tracking System

## API

- [User] List spots: `GET /pots` (PostgreSQL)
- [User] List spot items: `GET /spots/{id}/items` (PostgreSQL)
- [User] Create order: `POST /orders` (PostgreSQL)
- [Courier] Send location `POST /locations`
- [Courier] Get orders `GET /orders?`
- [Manager] Read location `GET /couriers/{id}/locations`
