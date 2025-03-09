export const URL = 'http://localhost:8080/api';

export function isLogged() {
    return fetch(URL + '/auto-login', {
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


export async function logout() {
    const response = await fetch(URL+'/logout', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        credentials: 'include'
    })
}