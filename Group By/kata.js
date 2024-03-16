/**
 * @param {Function} fn
 * @return {Object}
 */
Array.prototype.groupBy = function(fn) {
    let groups = {}
    this.forEach((element, index) => {
        const key = fn(element)
        if (groups[key]) {
            groups[key].push(element)    
        } else {
            groups[key] = []
            groups[key].push(element)
        }
    })
    return groups
};

/**
 * [1,2,3].groupBy(String) // {"1":[1],"2":[2],"3":[3]}
 */