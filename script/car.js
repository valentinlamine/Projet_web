//i have a list of cars that is generated with a range from golang backend 
//i want when i click on the button "Voir plus de détails" to show the absolute box of the details of the car
//i need to get the id of the car that i clicked on and then show the details of this car

//here is the html code of the list of cars
{/* <div class="cards">
    {{range .}}
    <div class="item">
        <img src={{.Image}} alt="ford mustang">
        <div class = "bottom-item">
            <div class="info">
                <div class="firstgroup">
                    <p class="name">{{.Brand}}</p>
                    <p class="km">{{.Kilometers}} KM</p>
                </div>
                <p class="price">{{.Price}} €</p>
            </div>
            <div class="details">
                <p>Voir plus de détails</p>
            </div>
        </div>
    </div> 
    {{end}} 
</div> */}

//here is the absolute car details
{/* <div class="pop-up">
    <img src="images/car_1.png" alt="ford mustang">
    <div class="bottom-pop-up">
        <p>Fiche technique</p>
        <div class="main-info-pop-up">
            <div class="first-line">
                <h1>{{.Brand}}</h1>
                <h1>{{.Price}} €</h1>
            </div>
            <div class="second-line">
                <p>{{.kilometers}} KM</p>
                <p>{{.Year}}</p>
            </div>
        </div>
        <div class="secondary-info-pop-up">
            <p>Moteur : <span class="bolder">{{.Engine}}</span></p>
            <p>Vitesse de pointe : <span class="bolder">{{.MaxSpeed}} km/h</span></p>
        </div>
        <input class="button" type="submit" value="Acheter">
    </div>
</div> */}

//Now make the script to get the id of the car that i clicked on and then show the details of this car

//make the function openPopUp() 

// function openPopUp() {
//     const popUp = document.querySelector('.pop-up');
//     popUp.style.display = 'block';
//     console.log('pop up is open');
//     //mask the item 
//     const item = document.querySelector('.item');
//     item.style.display = 'none';
//     console.log('item is hidden');
// }
