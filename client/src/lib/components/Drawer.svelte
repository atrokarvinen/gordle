<script lang="ts">
	import { page } from '$app/state';
	import MobileNavigation from '../../routes/MobileNavigation.svelte';
	import Backdrop from './Backdrop.svelte';

	interface Props {
		isOpen: boolean;
		onClose: () => void;
	}

	let { isOpen = false, onClose }: Props = $props();

	let drawer: HTMLElement | undefined;

	$effect(() => {
		if (!drawer) return;
		if (isOpen) {
			drawer.style.width = '12rem';
		} else {
			drawer.style.width = '0';
		}
	});

	let previousPath = '';
	$effect(() => {
		if (page.url.pathname !== previousPath) {
			onClose();
		}
		previousPath = page.url.pathname;
	});
</script>

<Backdrop visible={isOpen} {onClose} />
<div id="mySidenav" class="sidenav surface-500" bind:this={drawer}>
	<button class="closebtn" onclick={onClose} aria-label="Close navigation menu">&times;</button>
	<MobileNavigation />
</div>

<style>
	.sidenav {
		height: 100%;
		width: 0;
		position: fixed;
		z-index: 2000;
		top: 0;
		right: 0;
		overflow-x: hidden;
		padding-top: 60px;
		transition: 0.5s;
		background-color: var(--body-background-color-dark);
	}

	.sidenav .closebtn {
		position: absolute;
		top: 0;
		right: 25px;
		font-size: 36px;
		margin-left: 50px;
	}

	@media screen and (max-height: 450px) {
		.sidenav {
			padding-top: 15px;
		}
	}
</style>
