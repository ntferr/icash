import { writable } from "svelte/store";

export class card {
    ID;
    BankID;
    Number;
    ExpireAt;
    Bank;
}

export const cardStore = writable([])

export async function fetchCards() {
    const res = await fetch('http://127.0.0.1:8000/cards', {
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },    
    });

    const listOfCards = await res.json();
    cardStore.set(listOfCards);
}