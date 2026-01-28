# Amar Pathagar - Implementation Guide

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Go 1.21+ (for local development)
- Node.js 18+ (for local development)

### 1. Clone and Setup

```bash
# Navigate to project
cd online-library

# Copy environment files
cp backend/.env.example backend/.env
cp frontend/.env.example frontend/.env
```

### 2. Start with Docker (Easiest)

```bash
# Start all services
docker-compose up --build

# Or use simple mode
docker-compose -f docker-compose.simple.yml up --build
```

Access:
- Frontend: http://localhost:5173
- Backend: http://localhost:8080
- Database: localhost:5432

### 3. Manual Setup (Development)

#### Backend
```bash
cd backend
go mod tidy
go run cmd/api/main.go

# Or with hot reload
air
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

#### Database
```bash
# Start PostgreSQL
docker run -d \
  --name amar-pathagar-db \
  -e POSTGRES_USER=library \
  -e POSTGRES_PASSWORD=library123 \
  -e POSTGRES_DB=online_library \
  -p 5432:5432 \
  postgres:15

# Run migrations
psql -h localhost -U library -d online_library -f database/init.sql
```

## Testing the Features

### 1. Create Test Users

```bash
# Register users via API or UI
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "reader1",
    "email": "reader1@example.com",
    "password": "password123",
    "full_name": "Test Reader"
  }'
```

### 2. Test Success Score System

```bash
# Login and get token
TOKEN=$(curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"reader1","password":"password123"}' \
  | jq -r '.access_token')

# View profile (should show success_score: 100)
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8080/api/users/:userId/profile
```

### 3. Test Donations

```bash
# Create a donation
curl -X POST http://localhost:8080/api/donations \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "donation_type": "money",
    "amount": 50.00,
    "currency": "USD",
    "message": "Supporting the community!",
    "is_public": true
  }'

# Check success score increased by 10
```

### 4. Test Ideas & Voting

```bash
# Post an idea
curl -X POST http://localhost:8080/api/ideas \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "book_id": "book-uuid-here",
    "title": "Key Insights",
    "content": "This book taught me..."
  }'

# Vote on an idea
curl -X POST http://localhost:8080/api/ideas/:ideaId/vote \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"vote_type": "upvote"}'
```

### 5. Test Leaderboard

```bash
# Get leaderboard
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8080/api/leaderboard
```

## Feature Testing Checklist

### User System
- [ ] Register new user
- [ ] Login
- [ ] View public profile
- [ ] Update profile with location
- [ ] Add interests

### Success Score
- [ ] Check initial score (100)
- [ ] Post idea (+3 points)
- [ ] Receive upvote (+1 point)
- [ ] Donate book (+20 points)
- [ ] Donate money (+10 points)
- [ ] View score history

### Book Management
- [ ] Add book (admin)
- [ ] View book details
- [ ] See current holder
- [ ] Check availability

### Bookmarks
- [ ] Like a book
- [ ] Bookmark a book
- [ ] Add to priority list
- [ ] View all bookmarks

### Ideas & Knowledge
- [ ] Post reading idea
- [ ] View book ideas
- [ ] Upvote idea
- [ ] Downvote idea
- [ ] See vote counts

### Reviews
- [ ] Write user review
- [ ] Rate behavior (1-5)
- [ ] Rate book condition (1-5)
- [ ] Rate communication (1-5)
- [ ] View received reviews

### Donations
- [ ] Donate money
- [ ] Donate book
- [ ] View donation list
- [ ] Check donor badge

### Leaderboard
- [ ] View top readers
- [ ] View top sharers
- [ ] View top donors
- [ ] View highest scores
- [ ] View top idea writers

### Matching System
- [ ] Request a book
- [ ] See priority score
- [ ] Check distance calculation
- [ ] View interest match score

## Common Issues & Solutions

### Database Connection Error
```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Restart database
docker-compose restart db
```

### Backend Won't Start
```bash
# Check Go version
go version  # Should be 1.21+

# Clean and rebuild
cd backend
go clean
go mod tidy
go build cmd/api/main.go
```

### Frontend Build Error
```bash
# Clear node_modules
cd frontend
rm -rf node_modules package-lock.json
npm install
```

### CORS Issues
```bash
# Check backend CORS config in cmd/api/main.go
# Ensure frontend URL is in AllowOrigins
```

## Environment Variables

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=library
DB_PASSWORD=library123
DB_NAME=online_library
JWT_SECRET=your-secret-key-change-in-production
SERVER_PORT=8080
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080
```

## Database Schema Verification

```sql
-- Check all tables exist
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public';

-- Should show:
-- users, books, reading_history, book_requests
-- reading_ideas, idea_votes, user_reviews, donations
-- user_interests, user_bookmarks, notifications
-- success_score_history, waiting_queue, audit_logs

-- Check user table structure
\d users

-- Verify indexes
\di
```

## Performance Optimization

### Database Indexes
All critical indexes are created in `database/init.sql`:
- User lookups (username, email)
- Book status queries
- Success score sorting
- Location-based queries
- Notification filtering

### Backend Optimization
- Connection pooling configured
- Prepared statements used
- Efficient queries with proper JOINs
- Pagination support

### Frontend Optimization
- Code splitting with React Router
- Lazy loading components
- Optimized bundle size
- Efficient state management with Zustand

## Deployment Checklist

### Pre-deployment
- [ ] Update JWT_SECRET to strong random value
- [ ] Configure production database
- [ ] Set up SSL/TLS certificates
- [ ] Configure CORS for production domain
- [ ] Enable rate limiting
- [ ] Set up monitoring
- [ ] Configure backups

### Production Environment
```bash
# Build frontend
cd frontend
npm run build

# Build backend
cd backend
go build -o main cmd/api/main.go

# Run with production config
./main
```

### Docker Production
```bash
# Use production compose file
docker-compose -f docker-compose.prod.yml up -d
```

## Monitoring & Maintenance

### Health Checks
```bash
# Backend health
curl http://localhost:8080/health

# Database connection
psql -h localhost -U library -d online_library -c "SELECT 1"
```

### Logs
```bash
# Backend logs
docker-compose logs -f backend

# Database logs
docker-compose logs -f db

# All logs
docker-compose logs -f
```

### Backup Database
```bash
# Backup
docker exec amar-pathagar-db pg_dump -U library online_library > backup.sql

# Restore
docker exec -i amar-pathagar-db psql -U library online_library < backup.sql
```

## Support & Documentation

- Full feature list: `AMAR_PATHAGAR_FEATURES.md`
- Project summary: `PROJECT_SUMMARY.md`
- Quick start: `QUICKSTART.md`
- Detailed setup: `SETUP.md`

## Next Steps

1. **Customize**: Update branding, colors, logo
2. **Test**: Run through all features
3. **Seed Data**: Add initial books and users
4. **Deploy**: Choose hosting platform
5. **Monitor**: Set up analytics and error tracking
6. **Iterate**: Gather feedback and improve

---

**Happy Building! 📚✨**
