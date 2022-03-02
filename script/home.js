
// Alert types
const ALERT_TYPES = {
    "Success"   : "alert alert-success alert-dismissible",
    "Warning"   : "alert alert-warning alert-dismissible",
    "Error"     : "alert alert-danger alert-dismissible",
    "Info"      : "alert alert-info alert-dismissible",
};

// Send alert
function sendAlert(alertType, message) {

    const insertDiv = document.getElementById("insert-alert");

    // Remove all the alerts if there is more than 3 alerts present
    if (insertDiv.childElementCount >= 3) {
        
        const dummyList = []
        const children = insertDiv.childNodes;

        children.forEach(child => dummyList.push(child));
        dummyList.forEach(child => insertDiv.removeChild(child));
    }

    // Create and add the new alert
    const alertDiv = document.createElement("div");
    alertDiv.className = ALERT_TYPES[alertType];
    alertDiv.innerHTML = `<button type="button" class="close" data-dismiss="alert">&times;</button>`;

    const span = document.createElement("span");
    span.innerHTML = `<strong>${alertType}</strong> ${message}`;
    
    alertDiv.appendChild(span);
    insertDiv.appendChild(alertDiv);
}

// Get all users
function getAllUsers() {

    let xmlHttp = new XMLHttpRequest();

    let params = new URLSearchParams()
    params.append("command", "get-all-users")
    
    xmlHttp.open("GET", "/home?" + params.toString(), true);
    xmlHttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    
    xmlHttp.onreadystatechange = () => {
        if (xmlHttp.readyState == 4) {
            console.log(xmlHttp.status)
            console.log(xmlHttp.responseText)
        }
    }

    xmlHttp.send()
}