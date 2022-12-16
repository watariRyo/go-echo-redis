<script>
	import LoginPage from "./pages/LoginPage.svelte";
	import {auth} from "./store/store";
	import HomePage from "./pages/HomePage.svelte";
	import { Router, Route } from "svelte-routing";
	import CategoryPage from "./pages/CategoryPage.svelte";

</script>

<Router url={window.location.pathname}>
	<Route path="/">
		{#if $auth.isLoggedIn}
			<HomePage />
		{:else }
			<LoginPage />
		{/if}
	</Route>
	<Route path="/category">
		<CategoryPage />
	</Route>
	<Route path="*">
		<h1>404</h1>
	</Route>
	<h1>TODO：Docker上で動かす</h1>
</Router>