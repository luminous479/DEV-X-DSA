
const element = document.querySelector('h1');
element.style.color = 'red';
element.style.fontSize = '24px';
element.style.fontWeight = 'bold';
element.style.textAlign = 'center';
element.style.textShadow = '2px 2px 4px rgba(0, 0, 0, 0.5)';

const button = document.getElementById('btn');
button.onclick = function() {
   document.body.style.backgroundColor = 'lightblue';
}

const greenButton = document.getElementById('green-btn');
greenButton.addEventListener('click', function() {
    document.body.style.backgroundColor = 'lightgreen';
});

const dtextButton = document.getElementById('dtext-btn');
dtextButton.addEventListener('click', function() {
    const dtext = document.getElementById('dtext');
    dtext.textContent = 'Text has been changed!';
});

const inputButton = document.getElementById('input-btn');
inputButton.addEventListener('click', function() {
    const inputText = document.getElementById('input-text');
    const dtext = document.getElementById('dtext');
    dtext.textContent = inputText.value
    inputText.value = '';
});