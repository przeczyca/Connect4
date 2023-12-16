import Cell from './Cell';

interface GameProps {
    gameState: string;
    setGameState: any;
    playerTurn: string;
    setPlayerTurn: any;
    winner: string;
    setWinner: any;
    bot: string | null;
    setBot: any;
    gamePosition: string;
    setGamePosition: any;
}

function GameGrid(props: GameProps) {
    const cells = [];
    for (let i = 0; i < 42; i++) {
        cells.push((
            <Cell
                key={i}
                gameProps={props}
                index={i}
            />
        ))
    }

    return (
        <div className='gameWrapper'>
            {cells}
        </div>
    )
}

export default GameGrid;