import { createCard, renderFallback } from './card.js';
import { isLogged, logout } from './utils.js';
import { URL } from './utils.js';

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
    isLogged().then(logged => {
        if (logged) {
            document.getElementById('login-a').style.display = 'none';
    
            let logout_var = document.getElementById('logout-a');
            logout_var.style.display = 'block';
            logout_var.addEventListener('click', logout);
    
            document.getElementById('register-a').style.display = 'none';   
            document.getElementById('account-a').style.display = 'block';
        } else {
            document.getElementById('login-a').style.display = 'block';
            document.getElementById('logout-a').style.display = 'none';
            document.getElementById('register-a').style.display = 'block';
            document.getElementById('account-a').style.display = 'none';
        }
    });
    
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
