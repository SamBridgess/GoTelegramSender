const url = "request"

const inputForm = document.getElementById("inputForm");
inputForm.addEventListener("submit", (e)=>{
    e.preventDefault()
    $('.response_result').html("");
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
        msg = "<success>" + response.message + "</success>"
    else
        msg = "<error>" + response.message + "</error>"
    $('.response_result').html(msg);
}
function validateUsername(username) {
    let errorInfo = "";
    if(username.length === 0)
        errorInfo += "<error>Please, enter username</error>"
    else {
        if(username.length < 5)
            errorInfo += "<error>Username is too short</error>"
        else {
            const validTelegramNickname =  /^[A-Za-z\d_]*$/;
            if(!validTelegramNickname.test(username))
                errorInfo += "<error>You can only use a-z, 0-9 and underscores</error>"
        }
    }

    $('#username_error').html(errorInfo);
    return errorInfo === "";
}
function validateMessage(message){
    let errorInfo = "";
    if(message.length === 0)
        errorInfo += "<error>Please, enter message</error>"
    $('#message_error').html(errorInfo);
    return errorInfo === "";
}