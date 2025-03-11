import { createCard, renderFallback } from './card.js';
import { updateNavBar } from './utils.js';
import { URL } from './constants.js';

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

    fetch(URL + '/test-rooms', {
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
    updateNavBar();
    
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
