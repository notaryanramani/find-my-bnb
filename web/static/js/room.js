import { URL } from './constants.js';
import { updateNavBar } from './utils.js';
import { classes, fields } from './constants.js';


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
    heading.innerText = headingValue;

    var value = document.createElement('div');
    value.classList.add(...classes.classDetailsValueComponent.split(' '));
    value.innerText = valueValue;

    componentContainer.appendChild(heading);
    componentContainer.appendChild(value);

    return componentContainer;
}

document.addEventListener('DOMContentLoaded', function() {
    updateNavBar();

    const urlParams = new URLSearchParams(window.location.search);
    const roomId = urlParams.get('id');
    
    getRoom(roomId);
});
