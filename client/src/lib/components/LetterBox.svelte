<script lang="ts">
	import { LetterState } from '../models';

	export let letterState: LetterState = LetterState.CORRECT;
	export let letter: string = 'a';
	export let font: string = 'Verdana';
	export let readonly = true;
</script>

<div
	class={`letter-box ${letterState === LetterState.CORRECT ? 'success' : letterState === LetterState.CONTAINED ? 'warning' : letterState === LetterState.UNKNOWN ? 'unknown' : 'error'}`}
	style={`font-family: ${font};`}
>
	{#if readonly}
		<p>
			{letter}
		</p>
	{:else}
		<input type="text" class="letter-box error" value={letter} />
	{/if}
</div>

<style>
	.letter-box {
		display: flex;
		flex-direction: row;
		width: 50px;
		height: 50px;
		margin: 2px;
		border: 1px solid black;

		font-weight: bold;
		font-family: 'Verdana';
		font-size: 1.5rem;

		text-transform: uppercase;
		text-align: center;

		color: white;
		align-items: center;
		justify-content: center;
	}

	.success {
		background-color: green;
	}

	.warning {
		background-color: burlywood;
	}

	.error {
		background-color: black;
	}

	.unknown {
		background-color: gray;
	}
</style>
