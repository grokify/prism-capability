import { defineConfig } from 'vite';
import { resolve } from 'path';

export default defineConfig({
  build: {
    lib: {
      entry: resolve(__dirname, 'src/capability-stack.ts'),
      name: 'CapabilityStack',
      fileName: 'capability-stack',
      formats: ['es', 'umd']
    },
    rollupOptions: {
      external: [],
      output: {
        globals: {}
      }
    }
  }
});
