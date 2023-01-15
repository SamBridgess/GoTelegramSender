const url = "request"

var inputForm = document.getElementById("inputForm")

inputForm.addEventListener("submit", (e)=>{
    e.preventDefault()

    const formdata = new FormData(inputForm)
    fetch(url,{

        method:"POST",
        body:formdata,
    }).catch(
        error => console.error(error)
    )




})
