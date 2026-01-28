# Amar Pathagar - Quick Start Guide

## 🚀 Start Everything in 3 Commands

```bash
# 1. Navigate to project
cd online-library

# 2. Start all services
docker-compose up --build

# 3. Open in browser
# Frontend: http://localhost:3000
# Backend: http://localhost:8080
```

That's it! The entire platform is now running.

---

## 📱 Access the Application

### Frontend (Next.js)
**URL:** http://localhost:3000

**First Steps:**
1. Click "Register here"
2. Create an account (starts with 100 success score)
3. Login with your credentials
4. Explore the dashboard

### Backend API
**URL:** http://localhost:8080

**Health Check:**
```bash
curl http://localhost:8080/health
```

---

## 🎯 Test the Features

### 1. Register & Login
```
1. Go to http://localhost:3000
2. Click "Register here"
3. Fill in:
   - Full Name: Test User
   - Username: testuser
   - Email: test@example.com
   - Password: password123
4. Click "Register"
5. Login with username and password
```

### 2. Browse Books
```
1. Click "Books" in navigation
2. Use search to find books
3. Filter by status
4. Click on a book to see details
5. Like/Bookmark/Priority buttons
```

### 3. Request a Book
```
1. Go to book details
2. Click "Request This Book"
3. (Requires success score >= 20)
4. View request status
```

### 4. Post an Idea
```
1. Go to any book detail page
2. Click "Share Your Thoughts"
3. Write title and content
4. Submit (+3 points to your score)
```

### 5. Vote on Ideas
```
1. View ideas on book page
2. Click 👍 to upvote (+1 to author)
3. Click 👎 to downvote (-1 to author)
```

### 6. Make a Donation
```
1. Click "Donations" in navigation
2. Click "Make a Donation"
3. Choose type (money/book)
4. Fill in details
5. Submit (+10 or +20 points)
```

### 7. Check Leaderboard
```
1. Click "Leaderboard"
2. View top contributors
3. Switch between categories
4. See rankings and scores
```

### 8. View My Library
```
1. Click "My Library"
2. See your bookmarks
3. Filter by type
4. View reading stats
```

### 9. Admin Functions (Admin Only)
```
1. Login as admin
2. Click "Admin" in navigation
3. Add new books
4. Manage users
5. Approve requests
```

---

## 🎨 Design Features to Notice

### Classic Old-School Aesthetic
- **Grey tones** throughout
- **Bold borders** (2-4px)
- **Uppercase headings** with wide spacing
- **Serif fonts** for readability
- **Old paper texture** background
- **Stamp-like badges** (rotated)
- **Vintage feel** everywhere

### Interactive Elements
- **Hover effects** on all buttons
- **Shadow offsets** on cards
- **Border highlights** on focus
- **Smooth transitions**

---

## 📊 Success Score System

### Starting Score: 100

### Ways to Earn Points
| Action | Points |
|--------|--------|
| Post reading idea | +3 |
| Idea gets upvoted | +1 |
| Return book on time | +10 |
| Positive review | +5 |
| Donate money | +10 |
| Donate book | +20 |

### Ways to Lose Points
| Action | Points |
|--------|--------|
| Idea gets downvoted | -1 |
| Return book late | -15 |
| Negative review | -10 |
| Lose a book | -50 |

### Score Requirements
- **Minimum 20** to request books
- **Below 20** = Low priority in matching

---

## 🔧 Troubleshooting

### Frontend won't start
```bash
cd frontend
rm -rf node_modules .next
npm install
npm run dev
```

### Backend won't start
```bash
cd backend
go mod tidy
go run cmd/api/main.go
```

### Database issues
```bash
docker-compose down -v
docker-compose up -d postgres
# Wait 10 seconds
docker-compose up backend frontend
```

### Port already in use
```bash
# Kill process on port 3000
lsof -ti:3000 | xargs kill -9

# Kill process on port 8080
lsof -ti:8080 | xargs kill -9
```

### Can't login
```bash
# Check backend logs
docker-compose logs backend

# Check if database is ready
docker-compose logs postgres
```

---

## 📝 Default Test Data

### Create Admin User
```bash
# Register normally, then update in database
docker exec -it online-library-db psql -U library_user -d online_library

UPDATE users SET role = 'admin' WHERE username = 'testuser';
\q
```

### Add Sample Books (Admin Panel)
```
1. Login as admin
2. Go to Admin panel
3. Click "Add New Book"
4. Fill in:
   - Title: The Great Gatsby
   - Author: F. Scott Fitzgerald
   - Physical Code: BOOK001
   - Category: Fiction
5. Submit
```

---

## 🎯 Feature Checklist

Test these features:

- [ ] Register new account
- [ ] Login
- [ ] View dashboard
- [ ] Browse books
- [ ] Search books
- [ ] View book details
- [ ] Like a book
- [ ] Bookmark a book
- [ ] Add to priority
- [ ] Request a book
- [ ] Post an idea
- [ ] Upvote an idea
- [ ] View leaderboard
- [ ] Make a donation
- [ ] View my library
- [ ] Check success score
- [ ] Add book (admin)

---

## 📱 Pages Overview

| Page | URL | Purpose |
|------|-----|---------|
| Home | `/` | Redirects to dashboard/login |
| Login | `/login` | User authentication |
| Register | `/register` | New user signup |
| Dashboard | `/dashboard` | User overview |
| Books | `/books` | Browse collection |
| Book Detail | `/books/[id]` | Book info & ideas |
| My Library | `/my-library` | Personal bookmarks |
| Leaderboard | `/leaderboard` | Top contributors |
| Donations | `/donations` | Support platform |
| Admin | `/admin` | Management panel |

---

## 🔐 API Endpoints

### Authentication
```
POST /api/auth/register
POST /api/auth/login
GET  /api/me
```

### Books
```
GET    /api/books
GET    /api/books/:id
POST   /api/books (admin)
POST   /api/books/:id/request
```

### Ideas
```
POST   /api/ideas
GET    /api/books/:bookId/ideas
POST   /api/ideas/:id/vote
```

### Users
```
GET    /api/users/:id/profile
PUT    /api/users/profile
POST   /api/users/interests
GET    /api/leaderboard
```

### Donations
```
POST   /api/donations
GET    /api/donations
```

### Bookmarks
```
POST   /api/bookmarks
DELETE /api/bookmarks/:bookId
GET    /api/bookmarks
```

---

## 🎉 You're All Set!

The complete Amar Pathagar platform is now running with:

✅ **Next.js Frontend** - Classic old-school design
✅ **Go Backend** - All APIs working
✅ **PostgreSQL** - Database with full schema
✅ **Docker** - Everything containerized

**Start exploring:** http://localhost:3000

**Have fun building your reading community! 📚✨**

---

## 💡 Pro Tips

1. **Start with high score** - You begin with 100 points
2. **Post ideas early** - Easy +3 points each
3. **Donate to boost** - +20 for books, +10 for money
4. **Keep score above 20** - Required to request books
5. **Check leaderboard** - See top contributors
6. **Use bookmarks** - Organize your reading list
7. **Admin panel** - Add books to grow library

---

*Need help? Check NEXTJS_FRONTEND_COMPLETE.md for detailed documentation*
