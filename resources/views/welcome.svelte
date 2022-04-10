<svelte:options tag="gosvelte-welcome-page" />

<script>
	export let lang;
	export let dir;
	export let name;

	const src = "/www/public/gosvelte.png";

	console.log("Data from Go to Svelte:", lang, dir, name);
</script>

<main>
	<h1>Welcome to<br/>{name}</h1>
	<img {src} alt={name} />
	<h2>
		Golang + Svelte<br/>Starter Kit
	</h2>
	<p>
		Visit <a target="_blank" href="https://svelte.dev/">Svelte</a> to learn
		how to build Svelte apps.
	</p>
	<p>
		Visit <a target="_blank" href="https://go.dev/">Golang</a> to learn
		more.
	</p>
	<p>
		<a target="_blank" href="https://github.com/moadkey/gosvelte">GoSvelte Github Repository</a>
	</p>
	<h3>Happy Coding.</h3>	
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #29beb0;
		/*text-transform: uppercase;*/
		font-size: 4em;
		font-weight: 100;
	}

	img {
		width: 25%;
		height: auto;
	}

	p {
		font-size: 1.25em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
