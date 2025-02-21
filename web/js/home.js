const card_class = 'flex gap-3 bg-white border border-gray-300 rounded-xl overflow-hidden items-center justify-start w-4/5';
const card_div_image_class = 'relative w-32 h-32 flex-shrink-0';
const card_image_class = 'absolute left-0 top-0 w-full h-full object-cover object-center transition duration-50';
const card_div_content_class = 'flex flex-col gap-2 py-2';
const card_content_title_class = 'text-xl font-semibold';
const card_content_text_class = 'text-sm text-gray-500';

function fetchRooms(){
    const K = 10;
    fetch('http://localhost:8080/test-rooms', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            'k' : K
        })
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Something went wrong');
        }
        return response.json()
    })
    .then(data => {
        console.log('Success:', data);
        createCard(data, K);
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

function createCard(data, K) {
    const rooms = data['rooms'];
    const parentDiv = document.getElementById('page-content');
    for (let i = 0; i < K; i++){
        let room = rooms[i];

        console.log(room);

        // Main Card
        const card = document.createElement('div');
        card.className = card_class;
        

        // Card Image
        const cardImageDiv = document.createElement('div');
        cardImageDiv.className = card_div_image_class; 

        const cardImage = document.createElement('img');
        cardImage.src = room['picture_url'];
        cardImage.alt = 'Card Image';
        cardImage.className = card_image_class;

        cardImageDiv.appendChild(cardImage);
        card.appendChild(cardImageDiv);

        // Card Content 
        const cardContentDiv = document.createElement('div');
        cardContentDiv.className = card_div_content_class;

        const cardContentTitle = document.createElement('p');
        cardContentTitle.className = card_content_title_class;
        cardContentTitle.innerText = room['name'];

        const cardContentText = document.createElement('p');
        cardContentText.className = card_content_text_class;
        cardContentText.innerText = room['description'];

        cardContentDiv.appendChild(cardContentTitle);
        cardContentDiv.appendChild(cardContentText);

        card.appendChild(cardContentDiv);

        parentDiv.appendChild(card);
    }
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