import { URL } from './constants.js';

document.getElementById('registerForm').addEventListener('submit', function(e) {
    e.preventDefault();

    var name = document.getElementById('name').value;
    var username = document.getElementById('username').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;
    var confirmPassword = document.getElementById('confirmPassword').value;

    if (password != confirmPassword) {
        var error = document.getElementById('registerError');
        error.style.display = 'block';
        var errorDisply = document.getElementById('registerErrorDisplay');
        errorDisply.innerHTML = 'Passwords do not match';
        return;
    }

    const user = {
        name: name,
        username: username,
        email: email,
        password: password
    };

    register(user);
});

function register(user) {
    fetch(URL + '/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            user
        })
    })
    .then(response => {
        console.log(response);
        if (response.ok) {
            window.location.href = '/login';
        }
        else {
            var error = document.getElementById('registerError');
            error.style.display = 'block';
            var errorDisply = document.getElementById('registerErrorDisplay');
            response.text().then(text => {
                errorDisply.innerHTML = text;
            });
        }
    });
}