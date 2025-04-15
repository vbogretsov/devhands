# Courier Tracking System

## API

- [Admin] Create spot: `POST /spots` (PostgreSQL)
- [Admin] Create item: `POST /spots/{id}/items` (PostgreSQL)
- [User] List spots: `GET /spots` (PostgreSQL)
- [User] List spot items: `GET /spots/{id}/items` (PostgreSQL)
- [User] Create order: `POST /orders` (PostgreSQL)
- [Courier] Send location `POST /locations`
- [Courier] Get orders `GET /orders?`
- [Manager] Read location `GET /couriers/{id}/locations`
- [Manager] Read locations `GET /couriers&tl=x1,y1&br=x2,y2`
