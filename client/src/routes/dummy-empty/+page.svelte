<script lang="ts">
	import WordBoard from '$lib/components/WordBoard.svelte';
	import { guesses } from './dummyData';

	const data = guesses;
</script>

<WordBoard words={data} />
