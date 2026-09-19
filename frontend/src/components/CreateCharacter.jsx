import React, { useState } from 'react'
import Form from "react-bootstrap/Form"
import Button from 'react-bootstrap/esm/Button'
import loz_api from '../instances';

export default function CreateCharacter() {
    const [gameCount, setGameCount] = useState(1)

    function handleSubmit(e){
        e.preventDefault();

        const form = e.target;
        const formData = new FormData(form);
        const name = formData.get("name")
        const power = Number(formData.get("power"))
        let games = []
        for (let i = 0; i < gameCount; i++){
            games.push(formData.get(`game${i}`))
        }
        const newCharacter = {
            id: 10,
            name: name,
            games: games,
            power: power
        }
        console.log(newCharacter)
        loz_api.post("/characters", newCharacter)
        .then((response) => {
            console.log(response)
        }) .catch((err) => {
            console.log(err)
        }) .finally(() => {
            console.log("create character request complete")
        })

    }

    function GamesInputController(){
        let gameFields = []
        for (let i = 0; i < gameCount; i++){
            gameFields.push(
                <Form.Group key={i}>
                    <Form.Label>{`Character Game ${i + 1}`}</Form.Label>
                    <Form.Control placeholder='Ocarina Of Time' name={`game${i}`}/>
                </Form.Group>
            )
        }
        return (
            <>
                {gameFields}
            </>
        )
    }

  return (
    <Form onSubmit={handleSubmit}>
        <Form.Group >
            <Form.Label>Character Name</Form.Label>
            <Form.Control placeholder='Navi' name='name'/>
        </Form.Group>
        <Form.Group>
            <Form.Label>Character Power Level</Form.Label>
            <Form.Control placeholder='4' name='power'/>
        </Form.Group>
        <GamesInputController/>
        <Button onClick={() => {setGameCount(gameCount - 1)}} disabled={gameCount == 1}>Remove Game</Button>
        <Button onClick={() => {setGameCount(gameCount + 1)}} disabled={gameCount == 5}>Add Game</Button>
        <Button variant="primary" type='submit'>Create Character</Button>
    </Form>
  )
}
