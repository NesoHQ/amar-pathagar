# AI Agent Development Guide

This guide helps AI coding agents (Cursor, Windsurf, Claude, etc.) understand and extend this project efficiently.

## 🎯 Project Context

**What:** Community book sharing platform tracking physical book circulation
**Why:** Enable transparent book sharing with history, queues, and analytics
**How:** Go backend + React frontend + PostgreSQL database

## 🏗️ Architecture Overview

### Backend Pattern (Go + Clean Architecture)
```
Request → Handler → Service → Repository → Database
         ↓
      Middleware (Auth, Logging)
         ↓
      DTO (Validation, Transformation)
```

### Frontend Pattern (React + TypeScript)
```
Component → API Service → Backend
    ↓
  Zustand Store (State)
    ↓
  TanStack Query (Server State)
```

## 📝 Code Conventions

### Backend (Go)

**File Naming:**
- `user_repository.go` (snake_case)
- `auth_service.go`
- `book_handler.go`

**Function Naming:**
- Public: `CreateBook()` (PascalCase)
- Private: `validateInput()` (camelCase)

**Error Handling:**
```go
if err != nil {
    return nil, fmt.Errorf("failed to create book: %w", err)
}
```

**Repository Pattern:**
```go
type BookRepository struct {
    db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
    return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
    // Implementation
}
```

**Service Pattern:**
```go
type BookService struct {
    bookRepo *repository.BookRepository
}

func NewBookService(bookRepo *repository.BookRepository) *BookService {
    return &BookService{bookRepo: bookRepo}
}

func (s *BookService) CreateBook(req dto.CreateBookRequest) (*models.Book, error) {
    // Business logic
}
```

**Handler Pattern:**
```go
type BookHandler struct {
    bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
    return &BookHandler{bookService: bookService}
}

func (h *BookHandler) Create(c *gin.Context) {
    var req dto.CreateBookRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
        return
    }
    
    book, err := h.bookService.CreateBook(req)
    if err != nil {
        c.JSON(http.StatusBadRequest, dto.Error(err.Error()))
        return
    }
    
    c.JSON(http.StatusCreated, dto.SuccessResponse("Book created", book))
}
```

### Frontend (React + TypeScript)

**File Naming:**
- `BookCard.tsx` (PascalCase for components)
- `useBooks.ts` (camelCase for hooks)
- `api.ts` (lowercase for utilities)

**Component Pattern:**
```typescript
interface BookCardProps {
  book: Book
  onSelect?: (book: Book) => void
}

export default function BookCard({ book, onSelect }: BookCardProps) {
  return (
    <div className="card">
      {/* Component content */}
    </div>
  )
}
```

**API Service Pattern:**
```typescript
export const booksAPI = {
  getAll: (params?: BookQueryParams) =>
    api.get('/books', { params }),
  
  getById: (id: string) =>
    api.get(`/books/${id}`),
  
  create: (data: CreateBookRequest) =>
    api.post('/books', data),
}
```

**Custom Hook Pattern:**
```typescript
export function useBooks(params?: BookQueryParams) {
  return useQuery({
    queryKey: ['books', params],
    queryFn: async () => {
      const response = await booksAPI.getAll(params)
      return response.data.data
    },
  })
}
```

## 🔧 Adding New Features

### Step 1: Database (if needed)
```sql
-- Add to database/init.sql or create migration
ALTER TABLE books ADD COLUMN new_field VARCHAR(255);
```

### Step 2: Model
```go
// backend/internal/models/book.go
type Book struct {
    // ... existing fields
    NewField string `json:"new_field"`
}
```

### Step 3: DTO
```go
// backend/internal/dto/book_dto.go
type CreateBookRequest struct {
    // ... existing fields
    NewField string `json:"new_field" binding:"required"`
}
```

### Step 4: Repository
```go
// backend/internal/repository/book_repository.go
func (r *BookRepository) Create(book *models.Book) error {
    query := `INSERT INTO books (..., new_field) VALUES (..., $n)`
    // Implementation
}
```

### Step 5: Service
```go
// backend/internal/services/book_service.go
func (s *BookService) CreateBook(req dto.CreateBookRequest) (*models.Book, error) {
    // Business logic
    book := &models.Book{
        NewField: req.NewField,
    }
    return book, s.bookRepo.Create(book)
}
```

### Step 6: Handler
```go
// backend/internal/handlers/book_handler.go
func (h *BookHandler) Create(c *gin.Context) {
    // Already handles new field via DTO
}
```

### Step 7: Route
```go
// backend/cmd/api/main.go
books := api.Group("/books")
{
    books.POST("", bookHandler.Create)
    books.GET("", bookHandler.List)
    books.GET("/:id", bookHandler.GetByID)
}
```

### Step 8: Frontend API
```typescript
// frontend/src/services/api.ts
export const booksAPI = {
  create: (data: CreateBookRequest) =>
    api.post('/books', data),
}
```

### Step 9: Frontend Component
```typescript
// frontend/src/pages/Books.tsx
const { mutate: createBook } = useMutation({
  mutationFn: (data: CreateBookRequest) => booksAPI.create(data),
  onSuccess: () => {
    queryClient.invalidateQueries(['books'])
  },
})
```

## 🎨 UI Components

