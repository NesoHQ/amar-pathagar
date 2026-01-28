# Online Library - Setup Guide

## 🚀 Quick Start (Recommended)

The easiest way to get started is using Docker Compose:

```bash
# 1. Clone or navigate to the project
cd online-library

# 2. Start all services
make dev

# Or without make:
docker-compose up
```

That's it! The application will be available at:
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080
- Database: localhost:5432

## 📋 Prerequisites

### For Docker Setup (Recommended)
- Docker Desktop or Docker Engine
- Docker Compose

### For Local Development
- Go 1.21 or higher
- Node.js 18 or higher
- PostgreSQL 15
- npm or yarn

## 🔧 Detailed Setup

### Option 1: Docker Setup (Recommended)

1. **Environment Configuration**
   ```bash
   # Backend
   cp backend/.env.example backend/.env
   
   # Frontend
   cp frontend/.env.example frontend/.env
   ```

2. **Start Services**
   ```bash
   docker-compose up -d
   ```

3. **Check Logs**
   ```bash
   docker-compose logs -f
   ```

4. **Stop Services**
   ```bash
   docker-compose down
   ```

### Option 2: Local Development

#### Backend Setup

1. **Navigate to backend**
   ```bash
   cd backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Setup PostgreSQL**
   ```bash
   # Create database
   createdb online_library
   
   # Run init script
   psql online_library < ../database/init.sql
   ```

4. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env with your local database credentials
   ```

5. **Run backend**
   ```bash
   # With hot reload (requires Air)
   go install github.com/cosmtrek/air@latest
   air -c .air.toml
   
   # Or without hot reload
   go run cmd/api/main.go
   ```

#### Frontend Setup

1. **Navigate to frontend**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Configure environment**
   ```bash
   cp .env.example .env
   ```

4. **Run frontend**
   ```bash
   npm run dev
   ```

## 🧪 Testing the Application

### 1. Create an Admin User

First user should be created as admin. You can do this by:

**Option A: Direct Database Insert**
```sql
INSERT INTO users (username, email, password_hash, full_name, role)
VALUES (
  'admin',
  'admin@library.com',
  '$2a$10$YourHashedPasswordHere',
  'Admin User',
  'admin'
);
```

**Option B: Register via API then update role**
```bash
# Register
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "email": "admin@library.com",
    "password": "admin123",
    "full_name": "Admin User"
  }'

# Then update role in database
UPDATE users SET role = 'admin' WHERE username = 'admin';
```

### 2. Login and Test

1. Open http://localhost:5173
2. Login with your credentials
3. Explore the dashboard

## 📁 Project Structure

```
online-library/
├── backend/
│   ├── cmd/api/              # Application entry point
│   ├── internal/
│   │   ├── config/          # Configuration
│   │   ├── database/        # Database connection
│   │   ├── dto/             # Data Transfer Objects
│   │   ├── handlers/        # HTTP handlers
│   │   ├── middleware/      # HTTP middleware
│   │   ├── models/          # Domain models
│   │   ├── repository/      # Data access layer
│   │   └── services/        # Business logic
│   ├── go.mod
│   └── .air.toml           # Hot reload config
├── frontend/
│   ├── src/
│   │   ├── components/     # Reusable components
│   │   ├── pages/         # Page components
│   │   ├── services/      # API services
│   │   ├── stores/        # State management
│   │   └── App.tsx
│   ├── package.json
│   └── vite.config.ts
├── database/
│   └── init.sql           # Database schema
├── docker-compose.yml
├── Makefile
└── README.md
```

## 🔑 Default Credentials

After setup, create your first user via the registration page.

## 🐛 Troubleshooting

### Database Connection Issues
```bash
# Check if PostgreSQL is running
docker-compose ps

# View database logs
docker-compose logs postgres

# Restart database
docker-compose restart postgres
```

### Backend Not Starting
```bash
# Check backend logs
docker-compose logs backend

# Rebuild backend
docker-compose up --build backend
```

### Frontend Not Loading
```bash
# Check frontend logs
docker-compose logs frontend

# Rebuild frontend
docker-compose up --build frontend
```

### Port Already in Use
```bash
# Check what's using the port
lsof -i :8080  # Backend
lsof -i :5173  # Frontend
lsof -i :5432  # Database

# Kill the process or change ports in docker-compose.yml
```

## 📝 Next Steps

1. **Add Books**: Use the admin panel to add books to the library
2. **Invite Users**: Share registration link with community members
3. **Configure**: Customize settings in .env files
4. **Deploy**: See deployment guide for production setup

## 🤝 Contributing

This is a community project! Feel free to:
- Report bugs
- Suggest features
- Submit pull requests
- Improve documentation

## 📚 API Documentation

Once the backend is running, API documentation will be available at:
- Swagger UI: http://localhost:8080/swagger/index.html (coming soon)

## 🔒 Security Notes

- Change JWT_SECRET in production
- Use strong passwords
- Enable HTTPS in production
- Regular database backups
- Keep dependencies updated

## 📞 Support

If you encounter issues:
1. Check this guide
2. Review logs: `docker-compose logs`
3. Check GitHub issues
4. Ask in community discussions
