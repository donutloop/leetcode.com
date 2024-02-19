type ToBeOrNotToBe = {
    toBe: (val: any) => boolean;
    notToBe: (val: any) => boolean;
};

function expect(val: any): ToBeOrNotToBe {
    return new toBeOrNotToBe(val);
};


class toBeOrNotToBe implements ToBeOrNotToBe {
  private val: any;
  
  public constructor(val: any) {
    this.val = val;
  }

  public toBe(val: any) :boolean {
    let ok: boolean = this.val === val
    if (!ok) {
        throw new Error("Not Equal");
    }
    return ok
  }


  public notToBe(val: any) :boolean {
    let ok: boolean = this.val !== val
    if (!ok) {
        throw new Error("Equal");
    }
    return ok
  }

}

/**
 * expect(5).toBe(5); // true
 * expect(5).notToBe(5); // throws "Equal"
 */