import React from 'react'
import Card from "react-bootstrap/Card"

export default function CharacterCard({characterName, characterPower}) {
  return (
    <Card>
        <Card.Header>
            {characterName}
        </Card.Header>
        <Card.Subtitle>
            Power Level: {characterPower}
        </Card.Subtitle>

        <ul>
            {characterGames.map((game, key) => <li key={`${key}${game}`}>{game}</li>)}
        </ul>

    </Card>
  )
}
