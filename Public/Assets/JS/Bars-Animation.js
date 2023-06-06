function rotate_the_bars() {
    let nav_icon = document.getElementById("nav-icon");

    if (nav_icon.style.transform === "rotate(90deg)") {
        nav_icon.style.transform = "rotate(0)";
    } else {
        nav_icon.style.transform = "rotate(90deg)";
    };
};