import React, { useState } from 'react'
import CharacterCard from './cards/CharacterCard'
import axios from 'axios'
import loz_api from "../instances"
import Button from "react-bootstrap/Button"

export default function ViewCharacters() {
    const [ characters, setCharacters ] = useState([])

    async function GetCharacters(){
        const response = await loz_api.get("/characters")
        console.log(response.data)
        setCharacters(response.data)
        RenderCharacters()
    }

    function RenderCharacters(){
        return(
            <>
                {characters.map((character) => <CharacterCard 
                    key={character.id} 
                    characterName={character.name} 
                    characterGames={character.games} 
                    characterPower={character.power}/>)}
            </>
        )
    }

    return (
        <>
            <Button onClick={GetCharacters}>Refresh Characters</Button>
            <h1>All Characters</h1>
            <RenderCharacters/>
        </>
  )
}
