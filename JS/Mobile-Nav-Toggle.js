function toggleOverlay() {
    let overlay = document.getElementById("toggle-overlay");
    let toggle = document.getElementById("nav-toggle");

    if (overlay.style.gridTemplateRows === "1fr") {
        overlay.style.gridTemplateRows = "0fr";
    } else {
        overlay.style.gridTemplateRows = "1fr";
    };

    if (toggle.style.transform === "rotate(90deg)") {
        toggle.style.transform = "rotate(0)";
    } else {
        toggle.style.transform = "rotate(90deg)";
    };
};