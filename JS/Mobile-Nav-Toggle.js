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



let choice = localStorage.getItem('choice');
const themeSwitchToggle = document.getElementById("theme-switch");

const darkMode = () => {
    document.body.classList.add("dark-mode");
    localStorage.setItem("choice", "on");
};

const lightMode = () => {
    document.body.classList.remove("dark-mode");
    localStorage.setItem("choice", "off");
};

if(choice === "on") {
    darkMode();
};

themeSwitchToggle.addEventListener("click", () => {
    choice = localStorage.getItem("choice");
    if(choice !== "on") {
        darkMode();
    } else {
        lightMode();
    }
});