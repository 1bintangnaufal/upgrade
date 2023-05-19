function emptyFormAlert() {
    let projectName = document.getElementById("project-name-input").value;
    let startDate = document.getElementById("start-date-input").value;
    let finishDate = document.getElementById("finish-date-input").value;
    let description = document.getElementById("description-input").value;
    let image = document.getElementById("image-input").value;
    
    if(projectName == "") {
        return alert("Please input your project name or title");
    } else if(startDate == "") {
        return alert("When did you start this project?");
    } else if(finishDate == "") {
        return alert("When did you finish this project?");
    } else if(description == "") {
        return alert("Please describe this project");
    } else if(image == "") {
        return alert("Please attach an image of your project");
    }
};

let projectData = [];

function postProject(event) {
    event.preventDefault();

    let projectName = document.getElementById("project-name-input").value;
    let startDate = document.getElementById("start-date-input").value;
    let finishDate = document.getElementById("finish-date-input").value;
    let description = document.getElementById("description-input").value;
    let image = document.getElementById("image-input").files;

    const jsIcon = '<i class="fa-brands fa-square-js fa-xl fa-fw"></i>';
    const bootstrapIcon = '<i class="fa-brands fa-bootstrap fa-xl fa-fw"></i>';
    const goIcon = '<i class="fa-brands fa-golang fa-xl fa-fw"></i>';
    const reactIcon = '<i class="fa-brands fa-react fa-xl fa-fw"></i>';

    let jsIconDecide = document.getElementById("js-check").checked ? jsIcon : "";
    let bootstrapIconDecide = document.getElementById("bootstrap-check").checked ? bootstrapIcon : "";
    let goIconDecide = document.getElementById("go-check").checked ? goIcon : "";
    let reactIconDecide = document.getElementById("react-check").checked ? reactIcon : "";

    image = URL.createObjectURL(image[0]);
    console.log(image);

    let projectPreviewCard = {
        projectName,
        startDate,
        finishDate,
        description,
        jsIconDecide,
        bootstrapIconDecide,
        goIconDecide,
        reactIconDecide,
        image,
    };

    projectData.push(projectPreviewCard);
    console.log(projectData);

    renderPpc();
}

function renderPpc() {
    document.getElementById("wrap-up-the-cards").innerHTML = "";

    for (let index = 0; index < projectData.length; index++) {
        document.getElementById("wrap-up-the-cards").innerHTML += `
        <div class="project-preview-card">
                <img src="${projectData[index].image}" alt="Smartphone"/>
                <a href="#">
                    <h4>${projectData[index].projectName}</h4>
                </a>
                <p class="duration">${projectData[index].startDate} - ${projectData[index].finishDate}</p>
                <p class="description">${projectData[index].description}</p>
                <div class="tech-icons">
                    ${projectData[index].jsIconDecide}
                    ${projectData[index].bootstrapIconDecide}
                    ${projectData[index].goIconDecide}
                    ${projectData[index].reactIconDecide}
                </div>
                <div class="card-buttons">
                    <button class="edit-button">Edit</button>
                    <button class="delete-button">Delete</button>
                </div>
            </div>
        `
    }
}