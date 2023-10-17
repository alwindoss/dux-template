import { useState } from 'react'
import './App.css'
import HomePage from './pages/HomePage'

function App() {

  const [userLoggedIn, setUserLoggedIn] = useState(true);
  function toggle() {
    setUserLoggedIn(!userLoggedIn)
  }

  return (
    <>
      <input type='checkbox' onChange={toggle} />
      <HomePage loggedIn={userLoggedIn} />
    </>
  )
}



export default App
