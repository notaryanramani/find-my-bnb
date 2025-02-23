document.getElementById('loginForm').addEventListener('submit', function(e) {
    e.preventDefault();

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    login(username, password)
    .then(response => {
        console.log(response);
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
});

async function login(username, password) {
    const response = await fetch('http://localhost:8080/api/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ 
            "username" : username, 
            "password" : password 
        })
    });
    return response; 
}