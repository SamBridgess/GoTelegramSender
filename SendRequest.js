function sendRequest(){
    let message = $('#message')
    $.ajax({
        type: 'POST',
        url: 'main.go',
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