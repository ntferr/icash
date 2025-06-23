<script>
        import {onMount} from 'svelte'
        import {bankStore, fetchBanks} from '../../store/bank.js'

        onMount(async () => {
            await fetchBanks()
        });
</script>

<div class="listOfBanks">
    <ul>
        {#each $bankStore as bank}
            <li>{bank.name} - {bank.code}</li>
        {/each}
    </ul>
</div>