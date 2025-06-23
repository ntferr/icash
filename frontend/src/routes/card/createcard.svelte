<script>
    class Card {
        number
        bank_id
        expire_at
    }

    async function createCard() {
        let card = new Card();
        card.bank_id = document.getElementById("bankID").value;
        card.number = document.getElementById("cardNumber").value;
        card.expire_at = document.getElementById("expireAt").value;

        const res = await fetch('http://127.0.0.1:8000/cards', {
            method: 'POST',
            headers: {
                'Accept':'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(card)
        });
    
        location.reload(true)

    }
</script>

<button on:click={createCard}>Send</button>
