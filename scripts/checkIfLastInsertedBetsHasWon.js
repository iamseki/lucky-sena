const nextGameCode = db.results.count() + 1

const nextBets = db.bets.find({ code: nextGameCode })

const nextGameBets = []
while (nextBets.hasNext()) {
    nextGameBets.push(nextBets.next().numbers)
}

print(`checking if ${nextGameBets.length} bets of game [${nextGameCode}] is already won`)

const alreadyWon = db.results.find({ numbers: { $in: nextGameBets } })

let isWon = false

while (alreadyWon.hasNext()) {
    isWon = true
    const wonGame = alreadyWon.next()
    print(`repeated bet in \ngame code = ${wonGame.code}\nbet = ${wonGame.numbers}\n-----`)
}

if(!isWon) print(`no one of the bets generated has won, go ahead !`)
for (const bet of nextGameBets) {
    print(`game ${nextGameCode} generated - ${bet}`)
}