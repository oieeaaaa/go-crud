/** type {import('tailwindcss/tailwind-config').TailwindConfig} */
module.exports = {
  content: [
    './views/**/*.{html,css}',
    './client/**/*.css'
  ],
  theme: {
    extend: {
      colors: {
        'primary': '#FF6363',
        'secondary': '#4C4C6D',
      },
      container: {
        center: true,
      }
    },
    fontFamily: {
      'sans': ['Titillium Web', 'sans-serif'],
    }
  },
}
