type Counter = {
    increment: () => number,
    decrement: () => number,
    reset: () => number,
}

function createCounter(init: number): Counter {
    return new counter(init)
};

class counter implements Counter {
  private value: number;
  private init: number;
 
  public constructor(init: number) {
    this.value = init;
    this.init = init;
  }
 
  public increment() :number {
    this.value++
    return this.value
  }

  public decrement() :number {
    this.value--
    return this.value
  }

   public reset() :number {
    this.value = this.init
    return this.value
  }
}

/**
 * const counter = createCounter(5)
 * counter.increment(); // 6
 * counter.reset(); // 5
 * counter.decrement(); // 4
 */