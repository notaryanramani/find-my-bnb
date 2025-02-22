require('dotenv').config();

const express = require('express');
const path = require('path');

const app = express();


app.use(express.static(path.join(__dirname, 'static')));

app.get('/home', (req, res) => {
    res.sendFile(path.join(__dirname, 'static', 'home.html'));
});

app.get('/rooms', (req, res) =>{
    res.sendFile(path.join(__dirname, 'static', 'room.html'));
})

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});