var path = require('path');
var webpack = require('webpack');
const UglifyJsPlugin = require('uglifyjs-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const HtmlWebpackPlugin = require('html-webpack-plugin');
// const WebpackMd5Hash = require('webpack-md5-hash');

// Кэширование при сборке
var HardSourceWebpackPlugin = require('hard-source-webpack-plugin');

// копирование файлов из ХХХ ту ААА
//const CopyPlugin = require('copy-webpack-plugin');

// очистить папку перед новой сборкой

const {CleanWebpackPlugin} = require('clean-webpack-plugin');


module.exports = {
  watch: true,

  entry: {
    build: path.join(__dirname, 'staticDev', 'js', 'main.js'),
  },
  output: {
    path: path.resolve(__dirname, 'static/dist'),
    publicPath: process.env.NODE_ENV === 'development' ? '/static/dist/' : './static/dist/',
    filename: '[name].js'
  },

  plugins: [

    new MiniCssExtractPlugin({
      // Options similar to the same options in webpackOptions.output
      // all options are optional
      filename: '[name].css' ,
      // chunkFilename: '[id].css',
      // ignoreOrder: false, // Enable to remove warnings about conflicting order
    }),
    // new WebpackMd5Hash(),
  ],


  devServer: {
    host: 'localhost',
    port: 8012,
    hot: true,
    // inline: true,

    // watchOptions: {
    //   //poll: true
    // },
    // disableHostCheck: true,
    // allowedHosts: [
    //   'localhost:8000',
    //   'localhost:9001',
    //   'http://93d7bf6c.ngrok.io',
    // ],

  },

  module: {
    rules: [
      {
        test: /\.html$/,
        loader: 'ejs-loader',
        query: {
          interpolate: '<%=([\\s\\S]+?)%>'
        }
      },
      {
        test: require('path').resolve(__dirname, 'node_modules/leader-line/'),
        loader: 'skeleton-loader',
        options: {
          procedure: content => `${content} export default LeaderLine`
        }
      },
      {
        test: /\.vue$/,
        loader:
          'vue-loader',
        options:
          {
            //loaders: {}
            // other vue-loader options go here
          }
      },
      {
        test: /\.js$/,
        use: [
          {
            loader: 'cache-loader',
            options: {
              cacheDirectory: path.resolve(
                __dirname,
                'node_modules/.cache/cache-loader'
              ),
            },
          },
          {
            loader: 'babel-loader',
            options: {
              cacheDirectory: true
            }
          }
        ],
        exclude:
          /node_modules/
      },
      {
        test: /\.woff(2)?(\?v=[0-9]+\.[0-9]+\.[0-9]+)?$/,
        loader:
          "url-loader?limit=10000&mimetype=application/font-woff"
      }
      ,
      {
        test: /\.(mp3|svg|ttf|woff2|woff|eot)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
        loader:
          'file-loader',
        options:
          {
            name: '[name].[ext]'
            // name: '[name].[ext]?[hash]'
          }
      }
      ,
      {
        test: /\.(jpg|jpeg|png|gif)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
        loader:
          'file-loader',
        options:
          {
            name: 'img/[name].[ext]'
            // name: '[name].[ext]?[hash]'
          }
      }
      ,

      {
        test: /\.css$/i,
        use:
          [
            'vue-style-loader',
            process.env.NODE_ENV !== 'production' ? 'style-loader' : MiniCssExtractPlugin.loader,

            {
              loader: 'cache-loader',
              options: {
                cacheDirectory: path.resolve(
                  __dirname,
                  'node_modules/.cache/cache-loader'
                ),
              },
            },
            //'style-loader',
            {
              loader: 'css-loader',
              options: {sourceMap: true,}
            },
            {
              loader: 'postcss-loader',
              options: {
                sourceMap: true,
                config: {
                  path: path.join(__dirname, 'staticDev', 'config', 'postcss.config.js')
                }
              }
            },
          ],
      }
      ,
      {
        test: /\.scss$/,
        use:
          [
            //'style-loader',
            // {
            //     loader: MiniCssExtractPlugin.loader,
            //     options: {
            //         esModule: true,
            //     },
            // },
            // style-loader нужен чтобы на горячую обновлять Scss
            // process.env.NODE_ENV !== 'production' ? 'style-loader' : MiniCssExtractPlugin.loader,
            'style-loader',
            MiniCssExtractPlugin.loader,
            {
              loader: 'cache-loader',
              options: {
                cacheDirectory: path.resolve(
                  __dirname,
                  'node_modules/.cache/cache-loader'
                ),
              },
            },
            {
              loader: 'css-loader',
              options: {
                //modules: true,
                sourceMap: true,
                //importLoaders: 2,
                // localIdentName: '[name]__[local]___[hash:base64:5]'
              }
            },
            {
              loader: 'postcss-loader',
              options: {
                sourceMap: true,
                config: {
                  path: path.join(__dirname, 'staticDev', 'config', 'postcss.config.js')
                }
              }
            },
            {
              loader: 'sass-loader',
              options: {sourceMap: true}
            },
            // {
            //   loader: 'sass-resources-loader',
            //   options: {
            //     resources: [
            //       './src/main/resources/staticDev/style/scss/variables/variables.scss',
            //     ]
            //   },
            // },
          ],
      }
      ,
      {
        test: /\.sass$/,
        use:
          [
            'vue-style-loader',
            'css-loader',
            {
              loader: 'sass-loader',

              // Requires sass-loader@^8.0.0
              options: {
                implementation: require('sass'),
                sassOptions: {
                  fiber: require('fibers'),
                  indentedSyntax: true // optional
                },
              },
            },
          ],
      }
      ,
    ]
  },
  resolve: {
    modules: [
      path.join(__dirname, 'staticDev', 'js'),
      path.join(__dirname, 'node_modules')
    ],
    alias:
      {
        'vue$':
          'vue/dist/vue.esm.js',
        // myVariables:
        //   path.join(__dirname, 'src', 'main', 'resources', 'staticDev', 'style', 'scss', 'variables', 'variables.scss'),
      }
    ,
    extensions: ['*', '.js', '.vue', '.json']
  }
  ,

  performance: {
    hints: false
  }
  ,
  devtool: '#eval-source-map',
  optimization:
    {
      minimizer: [
        // we specify a custom UglifyJsPlugin here to get source maps in production

      ]
    }
};

if (process.env.NODE_ENV === 'production') {
  module.exports.optimization.minimizer = (module.exports.optimization.minimizer || []).concat(
    [
      new UglifyJsPlugin({
        cache: true,
        parallel: true,
        uglifyOptions: {
          compress: false,
          ecma: 6,
          mangle: true
        },
        sourceMap: true
      })
    ]
  );
}

if (process.env.NODE_ENV === 'production') {
  module.exports.devtool = '#source-map';
  // http://vue-loader.vuejs.org/en/workflow/production.html
  module.exports.plugins = (module.exports.plugins || []).concat([
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"'
      }
    }),
    // new webpack.optimize.UglifyJsPlugin({
    //     sourceMap: true,
    //     compress: {
    //         warnings: false
    //     }
    // }),
    new webpack.LoaderOptionsPlugin({
      minimize: true
    })
  ])
}
