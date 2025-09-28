<script>
	import { onMount } from 'svelte';
	import { loadWasm } from '$lib/wasm.js';
	let result = '';
	let status = 'loading'; // 'loading' | 'ready' | 'error'
	// Form inputs
	let grammar = `program : header+ function^;
header:'#include<'identifier'.h>\\n';
function:functionheader'{''\\n'functioncontent'}';
functionheader:datatype ' ' identifier '(' ')';
functioncontent:block+;
block:(statement+) | ifblock | whileblock ;
ifblock: 'if(' conditionalexpression '){\\n' statement+ '}\\n';
whileblock :'while(' conditionalexpression '){\\n' statement+ '}\\n';
conditionalexpression: conditionalexpone (conditionaljoin conditionalexpone)*;
conditionalexpone: (operand conditionaloperation operand);
conditionaloperation: ' < ' | ' > ';
conditionaljoin: ' && ' | ' || ';
statement:assignment;
assignment: datatype ' ' identifier ' = ' expression ';\\n';
expression: operand (operator (operand | '(' expression ')'))*;
operand: identifier | integer | float;
operator : ' + ' | ' - ' | ' * ' | ' / ';
datatype: 'int'|'float'|'double';
identifier: [A-Z];
integer:[0-9];
float:[0-9]'.'[0-9];
whitespace: ' ';`;

	let startpoint = 'program';
	let seed = 42;
	let tokenlen = 10000;

	// Stats tracking
	let executionTime = 0;
	let tokensGenerated = 0;
	let isGenerating = false;
	let lastGeneratedAt = null;

	onMount(async () => {
		try {
			await loadWasm();
			status = 'ready';
		} catch (error) {
			console.error('Failed to load WASM:', error);
			status = 'error';
		}
	});

	function callGenerateText() {
		if (globalThis.generateText) {
			isGenerating = true;

			// High-resolution time measurement
			const startTime = performance.now();

			try {
				result = globalThis.generateText(grammar, startpoint, seed, tokenlen);

				// Calculate execution time
				const endTime = performance.now();
				executionTime = endTime - startTime;

				// Count tokens (approximate - split by whitespace and filter empty)
				tokensGenerated = result
					.trim()
					.split(/\s+/)
					.filter((token) => token.length > 0).length;

				lastGeneratedAt = new Date().toLocaleTimeString();
			} catch (error) {
				result = '❌ Error generating text: ' + error.message;
				executionTime = 0;
				tokensGenerated = 0;
			} finally {
				isGenerating = false;
			}
		} else {
			result = '⚠️ generateText function not found';
			executionTime = 0;
			tokensGenerated = 0;
		}
	}

	// Calculate tokens per second
	$: tokensPerSecond =
		executionTime > 0 ? (tokensGenerated / (executionTime / 1000)).toFixed(2) : 0;
</script>

<main>
	<h1>SvelteKit + Go WASM</h1>
	<p class="subtitle">Try it yourself segment</p>

	{#if status === 'loading'}
		<p>⏳ Loading...</p>
	{:else if status === 'error'}
		<p>❌ Failed to load WASM</p>
	{/if}

	{#if status === 'ready'}
		<div class="container">
			<!-- Left Side: Grammar, Inputs, and Button -->
			<div class="left-panel">
				<label>Grammar</label>
				<textarea spellcheck="false" bind:value={grammar} rows="20" />

				<!-- Small inputs below grammar -->
				<div class="small-inputs">
					<div class="input-group">
						<label>Start Point</label>
						<input bind:value={startpoint} type="text" />
					</div>
					<div class="input-group">
						<label>Seed <small>(0 = random)</small></label>
						<input bind:value={seed} type="number" />
					</div>
					<div class="input-group">
						<label>Token Length</label>
						<input bind:value={tokenlen} type="number" />
					</div>
				</div>

				<button on:click={callGenerateText} disabled={isGenerating}>
					{isGenerating ? '⏳ Generating...' : 'Generate Text'}
				</button>
			</div>

			<!-- Right Side: Output and Stats -->
			<div class="right-panel">
				<!-- Output -->
				<div class="output-section">
					<label>Result</label>
					<textarea readonly value={result} rows="20" />
				</div>

				<!-- Simplified Stats -->
				{#if result && !isGenerating}
					<div class="stats">
						<h3>Ran using Go in WASM on your machine</h3>
						<div class="stats-headers">
							<div class="stat-header">Time taken</div>
							<div class="stat-header">Tokens/second</div>
							<div class="stat-header">Characters</div>
						</div>
						<div class="stats-values">
							<div class="stat-value">{executionTime.toFixed(2)} ms</div>
							<div class="stat-value">{tokensPerSecond}</div>
							<div class="stat-value">{result.length.toLocaleString()}</div>
						</div>
					</div>
				{/if}
			</div>
		</div>
	{/if}
</main>

<style>
	main {
		padding: 20px;
		font-family: sans-serif;
	}

	.subtitle {
		color: #666;
		font-style: italic;
		margin-bottom: 20px;
	}

	.container {
		display: flex;
		gap: 20px;
		margin-top: 20px;
	}

	.left-panel {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: 12px;
	}

	.right-panel {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: 12px;
	}

	.small-inputs {
		display: flex;
		gap: 12px;
	}

	.input-group {
		flex: 1;
		display: flex;
		flex-direction: column;
		gap: 4px;
	}

	.input-group label {
		font-size: 12px;
	}

	.output-section {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}

	.stats {
		background: #f8f9fa;
		border: 1px solid #e9ecef;
		border-radius: 8px;
		padding: 16px;
		margin-top: 8px;
	}

	.stats h3 {
		margin: 0 0 12px 0;
		font-size: 14px;
		font-weight: bold;
		color: #333;
	}

	.stats-headers {
		display: flex;
		gap: 20px;
		margin-bottom: 8px;
	}

	.stats-values {
		display: flex;
		gap: 20px;
	}

	.stat-header {
		flex: 1;
		font-size: 12px;
		color: #666;
		font-weight: 500;
		text-align: center;
	}

	.stat-value {
		flex: 1;
		font-size: 12px;
		font-weight: bold;
		color: #333;
		font-family: monospace;
		text-align: center;
	}

	label {
		font-weight: bold;
	}

	textarea,
	input {
		padding: 8px;
		border: 1px solid #ccc;
		border-radius: 6px;
		font-size: 14px;
		width: 100%;
		box-sizing: border-box;
		font-family: monospace;
	}

	textarea {
		resize: vertical;
	}

	button {
		background: #007acc;
		color: white;
		border: none;
		padding: 10px;
		border-radius: 6px;
		cursor: pointer;
		font-weight: bold;
		transition: background-color 0.2s;
	}

	button:hover:not(:disabled) {
		background: #005a9e;
	}

	button:disabled {
		background: #ccc;
		cursor: not-allowed;
	}

	/* Responsive adjustments */
	@media (max-width: 768px) {
		.container {
			flex-direction: column;
		}

		.stats-headers,
		.stats-values {
			flex-direction: column;
			gap: 8px;
		}

		.stat-header,
		.stat-value {
			text-align: left;
		}

		.small-inputs {
			flex-direction: column;
		}
	}
</style>
