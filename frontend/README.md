# Amar Pathagar - Frontend

Classic, old-school styled Next.js frontend for the Amar Pathagar community library platform.

## Features

- **Next.js 14** with App Router
- **TypeScript** for type safety
- **Classic Design** - Grey, old-school aesthetic
- **Tailwind CSS** for styling
- **Zustand** for state management
- **Axios** for API calls

## Getting Started

### Development

```bash
# Install dependencies
npm install

# Run development server
npm run dev

# Open http://localhost:3000
```

### Production

```bash
# Build for production
npm run build

# Start production server
npm start
```

## Environment Variables

Create a `.env.local` file:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
```

## Pages

- `/` - Home (redirects to dashboard or login)
- `/login` - Login page
- `/register` - Registration page
- `/dashboard` - User dashboard
- `/books` - Browse books
- `/books/[id]` - Book details
- `/my-library` - User's bookmarks and history
- `/leaderboard` - Top contributors
- `/donations` - Make and view donations
- `/admin` - Admin panel (admin only)

## Design Philosophy

- **Classic Typography** - Serif fonts, uppercase headings
- **Old Paper Texture** - Subtle background pattern
- **Bold Borders** - 2-4px borders throughout
- **Minimal Colors** - Black, grey, off-white palette
- **Stamp-like Elements** - Rotated badges and labels
- **Typewriter Feel** - Monospace for special elements

## API Integration

All API calls go through `src/lib/api.ts` which handles:
- Authentication tokens
- Request/response interceptors
- Error handling
- Automatic logout on 401

## State Management

Uses Zustand for:
- Authentication state
- User profile
- Token management
- Persistent storage

## Styling

Custom Tailwind classes:
- `.classic-card` - Standard card with border and shadow
- `.classic-button` - Primary button style
- `.classic-button-secondary` - Secondary button style
- `.classic-input` - Form input style
- `.classic-heading` - Section heading with underline
- `.stamp` - Rotated badge style
- `.vintage-badge` - Inline badge style

## Development Notes

- All pages use `'use client'` directive for client-side rendering
- Authentication state loads from localStorage on mount
- Protected routes redirect to login if not authenticated
- Admin routes check user role before rendering
