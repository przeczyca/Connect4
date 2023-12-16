function Winner({ winner }: { winner: string }) {
    let content;
    if (winner !== "0") {
        const player = winner === "1" ? "Red" : "Black";

        content = <h1>
            {player} wins!
        </h1>
    }

    return (
        <div>
            {content}
        </div>
    )
}

export default Winner;