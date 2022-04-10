import { exit } from 'process';
import { exec } from 'child_process';
import { watch, rm } from 'fs';
import { extname } from 'path';
import goSvelte from './gosvelte.js';

const goSvelteCfg = goSvelte();
const mainProc = process;
const argv = process.argv.slice(2);
let restartLock = false;
let svelteProc;
let svelteWtch;
let goProc;

(function () {
    switch (argv[0]) {
        case '-install':
            goProc = goSvelteExec("go mod tidy", nil);
            break;
        case '-dev':
            switch (argv[1]) {
                case '-go':
                    console.log("Starting dev go...");
                    goProc = goSvelteExec('go run main.go', nil);
                    break;
                case '-svelte':
                    console.log("Starting dev svelte...");
                    goSvelteClean();
                    svelteProc = goSvelteExec('rollup -c -w', goSvelteWatch);
                    break;
                default:
                    console.log("Starting dev...");
                    goSvelteClean();
                    svelteProc = goSvelteExec('rollup -c -w', goSvelteWatch);
                    goProc = goSvelteExec('go run main.go', nil);
                    break;
            }
            break;
        case '-build':
            console.log("Building...");
            goSvelteClean();
            svelteProc = goSvelteExec("rollup -c", nil);
            goProc = goSvelteExec("go build main.go", nil);
            break;
        case '-start':
            console.log("Starting prod...");
            goProc = goSvelteExec("start ./main", nil);
            setTimeout(function () { exit(); }, 1000);
            break;
        default:
            console.error("Command line flag not found! for more info please visit https://github.com/moadkey/gosvelte");
            exit();
    }
    mainProc.on('SIGTERM', goSvelteExit);
    mainProc.on('SIGKILL', goSvelteExit);
    mainProc.on('exit', goSvelteExit);
})();

function goSvelteExec(cmd, callback) {
    const proc = exec(cmd);
    proc.stdout.on("data", data => {
        console.log(`${data}`);
    });
    proc.stderr.on("data", data => {
        console.log(`${data}`);
    });
    proc.on('error', (error) => {
        console.error(`${error.message}`);
    });
    proc.on('close', (code) => {
        if (cmd.includes('rollup -c -w')) {
            goSvelteExit();
        }
    });
    callback();
    return proc;
}

function goSvelteWatch() {
    svelteWtch = watch(goSvelteCfg.SvelteWorkspacePath, (eventType, fileName) => {
        if (eventType == "rename" && extname(fileName).toLowerCase() === goSvelteCfg.SvelteExtension && !restartLock) {
            restartLock = !restartLock;
            console.log("Svelte files have been changed, restarting...\n");
            svelteProc.kill(0);
            //goSvelteClean();
            setTimeout(function () { svelteProc = goSvelteExec('rollup -c -w', function () { restartLock = !restartLock; }); }, 1000);
        }
    });
}

function goSvelteClean() {
    rm(goSvelteCfg.InternPublicPath.concat(goSvelteCfg.SvelteOutputPath), { recursive: true }, nil);
}

function goSvelteExit() {
    if (svelteWtch) svelteWtch.close();
    if (svelteProc) svelteProc.kill(0);
    if (goProc) goProc.kill(0);
}

function nil() { }