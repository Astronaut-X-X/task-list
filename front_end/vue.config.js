module.exports = {
  devServer: {
    open:true,
    host:'localhost',
    port:8080,
    https:false,
    // 跨域请求
    proxy: {
      '/api': {
        target: 'http://localhost:8080/api',
        ws:true,
        changeOrigin: true,
        pathRewrite: {
          '^api/': ''
        }
      }
    }
  }
}  