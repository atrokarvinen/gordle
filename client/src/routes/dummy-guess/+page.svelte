<script lang="ts">
	import { LETTERS_COUNT } from '$lib/constants';

	const letters = LETTERS_COUNT;
	let inputLetters = Array.from(Array(letters).keys()).map(() => '');
	let inputElements: (HTMLInputElement | undefined)[] = Array.from(Array(letters).keys()).map(
		() => undefined
	);

	const onChange = (value: string, index: number) => {
		const oldValue = inputLetters[index];
		const lastLetter = value.toUpperCase().split('')[value.length - 1];

		const alphabets = 'abcdefghijklmnopqrstuvwxyz';
		if (!alphabets.includes(lastLetter.toLowerCase())) {
			inputElements[index]!.value = oldValue;
			return;
		}

		inputLetters = inputLetters.map((l, i) => (i === index ? lastLetter : l));
		if (inputElements[index]) {
			inputElements[index]!.value = lastLetter;
		}

		const isLast = index === letters - 1;
		if (isLast) {
			inputElements[index]?.blur();
			return;
		}
		const next = inputElements[index + 1];
		next?.focus();
	};
</script>

<div>
	<h1>Guess:</h1>
	<p>{inputLetters.join(', ')}</p>
</div>
<div class="letters">
	{#each inputLetters as letter, i}
		<input
			bind:this={inputElements[i]}
			on:input={(e) => {
				console.log('input', e.currentTarget.value);
				onChange(e.currentTarget.value, i);
			}}
			value={letter}
		/>
	{/each}
</div>

<style>
	.letters {
		display: flex;
		flex-direction: row;
	}
	input {
		width: 2rem;
		height: 2rem;
		font-size: 1.5rem;
		text-align: center;
		margin: 0.5rem;
	}
</style>
