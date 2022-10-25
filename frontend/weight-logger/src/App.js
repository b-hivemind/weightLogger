import React, { useState } from 'react'

import { Login } from './components/Login';

import './App.css';
import Dashboard from './components/Dashboard';

function App() {
  const [user, setUser] = useState(null)

  return(
    <div className='App'>
      <div className="title">
            <h1>Archimedes</h1>
            <p><i>The simple weight logger</i></p>
      </div>
      {user == null && <Login stateHandler={setUser}/> }
      {user != null && <Dashboard user={user} setUser={setUser}/>}
      {console.log(user)}
    </div>

  )
}

export default App;
