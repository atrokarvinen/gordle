<script lang="ts">
	import { base } from '$app/paths';
</script>

<nav class="navi">
	<a href="{base}/">Home</a>
	<a href="{base}/game/2">Game</a>
	<a href="{base}/dummy-solved">Dummy solved</a>
	<a href="{base}/dummy-half-filled">Dummy half filled</a>
	<a href="{base}/dummy-empty">Dummy empty</a>
	<a href="{base}/dummy-guess">Guess</a>
	<a href="{base}/dummy-keyboard">Keyboard</a>
</nav>

<div class="container">
	<div class="content">
		<slot />
	</div>
</div>

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

	.container {
		width: 100vw;
		height: 100vh;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	.content {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		width: 400px;
	}
</style>
