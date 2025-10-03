<script lang="ts">
	import { LetterState } from '../models';

	// interface $$Props extends HTMLInputAttributes {
	// 	letterState: LetterState;
	// 	letter: string;
	// 	font?: string;
	// 	readonly?: boolean;
	// 	cursor?: string;
	// 	focused?: boolean;
	// 	icon?: string;
	// }

	export let icon: string | undefined = undefined;
	export let letterState: LetterState;
	export let letter: string;
	export let font: string = 'Verdana';
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
				return 'bg-surface-700';
			default:
				return 'bg-error-500';
		}
	};

	$: bgColor = getBackgroundColor(letterState);
	$: border = $$props.focused ? 'border-2 border-primary-500' : 'border-2 border-transparent';
	const height = 'md:h-12 h-7';
	$: width = icon ? 'md:w-24 w-14' : 'md:w-12 w-7';
	$: classNames = `font-verdana text-center font-bold uppercase outline-none md:text-2xl  ${width} ${height} ${cursor} ${bgColor} ${border}`;
</script>

<button
	aria-label="Letter {letter}"
	data-testid="guess-letter"
	class={classNames}
	style={`font-family: ${font};`}
	on:click
>
	{#if icon}
		<i class={icon}></i>
	{:else}
		{letter}
	{/if}
</button>
