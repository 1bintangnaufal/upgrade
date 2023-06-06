function send_message() {
    let nickname = document.getElementById("nickname").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("phone").value;
    let select_subject = document.getElementById("select-subject").value;
    let message = document.getElementById("message").value;

    const send_target = "1bintangnaufal@gmail.com";

    let a = document.createElement("a");
    a.href = `mailto:${send_target}?subject=${select_subject}&body=Hi Bintang! My name is ${nickname}. ${message}. Please kindly call or text me on ${phone} or reply to this email ${email} if you are interested. Thank you!`;
    a.click();
};