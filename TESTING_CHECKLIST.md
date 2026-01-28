# Amar Pathagar - Testing Checklist

## 🧪 Complete Testing Guide

Use this checklist to verify all features are working correctly.

---

## ✅ Setup & Installation

- [ ] Docker and Docker Compose installed
- [ ] `docker-compose up --build` runs successfully
- [ ] Frontend accessible at http://localhost:5173
- [ ] Backend accessible at http://localhost:8080
- [ ] Database accessible at localhost:5432
- [ ] No console errors on page load

---

## 🔐 Authentication & Authorization

### Registration
- [ ] Can register new user with username, email, password
- [ ] Email validation works
- [ ] Password requirements enforced
- [ ] Duplicate username/email rejected
- [ ] User starts with 100 success score
- [ ] User role defaults to 'member'

### Login
- [ ] Can login with username and password
- [ ] Can login with email and password
- [ ] Invalid credentials rejected
- [ ] JWT token received and stored
- [ ] Token included in subsequent requests
- [ ] Redirected to dashboard after login

### Authorization
- [ ] Protected routes require authentication
- [ ] Unauthenticated users redirected to login
- [ ] Admin routes only accessible to admins
- [ ] Token refresh works correctly

---

## 👤 User Profile System

### Public Profile
- [ ] Can view other users' public profiles
- [ ] Profile shows username, full name, avatar
- [ ] Success score displayed
- [ ] Books shared/received count shown
- [ ] Reviews received count shown
- [ ] Ideas posted count shown
- [ ] Donor badge shown if applicable
- [ ] Bio displayed if set

### Profile Update
- [ ] Can update full name
- [ ] Can update bio
- [ ] Can upload/change avatar
- [ ] Can set location (lat/lng/address)
- [ ] Changes saved successfully
- [ ] Changes reflected immediately

### Interests
- [ ] Can add interests/topics
- [ ] Interests saved to database
- [ ] Duplicate interests prevented
- [ ] Interests used in matching algorithm

---

## 📚 Book Management

### Book Listing
- [ ] All books displayed on /books page
- [ ] Book covers shown
- [ ] Book titles, authors visible
- [ ] Book status indicated (Available, Reading, etc.)
- [ ] Current holder shown if applicable
- [ ] Pagination works (if implemented)

### Book Details
- [ ] Can view individual book details
- [ ] Full description shown
- [ ] Category and tags displayed
- [ ] Topics listed
- [ ] Current holder information
- [ ] Availability status clear
- [ ] Request button shown if available

### Book Search & Filter
- [ ] Can search by title
- [ ] Can search by author
- [ ] Can filter by category
- [ ] Can filter by status
- [ ] Can filter by topics
- [ ] Search results accurate

### Admin Book Management
- [ ] Admin can add new books
- [ ] Admin can edit book details
- [ ] Admin can delete books
- [ ] Admin can mark books as donated
- [ ] Physical code uniqueness enforced

---

## 🔖 Bookmarks & Priorities

### Like Books
- [ ] Can like a book
- [ ] Like button toggles correctly
- [ ] Liked books saved
- [ ] Can view all liked books
- [ ] Can unlike a book

### Bookmark Books
- [ ] Can bookmark a book
- [ ] Bookmark button toggles correctly
- [ ] Bookmarked books saved
- [ ] Can view all bookmarks
- [ ] Can remove bookmark

### Priority List
- [ ] Can add book to priority list
- [ ] Can set priority level (0-10)
- [ ] Priority books sorted by level
- [ ] Can view priority list
- [ ] Can remove from priority list
- [ ] Can update priority level

### Availability Check
- [ ] Can see if liked book is available
- [ ] Availability status updates in real-time
- [ ] Current holder shown for unavailable books

---

## 🔄 Book Request Flow

### Request Book
- [ ] Can request available book
- [ ] Success score checked (minimum 20)
- [ ] Low score users blocked from requesting
- [ ] Request created successfully
- [ ] Book status changes to 'requested'
- [ ] Notification sent to current holder

### Request Queue
- [ ] Multiple users can request same book
- [ ] Queue position calculated
- [ ] Priority score computed correctly
- [ ] Distance calculated if location set
- [ ] Interest match score calculated
- [ ] Can view position in queue

