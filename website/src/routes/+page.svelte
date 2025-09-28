<script>
	import { onMount } from 'svelte';
	import { loadWasm } from '$lib/wasm.js';

	let result = '';
	let wasmLoaded = false;
	let loading = true;

	// Form inputs
	let grammar = 'simple';
	let startpoint = 'begin';
	let seed = 42;
	let tokenlen = 100;

	onMount(async () => {
		try {
			await loadWasm();
			wasmLoaded = true;
			loading = false;
			console.log('WASM loaded successfully!');
		} catch (error) {
			console.error('Failed to load WASM:', error);
			loading = false;
		}
	});

	function callGenerateText() {
		if (globalThis.generateText) {
			result = globalThis.generateText(grammar, startpoint, seed, tokenlen);
		} else {
			result = 'generateText function not found';
		}
	}

	function testGoFunction() {
		if (globalThis.goTest) {
			result = globalThis.goTest();
		}
	}
</script>

<main>
	<h1>SvelteKit + Go WASM</h1>

	{#if loading}
		<p>⏳ Loading WASM module...</p>
	{:else if wasmLoaded}
		<p>✅ WASM module loaded!</p>

		<div style="margin: 20px 0; padding: 20px; border: 1px solid #ccc; border-radius: 8px;">
			<h3>Generate Text Function</h3>

			<div style="margin: 10px 0;">
				<label>Grammar: <input bind:value={grammar} type="text" /></label>
			</div>

			<div style="margin: 10px 0;">
				<label>Start Point: <input bind:value={startpoint} type="text" /></label>
			</div>

			<div style="margin: 10px 0;">
				<label>Seed: <input bind:value={seed} type="number" /></label>
			</div>

			<div style="margin: 10px 0;">
				<label>Token Length: <input bind:value={tokenlen} type="number" /></label>
			</div>

			<button on:click={callGenerateText} style="padding: 10px 20px; margin: 10px 0;">
				Generate Text
			</button>
		</div>

		<button on:click={testGoFunction}>Test Go Function</button>

		{#if result}
			<div style="margin: 20px 0; padding: 15px; background: #f5f5f5; border-radius: 4px;">
				<strong>Result:</strong>
				{result}
			</div>
		{/if}
	{:else}
		<p>❌ Failed to load WASM module</p>
	{/if}
</main>

<style>
	label {
		display: block;
		margin: 5px 0;
	}

	input {
		margin-left: 10px;
		padding: 5px;
		width: 200px;
	}

	button {
		background: #007acc;
		color: white;
		border: none;
		padding: 8px 16px;
		border-radius: 4px;
		cursor: pointer;
		margin: 5px;
	}

	button:hover {
		background: #005a9e;
	}
</style>
