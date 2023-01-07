const buttonToTop = document.querySelector(".button-to-top");
buttonToTop.addEventListener("click", () => {
    window.scrollTo({
        top: 0,
        behavior: "smooth"
    });
});