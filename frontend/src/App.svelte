<script>
	import LoginPage from "./pages/LoginPage.svelte";
	import {auth} from "./store/store";
	import HomePage from "./pages/HomePage.svelte";
	import { Router, Route } from "svelte-routing";
	import CategoryPage from "./pages/CategoryPage.svelte";

</script>

<Router url={window.location.pathname}>
	<!-- 開発環境以外（yarn start）だとURL直打ちは失敗する
	rollupがルートパスのindex.htmlのみを提供しているため
	devはfallbackにルートを指定してどのURLでもアプリがロードされるように構成している
	dev以外も同じような設定を入れないといけないが不明
	-->
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