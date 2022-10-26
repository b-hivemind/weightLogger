import React, { useState, useEffect } from 'react'

import { Login } from './components/Login';

import './App.css';
import Dashboard from './components/Dashboard';

const getUserFromLocalStorage = () => {
  if(window.localStorage.getItem("token") == null) {
    return null
  }
  return (
    {
      token: window.localStorage.getItem("token"), 
      username: window.localStorage.getItem("username")
    }
  )
}

function App() {
  const [user, setUser] = useState(null)

  useEffect(() => {
    setUser(getUserFromLocalStorage())
  }, [])

  const saveUserCreds = (user) => {
    setUser(user)
    window.localStorage.setItem("token", user.token)
    window.localStorage.setItem("username", user.username)
  }

  return(
    <div className='App'>
      <div className="title">
            <h1>Archimedes</h1>
            <p><i>The simple weight logger</i></p>
            <hr/>
      </div>
      {user == null && <Login stateHandler={saveUserCreds}/> }
      {user != null && <Dashboard user={user} setUser={setUser}/>}
    </div>

  )
}

export default App;
