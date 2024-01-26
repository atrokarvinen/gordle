<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';
	import { LetterState } from '../models';

	interface $$Props extends HTMLInputAttributes {
		letterState: LetterState;
		letter: string;
		font?: string;
		readonly?: boolean;
		cursor?: string;
		focused?: boolean;
	}

	export let letterState: LetterState;
	export let letter: string;
	export let font: string = 'Verdana';
	export let readonly = true;
	export let cursor = 'default';

	const getBackgroundColor = (letterState: LetterState) => {
		switch (letterState) {
			case LetterState.CORRECT:
				return 'bg-success-500';
			case LetterState.CONTAINED:
				return 'bg-warning-500';
			case LetterState.UNKNOWN:
				return 'bg-surface-400';
			case LetterState.INCORRECT:
				return 'bg-surface-900';
			default:
				return 'bg-error-500';
		}
	};

	$: bgColor = getBackgroundColor(letterState);
	$: border = $$props.focused ? 'border-2 border-primary-500' : 'border-2 border-transparent';
	const height = 'md:h-16 h-8';
	const width = 'md:w-16 w-8';
</script>

<input
	{...$$restProps}
	class={`font-verdana text-center font-bold uppercase text-white outline-none md:text-2xl  ${width} ${height} ${cursor} ${bgColor} ${border}`}
	style={`font-family: ${font};`}
	value={letter}
	{readonly}
	on:click
/>
