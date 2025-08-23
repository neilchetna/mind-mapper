<script>
	import { Button } from '$lib/components/ui/button';
	import { toggleMode, theme } from 'mode-watcher';
	import { SignedIn, SignedOut, SignInButton, UserButton, UserProfile } from 'svelte-clerk';
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
				class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 !transition-all dark:-rotate-90 dark:scale-0"
			>
				🌙
			</span>
			<span
				class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 !transition-all dark:rotate-0 dark:scale-100"
			>
				☀️
			</span>
		</Button>
	</div>
</div>
