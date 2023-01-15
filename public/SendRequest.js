const url = "127.0.0.1:8080"

var inputForm = document.getElementById("inputForm")

inputForm.addEventListener("submit", (e)=>{
    e.preventDefault()

    const formdata = new FormData(inputForm)
    fetch(url,{

        method:"POST",
        body:formdata,
    }).then(
        response => response.text()
    ).catch(
        error => console.error(error)
    )




})
/*
function sendRequest(){
    let message = $('#message')
    $.ajax({
        type: 'POST',
        url: 'http://localhost:8080',
        //url: '../main.go/request',
        data: {
            'message': message
        },
        success: function(response){
            alert(response)
        },
        error: function (response){
            alert(response)
        }
    });
}

 */