<script lang="ts">
	import { SignedIn, SignedOut, SignInButton, UserButton } from 'svelte-clerk';
</script>

<SignedOut>
	<SignInButton />
</SignedOut>
<SignedIn>
	<UserButton />
</SignedIn>
