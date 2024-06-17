import { defineConfig } from 'vite';
import { resolve } from 'path';

export default defineConfig({
  base: '/public/', // Ensure base path is set correctly
  build: {
    manifest: true,
    outDir: 'public',
    rollupOptions: {
      input: {
        main: resolve(__dirname, 'views/scripts/index.ts'),
        moduleLoader: resolve(__dirname, 'views/scripts/module-loader.ts'),
      },
      output: {
        entryFileNames: (chunk) => {
          // Name the module-loader file without a hash
          if (chunk.name === 'moduleLoader') {
            return 'assets/module-loader.js';
          }
          // Use hashes for other entry files
          return 'assets/[name]-[hash].js';
        },
        chunkFileNames: 'assets/[name]-[hash].js',
        assetFileNames: 'assets/[name]-[hash][extname]'
      },
    },
  },
});