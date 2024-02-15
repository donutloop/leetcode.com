type Fn = (n: number, i: number) => any

function filter(arr: number[], fn: Fn): number[] {
    let i = -1;
    return arr.filter(number => {
        i++
        return fn(number, i)
    })
};