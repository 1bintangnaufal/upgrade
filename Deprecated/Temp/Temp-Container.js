// Dark mode toggle

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