/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      colors: {
        'old-paper': '#f4f1ea',
        'old-ink': '#2b2b2b',
        'old-grey': '#6b6b6b',
        'old-border': '#d4d4d4',
      },
      fontFamily: {
        'serif': ['Georgia', 'Cambria', 'Times New Roman', 'serif'],
        'mono': ['Courier New', 'Courier', 'monospace'],
      },
    },
  },
  plugins: [],
}
