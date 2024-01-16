export default class Helper {
    hasWinner(gameState: string, playerTurn: string, index: number): boolean {
        let numInARow = 1;
        //check rows
        let idxToCheck = index - 1;
        while (idxToCheck > -1 && idxToCheck % 7 >= 0 && idxToCheck % 7 < index % 7) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck -= 1;
            }
            else {
                break;
            }
        }

        idxToCheck = index + 1;
        while (idxToCheck < 42 && idxToCheck % 7 <= 6 && idxToCheck % 7 > index % 7) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck += 1;
            }
            else {
                break;
            }
        }

        if (numInARow >= 4) {
            return true;
        }

        //check columns
        numInARow = 1;
        idxToCheck = index - 7;
        while (idxToCheck >= 0) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck -= 7;
            }
            else {
                break;
            }
        }

        idxToCheck = index + 7;
        while (idxToCheck < 42) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck += 7;
            }
            else {
                break;
            }
        }

        if (numInARow >= 4) {
            return true;
        }

        //check diagonals
        //down and right
        numInARow = 1;
        idxToCheck = index - 8
        while (idxToCheck > -1 && idxToCheck % 7 >= 0 && idxToCheck % 7 === (idxToCheck + 8) % 7 - 1) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck -= 8;
            }
            else {
                break;
            }
        }

        idxToCheck = index + 8;
        while (idxToCheck < 42 && idxToCheck % 7 <= 8 && idxToCheck % 7 === (idxToCheck - 8) % 7 + 1) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck += 8;
            }
            else {
                break;
            }
        }

        if (numInARow >= 4) {
            return true;
        }

        //up and right
        numInARow = 1;
        idxToCheck = index - 6
        while (idxToCheck > -1 && idxToCheck % 7 <= 6 && idxToCheck % 7 === (idxToCheck + 6) % 7 + 1) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck -= 6;
            }
            else {
                break;
            }
        }

        idxToCheck = index + 6;
        while (idxToCheck < 42 && idxToCheck % 7 >= 0 && idxToCheck % 7 === (idxToCheck - 6) % 7 - 1) {
            if (gameState[idxToCheck] === playerTurn) {
                numInARow += 1;
                idxToCheck += 6;
            }
            else {
                break;
            }
        }

        if (numInARow >= 4) {
            return true;
        }

        return false;
    }

    makeMove(gameState: string, setGameState: any, gamePosition: string, setGamePosition: any, playerTurn: string, setPlayerTurn: any, column: number, setWinner: any): void {
        let index = column - 1;
        while (index < 35 && gameState[index + 7] === "0") {
            index += 7;
        }

        if (gameState[index] === "0" && gameState[index + 7] !== "0") {
            let stateArr = gameState.split("");
            stateArr[index] = playerTurn;
            setGameState(stateArr.join(""));

            if (this.hasWinner(gameState, playerTurn, index)) {
                setWinner(playerTurn);
            }
            setPlayerTurn(playerTurn === "1" ? "2" : "1");
            setGamePosition(gamePosition + column)
        }
    }
}