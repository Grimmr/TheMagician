package YGO

type YgoBackend struct {
}

func (this YgoBackend) GetUrlPrefix() string {
	return "https://yugipedia.com"
}

func (this YgoBackend) GetSetListUrl() string {
	return "https://yugipedia.com/wiki/Order_of_Set_Release"
}
