import { URL } from './utils.js';

// TODO: Clean this
const classContainer = 'flex w-full p-4 justify-center items-center gap-4';
const classImgContainer = 'w-1/2 overflow-hidde flex justify-center items-center';
const classImg = 'size-125 object-center rounded-lg';
const classDetailsContainer = 'w-1/2 p-2 flex flex-col gap-2';
const className = 'text-xl font-bold text-gray-800';
const classDetailComponent = 'flex gap-2';
const classDetailsHeadingComponent = 'text-l font-semibold text-gray-800 w-1/3';
const classDetailsValueComponent = 'text-sm text-gray-800 w-2/3';

const numericFields = ['price', 'bedrooms', 'beds'];
const textFields = ['description', 'neighborhood_overview', 'room_type', 'property_type', 'neighborhood', 'host_id'];

const fieldsToInclude = ['description', 'neighborhood_overview', 'price', 'bedrooms', 'beds', 'room_type', 'property_type', 'neighborhood', 'host_id'];

function getRoom(roomId) {
    fetch(URL + `/rooms/${roomId}`)
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
    container.classList.add(...classContainer.split(' '));
    
    // image container
    var imgContainer = document.createElement('div');
    imgContainer.classList.add(...classImgContainer.split(' '));

    var img = document.createElement('img');
    img.src = data['picture_url'];
    img.alt = 'Room Image';
    img.classList.add(...classImg.split(' '));

    imgContainer.appendChild(img);

    

    // details container
    var detailsContainer = document.createElement('div');
    detailsContainer.classList.add(...classDetailsContainer.split(' '));
    
    var name = document.createElement('div');
    name.classList.add(...className.split(' '));
    name.innerText = data['name'];
    detailsContainer.appendChild(name);  


    // details container loop
    fieldsToInclude.forEach(field => {
        var value = data[field];
        if(numericFields.includes(field)) {
            value = value != -1 ? value : 'Not Available';
            if (field == 'price') {
                value = '$' + value;
            }
        } else if(textFields.includes(field)) {
            value = value != 'N/A' ? value : 'Not Available';
        }

        var detailContainer = createComponentContainer(field, value);
        detailsContainer.appendChild(detailContainer);
    });

    // append all the elements
    container.appendChild(imgContainer);
    container.appendChild(detailsContainer);

    // append the container to the main div
    maindiv.appendChild(container);
}

function createComponentContainer(headingValue, valueValue) {

    var componentContainer = document.createElement('div');
    componentContainer.classList.add(...classDetailComponent.split(' '));

    var heading = document.createElement('div');
    heading.classList.add(...classDetailsHeadingComponent.split(' '));
    heading.innerText = headingValue;

    var value = document.createElement('div');
    value.classList.add(...classDetailsValueComponent.split(' '));
    value.innerText = valueValue;

    componentContainer.appendChild(heading);
    componentContainer.appendChild(value);

    return componentContainer;
}

document.addEventListener('DOMContentLoaded', function() {
    const urlParams = new URLSearchParams(window.location.search);
    const roomId = urlParams.get('id');
    
    getRoom(roomId);
});
