import React from 'react'
import Card from "react-bootstrap/Card"

export default function CharacterCard({characterName, characterGames, characterPower}) {
  return (
    <Card>
        <Card.Header>
            {characterName}
        </Card.Header>
        <Card.Subtitle>
            Power Level: {characterPower}
        </Card.Subtitle>
        <Card.Text>
            <ul>
                {characterGames.map((game) => <li>{game}</li>)}
            </ul>
        </Card.Text>
    </Card>
  )
}
