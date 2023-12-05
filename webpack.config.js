const path = require('path');
const { exec } = require('child_process');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

class RunCommandOnFileChangePlugin {
    constructor() {
        this.firstTime = true;
        this.emptyMap = new Map();
        this.previousTimestamps = new Map();
    }

    async apply(compiler) {
        console.debug('generating go templates');
        await this.generateTempl();

        compiler.hooks.watchRun.tapPromise('RunCommandOnTemplChangePlugin', async (compilation) => {
            const currentTimestamps = compilation.fileTimestamps || this.emptyMap;
            const changedFiles = new Set();

            for (const [path, timestamp] of currentTimestamps) {
                if (this.previousTimestamps.get(path) !== timestamp) {
                    changedFiles.add(path);
                }
            }

            this.previousTimestamps = new Map(currentTimestamps);

            if (this.firstTime) {
                this.firstTime = false;
                return;
            }

            const hasTemplChanges = !!changedFiles.size && Array.from(changedFiles).some(file => file.endsWith('.templ'));

            console.debug('generating go templates:', hasTemplChanges);

            if (hasTemplChanges) {
                await this.generateTempl();
            }
        });
    }

    generateTempl() {
        return new Promise((resolve) => {
            exec('templ generate', (err, stdout, stderr) => {
                if (err) {
                    console.error(`Error: ${stderr}`);
                } else {
                    console.log(`stdout: ${stdout}`);
                }
                resolve();
            });
        });
    }
}

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
        new RunCommandOnFileChangePlugin()
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
