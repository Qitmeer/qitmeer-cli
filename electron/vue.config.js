module.exports = {
  publicPath: '/app/',
  outputDir: 'dist',
  assetsDir: 'assets',
  css: {
    loaderOptions: {
      sass: {
        data: '@import "@/scss/settings.scss";'
      }
    }
  }
}
