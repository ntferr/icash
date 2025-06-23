<script>
        import { bankStore } from "../../store/bank.js"
        import CreateCard from "./createcard.svelte"
        import CardList from "./cardlist.svelte"
</script>

<div class=header>
    <h1>Morpheus</h1>
</div>

<div>
    <div class=content>
        <a href="/..">&#129184;</a>
        <h2>Cartões</h2>
        <form>
            <div>
                <select id="bankID">
                    {#each $bankStore as bank}
                        <option value="{bank.id}">{bank.name}</option>
                    {/each}
                </select>
                <div for="cardNumberDiv">
                    <label for="cardNumberLabel">Number: <abbr title="required"></abbr></label>
                    <br/>
                    <input type="text" id="cardNumber"/>
                </div>
                <div for="expireAtDiv">
                    <label for= "expireAtLabel">Expire Date: <abbr title="required"></abbr></label>
                    <br/>
                    <input type="month" id="expireAt" value="2023-01"/>
                </div>
                <br/>
                <CreateCard />
            </div>
        </form>
        <div class=cards>
            <CardList />
        </div>
    </div>
</div>
