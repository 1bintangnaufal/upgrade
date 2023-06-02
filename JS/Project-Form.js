let project_data = [];

function post_project(event) {
    event.preventDefault();

    let project_title = document.getElementById("project-title").value;
    let start_date = document.getElementById("start-date").value;
    let finish_date = document.getElementById("finish-date").value;
    let description = document.getElementById("description").value;
    let upload_image = document.getElementById("upload-image").files;

    let today = new Date().toLocaleDateString().split("/").join("-");
    if (finish_date > today) {
        return alert("Time travel is not yet invented.");
    };

    const js_i = '<i class="fa-brands fa-square-js fa-lg fa-fw"></i>';
    const bs_i = '<i class="fa-brands fa-bootstrap fa-lg fa-fw"></i>';
    const go_i = '<i class="fa-brands fa-golang fa-lg fa-fw"></i>';
    const react_i = '<i class="fa-brands fa-react fa-lg fa-fw"></i>';

    let form_check_input = document.querySelectorAll(".form-check-input:checked");
    if(form_check_input.length === 0) {
        return alert("Select at least one technology used.");
    };

    let js_check = document.getElementById("js-check").checked ? js_i : "";
    let bs_check = document.getElementById("bs-check").checked ? bs_i : "";
    let go_check = document.getElementById("go-check").checked ? go_i : "";
    let react_check = document.getElementById("react-check").checked ? react_i : "";

    upload_image = URL.createObjectURL(upload_image[0]);
    console.log(upload_image);

    const sd_validation = new Date(start_date);
    const fd_validation = new Date(finish_date);
    if (sd_validation > fd_validation) {
        return alert("Please input your dates correctly.");
    };

    let ppc = {
        project_title,
        start_date,
        finish_date,
        description,
        js_check,
        bs_check,
        go_check,
        react_check,
        upload_image,
    };

    project_data.push(ppc);
    console.log(project_data);

    render_ppc();

    
};

const textarea = document.getElementById("description");
textarea.addEventListener("input", function() {
    const min_char = 80;
    const input_length = this.value.length;

    if (input_length < min_char) {
        textarea.setCustomValidity("Minimum " + min_char + " characters required.")
    } else {
        textarea.setCustomValidity("");
    }
});

function render_ppc() {
    document.getElementById("ppc-container").innerHTML = "";

    for (let index = 0; index < project_data.length; index++) {
        const start_date = new Date(project_data[index].start_date);
        const finish_date = new Date(project_data[index].finish_date);
        const remainder = finish_date - start_date;
        const time_units = [
            {value: 365.25 * 24 * 60 * 60 * 1000, label: "year(s)"},
            {value: 30 * 24 * 60 * 60 * 1000, label: "month(s)"},
            {value: 7 * 24 * 60 * 60 * 1000, label: "week(s)"},
            {value: 24 * 60 * 60 * 1000, label: "day(s)"},
        ];

        let result = "";
        for (let calculation = 0; calculation < time_units.length; calculation++) {
            const {value, label} = time_units[calculation];
            const calculate = Math.floor(remainder / value);
            if (calculate > 0) {
                result = `${calculate} ${label}`;
                break;
            };
        };

        if (result === "") {
            result = "Less than a day";
        };

        document.getElementById("ppc-container").innerHTML += `
        <div class="card" style="width: 16em;">
            <img src="${project_data[index].upload_image}" class="card-img-top" alt="Mobile App">
            <div class="card-body">
              <a href="#" class="card-title" style="text-decoration: none;">
                <h6>${project_data[index].project_title}</h6>
              </a>
              <p class="text-muted" style="font-size: small; line-height: .5;">${result}</p>
              <p class="card-text lh-sm"
                style="font-size: small; display: -webkit-box; -webkit-box-orient: vertical; -webkit-line-clamp: 3; overflow: hidden;">
                ${project_data[index].description}</p>
              <div class="d-flex gap-2 my-4">
                ${project_data[index].js_check}
                ${project_data[index].bs_check}
                ${project_data[index].go_check}
                ${project_data[index].react_check}
              </div>
              <div class="d-flex flex-row gap-3">
                <button class="btn btn-outline-secondary btn-sm w-50">Edit</button>
                <button class="btn btn-outline-danger btn-sm w-50">Delete</button>
              </div>
            </div>
          </div>
        `;
    };

    document.getElementById("ppc-container").scrollIntoView({behavior: 'smooth'});
};