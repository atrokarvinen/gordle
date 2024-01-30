<script lang="ts">
	import LetterBox from '$lib/components/LetterBox.svelte';
	import { LetterState } from '$lib/models';

	export let letterState: LetterState;
	export let letter: string;
	export let text: string;
</script>

<div class="flex items-center gap-x-2">
	<div class="border-2 border-solid">
		<LetterBox {letter} {letterState} />
	</div>
	<p>{text}</p>
</div>
