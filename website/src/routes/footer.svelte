<script>
	const githubUrl = 'https://github.com/osdc/Resrap';
	const pkgUrl = 'https://pkg.go.dev/github.com/osdc/resrap';
	const osdcGithubUrl = 'https://github.com/osdc';
	const osdcWebUrl = 'https://osdc.dev';
</script>

<footer class="bg-black px-8 py-12 text-white">
	<div class="mx-auto max-w-4xl text-center">
		<!-- Main Message -->
		<div class="mb-8">
			<p class="mb-2 text-2xl text-gray-200">
				Made with <span class="text-2xl text-red-500">♥</span> by
				<a
					href="https://github.com/itsarnavsh"
					target="_blank"
					rel="noopener noreferrer"
					class="font-semibold text-[#00add8] transition-colors hover:text-[#0099c7]"
				>
					Arnav
				</a>
			</p>
			<p class="text-sm text-gray-400">© 2025 OSDC Organization</p>
		</div>

		<!-- Links -->
		<div class="flex flex-wrap justify-center gap-6 text-sm">
			<a
				href={githubUrl}
				target="_blank"
				rel="noopener noreferrer"
				class="text-gray-400 transition-colors hover:text-[#00add8]"
			>
				GitHub
			</a>
			<a
				href={pkgUrl}
				target="_blank"
				rel="noopener noreferrer"
				class="text-gray-400 transition-colors hover:text-[#00add8]"
			>
				pkg.go.dev
			</a>
			<a
				href={osdcWebUrl}
				target="_blank"
				rel="noopener noreferrer"
				class="text-gray-400 transition-colors hover:text-[#00add8]"
			>
				osdc.dev
			</a>
			<a
				href={osdcGithubUrl}
				target="_blank"
				rel="noopener noreferrer"
				class="text-gray-400 transition-colors hover:text-[#00add8]"
			>
				OSDC GitHub
			</a>
			<a
				href="https://golang.org"
				target="_blank"
				rel="noopener noreferrer"
				class="text-gray-400 transition-colors hover:text-[#00add8]"
			>
				Go Language
			</a>
		</div>
	</div>
</footer>
