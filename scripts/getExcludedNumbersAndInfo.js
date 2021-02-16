const apostasRealizadas = db.results.count()

const bets = db.results.aggregate([
    {
        "$group": {
            "_id": { numbers: "$numbers" },
            uniqueIds: { $addToSet: "$_id" },
            count: { $sum: 1 }
        }
    },
    { "$match": { count: { $gt: 1 } } }
]).count();

print('------------apostas------------')
print(`número de sorteios repetidos: ${bets}`)
print(`número de sorteios realizados: ${apostasRealizadas}`)

const ultimasApostasConsideradas = 7
const results = db.results.aggregate([
    { $match: { code: { $gt: apostasRealizadas - ultimasApostasConsideradas } } },
    { $unwind: "$numbers" },
    {
        $group: {
            "_id": { ball: "$numbers" },
            count: { $sum: 1 }
        }
    },
    { $match: { count: { $gte: 2 } } }
]).sort({ count: -1 })

const excludeds = []
while (results.hasNext()) {
    excludeds.push(results.next()._id.ball)
}

const lastWinner = db.results.findOne({code: apostasRealizadas})
for (const ball of lastWinner.numbers){
    excludeds.push(ball)
}

print('------------exclusão------------')
print(`${excludeds.length} números para excluir: ${excludeds.sort((a, b) => a - b)}`)

const nonRepeatNumbers = 60 - excludeds.length

// 720 = 6!
const defaultChance = (1 / ((60 * 59 * 58 * 57 * 56 * 55) / 720)) * 100

const withCriteriaChance = (1 / ((nonRepeatNumbers * (nonRepeatNumbers - 1) * (nonRepeatNumbers - 2)
    * (nonRepeatNumbers - 3) * (nonRepeatNumbers - 4) * (nonRepeatNumbers - 5)) / 720)) * 100

print('------------probabilidade------------')
print(`chance padrão = ${defaultChance}%`)
print(`chance com critério = ${withCriteriaChance}%`)
print(`chance aumentada em = ${withCriteriaChance / defaultChance}%`)

print('------------gastos------------')
const APOSTAS_POR_SEMANA = 3
const SEMANAS = 4
const priceMonth = 31.5 * APOSTAS_POR_SEMANA * SEMANAS
print(`${priceMonth}R$ por mês`)
print(`${priceMonth * 12}R$ por ano`)