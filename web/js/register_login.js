
function Register() {
    const inputName = document.getElementById("name").value;
    const inputEmail = document.getElementById("email").value;
    const inputPassword = document.getElementById("password").value;

    const payload = {
        name: inputName,
        email: inputEmail,
        password: inputPassword
    };
    
    fetch('http://localhost:8080/api/users/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (response.ok) {
            alert('Registration successful!');
            window.location.href = 'login';
        } else {
            alert('Registration Error');
            return response.text().then(error => {
                throw new Error(error);
            });
        }
    })
    .catch(error => {
        alert('Registration failed: ' + error.message);
        throw new Error(error);
    });
}
  
function Login() {
    const inputEmail = document.getElementById("email").value;
    const inputPassword = document.getElementById("password").value;

    const payload = {
        email: inputEmail,
        password: inputPassword
    };
    
    fetch('http://localhost:8080/api/users/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (response.ok) {
            alert('Login successful!');
            window.location.href = 'index';
        } else {
            alert('Login Error');
            return response.text().then(error => {
                throw new Error(error);
            });
        }
    })
    .catch(error => {
        alert('Login failed: ' + error.message);
        throw new Error(error);
    });
}