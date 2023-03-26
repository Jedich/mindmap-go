import {fileURLToPath, URL} from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default ({mode}) => {
  process.env = {...process.env, ...loadEnv(mode, process.cwd())}
  return defineConfig({
    server: {
      // host: true,
      proxy: {
        "^/api": {
          target: `http://${process.env.VITE_PROXY_IP}`,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
    },
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    }
  });
}
