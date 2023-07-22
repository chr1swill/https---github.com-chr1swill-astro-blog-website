import { defineConfig } from 'astro/config';

export default defineConfig({

    routes: [
        { path: '/', component: './src/pages/index.astro' },
        { path: '/blog/*', component: './src/pages/blog/[post].md' },
        // Add more routes as needed
      ],

    });