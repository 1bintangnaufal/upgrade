function sendMessage() {
    let name = document.getElementById("name-input").value;
    let email = document.getElementById("email-input").value;
    let phone = document.getElementById("phone-input").value;
    let subject = document.getElementById("subject-input").value;
    let message = document.getElementById("message-input").value;
    
    if(name == "") {
        return alert("What's your name?");
    } else if(email == "") {
        return alert("What's your email?");
    } else if(phone == "") {
        return alert("What's your phone number?");
    } else if(subject == "") {
        return alert("Please select a subject");
    }

    const sendTarget = "1bintangnaufal@gmail.com";

    let a = document.createElement("a");
    a.href = `mailto:${sendTarget}?subject=${subject}&body=Hi Bintang! My name is ${name}. ${message}. Please kindly call or text me on ${phone} if you are interested. Thank you!`;
    a.click();
};