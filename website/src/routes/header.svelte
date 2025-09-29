<script>
	const githubUrl = 'https://github.com/osdc/Resrap';
	const pkgUrl = 'https://pkg.go.dev/github.com/osdc/resrap';
</script>

<div class="flex flex-col justify-between p-4 lg:p-8 text-white h-auto lg:min-h-screen">
	<div class="flex flex-1 flex-col justify-center">
		<div class="mb-8 lg:mb-12">
			<h1 class="headfont mb-2 lg:mb-4 text-3xl lg:text-6xl leading-tight font-bold text-[#00add8]">Resrap</h1>
			<p class="headfont mb-4 lg:mb-6 text-xl lg:text-3xl text-gray-300">A parser...but in reverse</p>
			<div class="max-w-lg space-y-3 text-base lg:text-lg leading-relaxed text-gray-400">
				<p class="codefont font-medium text-gray-300">
					Generate infinitely long procedurally generated code with just a basic grammar, suitable
					for stress testing linters and parsers, generating realistic sample non copyrighted code
					for display, or just for playing around with grammar theory
				</p>
			</div>
		</div>

		<div class="mb-6 lg:mb-8 flex flex-col sm:flex-row gap-3 lg:gap-4">
			<a
				href={githubUrl}
				target="_blank"
				rel="noopener noreferrer"
				class="flex items-center justify-center gap-2 rounded-lg bg-[#00add8] px-4 lg:px-6 py-3 font-semibold text-white transition-colors duration-200 hover:bg-[#0099c7]"
			>
				<svg class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24">
					<path
						d="M12 2L15.09 8.26L22 9.27L17 14.14L18.18 21.02L12 17.77L5.82 21.02L7 14.14L2 9.27L8.91 8.26L12 2Z"
					/>
				</svg>
				Star on GitHub
			</a>
			<a
				href="#try"
				class="flex items-center justify-center gap-2 rounded-lg border-2 border-[#00add8] px-4 lg:px-6 py-3 font-semibold text-[#00add8] transition-all duration-200 hover:bg-[#00add8] hover:text-white"
			>
				<svg class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24">
					<path d="M8 5V16L12 13L16 16V5H8Z" />
				</svg>
				Try Now
			</a>
		</div>
	</div>

	<div class="flex flex-col gap-2 lg:gap-3 mt-8 lg:mt-0">
		<a
			href={githubUrl}
			target="_blank"
			rel="noopener noreferrer"
			class="flex w-fit items-center gap-2 text-sm text-gray-500 transition-colors duration-200 hover:text-[#00add8]"
		>
			<svg class="h-4 w-4" fill="currentColor" viewBox="0 0 24 24">
				<path
					d="M12 0C5.374 0 0 5.373 0 12 0 17.302 3.438 21.8 8.207 23.387c.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23A11.509 11.509 0 0112 5.803c1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576C20.566 21.797 24 17.3 24 12c0-6.627-5.373-12-12-12z"
				/>
			</svg>
			GitHub Repository
		</a>
		<a
			href={pkgUrl}
			target="_blank"
			rel="noopener noreferrer"
			class="flex w-fit items-center gap-2 text-sm text-gray-500 transition-colors duration-200 hover:text-[#00add8]"
		>
			<svg class="h-4 w-4" fill="currentColor" viewBox="0 0 24 24">
				<path
					d="M4.5 6.75a.75.75 0 00-.75.75v9c0 .414.336.75.75.75h15a.75.75 0 00.75-.75v-9a.75.75 0 00-.75-.75h-15zM3 7.5a1.5 1.5 0 011.5-1.5h15a1.5 1.5 0 011.5 1.5v9a1.5 1.5 0 01-1.5 1.5h-15A1.5 1.5 0 013 16.5v-9z"
				/>
				<path
					d="M6 10.5a.75.75 0 01.75-.75h2.5a.75.75 0 010 1.5h-2.5A.75.75 0 016 10.5zM6 12.75a.75.75 0 01.75-.75h6.5a.75.75 0 010 1.5h-6.5a.75.75 0 01-.75-.75z"
				/>
			</svg>
			pkg.go.dev Documentation
		</a>
	</div>
</div>
