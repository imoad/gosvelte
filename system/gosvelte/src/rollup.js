import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import css from 'rollup-plugin-css-only';
import { readdirSync } from 'fs';
import { extname } from 'path';
import goSvelte from './gosvelte.js';

const goSvelteCfg = goSvelte();

export default function () {
    let cfg = [];
    readdirSync(goSvelteCfg.SvelteWorkspacePath).filter(fileName => {
        return extname(fileName).toLowerCase() === goSvelteCfg.SvelteExtension;
    }).forEach(svelteFileName => cfg.push(goSvelteCompiler(svelteFileName.split('.')[0])));
    return cfg;
}

function goSvelteCompiler(svelteFileName) {
    return {
        input: goSvelteCfg.SvelteWorkspacePath.concat("/", svelteFileName, goSvelteCfg.SvelteExtension),
        output: {
            sourcemap: true,
            format: 'iife',
            name: svelteFileName.toLowerCase(),
            file: goSvelteCfg.InternPublicPath.concat(goSvelteCfg.SvelteOutputPath, "/", svelteFileName.toLowerCase(), "/", "bundle.js"),
        },
        plugins: [
            svelte({
                compilerOptions: {
                    dev: !goSvelteCfg.isProduction,
                    customElement: true,
                }
            }),
            css({ output: goSvelteCfg.InternPublicPath.concat(goSvelteCfg.SvelteOutputPath, "/", svelteFileName.toLowerCase(), "/", "bundle.css") }),
            resolve({
                browser: true,
                dedupe: ['svelte'],
            }),
            commonjs(),
            !goSvelteCfg.IsProduction && livereload(goSvelteCfg.InternPublicPath),
            goSvelteCfg.IsProduction && terser(),
        ],
        watch: {
            clearScreen: false,
        }
    };
}