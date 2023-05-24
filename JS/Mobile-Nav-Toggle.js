function toggleOverlay() {
    let overlay = document.getElementById("toggle-overlay");
    if(overlay.style.display === "flex") {
        overlay.style.display = "none";
    } else {
        overlay.style.display = "flex";
    }
};