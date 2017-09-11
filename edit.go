package excelize

const (
	//AutoFilter use AutoFilter
	EditActionAutoFilter EditAction = iota + 1
	//DeleteColumns DeleteColumns
	DeleteColumns
	//DeleteRows DeleteRows
	DeleteRows
	//FormatCells FormatCells
	FormatCells
	//FormatColumns FormatColumns
	FormatColumns
	//FormatRows FormatRows
	FormatRows

	//InsertColumns InsertColumns
	InsertColumns
	//InsertHyperlinks InsertHyperlinks
	InsertHyperlinks
	// InsertRows InsertRows
	InsertRows
	//PivotTables PivotTables
	PivotTables
	//Scenarios Scenarios
	Scenarios
	//SelectLockedCells SelectLockedCells
	SelectLockedCells
	//SelectUnlockedCell SelectUnlockedCell
	SelectUnlockedCell
	//Objects edit Objects
	Objects
	//Sort Sort
	Sort
)

//EditAction EditAction
type EditAction int
