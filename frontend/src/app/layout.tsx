import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Amar Pathagar - Community Library',
  description: 'A trust-based book sharing platform',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
