<nav class="navi">
	<a href="/">Home</a>
	<a href="/game">Game</a>
	<a href="/dummy-solved">Dummy solved</a>
	<a href="/dummy-half-filled">Dummy half filled</a>
	<a href="/dummy-empty">Dummy empty</a>
	<a href="/dummy-guess">Guess</a>
	<a href="/dummy-keyboard">Keyboard</a>
</nav>
<slot />

<style global>
	.navi {
		display: flex;
		flex-direction: row;
		justify-content: space-evenly;
		align-items: center;
	}

	:global(*) {
		box-sizing: border-box;
	}
</style>
