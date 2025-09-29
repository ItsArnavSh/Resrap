<script lang="ts">
	import ABNF from './ABNF.svelte';
	import About from './About.svelte';
	import Footer from './footer.svelte';
	import GenDel from './GenDel.svelte';
	import Header from './header.svelte';
	import Sample from './Sample.svelte';
</script>

<div class="h-full w-full bg-[#06081d]">
	<div class="flex w-full flex-col lg:flex-row lg:items-center lg:justify-center">
		<div class="w-full lg:w-[40%]">
			<Header />
		</div>
		<div class="w-full lg:w-[60%]">
			<GenDel />
		</div>
	</div>
	<About />
	<ABNF />
	<Sample />
	<Footer />
</div>
