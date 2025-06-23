<script>
    async function createBank() {
        const bankName = document.getElementById("bankName").value;
        const bankCode = document.getElementById("bankCode").value;

        const res = await fetch('http://127.0.0.1:8000/banks', {
            method: 'POST',
            headers: {
                'Accept':'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                "name": bankName,
                "code": bankCode
            })
        })
        
        if(!res.ok) {
            throw new Error(res.json());
        }

        location.reload(true)
    }
</script>

<button on:click={createBank}>Send</button>