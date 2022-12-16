<script>
    import Button from "../component/Button.svelte";
    import NavBar from "../component/NavBar.svelte";
    import { getUser } from "../api/user";
    import { auth } from "../store/store";
    import { onMount } from "svelte";
    import { navigate } from "svelte-routing";

    let apiProgress = false

    const onClick = async() => {
        apiProgress = true
        const response = await getUser($auth.header);
        console.log(response.name)
    }

</script>

<NavBar />
<h1>Category</h1>
<Button disabled=false {apiProgress} 
classes="group relative flex w-full justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2" 
{onClick}>
<span class="absolute inset-y-0 left-0 flex items-center pl-1">
    <svg class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400 ml-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
    <path fill-rule="evenodd" d="M10 1a4.5 4.5 0 00-4.5 4.5V9H5a2 2 0 00-2 2v6a2 2 0 002 2h10a2 2 0 002-2v-6a2 2 0 00-2-2h-.5V5.5A4.5 4.5 0 0010 1zm3 8V5.5a3 3 0 10-6 0V9h6z" clip-rule="evenodd" />
    </svg>
</span>
Get User
</Button>
