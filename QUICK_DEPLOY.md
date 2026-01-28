# 🚀 Quick Deploy - Amar Pathagar

## One-Command Deployment

```bash
chmod +x deploy.sh && sudo ./deploy.sh
```

That's it! Your app will be live at **http://3.73.214.80**

---

## What Gets Deployed

- ✅ Nginx reverse proxy on port 80
- ✅ Backend API (Go) on localhost:8080
- ✅ Frontend (Next.js) on localhost:3000
- ✅ PostgreSQL database
- ✅ All services in Docker containers

---

## Access Points

| Service | URL |
|---------|-----|
| **Frontend** | http://3.73.214.80 |
| **API** | http://3.73.214.80/api |
| **Health Check** | http://3.73.214.80/health |

---

## Login Credentials

**Username:** `kraken`  
**Password:** `12345678`

---

## Useful Commands

```bash
# View logs
docker compose logs -f

# Restart services
docker compose restart

# Stop everything
docker compose down

# Check status
docker compose ps

# Nginx logs
sudo tail -f /var/log/nginx/access.log
```

---

## Troubleshooting

**502 Bad Gateway?**
```bash
docker compose restart
```

**Can't access from outside?**
```bash
sudo systemctl status nginx
sudo ss -tulnp | grep :80
```

**Database issues?**
```bash
docker compose logs postgres
docker compose restart postgres
```

---

## Update Application

```bash
git pull
docker compose down
docker compose up -d --build
```

---

## Architecture

```
Internet (Port 80)
    ↓
Nginx Reverse Proxy
    ↓
    ├─→ Frontend (localhost:3000)
    └─→ Backend API (localhost:8080)
            ↓
        PostgreSQL (localhost:5432)
```

All services run on localhost and are only accessible through Nginx on port 80.

**No AWS Security Group changes needed!** ✨
