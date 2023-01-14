function openPopUp(element) {
    element = element.parentElement.parentElement;
    element.nextElementSibling.style.display = "flex";
    document.body.style.overflow = "hidden";
}

function closePopUp(element) {
    element = element.parentElement.parentElement.parentElement;
    document.body.style.overflow = "auto";
    element.style.display = "none";
}