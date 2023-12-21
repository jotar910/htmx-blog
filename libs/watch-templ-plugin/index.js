const { exec } = require('child_process');

class WatchTemplChangesPlugin {
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

module.exports = WatchTemplChangesPlugin;
