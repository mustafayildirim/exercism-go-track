package gross

var units = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	unitCount, isInUnitList := units[unit]
	if !isInUnitList {
		return false
	}

	itemCount, isInBill := bill[item]
	if isInBill {
		bill[item] = unitCount + itemCount
	} else {
		bill[item] = unitCount
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemCount, ok := bill[item]
	if !ok {
		return false
	}
	unitCount, ok := units[unit]
	if !ok {
		return false
	}
	if itemCount < unitCount {
		return false
	} else if itemCount == unitCount {
		delete(bill, item)
		return true
	}

	bill[item] = itemCount - unitCount

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {

	if val, ok := bill[item]; ok {
		return val, true
	}
	return 0, false
}
