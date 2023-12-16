export default class Bot {
    makeMove(gameState: string): number {
        let availableColumns: number[] = [];
        for (let i = 0; i < 7; i++) {
            if (gameState[i] === "0") {
                availableColumns.push(i);
            }
        }

        return availableColumns[Math.floor(Math.random() * (availableColumns.length))];
    }
}