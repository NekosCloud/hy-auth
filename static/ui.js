window.onload = function() {
    document.getElementById("copy").innerHTML = new Date().getFullYear();
    document.getElementById("lds").classList.remove("lds-visible");
    document.getElementById("header-btn").onclick = function(e) {
        document.getElementById("header-menu").classList.toggle("active");
        document.getElementById("nav-btn").classList.toggle("active");
    }
}