const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const WatchTemplChangesPlugin = require('watch-templ-plugin');

module.exports = {
    entry: {
        bundle: './src/index.ts',
        'tw-elements': './src/tw-elements.ts'
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader',
                exclude: /node_modules/,
            },
            {
                test: /\.css|\.scss|\.html|\.templ$/,
                use: [
                    MiniCssExtractPlugin.loader, // 3. Extract CSS into files
                    'css-loader',                // 2. Turn CSS into CommonJS
                    'sass-loader',               // 1. Compile SCSS into CSS
                    'postcss-loader',            // 0. Execute PostCSS
                ],
                exclude: /node_modules/,
            }
        ],
    },
    resolve: {
        extensions: ['.tsx', '.ts', '.js', '.scss', '.templ'],
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: '[name].css',
        }),
        new WatchTemplChangesPlugin()
    ],
    output: {
        filename: '[name].js',
        path: path.resolve(__dirname, 'dist'),
    },
    watch: false,
    watchOptions: {
        ignored: ['node_modules', 'dist'],
    }
};
