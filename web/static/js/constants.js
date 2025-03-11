export const classes = {
    classContainer: 'flex w-full p-4 justify-center items-center gap-4',
    classImgContainer: 'w-1/2 overflow-hidde flex justify-center items-center',
    classImg: 'size-125 object-center rounded-lg',
    classDetailsContainer: 'w-1/2 p-2 flex flex-col gap-2',
    className: 'text-xl font-bold text-gray-800',
    classDetailComponent: 'flex gap-2',
    classDetailsHeadingComponent: 'text-l font-semibold text-gray-800 w-1/3',
    classDetailsValueComponent: 'text-sm text-gray-800 w-2/3',
};

export const fields = {
    numericFields: ['price', 'bedrooms', 'beds'],
    textFields: ['description', 'neighborhood_overview', 'room_type', 'property_type', 'neighborhood', 'host_id'],
    fieldsToInclude: ['description', 'neighborhood_overview', 'price', 'bedrooms', 'beds', 'room_type', 'property_type', 'neighborhood', 'host_id']
};

export const URL = 'http://localhost:8080/api';