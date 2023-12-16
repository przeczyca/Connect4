import Helper from './Helpers';

interface GameProps {
    gameState: string;
    setGameState: any;
    playerTurn: string;
    setPlayerTurn: any;
    winner: string;
    setWinner: any;
    gamePosition: string;
    setGamePosition: any;
    bot: string | null;
}

function Cell({ gameProps, index }: { gameProps: GameProps, index: number }) {

    const setColor = {
        backgroundColor: "white"
    };

    if (gameProps.gameState[index] === '1') {
        setColor.backgroundColor = "red";
    }
    else if (gameProps.gameState[index] === '2') {
        setColor.backgroundColor = "black";
    }

    const helper = new Helper();

    function makeMove() {
        if (gameProps.gameState[index] === "0" && gameProps.gameState[index + 7] !== "0" && gameProps.winner === "0" && gameProps.playerTurn !== gameProps.bot) {
            let stateArr = gameProps.gameState.split("");
            stateArr[index] = gameProps.playerTurn;
            gameProps.setGameState(stateArr.join(""));

            if (helper.hasWinner(gameProps.gameState, gameProps.playerTurn, index)) {
                gameProps.setWinner(gameProps.playerTurn);
            }
            gameProps.setPlayerTurn(gameProps.playerTurn === "1" ? "2" : "1");
            gameProps.setGamePosition(gameProps.gamePosition.concat(String(index % 7 + 1)))
        }
    }

    return (
        <div className='gameCell'>
            <div
                className='gameCellCircle'
                style={setColor}
                onClick={makeMove}>
                {index}
            </div>
        </div>
    )
}

export default Cell;