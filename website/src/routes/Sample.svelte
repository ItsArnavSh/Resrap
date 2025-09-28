<script>
	import { onMount } from 'svelte';
	import { loadWasm } from '$lib/wasm.js';
	import { samplegrammer } from '$lib/template';
	let result = '';
	let status = 'loading'; // 'loading' | 'ready' | 'error'
	// Form inputs
	let grammar = samplegrammer;

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

<div class="flex h-screen flex-col justify-center">
	<main id="try" class="p-8 text-center text-white">
		<div class="mb-8">
			<h1 class="mb-2 text-4xl font-bold text-[#00add8]">Try it yourself</h1>
			<p class="text-lg text-gray-400">Test Resrap's power directly in your browser</p>
		</div>

		{#if status === 'loading'}
			<div class="flex h-96 items-center justify-center">
				<div class="flex items-center gap-3">
					<div
						class="h-8 w-8 animate-spin rounded-full border-2 border-[#00add8] border-t-transparent"
					></div>
					<span class="text-xl text-gray-300">Loading engine...</span>
				</div>
			</div>
		{:else if status === 'error'}
			<div class="flex h-96 items-center justify-center">
				<div class="text-center">
					<div class="mb-4 text-4xl">❌</div>
					<span class="text-xl text-red-400">Failed to load WASM engine</span>
				</div>
			</div>
		{/if}

		{#if status === 'ready'}
			<div class="flex h-full flex-col gap-8 lg:flex-row">
				<!-- Left Side: Grammar & Controls -->
				<div class="flex flex-1 flex-col gap-4">
					<div class="flex flex-col gap-2">
						<label class="text-lg font-semibold text-gray-300">ABNF Grammar</label>
						<textarea
							spellcheck="false"
							bind:value={grammar}
							rows="16"
							class="w-full resize-y rounded-lg border border-gray-600 bg-gray-900 p-4 font-mono text-sm text-gray-100 focus:border-[#00add8] focus:ring-1 focus:ring-[#00add8] focus:outline-none"
							placeholder="Enter your grammar here..."
						/>
					</div>

					<!-- Controls -->
					<div class="flex flex-col gap-3">
						<div class="grid grid-cols-1 gap-3 md:grid-cols-3">
							<div class="flex flex-col gap-1">
								<label class="text-sm text-gray-400">Start Point</label>
								<input
									bind:value={startpoint}
									type="text"
									class="w-full rounded-md border border-gray-600 bg-gray-900 p-3 font-mono text-sm text-gray-100 focus:border-[#00add8] focus:ring-1 focus:ring-[#00add8] focus:outline-none"
									placeholder="program"
								/>
							</div>
							<div class="flex flex-col gap-1">
								<label class="text-sm text-gray-400"
									>Seed <span class="text-xs">(0 = random)</span></label
								>
								<input
									bind:value={seed}
									type="number"
									class="w-full rounded-md border border-gray-600 bg-gray-900 p-3 font-mono text-sm text-gray-100 focus:border-[#00add8] focus:ring-1 focus:ring-[#00add8] focus:outline-none"
									placeholder="0"
								/>
							</div>
							<div class="flex flex-col gap-1">
								<label class="text-sm text-gray-400">Token Length</label>
								<input
									bind:value={tokenlen}
									type="number"
									class="w-full rounded-md border border-gray-600 bg-gray-900 p-3 font-mono text-sm text-gray-100 focus:border-[#00add8] focus:ring-1 focus:ring-[#00add8] focus:outline-none"
									placeholder="150"
								/>
							</div>
						</div>

						<button
							on:click={callGenerateText}
							disabled={isGenerating}
							class="flex items-center justify-center gap-2 rounded-lg bg-[#00add8] px-6 py-4 font-semibold text-white transition-all duration-200 hover:bg-[#0099c7] disabled:cursor-not-allowed disabled:bg-gray-600"
						>
							{#if isGenerating}
								<div
									class="h-5 w-5 animate-spin rounded-full border-2 border-white border-t-transparent"
								></div>
								Generating...
							{:else}
								<svg class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24">
									<path d="M13 10V3L4 14h7v7l9-11h-7z" />
								</svg>
								Generate Text
							{/if}
						</button>
					</div>
				</div>

				<!-- Right Side: Output & Stats -->
				<div class="flex flex-1 flex-col gap-4">
					<div class="flex flex-col gap-2">
						<label class="text-lg font-semibold text-gray-300">Generated Output</label>
						<div class="relative flex-1">
							<textarea
								readonly
								value={result}
								rows="16"
								class="h-full w-full resize-none rounded-lg border border-gray-600 bg-gray-900 p-4 font-mono text-sm text-gray-100 focus:outline-none"
								placeholder="Generated text will appear here..."
							/>
							{#if !result && !isGenerating}
								<div class="absolute inset-0 flex items-center justify-center text-gray-500">
									<div class="text-center">
										<div class="mb-2 text-3xl">⚡</div>
										<p>Click Generate to see results</p>
									</div>
								</div>
							{/if}
						</div>
					</div>

					<!-- Performance Stats -->
					{#if result && !isGenerating}
						<div class="rounded-lg border border-gray-600 bg-gray-900 p-4">
							<div class="mb-3 flex items-center gap-2">
								<svg class="h-4 w-4 text-[#00add8]" fill="currentColor" viewBox="0 0 24 24">
									<path
										d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"
									/>
								</svg>
								<span class="text-sm font-medium text-[#00add8]">Performance Metrics</span>
								<div class="h-px flex-1 bg-gray-700"></div>
							</div>

							<div class="grid grid-cols-3 gap-4">
								<div class="text-center">
									<div class="mb-1 text-xs text-gray-400">Execution Time</div>
									<div class="font-mono text-lg font-bold text-white">
										{executionTime.toFixed(2)}<span class="text-sm text-gray-400">ms</span>
									</div>
								</div>
								<div class="text-center">
									<div class="mb-1 text-xs text-gray-400">Throughput</div>
									<div class="font-mono text-lg font-bold text-[#00add8]">
										{tokensPerSecond}<span class="text-sm text-gray-400">k/s</span>
									</div>
								</div>
								<div class="text-center">
									<div class="mb-1 text-xs text-gray-400">Output Size</div>
									<div class="font-mono text-lg font-bold text-white">
										{result.length.toLocaleString()}<span class="text-sm text-gray-400">
											chars</span
										>
									</div>
								</div>
							</div>

							<div class="mt-3 border-t border-gray-700 pt-3">
								<div class="text-center text-xs text-gray-500">
									🚀 Powered by Go WASM running in your browser
								</div>
							</div>
						</div>
					{/if}
				</div>
			</div>
		{/if}
	</main>
</div>
