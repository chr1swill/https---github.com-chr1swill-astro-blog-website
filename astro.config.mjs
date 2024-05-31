import { defineConfig } from "astro/config";

export default defineConfig({
	routes: [
		{ path: "/", component: "./src/pages/index.astro" },
		{ path: "/blog/*", component: "./src/pages/blog/[post].md" },
		// Add more routes as needed
	],
	mount: {
		public: "/",
		src: "/_dist_",
		// Add more mount points as needed
	},
	optimize: {
		bundle: true,
		minify: true,
		target: "es2018",
	},
	build: {
		sitemap: true,
	},
	// Add any additional configuration options here
});
