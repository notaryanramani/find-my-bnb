import { createCard, renderFallback } from "./card.js";

let scrollTimeout;
const k = 10;
var offset = 0;
var queryID = "";
const URL = 'http//localhost:8080/api/vector-search';

function vectorSearch(query, offset) {
    var body = {
        'text' : query,
        'k' : k,
        'offset' : offset,
        'query_id' : queryID,
    }

    fetch(URL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Something went wrong');
        }
        return response.json()
    })
    .then(data => {
        console.log('Success:', data);
        queryID = data['query_id'];
        offset += k;
        createCard(data, k);
    })
    .catch((error) => {
        console.error('Error:', error);
        renderFallback();
    });
}

document.getElementById('search-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const query = document.getElementById('search-input').value;
    vectorSearch(query, offset);
    offset += k;
});

document.addEventListener('DOMContentLoaded', function() {
    document.addEventListener('scroll', function () {
        clearTimeout(scrollTimeout); 

        scrollTimeout = setTimeout(() => {
            if (window.scrollY + window.innerHeight >= document.documentElement.scrollHeight - 500) {
                const query = document.getElementById('search-input').value;
                vectorSearch(query, offset);
                console.log(offset);
                offset += k;
            }
        }, 200); 
    });
});