const classCard = 'flex gap-3 rounded-xl items-center justify-start w-4/5';
const classCardDivImage = 'relative w-32 h-32 flex-shrink-0 overflow-hidden';
const classCardImage = 'absolute left-0 top-0 h-full w-full object-cover object-center transition duration-50 rounded-xl';
const classCardDivContent = 'flex flex-col gap-2 py-2';
const classCardContentTitle = 'text-xl font-semibold';
const classCardContentText = 'text-sm overflow-hidden';

function createCard(data, K) {
    const rooms = data['rooms'];
    const parentDiv = document.getElementById('page-content');
    var roomIds = [];
    
    for (let i = 0; i < K; i++) {
        let room = rooms[i];
        roomIds.push(room['id']);

        console.log(room);

        // Main Card
        const card = document.createElement('div');
        card.classList.add(...classCard.split(' '));

        // Card Image
        const cardImageDiv = document.createElement('div');
        cardImageDiv.classList.add(...classCardDivImage.split(' ')); 

        const cardImage = document.createElement('img');
        cardImage.src = room['picture_url'];
        cardImage.alt = 'Card Image';
        cardImage.classList.add(...classCardImage.split(' '));

        cardImageDiv.appendChild(cardImage);
        card.appendChild(cardImageDiv);

        // Card Content 
        const cardContentDiv = document.createElement('div');
        cardContentDiv.classList.add(...classCardDivContent.split(' '));

        const cardContentTitle = document.createElement('a');
        cardContentTitle.classList.add(...classCardContentTitle.split(' '));
        cardContentTitle.innerText = room['name'];
        cardContentTitle.href = `/rooms?id=${room['id_string']}`;

        const cardContentText = document.createElement('p');
        cardContentText.innerText = room['description'];
        cardContentText.classList.add(...classCardContentText.split(' '));

        cardContentDiv.appendChild(cardContentTitle);
        cardContentDiv.appendChild(cardContentText);

        card.appendChild(cardContentDiv);

        parentDiv.appendChild(card);
    }

    return roomIds;
}

export { createCard };