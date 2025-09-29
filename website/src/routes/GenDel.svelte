<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { loadWasm } from '$lib/wasm.js';
	import { samplegrammer } from '$lib/template';

	import materialPalenight from 'svelte-highlight/styles/material-palenight';
	import c from 'svelte-highlight/languages/c';
	import Highlight from 'svelte-highlight';

	let displayText = '';
	let status: 'loading' | 'ready' | 'error' = 'loading';
	let isGenerating = false;
	let isTyping = false;

	let grammar = samplegrammer;
	let startpoint = 'program';
	let tokenlen = 150;

	let executionTime = 0;
	let tokensGenerated = 0;
	let generationCount = 0;

	let generateInterval: NodeJS.Timeout | undefined;
	let typeInterval: NodeJS.Timeout | undefined;
	let currentGeneratedText = '';
	let typeIndex = 0;

	onMount(async () => {
		try {
			await loadWasm();
			status = 'ready';
			startShowcase();
		} catch (error) {
			console.error('Failed to load WASM:', error);
			status = 'error';
		}
	});

	onDestroy(() => {
		if (generateInterval) clearInterval(generateInterval);
		if (typeInterval) clearInterval(typeInterval);
	});

	function startShowcase() {
		generateNewText();
		generateInterval = setInterval(generateNewText, 4000);
	}

	function generateNewText() {
		if (isGenerating || isTyping || !globalThis.generateText) return;

		isGenerating = true;
		generationCount++;

		const startTime = performance.now();

		try {
			currentGeneratedText = globalThis.generateText(grammar, startpoint, 0, tokenlen);

			const endTime = performance.now();
			executionTime = endTime - startTime;
			tokensGenerated = currentGeneratedText
				.trim()
				.split(/\s+/)
				.filter((token) => token.length > 0).length;

			displayText = '';
			typeIndex = 0;
			startTypingAnimation();
		} catch (error) {
			currentGeneratedText = '❌ Error generating text: ' + (error as Error).message;
			displayText = currentGeneratedText;
			executionTime = 0;
			tokensGenerated = 0;
		} finally {
			isGenerating = false;
		}
	}

	function startTypingAnimation() {
		isTyping = true;
		typeInterval = setInterval(() => {
			if (typeIndex < currentGeneratedText.length) {
				displayText += currentGeneratedText[typeIndex];
				typeIndex++;
			} else {
				if (typeInterval) clearInterval(typeInterval);
				isTyping = false;
			}
		}, 15);
	}

	// Calculate tokens per second
	$: tokensPerSecond =
		executionTime > 0 ? (tokensGenerated / (executionTime / 1000)).toFixed(2) : 0;
</script>

<svelte:head>
	{@html materialPalenight}
</svelte:head>
<div class="flex flex-col p-4 lg:p-8 pt-8 lg:pt-20 pb-8 lg:pb-20 text-white h-auto lg:min-h-screen">
	{#if status === 'loading'}
		<div class="flex h-96 items-center justify-center">
			<div class="flex items-center gap-3">
				<div
					class="h-8 w-8 animate-spin rounded-full border-2 border-[#00add8] border-t-transparent"
				></div>
				<span class="text-lg lg:text-xl text-gray-300">Loading Resrap engine...</span>
			</div>
		</div>
	{:else if status === 'error'}
		<div class="flex h-96 items-center justify-center">
			<div class="text-center">
				<div class="mb-4 text-4xl">❌</div>
				<span class="text-lg lg:text-xl text-red-400">Failed to load Resrap engine</span>
			</div>
		</div>
	{/if}

	{#if status === 'ready'}
		<div class="flex flex-col">
			<div class="flex flex-col">
				<div
					class="relative h-96 lg:h-[32rem] overflow-hidden rounded-lg border border-gray-700 bg-[#292d3e] p-3 lg:p-6"
				>
					<div
						class="h-full overflow-y-auto font-mono text-xs lg:text-sm leading-relaxed whitespace-pre-wrap text-gray-100"
					>
						<div class="overflow-x-auto">
							<Highlight language={c} code={displayText} />
						</div>
						{#if isTyping}
							<span class="ml-1 inline-block h-4 lg:h-5 w-2 animate-pulse bg-[#00add8]"></span>
						{/if}
					</div>

					{#if displayText.length === 0 && !isTyping}
						<div class="absolute inset-0 flex items-center justify-center text-gray-500">
							<div class="text-center">
								<div class="mb-2 text-3xl lg:text-4xl">⚡</div>
								<p class="text-sm lg:text-base">Waiting for generation...</p>
							</div>
						</div>
					{/if}
				</div>

				<div class="mt-3 lg:mt-4 flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2 lg:gap-0 text-xs lg:text-sm text-gray-400">
					<div class="flex gap-3 lg:gap-4">
						<span>{displayText.length.toLocaleString()} chars</span>
						<span>{tokensGenerated} tokens</span>
					</div>
					<div class="flex gap-3 lg:gap-4">
						<span>{tokensPerSecond}k tokens/sec</span>
						<span>{executionTime.toFixed(1)}ms</span>
					</div>
				</div>
			</div>
		</div>
	{/if}
</div>

<style>
	/* Custom scrollbar */
	:global(.h-full.overflow-y-auto::-webkit-scrollbar) {
		width: 6px;
	}

	:global(.h-full.overflow-y-auto::-webkit-scrollbar-track) {
		background: #374151;
	}

	:global(.h-full.overflow-y-auto::-webkit-scrollbar-thumb) {
		background: #00add8;
		border-radius: 3px;
	}

	:global(.h-full.overflow-y-auto::-webkit-scrollbar-thumb:hover) {
		background: #0099c7;
	}
</style>
