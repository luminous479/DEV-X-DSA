
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