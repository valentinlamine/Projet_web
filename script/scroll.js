function scrollToBottom() {
    console.log("scrolling to bottom");
    window.scrollTo({
        top: window.innerHeight,
        behavior: 'smooth'
    });
}

function scrollToTop() {
    console.log("scrolling to top");
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    });
}