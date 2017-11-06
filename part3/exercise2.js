const fibonacci = () => {
    let sum1 = 0
    let sum2 = 1
    return () => {
        let res = sum1 + sum2
        sum1 = sum2
        sum2 = res
        return res
    }
}
let f = fibonacci()
for (let i = 1; i <= 10; i++) {
    console.log(f())
}

//es6版fibonacci function