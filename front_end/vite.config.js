import { defineConfig } from 'vite'
import { createVuePlugin } from 'vite-plugin-vue2'
import WindiCSS from 'vite-plugin-windicss'

export default () => {
  return defineConfig({
    root: './',
    plugins: [
      createVuePlugin(),
      WindiCSS(),
    ],
    server: {
      port: '3000',
      proxy: {
        // 字符串简写写法
        '/api': 'http://localhost:8081',
      }
    }
  })
}
