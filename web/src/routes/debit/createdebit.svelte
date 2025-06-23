<script>
    class Installments {
        number
    }
    class Ticket {
        code_bars
    }
    class Debit {
        card_id
        holder
        description
        value
        recurrent
        paid
        due_date
        installments
        ticket
    }

    function mapper() {
        let debit = new Debit;
        debit.installments = new Installments;
        debit.ticket = new Ticket;

        debit.card_id = document.getElementById("cardID").value;
        debit.holder = document.getElementById("debitHolder").value;
        debit.description = document.getElementById("debitDescription").value;
        debit.value = document.getElementById("debitValue").value;
        debit.recurrent = document.getElementById("debitRecurrent").value;
        debit.paid = document.getElementById("debitPaid").value;
        debit.due_date = document.getElementById("debitDate").value;
        debit.installments.number = document.getElementById("instalmentNumber").value
        debit.ticket.code_bars = document.getElementById("ticketBars").value;

        console.log("AQUI", debit)


        return debit;
    }

    async function createDebit() {
        const debit = mapper()


        const res = await fetch('http://127.0.0.1:8000/debts', {
            method: 'POST',
            headers: {
                'Accept':'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(debit)
        });
    
        location.reload(true)
    }

</script>

<button on:click={createDebit}>Send</button>