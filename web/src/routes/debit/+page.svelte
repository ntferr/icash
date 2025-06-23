<script>
    import { cardStore } from "../../store/card"
    import Createdebit from "./createdebit.svelte";
</script>

<div class=header>
    <h1>Morpheus</h1>
</div>

<div>
    <div class=content>
        <a href="/..">&#129184;</a>
        <h2>Débitos</h2>
        <form>
            <div>
                <select id="cardID">
                    {#each $cardStore as card }
                        <option value="{card.id}">{card.number}</option>
                    {/each}
                </select>
                <div for="debitDiv">
                    <div for="debitNameDiv">
                        <label for="debitHolderLabel">Name<abbr title="required"></abbr></label>
                        <br/>
                        <input type="text" name="debitHolder" id="debitHolder"/>
                    </div>
                    <div for="debitDescriptionDiv">
                        <label for="debitDescriptionLabel">Description</label>
                        <br/>
                        <input type="text" name="debitDescription" id="debitDescription"/>
                    </div>
                    <div for="debitValueDiv">
                        <label for="debitValueLabel">Value<abbr title="required"></abbr></label>
                        <br/>
                        <input type="number" step="0.01" name="debitValue" id="debitValue">
                    </div>
                    <div for="debitRecurrentDiv">
                        <label for="debitRecurrentLabel">Recurrent</label>
                        <br/>
                        <input type="checkbox" name="debitRecurrent" id="debitRecurrent"/>
                    </div>
                    <div for="debitPaidDiv">
                        <label for="debitPaidLabel">Paid</label>
                        <br/>
                        <input type="checkbox" name="debitPaid" id="debitPaid"/>
                    </div>
                    <div for="debitDueDateDiv">
                        <label for="debitDueDateLabel">Due Date<abbr title="required"></abbr></label>
                        <br/>
                        <input type="date" name="debitDate" id="debitDate" />
                    </div>
                </div>
                <div for="installmentDiv">
                    <div for="installmentNumberDiv">
                        <label for="installmentNumberDiv">Instalment Number</label>
                        <br/>
                        <input type="number" step="0" name="instalmentNumber" id="instalmentNumber">
                    </div>
                </div>
                <div for="ticketDiv">
                    <label for="ticketBars">Ticket Bars</label>
                    <br/>
                    <input type="text" name="ticketBars" id="ticketBars"/>
                </div>
            </div>
            <Createdebit/>
        </form>
    </div>
</div>