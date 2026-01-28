#!/bin/bash

# Amar Pathagar Deployment Script
# This script sets up Nginx reverse proxy and deploys the application

set -e

echo "🚀 Amar Pathagar Deployment Script"
echo "===================================="

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    echo "❌ Please run as root (use sudo)"
    exit 1
fi

# 1. Install Nginx if not installed
echo ""
echo "📦 Step 1: Checking Nginx installation..."
if ! command -v nginx &> /dev/null; then
    echo "Installing Nginx..."
    apt update
    apt install -y nginx
else
    echo "✅ Nginx is already installed"
fi

# 2. Copy Nginx configuration
echo ""
echo "⚙️  Step 2: Setting up Nginx configuration..."
cp nginx.conf /etc/nginx/sites-available/amar-pathagar

# Remove default site if exists
if [ -f /etc/nginx/sites-enabled/default ]; then
    rm /etc/nginx/sites-enabled/default
    echo "Removed default Nginx site"
fi

# Enable the site
ln -sf /etc/nginx/sites-available/amar-pathagar /etc/nginx/sites-enabled/
echo "✅ Nginx configuration installed"

# 3. Test Nginx configuration
echo ""
echo "🧪 Step 3: Testing Nginx configuration..."
nginx -t

# 4. Reload Nginx
echo ""
echo "🔄 Step 4: Reloading Nginx..."
systemctl reload nginx
systemctl enable nginx
echo "✅ Nginx reloaded and enabled"

# 5. Check if Docker is installed
echo ""
echo "🐳 Step 5: Checking Docker installation..."
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

echo "✅ Docker is installed"

# 6. Start the application
echo ""
echo "🚀 Step 6: Starting the application..."
docker compose down
docker compose up -d --build

# 7. Wait for services to be ready
echo ""
echo "⏳ Waiting for services to start..."
sleep 10

# 8. Check service status
echo ""
echo "📊 Service Status:"
echo "==================="
docker compose ps

# 9. Test endpoints
echo ""
echo "🧪 Testing endpoints..."
echo ""

# Test backend health
if curl -s http://127.0.0.1:8080/health > /dev/null; then
    echo "✅ Backend is running on http://127.0.0.1:8080"
else
    echo "⚠️  Backend health check failed"
fi

# Test frontend
if curl -s http://127.0.0.1:3000 > /dev/null; then
    echo "✅ Frontend is running on http://127.0.0.1:3000"
else
    echo "⚠️  Frontend check failed"
fi

# Test Nginx proxy
if curl -s http://127.0.0.1 > /dev/null; then
    echo "✅ Nginx proxy is working on http://127.0.0.1"
else
    echo "⚠️  Nginx proxy check failed"
fi

echo ""
echo "✨ Deployment Complete!"
echo "======================="
echo ""
echo "🌍 Your application is now accessible at:"
echo "   👉 http://3.73.214.80"
echo ""
echo "📝 Useful commands:"
echo "   - View logs: docker compose logs -f"
echo "   - Restart: docker compose restart"
echo "   - Stop: docker compose down"
echo "   - Nginx logs: tail -f /var/log/nginx/access.log"
echo "   - Nginx errors: tail -f /var/log/nginx/error.log"
echo ""
