function toggleOverlay() {
    let toggle = document.getElementById("nav-toggle");

    if (toggle.style.transform === "rotate(90deg)") {
        toggle.style.transform = "rotate(0)";
    } else {
        toggle.style.transform = "rotate(90deg)";
    };
};