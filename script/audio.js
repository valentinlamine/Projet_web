function PlaySoundStart() {
    var sound = document.getElementById('thesecretsound');
    console.log(sound);
    sound.volume = 0.25;
    if (sound.paused) {
        sound.play();
    }
    else {
        sound.pause();
    }
}