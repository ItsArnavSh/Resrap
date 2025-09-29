<section id="abnf" class="p-4 lg:p-8 text-white">
	<div class="mx-auto max-w-3xl">
		<h1 class="mb-4 text-2xl lg:text-4xl font-bold text-[#00add8]">ABNF – Awesome BNF</h1>
		<p class="mb-8 lg:mb-12 text-base lg:text-lg leading-relaxed text-gray-300">
			ABNF (Awesome BNF) is a lightweight, custom grammar format designed for the Resrap code
			generation tool. It extends standard EBNF/BNF with probabilities and infinite generation.
		</p>

		<h2 class="mb-4 lg:mb-6 text-xl lg:text-2xl font-bold text-[#00add8]">Quick Revision: EBNF and Grammar</h2>
		<p class="mb-4 text-gray-300 text-sm lg:text-base">ABNF inherits the basics of EBNF:</p>
		<pre class="mb-6 lg:mb-8 rounded bg-gray-900 p-3 font-mono text-xs lg:text-sm text-gray-100 overflow-x-auto">function : rules ;</pre>

		<h3 class="mb-4 text-lg lg:text-xl font-semibold text-gray-200">ABNF Operators</h3>
		<div class="mb-8 lg:mb-12 space-y-2 lg:space-y-3">
			<p class="font-mono text-sm lg:text-base">
				<span class="font-bold text-[#00add8]">+</span>
				<span class="text-gray-300">One or more repetitions</span>
			</p>
			<p class="font-mono text-sm lg:text-base">
				<span class="font-bold text-[#00add8]">*</span>
				<span class="text-gray-300">Zero or more repetitions</span>
			</p>
			<p class="font-mono text-sm lg:text-base">
				<span class="font-bold text-[#00add8]">?</span>
				<span class="text-gray-300">Optional element</span>
			</p>
			<p class="font-mono text-sm lg:text-base">
				<span class="font-bold text-[#00add8]">()</span> <span class="text-gray-300">Grouping</span>
			</p>
			<p class="font-mono text-sm lg:text-base">
				<span class="font-bold text-[#00add8]">^</span>
				<span class="text-gray-300">Infinite generation - loops back infinitely</span>
			</p>
		</div>

		<div class="mb-6 lg:mb-8">
			<p class="mb-3 text-gray-300 text-sm lg:text-base">Example with infinite generation:</p>
			<pre class="rounded bg-gray-900 p-3 font-mono text-xs lg:text-sm text-gray-100 overflow-x-auto">program : function^;</pre>
			<p class="mt-2 text-xs lg:text-sm text-gray-400">
				Generates functions endlessly until token limit is reached
			</p>
		</div>

		<h2 class="mb-4 lg:mb-6 text-xl lg:text-2xl font-bold text-[#00add8]">Probabilities (&lt;...&gt;)</h2>
		<p class="mb-4 text-gray-300 text-sm lg:text-base">
			ABNF allows specifying weighted probabilities for finer control:
		</p>
		<pre
			class="mb-4 rounded bg-gray-900 p-3 font-mono text-xs lg:text-sm text-gray-100 overflow-x-auto">char : a&lt;0.2&gt; | b&lt;0.8&gt;;
a    : 'A';
b    : 'B';</pre>
		<p class="mb-4 text-gray-300 text-sm lg:text-base">
			<span class="text-[#00add8]">a</span> is chosen 20% of the time,
			<span class="text-[#00add8]">b</span> 80%. Probabilities are normalized automatically.
		</p>

		<div class="mb-8 lg:mb-12">
			<p class="mb-3 text-gray-300 text-sm lg:text-base">Probabilities work with all operators:</p>
			<pre
				class="rounded bg-gray-900 p-3 font-mono text-xs lg:text-sm text-gray-100 overflow-x-auto">statement : loop+&lt;0.3&gt; | condition?&lt;0.7&gt;;</pre>
		</div>
	</div>
</section>
