<script>
    export let disabled, apiProgress, onClick, classes = ""
</script>

<button class={classes}
    disabled={disabled | apiProgress}
    on:click|preventDefault={onClick}>
    <slot></slot>
</button>
