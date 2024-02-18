function map(arr: number[], fn: (n: number, i: number) => number): number[] {
    if (arr.length == 0) {
        return arr
    }

    let returnedArray: number[] = new Array(arr.length-1);
    for(let i: number = 0; i < arr.length; i++) {
        returnedArray[i] = fn(arr[i], i)
    }
    return returnedArray
};