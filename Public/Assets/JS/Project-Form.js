let project_data = [];

function post_project(event) {
    event.preventDefault();

    let project_title = document.getElementById("project-title").value;
    if (project_title.length > 30) {
        return alert("Maximum Characters For Project Title Exceeds 30.")
    };

    let start_date = document.getElementById("start-date").value;
    let finish_date = document.getElementById("finish-date").value;
    let description = document.getElementById("description").value;
    if(description.length < 80) {
        return alert("Minimum Characters For Project Description (80) Not Fullfilled")
    };

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
    if (form_check_input.length === 0) {
        return alert("Select at least one technology used.");
    };

    let js_check = document.getElementById("js-check").checked ? js_i : "";
    let bs_check = document.getElementById("bs-check").checked ? bs_i : "";
    let go_check = document.getElementById("go-check").checked ? go_i : "";
    let react_check = document.getElementById("react-check").checked ? react_i : "";

    upload_image = URL.createObjectURL(upload_image[0]);

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

    const get_toast = document.getElementById("toast-container");
    const show_toast = new bootstrap.Toast(get_toast);
    show_toast.show();

    project_data.push(ppc);
    console.log(project_data);

    render_ppc();

    close_modal();
};

function render_ppc() {
    document.getElementById("ppc-container").innerHTML = "";

    for (let index = 0; index < project_data.length; index++) {
        const start_date = new Date(project_data[index].start_date);
        const finish_date = new Date(project_data[index].finish_date);
        const remainder = finish_date - start_date;
        const time_units = [
            { value: 365.25 * 24 * 60 * 60 * 1000, label: "year(s)" },
            { value: 30 * 24 * 60 * 60 * 1000, label: "month(s)" },
            { value: 7 * 24 * 60 * 60 * 1000, label: "week(s)" },
            { value: 24 * 60 * 60 * 1000, label: "day(s)" },
        ];

        let result = "";
        for (let calculation = 0; calculation < time_units.length; calculation++) {
            const { value, label } = time_units[calculation];
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
        <div class="rounded-4 border-0 shadow-sm ppc" style="width: 16em;">
            <img src="${project_data[index].upload_image}" class="card-img-top rounded-top-4" alt="Mobile App" style="height: 10.25em; object-fit: cover;">
            <div class="card-body p-3">
              <a href="/Project-Detail/:id" class="card-title ppc-title" style="text-decoration: none;">
                <h6 class="text-truncate">${project_data[index].project_title}</h6>
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
                <button class="btn rounded-pill btn-outline-secondary btn-sm w-50">Edit</button>
                <button class="btn rounded-pill btn-outline-danger btn-sm w-50">Delete</button>
              </div>
            </div>
          </div>
        `;
    };

    document.getElementById("ppc-container").scrollIntoView({ behavior: 'smooth' });
};

function close_modal() {
    const project_form_modal = document.getElementById("add-new");

    const bootstrap_modal = bootstrap.Modal.getInstance(project_form_modal);
    bootstrap_modal.hide();
};

