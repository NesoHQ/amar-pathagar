# Amar Pathagar - Quick Reference Card

## 🚀 Quick Start Commands

```bash
# Start everything with Docker
docker-compose up --build

# Access the app
open http://localhost:5173
```

## 📊 Success Score Quick Reference

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

**Minimum Score to Request Books: 20**

## 🎯 Matching Algorithm

```
Priority Score = (Success Score × 0.4) + (Interest Match × 0.3) + (Distance × 0.3)
```

**Factors:**
- **40%** - Your success score (reputation)
- **30%** - Interest match (topics you like)
- **30%** - Distance (closer is better)

## 🔖 Bookmark Types

1. **Like** ❤️ - Simple favorite
2. **Bookmark** 🔖 - Save for later
3. **Priority** ⭐ - High-priority want (0-10 scale)

## 🏆 Leaderboard Categories

1. 📚 **Top Readers** - Most books read
2. 🤝 **Top Sharers** - Most books shared
3. 🎁 **Top Donors** - Most donations
4. ⭐ **Highest Scores** - Best reputation
5. 💡 **Top Idea Writers** - Most ideas posted

## 📱 Main Pages

| Page | URL | Purpose |
|------|-----|---------|
| Dashboard | `/dashboard` | Overview & stats |
| Books | `/books` | Browse all books |
| My Library | `/my-library` | Your books |
| Leaderboard | `/leaderboard` | Top users |
| Donations | `/donations` | Support platform |
| Profile | `/users/:id` | User profiles |
| Admin | `/admin` | Admin panel |

## 🔌 Key API Endpoints

### Authentication
```bash
POST /api/auth/register
POST /api/auth/login
GET  /api/me
```

### Books
```bash
GET    /api/books
GET    /api/books/:id
POST   /api/books (admin)
POST   /api/books/:id/request
```

### Social Features
```bash
POST   /api/ideas
POST   /api/ideas/:id/vote
POST   /api/reviews
POST   /api/bookmarks
```

### Community
```bash
GET    /api/leaderboard
POST   /api/donations
GET    /api/users/:id/profile
```

## 🎨 Color Scheme

- **Background**: Black (#000000)
- **Secondary**: Dark Grey (#1a1a1a, #2a2a2a)
- **Text**: Off-white (#f5f5f5)
- **Accent**: Grey (#6b7280)
- **Highlights**: White (#ffffff)

## 📝 Book Statuses

- **Available** 🟢 - Ready to borrow
- **Reading** 📖 - Someone has it
- **Reserved** 🔒 - Assigned to someone
- **Requested** 📬 - Multiple people want it

## 🔔 Notification Types

- 📚 Book available
- ✅ Request approved
- ❌ Request rejected
- ⏰ Return reminder
- ⭐ Review received
- 📊 Success score changed
- 👍 Idea voted

## 🛠️ Development Commands

### Backend
```bash
cd backend
go run cmd/api/main.go    # Run
air                        # Hot reload
go test ./...             # Test
go build cmd/api/main.go  # Build
```

### Frontend
```bash
cd frontend
npm run dev               # Development
npm run build            # Production build
npm run preview          # Preview build
```

### Database
```bash
# Connect
psql -h localhost -U library -d online_library

# Backup
pg_dump -U library online_library > backup.sql

# Restore
psql -U library online_library < backup.sql
```

## 🐛 Troubleshooting

### Backend won't start
```bash
cd backend
go mod tidy
go clean
```

### Frontend errors
```bash
cd frontend
rm -rf node_modules
npm install
```

### Database issues
```bash
docker-compose down -v
docker-compose up -d
```

### CORS errors
Check `cmd/api/main.go` - ensure frontend URL in AllowOrigins

## 📊 Database Tables

**Core Tables:**
- users, books, reading_history

**Request System:**
- book_requests, waiting_queue

**Social Features:**
- reading_ideas, idea_votes, user_reviews

**Community:**
- donations, user_interests, user_bookmarks

**System:**
- notifications, success_score_history, audit_logs

## 🔐 Environment Variables

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=library
DB_PASSWORD=library123
DB_NAME=online_library
JWT_SECRET=change-me-in-production
SERVER_PORT=8080
```

### Frontend (.env)
```env
VITE_API_URL=http://localhost:8080
```

## 📈 Success Score Thresholds

- **0-19**: ❌ Cannot request books
- **20-49**: ⚠️ Low priority in matching
- **50-99**: ✅ Normal priority
- **100-149**: ⭐ Good standing
- **150+**: 🌟 Excellent reputation

## 🎯 User Roles

### Member (Default)
- Browse books
- Request books
- Post ideas
- Write reviews
- Make donations

### Admin
- All member permissions
- Add/edit/delete books
- Manage requests
- Adjust success scores
- View audit logs

## 📚 Best Practices

### For Users
1. Return books on time (+10 points)
2. Write thoughtful reviews (+5 points)
3. Share reading ideas (+3 points)
4. Keep books in good condition
5. Communicate clearly

### For Admins
1. Review requests fairly
2. Monitor success scores
3. Handle disputes promptly
4. Add quality books
5. Engage with community

## 🔗 Useful Links

- **Documentation**: See `AMAR_PATHAGAR_FEATURES.md`
- **Setup Guide**: See `IMPLEMENTATION_GUIDE.md`
- **Project Summary**: See `PROJECT_SUMMARY.md`
- **Completion Status**: See `COMPLETION_SUMMARY.md`

## 💡 Pro Tips

1. **Build Reputation Early**: Post ideas and return books on time
2. **Add Interests**: Better matching for book requests
3. **Set Location**: Get priority for nearby books
4. **Use Priority Bookmarks**: Track books you really want
5. **Engage with Ideas**: Upvote quality content
6. **Donate**: Big reputation boost (+20 for books)
7. **Write Reviews**: Help build trust in community
8. **Check Leaderboard**: See top contributors
9. **Monitor Score**: Keep above 20 to request books
10. **Be Active**: Regular engagement builds reputation

## 🎉 Quick Wins

Want to boost your score fast?
1. Donate a book (+20)
2. Post 3 reading ideas (+9)
3. Return a book on time (+10)
4. Write a positive review (+5 to recipient)
5. Get 5 upvotes on ideas (+5)

**Total: +49 points in one day!**

---

**Need Help?** Check the full documentation or open an issue.

**Ready to Share?** Start with `docker-compose up --build`

📚✨ Happy Reading!
