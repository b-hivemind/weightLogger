import React from 'react'
import axios from 'axios'

export const Login = () => {
    axios.defaults.withCredentials = true
    let token = ""

    const handleLogin = (event) => {
        let uname = document.getElementsByName("uname")[0].value;
        let psw = document.getElementsByName("psw")[0].value;
        if(uname === "" || psw === "") {
            console.log("All details are required")
            return
        }
        axios
            .post('http://10.0.0.184:8081/login', {
                username: uname,
                password: psw
            })
            .then(function(response) {
                token = response.data["token"]
                axios.get("http://10.0.0.184:8081/entries/2", {
                    headers: {
                        bearer: token,
                    }, 
                    })
                    .then(function(response) {
                        console.log(response.data)
                    })
                    .catch(function(error) {
                        console.log(error)
                    })
            })
            .catch(function(error) {
                console.log(error)
            })
    }
  
    return (
    <div>
        <label for="uname"><b>Username</b></label>
        <input type="text" placeholder="Enter Username" name="uname" required/>

        <label for="psw"><b>Password</b></label>
        <input type="password" placeholder="Enter Password" name="psw" required/>

        <button onClick={handleLogin}>Login</button>
    </div>
  )
}