import { createCard, renderFallback } from './card.js';

var roomIds = [];
let scrollTimeout;

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

document.addEventListener('DOMContentLoaded', function() {
    fetchRooms();

    window.addEventListener('scroll', function() {
        clearTimeout(scrollTimeout);
        
        scrollTimeout = setTimeout(function() {
            document.addEventListener('scroll', function() {
                if (window.scrollY + window.innerHeight >= document.documentElement.scrollHeight - 500) {
                    fetchRooms();
                }
            });    
        }, 200);
    });
});
