<script>
    import {onMount} from 'svelte'
    import {cardStore, fetchCards} from '../../store/card.js'

    onMount(async () => {
        await fetchCards()
    });
</script>

<div class="listOfCards">
    <ul>
        <ul> <h5><b>Número - Expira em</b></h5> </ul>
        {#each $cardStore as card}
            <li>{card.number} - {card.expire_at}</li>
        {/each}
    </ul>
</div>