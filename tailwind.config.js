/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        './views/**/*.{templ,html,ts}',
        './src/components/**/*.{html,ts}',
        './node_modules/tw-elements/dist/js/**/*.js',
    ],
    plugins: [
        require('@tailwindcss/typography'),
        require("tw-elements/dist/plugin.cjs"),
    ],
    darkMode: "class",
    theme: {
        extend: {
            colors: {
                'base-100': '#F6F8FF',
                'base-200': '#fefdfd',
                'base-300': '#ffffff',
            },

        },
    }
}

