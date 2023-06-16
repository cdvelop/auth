package auth

func (a auth) SearchProvider(name string) provider {
	for _, provider := range a.providers {
		if provider.Name() == name {
			return provider
		}
	}
	return nil
}
