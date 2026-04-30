
// const element = document.querySelector('h1');
// element.style.color = 'red';
// element.style.fontSize = '24px';
// element.style.fontWeight = 'bold';
// element.style.textAlign = 'center';
// element.style.textShadow = '2px 2px 4px rgba(0, 0, 0, 0.5)';

// const button = document.getElementById('btn');
// button.onclick = function() {
//    document.body.style.backgroundColor = 'lightblue';
// }

// const greenButton = document.getElementById('green-btn');
// greenButton.addEventListener('click', function() {
//     document.body.style.backgroundColor = 'lightgreen';
// });

// const dtextButton = document.getElementById('dtext-btn');
// dtextButton.addEventListener('click', function() {
//     const dtext = document.getElementById('dtext');
//     dtext.textContent = 'Text has been changed!';
// });

// const inputButton = document.getElementById('input-btn');
// inputButton.addEventListener('click', function() {
//     const inputText = document.getElementById('input-text');
//     const dtext = document.getElementById('dtext');
//     dtext.textContent = inputText.value
//     inputText.value = '';
// });



// const cmntContainer = document.getElementById('comment-container');

// const commentInput = document.getElementById('comment-input');
// const submitButton = document.getElementById('submit-btn');

// submitButton.addEventListener('click', function() {
//     const commentText = commentInput.value.trim();
//     if (commentText !== '') {
//         const commentElement = document.createElement('p');
//         commentElement.textContent = commentText;
//         cmntContainer.appendChild(commentElement);
//         commentInput.value = '';
//     }
// });

// const inputText = document.getElementById('delete-input');
// const deleteButton = document.getElementById('btnDelete');
// const text = document.getElementById('Text');


// inputText.addEventListener('keyup', function(event) {
//     const value = event.target.value.trim().toLowerCase();
//     if(value === 'delete') 
//         deleteButton.removeAttribute('disabled');
//     else {
//         deleteButton.setAttribute('disabled', true);
//         alert('Please enter "delete" to enable the button.');
//     }

// });
// deleteButton.addEventListener('click', function() {    text.style.display = 'none';
// });

document.getElementsByClassName("li-container")[0].addEventListener('click',function(){
     console.log("container is clicked");
});

document.getElementById("ul-list").addEventListener('click',function(event){
     console.log("ul-list is clicked");
});
document.getElementById("item-1").addEventListener('click',function(event){
     console.log("item-1 is clicked");
});

document.getElementById("item-2").addEventListener('click',function(event){
     console.log("item-2 is clicked");
     event.stopPropagation();
});
