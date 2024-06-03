/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./src/**/*.{html,js,templ}"],
	theme: {
		extend: {
			colors: {
				text: "var(--text)",
				"text-op": "var(--text-op)",
				primary: "var(--primary)",
				secondary: "var(--secondary)",
				surface: "var(--surface)",
				accent: "var(--accent)",
				background: "var(--background)",
			},
			fontFamily: {
				base: "var(--font-family)",
			},
		},
	},
	plugins: [],
};
