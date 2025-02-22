import { createCard } from './card.js';

var roomIds = [];

function fetchRooms(){
    const K = 10;

    let requestBody;
    if (roomIds.length == 0){
        requestBody = JSON.stringify({
            'k' : K
        });
    }
    else {
        requestBody = JSON.stringify({
            'k' : K,
            'ids' : roomIds
        });
    }

    fetch('http://localhost:8080/api/test-rooms', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: requestBody
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Something went wrong');
        }
        return response.json()
    })
    .then(data => {
        console.log('Success:', data);
        var newRoomIds = createCard(data, K);
        roomIds = roomIds.concat(newRoomIds);
        console.log(`Room Ids Length: ${roomIds.length}`);
    })
    .catch((error) => {
        console.error('Error:', error);
        renderFallback();
    });
}

function renderFallback(){
    const parentDiv = document.getElementById('page-content');
    const fallback = document.createElement('p');
    fallback.innerText = 'Something went wrong. Please try again later. API is not available. :(';
    parentDiv.appendChild(fallback);
}

document.addEventListener('DOMContentLoaded', function() {
    fetchRooms();
    let count = 10;
    document.addEventListener('scroll', function() {
        if (count < 100 && window.scrollY + window.innerHeight >= document.documentElement.scrollHeight) {
            fetchRooms();
            count += 10;
        }
    });    
});