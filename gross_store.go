package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int {
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen": 	      12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, existsU := units[unit]
      if !existsU {
	return false
      }
	_, existsB := bill[item]
      if existsB {
	bill[item] += units[unit]
	return true
      }
      bill[item] = units[unit]
      return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	  _, existsB := bill[item]
	  _, existsU := units[unit]

	  if !existsB || !existsU {
		return false
	  }
	  quant := bill[item] - units[unit]
	  if quant == 0 {
	  	delete(bill, item)
		return true
	   } else if quant < 0 {
		return false
	   } else {
		bill[item] = quant
		return true
	   }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quant, ok := bill[item]
        return quant, ok
}
