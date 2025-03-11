import { URL } from './constants.js';

async function checkLogin() {
    return await fetch(URL + '/auto-login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include'
    })
    .then(response => {
        if (!response.ok) {
            return response.text().then(errorMessage => {
                console.log("Error:", errorMessage);
                return false; 
            });
        }
        return true; 
    })
    .catch(error => {
        console.error("Fetch error:", error);
        return false;
    });
}

export function updateNavBar() {
    checkLogin().then(logged => {
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
}


export async function logout() {
    await fetch(URL+'/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include'
    })
}