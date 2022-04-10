import { exit } from 'process';
import { readFileSync } from 'fs';
import { load } from 'js-yaml';

export default function() {
	let data;
	try {
		data = load(readFileSync('./system/gosvelte/src/gosvelte.yaml', 'utf8'));
	} catch (e) {
		console.log(e);
		exit();
	}
	return {
		InternPublicPath: data.gosvelte.internpublicpath,
		ForwardPublicPath: data.gosvelte.forwardpublicpath,
		SvelteWorkspacePath: data.gosvelte.svelteworkspacepath,
		SvelteOutputPath: data.gosvelte.svelteoutputpath,
		SvelteExtension: data.gosvelte.svelteextension,
		IsProduction: !process.env.ROLLUP_WATCH,
	};
}