### Request Management
- [ ] Can cancel own request
- [ ] Admin can view all requests
- [ ] Admin can approve request
- [ ] Admin can reject request
- [ ] Approved user gets book
- [ ] Rejected users notified
- [ ] Due date set on approval

---

## 🤝 Matching Algorithm

### Priority Calculation
- [ ] Success score weighted at 40%
- [ ] Interest match weighted at 30%
- [ ] Distance weighted at 30%
- [ ] Priority score calculated correctly
- [ ] Highest priority selected first

### Distance Calculation
- [ ] Haversine formula used
- [ ] Distance in kilometers
- [ ] Closer users get higher score
- [ ] Works with missing location data

### Interest Matching
- [ ] User interests compared to book topics
- [ ] Match score calculated (0-100)
- [ ] Higher match = higher priority
- [ ] Works with no interests set

### Best Match Selection
- [ ] Highest priority score wins
- [ ] Tie-breaker: earliest request
- [ ] Selection algorithm fair
- [ ] All factors considered

---

## 📊 Success Score System

### Score Display
- [ ] Current score shown on profile
- [ ] Score visible to other users
- [ ] Score history accessible
- [ ] Score changes logged

### Score Increases
- [ ] Return on time: +10 points
- [ ] Positive review: +5 points
- [ ] Post idea: +3 points
- [ ] Idea upvoted: +1 point
- [ ] Donate book: +20 points
- [ ] Donate money: +10 points

### Score Decreases
- [ ] Return late: -15 points
- [ ] Negative review: -10 points
- [ ] Idea downvoted: -1 point
- [ ] Lost book: -50 points

### Score History
- [ ] All changes recorded
- [ ] Reason for change shown
- [ ] Timestamp recorded
- [ ] Reference ID included
- [ ] History viewable by user

### Score Thresholds
- [ ] Minimum 20 to request books
- [ ] Low score warning shown
- [ ] Score affects matching priority
- [ ] Admin can adjust scores

---

## 💡 Reading Ideas & Knowledge

### Create Idea
- [ ] Can post idea after reading
- [ ] Title and content required
- [ ] Linked to specific book
- [ ] Author credited
- [ ] +3 points awarded
- [ ] Idea saved successfully

### View Ideas
- [ ] Can view ideas for a book
- [ ] Ideas sorted by votes
- [ ] Author name shown
- [ ] Vote counts displayed
- [ ] Timestamps shown

### Vote on Ideas
- [ ] Can upvote idea
- [ ] Can downvote idea
- [ ] Can change vote
- [ ] Cannot vote on own idea
- [ ] Vote counts update immediately
- [ ] Author score updated (+1 or -1)

### Idea Quality
- [ ] High-quality ideas upvoted
- [ ] Low-quality ideas downvoted
- [ ] Net score calculated (upvotes - downvotes)
- [ ] Ideas with negative score visible

---

## ⭐ Review System

### Write Review
- [ ] Can review user after book exchange
- [ ] Behavior rating (1-5 stars)
- [ ] Book condition rating (1-5 stars)
- [ ] Communication rating (1-5 stars)
- [ ] Written comment optional
- [ ] Review saved successfully

### Review Impact
- [ ] Average rating calculated
- [ ] High rating (≥4): +5 points to reviewee
- [ ] Low rating (<3): -10 points to reviewee
- [ ] Score updated immediately
- [ ] Reviewee notified

### View Reviews
- [ ] Can view received reviews
- [ ] All ratings displayed
- [ ] Comments shown
- [ ] Reviewer name visible
- [ ] Timestamp shown
- [ ] Average ratings calculated

---

## 🎁 Donation System

### Book Donations
- [ ] Can donate book to system
- [ ] Book marked as donated
- [ ] Donor credited
- [ ] +20 points awarded
- [ ] Donor badge added
- [ ] Donation recorded in history

### Money Donations
- [ ] Can make financial donation
- [ ] Amount and currency specified
- [ ] Optional message included
- [ ] +10 points awarded
- [ ] Donor badge added
- [ ] Public/private option works

### Donation List
- [ ] Public donations displayed
- [ ] Donor names shown
- [ ] Amounts shown (if public)
- [ ] Messages displayed
-