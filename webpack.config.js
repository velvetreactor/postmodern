module.exports = {
  context: __dirname + '/src',
  entry: './app',
  output: {
    path: __dirname + '/dist',
    filename: 'bundle.js'
  },
  watch: true,
  module: {
    rules: [
      {
        test: /(\.js$|\.jsx$)/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['env', 'react']
          }
        }
      }
    ]
  }
}
