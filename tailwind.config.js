/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: ["./templates/*.html"],
  theme: {
    extend: {},
  },
  plugins: [],
  browserslist: {
    // Add Firefox version 75 and above
    'firefox >= 75'
  },
}

