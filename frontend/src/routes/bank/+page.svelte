<script>
    import Banklist from "./banklist.svelte";'./baklist.svelte'
    import Createbank from "./createbank.svelte";
</script>

<div class=header>
    <h1>Morpheus</h1>
</div>    
<div>
<div class=content>
        <a href="/..">&#129184;</a>
        <h2>Bancos</h2>
        <Banklist/>
        <form>
            <div for="bankNameDiv">
                <label for= "bankNameLabel">Name: <abbr title="required"></abbr></label>
                <br/>
                <input type="text" name="bankName" id="bankName"/>
            </div>

            <div for="bankCodeDiv">
                <label for= "bankCodeLabel">Code: <abbr title="required"></abbr></label>
                <br/>
                <input type="text" name="bankCode" id="bankCode"/>
            </div>
            <Createbank/>
        </form>
    </div>
</div>
