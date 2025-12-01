<script>
	// import { goto } from '$app/navigation';
	let usernameOrEmail = $state('');
	let password = $state('');
	let confirmPassword = $state('');
	let loginSelected = $state(true);
	let cardHeader = $derived(loginSelected ? 'Welcome back' : 'Welcome');
	let cardDesc = $derived(
		loginSelected
			? 'Enter your credentials to access your notes.'
			: 'Set up your credentials and start taking notes.'
	);
	let submitText = $derived(loginSelected ? 'Log In' : 'Create Account');

	import { resolve } from '$app/paths';
</script>

<div class="font-default h-screen flex items-center justify-center text-ink-grey bg-void-slate z-1">
	<div class="max-w-[400px] p-[20px]">
		<div
			class="text-center mb-[30px] text-white font-bold text-[1.5rem] items-center justify-center"
		>
			Bilig
		</div>
		<div class="bg-white p-[40px] rounded-xl border-t-4 border-t-solid border-t-logic-teal">
			<div class="flex mb-[30px] border-b border-b-solid border-b-light-grey">
				<button
					class="flex-1 outline-none text-center pb-[10px] text-[0.9rem] font-bold cursor-pointer hover:text-ink-grey {loginSelected
						? 'text-logic-teal border-b-2 border-b-solid border-b-logic-teal'
						: 'text-slate-mist '}"
					onclick={() => (loginSelected = true)}>Log In</button
				>
				<button
					class="flex-1 outline-none text-center pb-[10px] text-[0.9rem] font-bold cursor-pointer hover:text-ink-grey {!loginSelected
						? 'text-logic-teal border-b-2 border-b-solid border-b-logic-teal'
						: 'text-slate-mist '}"
					onclick={() => (loginSelected = false)}>Create Account</button
				>
			</div>
			<h2 class="text-[1.25rem] mb-[8px] text-ink-grey font-bold">{cardHeader}</h2>
			<p class="text-[0.85rem] text-slate-mist mb-[25px]">{cardDesc}</p>
			<form action="">
				<div class="mb-[20px]">
					<label class="block text-[0.85rem] font-semibold mb-[6px] text-ink-grey" for="email"
						>Username or Email</label
					>
					<input
						class="w-full p-[12px] border border-solid border-light-grey rounded-xl bg-ice-white text-ink-grey transition-all duration-200 focus:outline-none focus:border-logic-teal focus:shadow-[0 0 0 3px rgba(13, 148, 136, 0.1)] focus:bg-white"
						type="text"
						name="email"
						id="email"
						placeholder="user@local.host"
						bind:value={usernameOrEmail}
						required
					/>
				</div>
				<div class="mb-[20px]">
					<label class="block text-[0.85rem] font-semibold mb-[6px] text-ink-grey" for="password"
						>Password</label
					>
					<input
						class="w-full p-[12px] border border-solid border-light-grey rounded-xl bg-ice-white text-ink-grey transition-all duration-200 focus:outline-none focus:border-logic-teal focus:shadow-[0 0 0 3px rgba(13, 148, 136, 0.1)] focus:bg-white"
						type="password"
						name="password"
						id="password"
						placeholder="••••••••"
						bind:value={password}
						required
					/>
				</div>
				{#if !loginSelected}
					<div class="mb-[20px]">
						<label
							class="block text-[0.85rem] font-semibold mb-[6px] text-ink-grey"
							for="repeat-password">Repeat Password</label
						>
						<input
							class="w-full p-[12px] border border-solid border-light-grey rounded-xl bg-ice-white text-ink-grey transition-all duration-200 focus:outline-none focus:border-logic-teal focus:shadow-[0 0 0 3px rgba(13, 148, 136, 0.1)] focus:bg-white"
							type="repeat-password"
							name="repeat-password"
							id="repeat-password"
							placeholder="••••••••"
							bind:value={confirmPassword}
							required
						/>
					</div>
				{/if}
				<button
					class="w-full p-[12px] bg-logic-teal text-white border-none rounded-xl font-semibold cursor-pointer transition-[background-color] duration-200 mt-[10px] hover:bg-accent-hover"
					type="submit"
					onclick={() =>
						console.log(`Email or username: ${usernameOrEmail}, password: ${password}`)}
					>{submitText}</button
				>
			</form>
			<div
				class="mt-[20px] text-center text-[0.85rem] text-slate-mist no-underline hover:underline"
			>
				<a href={resolve('/')}>Forgot password?</a>
			</div>
		</div>
	</div>
</div>
