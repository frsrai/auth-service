# Auth Service

Authentication and authorization service designed for modern backend systems.

## Overview
This service handles user authentication and authorization using JWT and role-based access control (RBAC).  
It is designed to be secure, scalable, and easy to integrate with other services.

## Features
- User registration and login
- JWT access & refresh token
- Role-based access control (RBAC)
- Token rotation
- Rate-limited authentication endpoints

## Tech Stack
- Backend: Go 
- Database: PostgreSQL
- Cache: Redis
- Auth: JWT

## Architecture
- Stateless authentication
- Middleware-based authorization
- Token validation at API gateway or service level

## Getting Started
```bash
docker-compose up
