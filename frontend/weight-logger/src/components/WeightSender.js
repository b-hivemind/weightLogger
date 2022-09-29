import React from 'react'
import axios from 'axios'

export const WeightSender = ({ stateHandler }) => {

  const createEntry = (event) => {
    let weight = document.getElementById("weightEntryBox").value;
    if(isNaN(weight) || isNaN(parseFloat(weight))) {
        alert("Invalid weight: " + weight)
        document.getElementById("weightEntryBox").value = null;
    }
    weight = parseFloat(weight)
    axios
        .post('http://10.0.0.184:8081/entries/new', {
            weight: weight,
            force: false
        })
        .then(function(response) {
            stateHandler()
            document.getElementById("weightEntryBox").value = null;
        })
        .catch(function(error) {
            if (error.response.status === 300) {
                document.getElementById('send').style = "display: none"
                document.getElementById('force').style = "display: inline-block"
            }
        })
}
const forceCreateEntry = (event) => {
    let weight = document.getElementById("weightEntryBox").value;
    if(isNaN(weight) || isNaN(parseFloat(weight))) {
        alert("Invalid weight: " + weight)
        document.getElementById("weightEntryBox").value = null;
    }
    weight = parseFloat(weight)
    axios
        .post('http://10.0.0.184:8081/entries/new', {
            weight: weight,
            force: true
        })
        .then(function(response) {

            stateHandler()
            document.getElementById("weightEntryBox").value = null;
            document.getElementById('send').style = "display: inline-block"
            document.getElementById('force').style = "display: none"
        })
        .catch(function(error) {
            console.log(error)
            document.getElementById("weightEntryBox").value = null;
        })
}

  return (
    <div id="sender">
        <input id="weightEntryBox" type="text"/>&nbsp;
        <button id="send" onClick={createEntry}>🚀</button>
        <button id="force" onClick={forceCreateEntry}>⚠️</button>
    </div>
  )
}
