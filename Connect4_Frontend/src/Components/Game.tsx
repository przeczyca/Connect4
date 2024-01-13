import GameGrid from './GameGrid';
import Helper from './Helpers';
import Winner from './Winner';

import { useEffect, useState } from 'react';

function Game() {
    const [gameState, setGameState] = useState("000000000000000000000000000000000000000000");
    const [gamePosition, setGamePosition] = useState("");
    const [playerTurn, setPlayerTurn] = useState("1"); //1 is red, 2 is black
    const [winner, setWinner] = useState("0");
    const [bot, setBot] = useState("0");

    const helper = new Helper

    function restartGame() {
        setGameState("000000000000000000000000000000000000000000");
        setGamePosition("");
        setPlayerTurn("1");
        setWinner("0");
        setBot("0");
    }

    function useBot() {
        if (bot === "0") {
            setBot("2");
        }
        else {
            setBot("0");
        }
    }

    useEffect(() => {
        if (playerTurn === bot) {
            const requestOptions = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ position: gamePosition })
            };

            const url = 'http://localhost:8080'

            fetch(url + '/getOneBestMove', requestOptions)
                .then(response => response.json())
                .then(data => helper.makeMove(gameState, setGameState, gamePosition, setGamePosition, playerTurn, setPlayerTurn, data.column, setWinner));
        }
    }, [playerTurn]);

    return (
        <div className='everything'>
            <h1>Connect 4</h1>
            <GameGrid
                gameState={gameState}
                setGameState={setGameState}
                playerTurn={playerTurn}
                setPlayerTurn={setPlayerTurn}
                winner={winner}
                setWinner={setWinner}
                bot={bot}
                setBot={setBot}
                gamePosition={gamePosition}
                setGamePosition={setGamePosition}
            />
            {bot === playerTurn && <h1>Bot Thinking</h1>}
            <Winner winner={winner} />
            <button onClick={restartGame}>restart</button>
            {gameState === "000000000000000000000000000000000000000000" && <button onClick={useBot}>{bot === "0" ? "set bot" : "bot is black"}</button>}
        </div>
    )
}

export default Game;