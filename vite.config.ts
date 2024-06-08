import { defineConfig } from 'vite';
import { resolve } from 'path';
import tsconfigPaths from 'vite-tsconfig-paths';
import { chunkSplitPlugin } from 'vite-plugin-chunk-split';

export default defineConfig({
  plugins: [tsconfigPaths(), chunkSplitPlugin()],
  base: '/public/', // Ensure base path is set correctly
  build: {
    manifest: true,
    outDir: 'public',
    rollupOptions: {
      input: resolve(__dirname, 'views/scripts/index.ts'),
    },
  },
});