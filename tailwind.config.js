/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './views/**/*.{templ,html,ts}',
        './node_modules/tw-elements/dist/js/**/*.js',
    ],
    plugins: [
        require('@tailwindcss/typography'),
        require("tw-elements/dist/plugin.cjs"),
    ],
    darkMode: "class",
    theme: {
        extend: {},
    }
}

