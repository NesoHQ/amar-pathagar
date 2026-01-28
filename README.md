# Amar Pathagar - Community Book Sharing Platform

> **"আমার পাঠাগার" (My Library)** - A trust-based, community-driven book sharing platform

[![Status](https://img.shields.io/badge/status-production--ready-success)]()
[![Backend](https://img.shields.io/badge/backend-Go%201.23-00ADD8)]()
[![Frontend](https://img.shields.io/badge/frontend-Next.js%2014-000000)]()
[![Database](https://img.shields.io/badge/database-PostgreSQL%2015-336791)]()

---

## 🎯 Vision

**Amar Pathagar** is more than a book tracking system—it's a living, breathing community where:
- 📚 Books travel from reader to reader
- 🤝 Trust and reputation matter
- 💡 Knowledge is shared, not hoarded
- ⭐ Readers become living libraries
- 🏆 Every exchange builds community

---

## ✨ Key Features

### 🏆 Trust & Reputation System
- **Success Score**: Dynamic reputation (100 starting points)
- Earn points: Return on time (+10), Post ideas (+3), Donate (+20)
- Lose points: Late returns (-15), Negative reviews (-10), Lost books (-50)
- Minimum score required to request books

### 🤝 Smart Matching Algorithm
When multiple users want a book, the system selects based on:
- **Location** (30%): Closest reader gets priority
- **Success Score** (40%): Higher reputation wins
- **Interest Match** (30%): Topic relevance matters

### 💡 Knowledge Layer
- Post reading ideas and reflections
- Upvote/downvote community insights
- Build a knowledge archive
- Votes affect success scores

### ⭐ Review System
- Rate behavior, book condition, communication (1-5 stars)
- Reviews impact success scores
- Build trust through transparency

### 🎁 Donation System
- Donate books or money
- Public recognition with donor badge
- Transparent contribution tracking
- Success score bonuses

### 🏅 Leaderboards
- Top readers, sharers, donors
- Highest success scores
- Most impactful idea writers

### 📚 Book Management
- Track physical books with unique codes
- See current holder and availability
- Request books with smart queue
- Time-bound borrowing with due dates

### 🔖 Bookmarks & Priorities
- Like, bookmark, or prioritize books
- See if liked books are available
- Manage your reading wishlist

---

## 🎨 Design Philosophy

### Classic Old-School Aesthetic
- **Grey tones** - Black ink on old paper
- **Bold typography** - Uppercase headings, wide tracking
- **Serif fonts** - Georgia for readability
- **Vintage elements** - Stamps, borders, shadows
- **Minimal UI** - Distraction-free reading focus

### Trust-Based Network
- Reputation through contribution
- Community moderation
- Transparent scoring system
- Identity and trust matter

### Knowledge > Hoarding
- Share insights after reading
- Build collective wisdom
- Reward thoughtful contributions
- Books as moving assets

---

## 🛠️ Tech Stack

### Backend
- **Language:** Go 1.23
- **Framework:** Gin
- **Database:** PostgreSQL 15
- **Auth:** JWT tokens
- **Architecture:** Clean Architecture

### Frontend
- **Framework:** Next.js 14
- **Language:** TypeScript
- **Styling:** Tailwind CSS
- **State:** Zustand
- **HTTP:** Axios

### DevOps
- **Containers:** Docker & Docker Compose
- **Hot Reload:** Air (backend), Next.js (frontend)
- **Database:** PostgreSQL with migrations

---

## 🚀 Quick Start

### Prerequisites
- Docker & Docker Compose
- (Optional) Go 1.23+ for local backend dev
- (Optional) Node.js 18+ for local frontend dev

### Start Everything (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd online-library

# Start all services
docker-compose up --build

# Access the application
# Frontend: http://localhost:3000
# Backend: http://localhost:8080
# Database: localhost:5432
```

That's it! The entire platform is now running.

### First Steps

1. **Register:** Go to http://localhost:3000 and create an account
2. **Explore:** Browse books, check leaderboard, view donations
3. **Engage:** Like books, post ideas, make donations
4. **Admin:** Update a user to admin role in database to access admin panel

---

## 📖 Documentation

### Quick Guides
- **[QUICK_START_NEXTJS.md](./QUICK_START_NEXTJS.md)** - Get started in 5 minutes
- **[QUICK_REFERENCE.md](./QUICK_REFERENCE.md)** - Quick reference card

### Detailed Documentation
- **[FINAL_IMPLEMENTATION_SUMMARY.md](./FINAL_IMPLEMENTATION_SUMMARY.md)** - Complete implementation overview
- **[AMAR_PATHAGAR_FEATURES.md](./AMAR_PATHAGAR_FEATURES.md)** - All features detailed
- **[NEXTJS_FRONTEND_COMPLETE.md](./NEXTJS_FRONTEND_COMPLETE.md)** - Frontend documentation
- **[ARCHITECTURE.md](./ARCHITECTURE.md)** - System architecture
- **[IMPLEMENTATION_GUIDE.md](./IMPLEMENTATION_GUIDE.md)** - Setup and testing

---

## 📊 Features Overview

### ✅ Fully Implemented (17/17)

1. **Design & Branding** - Classic grey/old-school theme
2. **User System** - Registration, login, profiles, success scores
3. **Book Management** - Add, browse, search, track books
4. **Search & Discovery** - By name, author, topic, category
5. **Bookmark System** - Like, bookmark, priority lists
6. **Request Flow** - Smart queue with time-bound borrowing
7. **Matching Engine** - Location + Score + Interest algorithm
8. **Success Score** - 11 ways to earn/lose points
9. **Ideas & Knowledge** - Post, vote, build archive
10. **Review System** - User-to-user ratings and feedback
11. **Book Donations** - Contribute books to community
12. **Money Contributions** - Financial support system
13. **Leaderboard** - 5 categories of top contributors
14. **Availability** - Clear status and holder info
15. **Logistics** - Track, countdown, auto-updates
16. **Notifications** - System-wide alerts (ready)
17. **Moderation** - Admin panel and controls

---

## 🎯 Success Score System

### Starting Score: 100

| Action | Points | Notes |
|--------|--------|-------|
| ✅ Return on time | +10 | Builds trust |
| ⏰ Return late | -15 | Hurts reputation |
| ⭐ Positive review | +5 | 4-5 stars |
| 👎 Negative review | -10 | <3 stars |
| 💡 Post idea | +3 | Share knowledge |
| 👍 Idea upvoted | +1 | Quality content |
| 👎 Idea downvoted | -1 | Poor content |
| 📚 Donate book | +20 | Big contribution |
| 💰 Donate money | +10 | Support platform |
| ❌ Lost book | -50 | Major penalty |

**Minimum Score:** 20 to request books

---

## 🗄️ Database Schema

### 15 Tables
- **users** - Enhanced with success scores, location, stats
- **books** - With donation tracking, topics, ratings
- **reading_history** - Complete reading timeline
- **book_requests** - Smart request queue with priority scores
- **reading_ideas** - Knowledge sharing posts
- **idea_votes** - Community voting
- **user_reviews** - Peer reviews
- **donations** - Book and money contributions
- **user_interests** - For matching algorithm
- **user_bookmarks** - Likes, bookmarks, priorities
- **notifications** - System alerts
- **success_score_history** - Full audit trail
- **waiting_queue** - Legacy support
- **audit_logs** - System-wide tracking

### 19 Indexes
All optimized for query performance

---

## 🔌 API Endpoints

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

### Ideas & Knowledge
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

### Reviews
```
POST   /api/reviews
GET    /api/users/:id/reviews
```

---

## 📱 Pages

| Page | URL | Purpose |
|------|-----|---------|
| Home | `/` | Redirects to dashboard/login |
| Login | `/login` | User authentication |
| Register | `/register` | New user signup |
| Dashboard | `/dashboard` | User overview & stats |
| Books | `/books` | Browse collection |
| Book Detail | `/books/[id]` | Book info & ideas |
| My Library | `/my-library` | Personal bookmarks |
| Leaderboard | `/leaderboard` | Top contributors |
| Donations | `/donations` | Support platform |
| Admin | `/admin` | Management panel |

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────┐
│         FRONTEND (Next.js 14)           │
│  Dashboard │ Books │ Ideas │ Reviews    │
└─────────────────────────────────────────┘
                    ↕ REST API
┌─────────────────────────────────────────┐
│          BACKEND (Go + Gin)             │
│  Handlers → Services → Repositories     │
│  • Success Score  • Matching            │
│  • Notifications  • Auth                │
└─────────────────────────────────────────┘
                    ↕
┌─────────────────────────────────────────┐
│       DATABASE (PostgreSQL 15)          │
│  15 Tables │ 19 Indexes │ Triggers      │
└─────────────────────────────────────────┘
```

---

## 🔧 Development

### Backend Development
```bash
cd backend
go mod tidy
air  # Hot reload
```

### Frontend Development
```bash
cd frontend
npm install
npm run dev  # http://localhost:3000
```

### Database
```bash
# Connect
docker exec -it online-library-db psql -U library_user -d online_library

# Backup
docker exec online-library-db pg_dump -U library_user online_library > backup.sql
```

---

## 🎨 Design System

### Colors
- **Background:** `#f4f1ea` (Old Paper)
- **Text:** `#2b2b2b` (Old Ink)
- **Secondary:** `#6b6b6b` (Old Grey)
- **Borders:** `#d4d4d4` (Old Border)

### Typography
- **Headings:** Bold, uppercase, wide tracking
- **Body:** Georgia, serif
- **Special:** Courier New (typewriter)

### Components
- Classic cards with 2-4px borders
- Offset shadows (4px, 4px)
- Stamp-like rotated badges
- Vintage inline labels
- Old paper texture background

---

## 🤝 Contributing

This is a community project! Ways to contribute:

1. **Code:** Submit PRs for new features
2. **Design:** Improve UI/UX
3. **Documentation:** Write guides
4. **Testing:** Report bugs
5. **Ideas:** Suggest improvements

---

## 📄 License

MIT License - Feel free to use, modify, and distribute.

---

## 🙏 Acknowledgments

Built with ❤️ for the community. Special thanks to everyone who believes in the power of shared knowledge and trust-based systems.

---

## 🎉 Get Started Now!

```bash
docker-compose up --build
```

**Access:** http://localhost:3000

**Start building your reading community! 📚✨**

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

## 📞 Support

- **Documentation:** See docs folder
- **Issues:** GitHub Issues
- **Discussions:** GitHub Discussions

---

**"Books are meant to travel, knowledge is meant to be shared."**

*Let's build a thriving reading community together! 📚✨*
