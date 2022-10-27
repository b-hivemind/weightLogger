import React, { useState } from 'react'
import axios from 'axios'

export const Signup = ({ stateHandler }) => {
    const [method, setMethod] = useState("Login")

    axios.defaults.withCredentials = true
    let token = ""

    const handleSignup = (event) => {
        let uname = document.getElementsByName("uname")[0].value;
        let psw = document.getElementsByName("psw")[0].value;
        if(uname === "" || psw === "") {
            console.log("All details are required")
            return
        }
        axios
            .post('http://10.0.0.184:8081/' + method.toLowerCase(), {
                username: uname,
                password: psw
            })
            .then(function(response) {
                token = response.data["token"]
                let profileName = response.data["profile"]["username"]
                stateHandler({
                    "token": token,
                    "username": profileName
                })
            })
            .catch(function(error) {
                console.log(error)
            })
    }
  
    return (
    <div>
        <div className="loginForm">
            <div className="loginFormContainer">
                <div className="methodChooser">
                    <button onClick={() => setMethod("Login")}>Login</button>
                    <button onClick={() => setMethod("Register")}>Register</button>
                </div>
                <h3 className="title">{method}</h3>
                <label for="uname"><b>Username</b></label>
                <input type="text" placeholder="Enter Username" name="uname" required/>

                <label for="psw"><b>Password</b></label>
                <input type="password" placeholder="Enter Password" name="psw" required/>

                <button id="submitButton" onClick={handleSignup}><b>{method}</b></button>
            </div>
        </div>
    </div>
  )
}
