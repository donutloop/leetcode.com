class Calculator {
    
    /** 
     * @param {number} value
     */
    constructor(value) {
        this.currentValue = value
        this.error = null
    }
    
    /** 
     * @param {number} value
     * @return {Calculator}
     */
    add(value){
        this.currentValue += value
        return this  
    }
    
    /** 
     * @param {number} value
     * @return {Calculator}
     */
    subtract(value){
        this.currentValue -= value
        return this
    }
    
    /** 
     * @param {number} value
     * @return {Calculator}
     */  
    multiply(value) {
         this.currentValue *= value
         return this
    }
    
    /** 
     * @param {number} value
     * @return {Calculator}
     */
    divide(value) {
        if (value == 0) {
            this.error = "Division by zero is not allowed"
            return this
        }

        this.currentValue /= value
        return this
    }
    
    /** 
     * @param {number} value
     * @return {Calculator}
     */
    power(value) {
        this.currentValue = Math.pow(this.currentValue, value);
        return this
    }
    
    /** 
     * @return {number}
     */
    getResult() {
        if (this.error != null) {
            return this.error
        }    

        return this.currentValue
    }
}