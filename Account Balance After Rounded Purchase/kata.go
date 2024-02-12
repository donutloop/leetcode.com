package kata

func accountBalanceAfterPurchase(purchaseAmount int) int {
    initAccount := 100
    point := purchaseAmount % 10
    if point <= 4 {
      purchaseAmount = purchaseAmount - point
    } else {
      purchaseAmount = purchaseAmount + (10 - point)  
    }
    return initAccount - purchaseAmount
}