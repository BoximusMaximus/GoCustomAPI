import React, { useState } from 'react'
import axios from 'axios'
import Button from "react-bootstrap/Button"

export default function ViewCharacters() {
    const [ characters, setCharacters ] = useState([])

    async function GetCharacters(){
        const response = await axios.get("http://localhost:8080/characters/")
        console.log(response.data)
    }

    function RenderCharacters(){
        return
    }

    return (
        <>
            <Button onClick={GetCharacters}>Refresh Characters</Button>
            <h1>All Characters</h1>

        </>
  )
}
