package YGO

type YgoBackend struct {
}

func (this *YgoBackend) GetUrlPrefix() string {
	return "https://yugipedia.com"
}
