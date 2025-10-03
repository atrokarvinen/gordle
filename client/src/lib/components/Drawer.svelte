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
			// drawer.style.width = '0';
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

<Backdrop visible={isOpen} {onClose}>
	<div id="mySidenav" class="sidenav surface-500" bind:this={drawer}>
		<button class="closebtn" onclick={onClose} aria-label="Close navigation menu">&times;</button>
		<MobileNavigation />
	</div>
</Backdrop>

<style>
	/* The side navigation menu */
	.sidenav {
		height: 100%; /* 100% Full-height */
		width: 0; /* 0 width - change this with JavaScript */
		position: fixed; /* Stay in place */
		z-index: 2000; /* Stay on top */
		top: 0; /* Stay at the top */
		right: 0;
		overflow-x: hidden; /* Disable horizontal scroll */
		padding-top: 60px; /* Place content 60px from the top */
		transition: 0.5s; /* 0.5 second transition effect to slide in the sidenav */
		background-color: var(--body-background-color-dark);
	}

	/* Position and style the close button (top right corner) */
	.sidenav .closebtn {
		position: absolute;
		top: 0;
		right: 25px;
		font-size: 36px;
		margin-left: 50px;
	}

	/* On smaller screens, where height is less than 450px, change the style of the sidenav (less padding and a smaller font size) */
	@media screen and (max-height: 450px) {
		.sidenav {
			padding-top: 15px;
		}
	}
</style>
