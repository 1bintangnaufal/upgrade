async function Post_Project(event) {
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
  let Today = new Date().toLocaleDateString().split("/").join("-");
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
  let Form_Data = new FormData(event.target);
  try {
    const Response = await fetch("/", {
      method: "POST",
      body: Form_Data,
    });
    if (Response.ok) {
      localStorage.setItem("Tasty_Toasty_Toast", "true");
      Close_Modal();
      location.reload();
    } else {
      alert("Fetch failed.");
    }
  } catch (error) {
    alert("Oh no....");
    console.error(error);
  }
}

function Close_Modal() {
  const Project_Form_Modal = document.getElementById("add-new");
  const Bootstrap_Modal = bootstrap.Modal.getInstance(Project_Form_Modal);
  Bootstrap_Modal.hide();
}

window.onload = () => {
  if (localStorage.getItem("Tasty_Toasty_Toast") === "true") {
    const Get_Toast = document.getElementById("toast-container");
    const Show_Toast = new bootstrap.Toast(Get_Toast);
    Show_Toast.show();
    localStorage.removeItem("Tasty_Toasty_Toast");
    Auto_Scroll();
  }
};

function Auto_Scroll() {
  const PPC_Container = document.getElementById("ppc-container");
  const PPC = PPC_Container.getElementsByClassName("ppc");
  const Recent_Card = PPC[PPC.length - 1];
  Recent_Card.scrollIntoView({ behavior: "smooth" });
}
