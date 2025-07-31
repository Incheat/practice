package abstractfactory

// === 抽象產品介面 ===
type Button interface {
	Paint() string
}

type Checkbox interface {
	Paint() string
}

// === 具體產品 ===
type WinButton struct{}
func (WinButton) Paint() string { return "Render Windows Button" }

type MacButton struct{}
func (MacButton) Paint() string { return "Render Mac Button" }

type WinCheckbox struct{}
func (WinCheckbox) Paint() string { return "Render Windows Checkbox" }

type MacCheckbox struct{}
func (MacCheckbox) Paint() string { return "Render Mac Checkbox" }

// === 抽象工廠介面 ===
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// === 具體工廠 ===
type WinFactory struct{}
func (WinFactory) CreateButton() Button   { return WinButton{} }
func (WinFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

type MacFactory struct{}
func (MacFactory) CreateButton() Button   { return MacButton{} }
func (MacFactory) CreateCheckbox() Checkbox { return MacCheckbox{} }

// === 工廠選擇器 ===
func GetFactory(os string) GUIFactory {
	switch os {
	case "windows":
		return WinFactory{}
	case "mac":
		return MacFactory{}
	default:
		return nil
	}
}
