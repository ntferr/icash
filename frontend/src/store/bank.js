import { writable } from "svelte/store";

export const bankStore = writable([])

export async function fetchBanks() {
    const res = await fetch('http://127.0.0.1:8000/banks', {
                method: 'GET',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                },
            });
           const listOfBanks = await res.json();
           bankStore.set(listOfBanks)
}