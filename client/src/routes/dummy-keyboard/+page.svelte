<script lang="ts">
	import Keyboard from '$lib/components/Keyboard.svelte';
	import WordBoard from '$lib/components/WordBoard.svelte';
	import { guesses } from '../dummy-half-filled/dummyData';
</script>

<WordBoard words={guesses} />
<Keyboard {guesses} onKeyDown={() => {}} />
