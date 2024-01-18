<script lang="ts">
	export let inputLetters: string[];
	const len = inputLetters.length;
	let inputElements: (HTMLInputElement | undefined)[] = Array.from(Array(len).keys()).map(
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

		const isLast = index === len - 1;
		if (isLast) {
			inputElements[index]?.blur();
			return;
		}
		const next = inputElements[index + 1];
		next?.focus();
	};
</script>

<div>
	<span>
		<span>Guess:</span>
		<span>{inputLetters.join('')}</span>
	</span>
</div>
<div class="letters">
	{#each inputLetters as letter, i}
		<!-- <LetterBox {letter} letterState={LetterState.INCORRECT} readonly={false} /> -->
		<input
			class="letter-box"
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

	.letter-box {
		display: flex;
		flex-direction: row;
		width: 50px;
		height: 50px;
		margin: 2px;
		border: 1px solid black;
		background-color: black;

		font-weight: bold;
		font-family: 'Verdana';
		font-size: 1.5rem;

		text-transform: uppercase;
		text-align: center;

		color: white;
		align-items: center;
		justify-content: center;
	}
</style>
