# Amar Pathagar - Deployment Guide

## 🌍 Deploying to AWS VPS (Nginx Reverse Proxy)

This guide explains how to deploy Amar Pathagar on your AWS VPS using Nginx as a reverse proxy.

### Architecture

```
Internet (Port 80/443)
    ↓
Nginx (Reverse Proxy)
    ↓
    ├─→ Frontend (localhost:3000) - Next.js
    └─→ Backend (localhost:8080) - Go API
```

### Prerequisites

- AWS VPS with Ubuntu
- Ports 80 and 443 open in AWS Security Group
- Docker and Docker Compose installed
- Root/sudo access

---

## 🚀 Quick Deployment

### Option 1: Automated Deployment (Recommended)

```bash
# Make the script executable
chmod +x deploy.sh

# Run the deployment script
sudo ./deploy.sh
```

This script will:
1. Install Nginx (if not installed)
2. Configure Nginx reverse proxy
3. Build and start Docker containers
4. Test all endpoints
5. Display access information

---

### Option 2: Manual Deployment

#### Step 1: Install Nginx

```bash
sudo apt update
sudo apt install -y nginx
```

#### Step 2: Configure Nginx

```bash
# Copy the configuration
sudo cp nginx.conf /etc/nginx/sites-available/amar-pathagar

# Remove default site
sudo rm /etc/nginx/sites-enabled/default

# Enable our site
sudo ln -s /etc/nginx/sites-available/amar-pathagar /etc/nginx/sites-enabled/

# Test configuration
sudo nginx -t

# Reload Nginx
sudo systemctl reload nginx
sudo systemctl enable nginx
```

#### Step 3: Update Environment Variables

Edit `backend/.env.example` and create `backend/.env`:

```bash
DB_HOST=postgres
DB_PORT=5432
DB_USER=library_user
DB_PASSWORD=library_pass
DB_NAME=online_library
JWT_SECRET=your-secret-key-change-in-production
PORT=8080
```

Edit `frontend/.env.local`:

```bash
NEXT_PUBLIC_API_URL=http://3.73.214.80
```

#### Step 4: Start the Application

```bash
# Build and start containers
docker compose down
docker compose up -d --build

# Check status
docker compose ps

# View logs
docker compose logs -f
```

---

## 🔍 Verification

### Check Services

```bash
# Backend health
curl http://127.0.0.1:8080/health

# Frontend
curl http://127.0.0.1:3000

# Through Nginx
curl http://127.0.0.1
```

### Access the Application

Open in your browser:
- **Frontend**: http://3.73.214.80
- **API**: http://3.73.214.80/api/health

---

## 📊 Monitoring & Logs

### Docker Logs

```bash
# All services
docker compose logs -f

# Specific service
docker compose logs -f backend
docker compose logs -f frontend
```

### Nginx Logs

```bash
# Access logs
sudo tail -f /var/log/nginx/access.log

# Error logs
sudo tail -f /var/log/nginx/error.log
```

### Service Status

```bash
# Docker containers
docker compose ps

# Nginx status
sudo systemctl status nginx
```

---

## 🔧 Troubleshooting

### Issue: Cannot access from outside

**Check Nginx is running:**
```bash
sudo systemctl status nginx
```

**Check ports are listening:**
```bash
sudo ss -tulnp | grep -E ':(80|3000|8080)'
```

**Check Docker containers:**
```bash
docker compose ps
```

### Issue: 502 Bad Gateway

This means Nginx can't reach the backend/frontend.

**Check containers are running:**
```bash
docker compose ps
```

**Check backend health:**
```bash
curl http://127.0.0.1:8080/health
```

**Restart services:**
```bash
docker compose restart
```

### Issue: Database connection failed

**Check PostgreSQL container:**
```bash
docker compose logs postgres
```

**Restart database:**
```bash
docker compose restart postgres
```

---

## 🔐 Security Recommendations

### 1. Change Default Credentials

Update in `docker-compose.yml`:
- Database password
- JWT secret

### 2. Enable HTTPS (Optional but Recommended)

Install Certbot for Let's Encrypt SSL:

```bash
sudo apt install -y certbot python3-certbot-nginx
sudo certbot --nginx -d yourdomain.com
```

### 3. Firewall Configuration

```bash
# Allow only necessary ports
sudo ufw allow 22/tcp   # SSH
sudo ufw allow 80/tcp   # HTTP
sudo ufw allow 443/tcp  # HTTPS
sudo ufw enable
```

---

## 🔄 Updates & Maintenance

### Update Application

```bash
# Pull latest changes
git pull

# Rebuild and restart
docker compose down
docker compose up -d --build
```

### Backup Database

```bash
# Create backup
docker compose exec postgres pg_dump -U library_user online_library > backup.sql

# Restore backup
docker compose exec -T postgres psql -U library_user online_library < backup.sql
```

### Clean Up

```bash
# Remove unused Docker resources
docker system prune -a

# Remove old images
docker image prune -a
```

---

## 📝 Default Credentials

**Test User:**
- Username: `kraken`
- Password: `12345678`

**Database:**
- User: `library_user`
- Password: `library_pass`
- Database: `online_library`

---

## 🆘 Support

If you encounter issues:

1. Check logs: `docker compose logs -f`
2. Check Nginx logs: `sudo tail -f /var/log/nginx/error.log`
3. Verify all services are running: `docker compose ps`
4. Test endpoints individually

---

## 📚 Additional Resources

- [Docker Documentation](https://docs.docker.com/)
- [Nginx Documentation](https://nginx.org/en/docs/)
- [Next.js Deployment](https://nextjs.org/docs/deployment)
- [Go Deployment Best Practices](https://golang.org/doc/)
