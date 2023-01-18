const url = "request"

const inputForm = document.getElementById("inputForm");

inputForm.addEventListener("submit", (e)=>{
    e.preventDefault()

    let formData = new FormData(inputForm)
    let userVal = validateUsername(formData.get("username"));
    let mesVal = validateMessage(formData.get("message"));

    if(userVal && mesVal) {
        fetch(url,{
            method:"POST",
            body:formData,
        }).then(
            function(response) {
                response.json().then(function(data) {
                    addResponseResult(data)
                });
            }
        )
    }
})
function addResponseResult(response) {
    let msg = ""
    if(response.status === "success")
        msg = "<span style=\"color:green\">" + response.message + "</span>"
    else
        msg = "<span style=\"color:red\">" + response.message + "</span>"
    $('.response_result').html(msg);
}
function validateUsername(username) {
    let errorInfo = "";

    if(username.length === 0)
        errorInfo += "<span>Please, enter username</span>"
    else {
        if(username.length < 5)
            errorInfo += "<span>Username is too short</span>"
        else {
            const validTelegramNickname =  /^[A-Za-z\d_]*$/;
            if(!validTelegramNickname.test(username))
                errorInfo += "<span>You can only use a-z, 0-9 and underscores</span>"
        }
    }

    $('.username_error').html(errorInfo);
    return errorInfo === "";
}
function validateMessage(message){
    let errorInfo = "";
    if(message.length === 0)
        errorInfo += "<span>Please, enter message</span>"
    $('.message_error').html(errorInfo);
    return errorInfo === "";
}