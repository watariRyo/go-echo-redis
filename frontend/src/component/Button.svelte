<script>
    export let disabled, onClick, classes = ""
</script>

<button class={classes}
    disabled={disabled} 
    on:click|preventDefault={onClick}>
    <slot></slot>
</button>