### Reusable Components Location
```
frontend/src/components/
├── Layout.tsx          # Main layout with nav
├── Modal.tsx           # Modal wrapper
├── Button.tsx          # Button variants
├── Input.tsx           # Form inputs
└── Card.tsx            # Card container
```

### TailwindCSS Classes
```typescript
// Buttons
"btn btn-primary"           // Primary action
"btn btn-secondary"         // Secondary action

// Inputs
"input"                     // Standard input

// Cards
"card"                      // White card with shadow

// Layout
"space-y-6"                 // Vertical spacing
"grid grid-cols-1 md:grid-cols-3 gap-6"  // Responsive grid
```

## 🔐 Authentication Flow

### Protected Routes (Backend)
```go
api := router.Group("/api")
api.Use(middleware.AuthMiddleware(authService))
{
    api.GET("/books", bookHandler.List)
}

// Admin only
admin := api.Group("/admin")
admin.Use(middleware.AdminOnly())
{
    admin.POST("/books", bookHandler.Create)
}
```

### Protected Routes (Frontend)
```typescript
<Route 
  path="/books" 
  element={isAuthenticated ? <Books /> : <Navigate to="/login" />} 
/>
```

### Getting User Info
```go
// In handler
userID, _ := c.Get("user_id")
role, _ := c.Get("user_role")
```

```typescript
// In component
const { user } = useAuthStore()
```

## 📊 Database Queries

### Pagination Pattern
```go
func (r *BookRepository) List(page, pageSize int) ([]*models.Book, int, error) {
    offset := (page - 1) * pageSize
    
    // Get total count
    var total int
    r.db.QueryRow("SELECT COUNT(*) FROM books").Scan(&total)
    
    // Get paginated results
    rows, err := r.db.Query(
        "SELECT * FROM books ORDER BY created_at DESC LIMIT $1 OFFSET $2",
        pageSize, offset,
    )
    // ... process rows
    
    return books, total, nil
}
```

### Join Pattern
```go
query := `
    SELECT 
        b.*,
        u.id as holder_id,
        u.username as holder_username
    FROM books b
    LEFT JOIN users u ON b.current_holder_id = u.id
    WHERE b.id = $1
`
```

## 🧪 Testing Approach

### Backend Tests (TODO)
```go
// backend/internal/services/book_service_test.go
func TestCreateBook(t *testing.T) {
    // Setup
    // Execute
    // Assert
}
```

### Frontend Tests (TODO)
```typescript
// frontend/src/components/BookCard.test.tsx
describe('BookCard', () => {
  it('renders book information', () => {
    // Test implementation
  })
})
```

## 🚀 Common Tasks

### Add New API Endpoint
1. Create DTO in `backend/internal/dto/`
2. Add method to repository
3. Add method to service
4. Create handler function
5. Register route in `main.go`
6. Add to frontend API service
7. Create/update component

### Add New Page
1. Create component in `frontend/src/pages/`
2. Add route in `App.tsx`
3. Add navigation link in `Layout.tsx`
4. Implement page logic

### Add Database Table
1. Add CREATE TABLE to `database/init.sql`
2. Create model in `backend/internal/models/`
3. Create repository
4. Create service
5. Create handlers
6. Register routes

## 🐛 Debugging

### Backend Logs
```bash
docker-compose logs -f backend
```

### Frontend Logs
```bash
docker-compose logs -f frontend
```

### Database Queries
```bash
docker-compose exec postgres psql -U library_user -d online_library
```

## 📦 Dependencies

### Adding Go Package
```bash
cd backend
go get github.com/package/name
go mod tidy
```

### Adding npm Package
```bash
cd frontend
npm install package-name
```

## 🎯 Next Features to Implement

### Priority 1: Book Management
- [ ] Complete book CRUD in backend
- [ ] Book list page with pagination
- [ ] Book detail page
- [ ] Add/Edit book modal
- [ ] Book cover upload

### Priority 2: Circulation
- [ ] Assign book to user
- [ ] Start reading flow
- [ ] Finish reading flow
- [ ] Transfer to next reader

### Priority 3: Queue System
- [ ] Join queue endpoint
- [ ] Leave queue endpoint
- [ ] Queue display component
- [ ] Position tracking

## 💡 Tips for AI Agents

1. **Follow Existing Patterns**: Look at auth implementation as reference
2. **Maintain Clean Architecture**: Keep layers separated
3. **Use Type Safety**: Leverage Go and TypeScript types
4. **Handle Errors Properly**: Always return meaningful errors
5. **Keep It Simple**: Don't over-engineer
6. **Test Incrementally**: Build and test one feature at a time
7. **Update Documentation**: Keep this guide current

## 🔗 Key Files Reference

**Backend Entry:** `backend/cmd/api/main.go`
**Frontend Entry:** `frontend/src/main.tsx`
**Database Schema:** `database/init.sql`
**API Service:** `frontend/src/services/api.ts`
**Auth Store:** `frontend/src/stores/authStore.ts`
**Config:** `backend/internal/config/config.go`

## 📚 Useful Commands

```bash
# Start everything
make dev

# Backend only
cd backend && air

# Frontend only
cd frontend && npm run dev

# Database shell
docker-compose exec postgres psql -U library_user -d online_library

# View logs
docker-compose logs -f [service]

# Rebuild
docker-compose up --build

# Clean everything
make clean
```

---

**Remember:** This is a community project. Keep code clean, well-documented, and easy to understand. Happy coding! 🚀
