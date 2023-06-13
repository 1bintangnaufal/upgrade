function Post_Project(event) {
  event.preventDefault();

  let Project_Title = document.getElementById("Project_Title").value;
  if (Project_Title.length > 30) {
    alert("Maximum Characters For Project Title Exceeds 30.");
    return;
  }

  let Start_Date = document.getElementById("Start_Date").value;
  let Finish_Date = document.getElementById("Finish_Date").value;

  let Description = document.getElementById("Description").value;
  if (Description.length < 80) {
    alert("Minimum Characters For Project Description (80) Not Fullfilled.");
    return;
  }

  let Current_Date = new Date();
  let Selected_Finish_Date = new Date(Finish_Date);
  if (Selected_Finish_Date > Current_Date) {
    alert("Time travel is not yet invented.");
    return;
  }

  let Switch_Toggle = document.querySelectorAll(".form-check-input:checked");
  if (Switch_Toggle.length === 0) {
    alert("Select at least one technology used.");
    return;
  }

  let SD_Validation = new Date(Start_Date);
  let FD_Validation = new Date(Finish_Date);
  if (
    isNaN(SD_Validation) ||
    isNaN(FD_Validation) ||
    SD_Validation > FD_Validation
  ) {
    alert("Please input your dates correctly.");
    return;
  }

  event.target.submit();
}

const Form = document.getElementById("edit-project-form");
Form.addEventListener("submit", Post_Project)