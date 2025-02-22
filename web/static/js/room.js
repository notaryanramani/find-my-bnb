const container_class = 'flex w-full p-4 justify-center items-center gap-4';
const img_container_class = 'w-1/2 overflow-hidde flex justify-center items-center';
const img_class = 'size-150 object-center rounded-lg';
const details_container_class = 'w-1/2 p-2 flex flex-col gap-2';
const name_class = 'text-xl font-bold text-gray-800';
const detail_component_class = 'flex gap-2';
const details_heading_component = 'text-l font-semibold text-gray-800 w-1/3';
const details_value_component = 'text-sm text-gray-800 w-2/3';

function fetchRoom(roomId) {
    fetch(`http://localhost:8080/api/rooms/${roomId}`)
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
    container.classList.add(...container_class.split(' '));
    
    // image container
    var img_container = document.createElement('div');
    img_container.classList.add(...img_container_class.split(' '));
    var img = document.createElement('img');
    img.src = data['picture_url'];
    img.alt = 'Room Image';
    img.classList.add(...img_class.split(' '));
    img_container.appendChild(img);
    

    // details container
    var details_container = document.createElement('div');
    details_container.classList.add(...details_container_class.split(' '));
    
    var name = document.createElement('div');
    name.classList.add(...name_class.split(' '));
    name.innerText = data['name'];
    details_container.appendChild(name);  

    var description_container = createComponentContainer('Description', data['description']);
    details_container.appendChild(description_container);

    var neigh_overview_container = createComponentContainer('Neighbourhood', data['neighborhood_overview']);
    details_container.appendChild(neigh_overview_container);

    var price = data['price'] != -1 ? '$' + data['price'] : 'Not Available';
    var price_container = createComponentContainer('Price',  price);
    details_container.appendChild(price_container);

    var bedrooms = data['bedrooms'] != -1 ? data['bedrooms'] : 'Not Available';
    var bedrooms_container = createComponentContainer('Bedrooms', bedrooms);
    details_container.appendChild(bedrooms_container);

    var beds = data['beds'] != -1 ? data['beds'] : 'Not Available';
    var beds_container = createComponentContainer('Beds', beds);
    details_container.appendChild(beds_container);

    var roomType_container = createComponentContainer('Room Type', data['room_type']);
    details_container.appendChild(roomType_container);

    var propertyType_container = createComponentContainer('Property Type', data['property_type']);
    details_container.appendChild(propertyType_container);

    var neighborhood_container = createComponentContainer('Neighbourhood', data['neighborhood']);
    details_container.appendChild(neighborhood_container);

    var host_container = createComponentContainer('Host ID: ', data['host_id']);
    details_container.appendChild(host_container);

    // append all the elements
    container.appendChild(img_container);
    container.appendChild(details_container);

    // append the container to the main div
    maindiv.appendChild(container);
}

function createComponentContainer(headingValue, valueValue) {

    var component_container = document.createElement('div');
    component_container.classList.add(...detail_component_class.split(' '));

    var heading = document.createElement('div');
    heading.classList.add(...details_heading_component.split(' '));
    heading.innerText = headingValue;

    var value = document.createElement('div');
    value.classList.add(...details_value_component.split(' '));
    value.innerText = valueValue;

    component_container.appendChild(heading);
    component_container.appendChild(value);

    return component_container;
}


document.addEventListener('DOMContentLoaded', function() {
    const urlParams = new URLSearchParams(window.location.search);
    const roomId = urlParams.get('id');
    
    fetchRoom(roomId);
});
