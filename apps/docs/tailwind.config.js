/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class', // Enable class-based dark mode
  content: [
    './src/**/*.{html,js,svelte,ts}'
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['Roboto', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'sans-serif'],
        'headers': ['Leksa Sans', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'sans-serif'],
        'brand': ['Leksa Sans Demi', 'Leksa Sans', '-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'sans-serif'],
        'mono': ['Roboto Mono', 'SF Mono', 'Monaco', 'Cascadia Code', 'Consolas', 'Courier New', 'monospace']
      }
    }
  }
};