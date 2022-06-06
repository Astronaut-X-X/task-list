import { defineConfig } from 'vite'
import { createVuePlugin } from 'vite-plugin-vue2'

const HOST = "localhost"
const PORT = "8080"

export default () => {
  return defineConfig({
    base: "./",
    server: {
      host: HOST,
      port: PORT,
    },
    plugins: [
      createVuePlugin(),
    ],
  })
}
