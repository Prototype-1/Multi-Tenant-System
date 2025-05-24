# Multi-Tenant User Management Service

This is a backend service for managing users and their locations in a multi-tenant environment. It supports role-based access for `admin` and `user`, tenant-specific data isolation, and JWT-based authentication.

---

##  How to Run the Project Locally

### Prerequisites

- Go (v1.20+)
- PostgreSQL
- Make sure the following environment variables are set:
  - `DB_URL`
  - `JWT_SECRET`

### 1. Clone the repository

```bash
git clone https://github.com/Prototype-1/Multi-Tenant-System.git
cd Multi-Tenant-System
```

### Run PostgreSQL (locally or via Docker)

docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=yourpassword -p 5432:5432 -d postgres


### Start the Server

go run main.go



### API Endpoints

| Method | Endpoint          | Description             | Auth Required           |
| ------ | ----------------- | ----------------------- | -----------------------|
| POST   | `/create/tenants` | Create a new tenant     | NA                       |
| POST   | `/users/signup`   | Sign up (admin/user)    | (tenant_id)           |
| POST   | `/users/login`    | Login and get JWT token | NA                      |

| Method | Endpoint     | Description             | Role Required |
| ------ | ------------ | ----------------------- | ------------- |
| GET    | `/get/users` | Get all users in tenant | admin         |

| Method | Endpoint            | Description           | Role Required |
| ------ | ------------------- | --------------------- | ------------- |
| POST   | `/create/locations` | Create user location  | user          |
| GET    | `/get/me`           | Get current user info | user          |

