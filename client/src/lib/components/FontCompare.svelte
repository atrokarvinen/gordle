<script lang="ts">
	import { LetterState } from '../models';
	import LetterBox from './LetterBox.svelte';

	const alphabets = 'abcdefghijklmnopqrstuvwxyz'.split('');
	const fonts = [
		'Arial',
		'Times New Roman',
		'Courier New',
		'Verdana',
		'Georgia',
		'Comic Sans MS',
		'Trebuchet MS',
		'Impact',
		'Lucida Console',
		'Tahoma',
		'Lucida Sans Unicode',
		'Palatino Linotype',
		'Book Antiqua',
		'Arial Black',
		'Gadget',
		'Copperplate Gothic Light',
		'Lucida Sans',
		'Brush Script MT',
		'Helvetica',
		'Courier',
		'Lucida Grande',
		'Geneva',
		'Helvetica Neue',
		'sans-serif',
		'serif',
		'monospace',
		'cursive',
		'fantasy'
	];
</script>

{#each fonts as font}
	<div class="flex">
		<span>
			<span>Font:</span>
			<span>{font}</span>
		</span>
		<div class="flexRow">
			{#each alphabets as alphabet}
				<LetterBox letter={alphabet} letterState={LetterState.CORRECT} {font} />
			{/each}
		</div>
	</div>
{/each}

<style>
	.flex {
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.flexRow {
		display: flex;
		flex-direction: row;
		align-items: center;
	}
</style>
