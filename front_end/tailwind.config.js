module.exports = {
  purge: ['./public/index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  plugins: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        // 使用element-plus的颜色变量
        primary: 'var(--el-color-primary)',
      },
    },
  },
  variants: {
    extend: {},
  },
  important: true,
}
