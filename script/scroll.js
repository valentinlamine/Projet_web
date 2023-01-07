//footer button
const buttonToTop = document.querySelector(".button-to-top");
buttonToTop.addEventListener("click", () => {
    window.scrollTo({
        top: 0,
        behavior: "smooth"
    });
});


//index page
function scrollToBottom() {
    window.scrollTo({
        top: window.innerHeight,
        behavior: 'smooth'
    });
}

function scrollToTop() {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    });
}