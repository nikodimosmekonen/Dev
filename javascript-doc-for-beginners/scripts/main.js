const myHeading = document.querySelector("h1");
myHeading.textContent = "Hello world!";

const myImage = document.querySelector("img");
myImage.addEventListener("click", () => {
  const mySrc = myImage.getAttribute("src");
  if (mySrc === "image/hero-image.jpg") {
    myImage.setAttribute("src", "image/Firefox_logo,_2019.svg.png");
  } else {
    myImage.setAttribute("src", "image/hero-image.jpg");
  } 
});
let myButton = document.querySelector("button");
function setUserName(){
    const myName = prompt('Please enter your name.');
    if(!myName || myName === null){
        setUserName();
    }
    localStorage.setItem('name', myName);
    myHeading.textContent = 'Mozilla is cool, ' + myName;
}
if(!localStorage.getItem('name')){
    setUserName();
} else {
    let storedName = localStorage.getItem('name');
    myHeading.textContent = 'Mozilla is cool, ' + storedName;
}
myButton.onclick = function(){
    setUserName();
}