import { defineConfig } from "vite";

export default defineConfig({
  base: '/', // Use relative paths for GitHub Pages deployment
  server: {
    port: 3000, // Development server port
  },
  build: {
    outDir: "dist", // Production build directory
  },
});
