const card_class = 'flex gap-3 rounded-xl items-center justify-start w-4/5';
const card_div_image_class = 'relative w-32 h-32 flex-shrink-0 overflow-hidden';
const card_image_class = 'absolute left-0 top-0 h-full w-full object-cover object-center transition duration-50';
const card_div_content_class = 'flex flex-col gap-2 py-2';
const card_content_title_class = 'text-xl font-semibold';
const card_content_text_class = 'text-sm overflow-hidden';

function createCard(data, K) {
    const rooms = data['rooms'];
    const parentDiv = document.getElementById('page-content');
    var roomIds = []
    for (let i = 0; i < K; i++){
        let room = rooms[i];
        roomIds.push(room['id']);

        console.log(room);

        // Main Card
        const card = document.createElement('div');
        card.classList.add(...card_class.split(' '));
        

        // Card Image
        const cardImageDiv = document.createElement('div');
        cardImageDiv.classList.add(...card_div_image_class.split(' ')); 

        const cardImage = document.createElement('img');
        cardImage.src = room['picture_url'];
        cardImage.alt = 'Card Image';
        cardImage.classList.add(...card_image_class.split(' '));

        cardImageDiv.appendChild(cardImage);
        card.appendChild(cardImageDiv);

        // Card Content 
        const cardContentDiv = document.createElement('div');
        cardContentDiv.classList.add(...card_div_content_class.split(' '));

        const cardContentTitle = document.createElement('a');
        cardContentTitle.classList.add(...card_content_title_class.split(' '));
        cardContentTitle.innerText = room['name'];
        cardContentTitle.href = `/rooms?id=${room['id_string']}`;

        const cardContentText = document.createElement('p');
        cardContentText.innerText = room['description'];
        cardContentText.classList.add(...card_content_text_class.split(' '));
        

        cardContentDiv.appendChild(cardContentTitle);
        cardContentDiv.appendChild(cardContentText);

        card.appendChild(cardContentDiv);

        parentDiv.appendChild(card);
    }

    return roomIds;
}

export { createCard };