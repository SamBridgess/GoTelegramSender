var dark = false;
function darkLight() {
    if (!dark) {
        //light
        document.getElementById('main').style.background = "#ebebf3";//background
        document.getElementById('inputForm').style.background = "white";//input back
        document.getElementById('inputForm').style.borderColor = "#dce1e6";//input border

        document.getElementById('username').style.background = "white";//back
        document.getElementById('username').style.borderColor = "#d3d9de";//border
        document.getElementById('username').style.color = "black";//text

        document.getElementById('message').style.background = "white";//back
        document.getElementById('message').style.borderColor = "#d3d9de";//border
        document.getElementById('message').style.color = "black";//text

        document.getElementById('theme').value = "dark theme";
    } else {
        //dark
        document.getElementById('main').style.background = "#202226";//background
        document.getElementById('inputForm').style.background = "#2b2e34";//input back
        document.getElementById('inputForm').style.borderColor = "#3e424b";//input border

        document.getElementById('username').style.background = "#2b2e34";//back
        document.getElementById('username').style.borderColor = "#474c56";//border
        document.getElementById('username').style.color = "white";//input name

        document.getElementById('message').style.background = "#2b2e34";//back
        document.getElementById('message').style.borderColor = "#474c56";//border
        document.getElementById('message').style.color = "white";//text

        document.getElementById('theme').value = "light theme";
    }
    dark = !dark;
}
