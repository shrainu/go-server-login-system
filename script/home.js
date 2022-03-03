
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
        
        const dummyList = [];
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

    let params = new URLSearchParams();
    params.append("command", "get-all-users");
    
    xmlHttp.open("GET", "/home?" + params.toString(), true);
    xmlHttp.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4) {
            console.log(`${xmlHttp.status} : \n${xmlHttp.responseText}`);
        }
    };

    xmlHttp.send();
}

// Register user
function registerUser() {

    const username = document.getElementById("username");
    const password = document.getElementById("password");

    let xmlHttp = new XMLHttpRequest();
    
    xmlHttp.open("POST", "/home" + "?command=register", true);
    xmlHttp.setRequestHeader("Content-Type", "application/json");
    
    xmlHttp.onreadystatechange = () => {
        if (xmlHttp.readyState == 4) {
            
            if (xmlHttp.status == 200) {

                const user = JSON.parse(xmlHttp.responseText);
                console.log(`${xmlHttp.status} : \n${xmlHttp.responseText}`);
                sendAlert(
                    "Success", 
                    'Welcome ' +
                    '<span class="font-italic">' + user.username +'</span>'+
                    ', you have successfully registered.'
                );
            }
            else {

                console.log(`${xmlHttp.status} : \n${xmlHttp.responseText}`);
                sendAlert("Error", xmlHttp.responseText);
            }
        }
    };

    const content = JSON.stringify(
        {
            "id": "00000000-0000-0000-0000-000000000000", 
            "username": username.value, 
            "password": password.value
        }
    );

    xmlHttp.send(content);
}

// Login User
function loginUser() {

    const username = document.getElementById("username");
    const password = document.getElementById("password");

    let xmlHttp = new XMLHttpRequest();

    xmlHttp.open("POST", "/home" + "?command=login", true);
    xmlHttp.setRequestHeader("Content-Type", "application/json");

    xmlHttp.onreadystatechange = () => {
        if (xmlHttp.readyState == 4) {

            if (xmlHttp.status == 200) {

                const user = JSON.parse(xmlHttp.responseText);
                console.log(`${xmlHttp.status} : \n${xmlHttp.responseText}`);
                sendAlert(
                    "Success", 
                    'Welcome back ' +
                    '<span class="font-italic">' + user.username +'</span>'
                );
            }
            else {

                console.log(`${xmlHttp.status} : \n${xmlHttp.responseText}`);
                sendAlert("Error", xmlHttp.responseText);
            }
        }
    };

    const content = JSON.stringify(
        {
            "id": "00000000-0000-0000-0000-000000000000", 
            "username": username.value, 
            "password": password.value
        }
    );

    xmlHttp.send(content);
}