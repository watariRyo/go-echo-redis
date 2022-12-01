<script>
    export let label = "", required = false, classes="", id = "", type="text", placeholder="", help = "";
</script>

<div class="form-group">
    <div>
        <label for={id} class="sr-only">{label}</label>
        <input type={type} id={id} required={required} class={classes} placeholder={placeholder} on:input>
    </div>
    {#if help}
        <span class="invalid-feedback" role="alert">{help}</span>
    {/if}
</div>