package inventory

type RouterGroup struct {
	GoodTypeRouter
	MyApi
	OutInventoryRouter
	CheckInventoryRouter
	InventoryDisplayRouter
	InInventoryRouter
}
