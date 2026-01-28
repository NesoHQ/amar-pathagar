# Quick Start Guide

## 🚀 Fastest Way to Get Started

If you're having issues with the hot-reload setup, use the simple version:

```bash
docker compose -f docker-compose.simple.yml up --build
```

Or with make:
```bash
make dev-simple
```

This skips the Air hot-reload tool and just runs the compiled Go binary.

## 📋 Two Options

### Option 1: With Hot Reload (Recommended for Development)
```bash
docker compose up --build
```

**Pros:**
- Code changes reload automatically
- Faster development cycle

**Cons:**
- Slightly longer initial build
- Requires Air tool

### Option 2: Simple Build (Recommended for Quick Testing)
```bash
docker compose -f docker-compose.simple.yml up --build
```

**Pros:**
- Faster initial build
- More stable
- No extra dependencies

**Cons:**
- Need to rebuild on code changes

## 🔧 If Build Fails

### Issue: Air installation fails

**Solution:** Use the simple version
```bash
docker compose -f docker-compose.simple.yml up --build
```

### Issue: Port already in use

**Solution:** Stop conflicting services
```bash
# Check what's using the port
lsof -i :8080  # Backend
lsof -i :5173  # Frontend
lsof -i :5432  # Database

# Or change ports in docker-compose.yml
```

### Issue: Database connection fails

**Solution:** Wait for database to be ready
```bash
# Check database status
docker compose ps

# View database logs
docker compose logs postgres

# Restart if needed
docker compose restart postgres
```

## ✅ Verify It's Working

1. **Check all containers are running:**
   ```bash
   docker compose ps
   ```
   You should see 3 services: postgres, backend, frontend

2. **Check backend health:**
   ```bash
   curl http://localhost:8080/health
   ```
   Should return: `{"status":"ok"}`

3. **Open frontend:**
   Open http://localhost:5173 in your browser

4. **Register a user:**
   - Click "Register here"
   - Fill in the form
   - Login with your credentials

## 🎯 First Steps After Setup

1. **Create an admin user:**
   ```bash
   # Connect to database
   docker compose exec postgres psql -U library_user -d online_library
   
   # Update your user to admin
   UPDATE users SET role = 'admin' WHERE username = 'your_username';
   
   # Exit
   \q
   ```

2. **Explore the app:**
   - Dashboard: Overview of your library
   - Books: Browse and manage books (admin can add)
   - My Library: Your reading history
   - Admin: Admin panel (admin only)

## 🛑 Stop Everything

```bash
docker compose down
```

Or with volumes (clean slate):
```bash
docker compose down -v
```

## 📝 Development Workflow

### With Hot Reload
```bash
# Start
docker compose up

# Edit code in backend/ or frontend/
# Changes reload automatically

# Stop
Ctrl+C
```

### Without Hot Reload
```bash
# Start
docker compose -f docker-compose.simple.yml up

# Edit code
# Rebuild
docker compose -f docker-compose.simple.yml up --build

# Stop
Ctrl+C
```

## 🐛 Common Issues

### "Cannot connect to database"
Wait 10-15 seconds for PostgreSQL to initialize on first run.

### "Port 5432 already in use"
You have PostgreSQL running locally. Either:
- Stop local PostgreSQL: `sudo systemctl stop postgresql`
- Change port in docker-compose.yml: `"5433:5432"`

### "Frontend shows blank page"
Check browser console for errors. Usually means backend isn't running.

### "401 Unauthorized"
Your token expired. Logout and login again.

## 💡 Pro Tips

1. **View logs in real-time:**
   ```bash
   docker compose logs -f
   ```

2. **Rebuild specific service:**
   ```bash
   docker compose up --build backend
   ```

3. **Access database directly:**
   ```bash
   docker compose exec postgres psql -U library_user -d online_library
   ```

4. **Clean everything and start fresh:**
   ```bash
   docker compose down -v
   docker compose up --build
   ```

## 🎉 You're Ready!

Once you see:
- ✅ Database connected successfully
- ✅ Server starting on port 8080
- ✅ Frontend running on port 5173

You're good to go! Open http://localhost:5173 and start building your library! 📚
