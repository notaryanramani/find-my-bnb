import { URL } from './constants.js';
import { updateNavBar } from './utils.js';
import { classes, fields } from './constants.js';


async function getRoom(roomId) {
    await fetch(URL + `/rooms/${roomId}`)
    .then(response => response.json())
    .then(data => {
        console.log('Success:', data);
        createRoomCard(data);
    })
}

function createRoomCard(data) {
    var maindiv = document.getElementById('target-div');

    // main container
    var container = document.createElement('div');
    container.classList.add(...classes.classContainer.split(' '));
    
    // image container
    var imgContainer = document.createElement('div');
    imgContainer.classList.add(...classes.classImgContainer.split(' '));

    var img = document.createElement('img');
    img.src = data['picture_url'];
    img.alt = 'Room Image';
    img.classList.add(...classes.classImg.split(' '));

    imgContainer.appendChild(img);

    // details container
    var detailsContainer = document.createElement('div');
    detailsContainer.classList.add(...classes.classDetailsContainer.split(' '));
    
    var name = document.createElement('div');
    name.classList.add(...classes.className.split(' '));
    name.innerText = data['name'];
    detailsContainer.appendChild(name);  


    // details container loop
    fields.fieldsToInclude.forEach(field => {
        var value = data[field];
        if(fields.numericFields.includes(field)) {
            value = value != -1 ? value : 'Not Available';
            if (field == 'price') {
                value = '$' + value;
            }
        } else if(fields.textFields.includes(field)) {
            value = value != 'N/A' ? value : 'Not Available';
        }

        var detailContainer = createComponentContainer(field, value);
        detailsContainer.appendChild(detailContainer);
    });

    // Create a like button
    var likeSpan = document.createElement('div');
    likeSpan.className = "flex justify-center items-center";
    var likeButton = document.createElement('i');
    likeButton.id = 'like-button';
    likeButton.className = "fa-regular fa-heart hover:cursor-pointer";
    likeSpan.appendChild(likeButton);
    detailsContainer.appendChild(likeSpan);

    likeSpan.addEventListener('click', function() {
        var likeButton = document.getElementById('like-button');
        
        if (likeButton.classList.contains("fa-regular")){
            likeButton.classList.remove("fa-regular");
            likeButton.classList.add("fa-solid");

            // Backend Logic 
        } else {
            likeButton.classList.remove("fa-solid");
            likeButton.classList.add("fa-regular");

            // Backend Logic
        }
    });


    // append all the elements
    container.appendChild(imgContainer);
    container.appendChild(detailsContainer);

    // append the container to the main div
    maindiv.appendChild(container);
}

function createComponentContainer(headingValue, valueValue) {

    var componentContainer = document.createElement('div');
    componentContainer.classList.add(...classes.classDetailComponent.split(' '));

    var heading = document.createElement('div');
    heading.classList.add(...classes.classDetailsHeadingComponent.split(' '));
    headingValue = formatHeading(headingValue);
    heading.innerText = headingValue;

    var value = document.createElement('div');
    value.classList.add(...classes.classDetailsValueComponent.split(' '));
    value.innerText = valueValue;

    componentContainer.appendChild(heading);
    componentContainer.appendChild(value);

    return componentContainer;
}


// Formats heading to be more readable: 'neighbourhood_overview' -> 'Neighbourhood Overview'
function formatHeading(heading) {
    var headingArray = heading.split('_');
    headingArray.forEach((word, index) => {
        headingArray[index] = word.charAt(0).toUpperCase() + word.slice(1);
    });
    return headingArray.join(' ');
}

document.addEventListener('DOMContentLoaded', function() {
    updateNavBar();

    const urlParams = new URLSearchParams(window.location.search);
    const roomId = urlParams.get('id');
    
    getRoom(roomId);
});
