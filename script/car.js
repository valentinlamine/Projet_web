function openPopUp(element) {
    element = element.parentElement.parentElement;
    element.nextElementSibling.style.display = "flex";
    document.body.style.overflow = "hidden";
    //element.style.display = "none";
    //element.nextElementSibling.scrollIntoView({behavior: "smooth", block: "center", inline: "center"});

}
function closePopUp(element) {
    element = element.parentElement.parentElement.parentElement;
    document.body.style.overflow = "auto";
    element.style.display = "none";
    //element.previousElementSibling.style.display = "flex";
    //element.previousElementSibling.scrollIntoView({behavior: "smooth", block: "center", inline: "center"});

}