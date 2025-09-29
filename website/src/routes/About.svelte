<section id="about" class="overflow-y-auto p-4 lg:p-8 text-white">
	<div class="mx-auto max-w-4xl">
		<h1 class="mb-4 text-2xl lg:text-4xl font-bold text-[#00add8]">What is Resrap?</h1>
		<p class="mb-6 lg:mb-8 text-base lg:text-lg leading-relaxed text-gray-300">
			Resrap is a <strong class="text-[#00add8]"
				>seedable, grammar-based code snippet generator</strong
			>. Instead of parsing code, it <strong>generates code</strong> from formal grammars — producing
			endless, realistic-looking (or hilariously nonsensical) snippets.
		</p>

		<p class="mb-6 lg:mb-8 text-gray-300 text-sm lg:text-base">
			It works with <strong class="text-[#00add8]">any language</strong> that can be described with a
			grammar (even English if you like!) and is perfect for:
		</p>

		<div class="mb-8 lg:mb-12 ml-4 lg:ml-6 space-y-2 lg:space-y-3">
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base">Typing practice with realistic-looking snippets</span>
			</div>
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base">Stress-testing parsers, syntax highlighters, or linters</span>
			</div>
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base">Fun exploration of procedural code generation</span>
			</div>
		</div>

		<p class="mb-8 lg:mb-12 text-gray-300 text-sm lg:text-base">
			Resrap now also supports <strong class="text-[#00add8]"
				>probabilistic and infinitely repeatable grammars</strong
			>
			via the <strong>ABNF (Awesome BNF) format</strong> —
			<a href="#abnf" class="text-[#00add8] underline transition-colors hover:text-[#0099c7]"
				>see full reference</a
			>.
		</p>

		<h2 class="mb-4 lg:mb-6 text-xl lg:text-2xl font-bold text-[#00add8]">How?</h2>
		<p class="mb-6 lg:mb-8 text-gray-300 text-sm lg:text-base">
			Resrap reads a grammar and builds a <strong class="text-[#00add8]">graph of expansions</strong
			>. It then randomly traverses the graph (or deterministically with a seed) to produce snippets
			that follow the grammar's syntax rules.
		</p>

		<h2 class="mb-4 lg:mb-6 text-xl lg:text-2xl font-bold text-[#00add8]">Why Resrap?</h2>
		<div class="mb-8 lg:mb-12 ml-4 lg:ml-6 space-y-2 lg:space-y-3">
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base"
					>Generate <strong>unlimited, realistic code snippets</strong></span
				>
			</div>
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base">Avoid copyright issues from using real code</span>
			</div>
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base"
					>Provide a <strong>fun, deterministic, and probabilistic code generator</strong></span
				>
			</div>
			<div class="flex items-start gap-2 lg:gap-3">
				<div class="mt-2 h-1.5 w-1.5 rounded-full bg-[#00add8]"></div>
				<span class="text-gray-300 text-sm lg:text-base"
					>Give programmers a playground for <strong>syntax, speed, and randomness</strong></span
				>
			</div>
		</div>

		<div class="text-center">
			<p class="text-lg lg:text-xl text-gray-400 italic">"Just a parser… in reverse."</p>
		</div>
	</div>
</section>
