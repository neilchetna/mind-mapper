<script>
	import { Button } from '$lib/components/ui/button';
	import { toggleMode } from 'mode-watcher';
	import { SignedIn, SignedOut, SignInButton, UserButton } from 'svelte-clerk';
</script>

<div class="m-auto w-fit pt-32">
	<h1 class="text-7xl">Mind Mapper</h1>
	<div class="flex w-full items-center justify-between">
		<SignedOut>
			<Button>
				<SignInButton />
			</Button>
		</SignedOut>
		<SignedIn>
			<UserButton />
		</SignedIn>

		<Button variant="outline" href="/map">Open Maps</Button>

		<Button onclick={toggleMode} variant="outline" size="icon">
			<span
				class="h-[1.2rem] w-[1.2rem] scale-100 rotate-0 !transition-all dark:scale-0 dark:-rotate-90"
			>
				🌙
			</span>
			<span
				class="absolute h-[1.2rem] w-[1.2rem] scale-0 rotate-90 !transition-all dark:scale-100 dark:rotate-0"
			>
				☀️
			</span>
		</Button>
	</div>
</div>
