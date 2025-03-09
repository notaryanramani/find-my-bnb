import { URL } from './utils.js';

document.getElementById('loginForm').addEventListener('submit', function(e) {
    e.preventDefault();

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    login(username, password);
});

function login(username, password) {
    fetch(URL + '/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
            "username" : username, 
            "password" : password 
        }),
        credentials: 'include'
    })
    .then(response => {
        if (response.ok) {
           window.location.href = '/home';
        }
        else {
            var error = document.getElementById('loginError');
            error.style.display = 'block';
            var errorDisply = document.getElementById('loginErrorDisplay');
            response.text().then(text => {
                errorDisply.innerHTML = text;
            });
        }
    });
}