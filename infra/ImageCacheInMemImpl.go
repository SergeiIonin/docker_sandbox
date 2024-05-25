package infra


type ImageCacheInMemImpl struct {
	images map[string][]string
}

func NewImageCacheInMemImpl() *ImageCacheInMemImpl {
	return &ImageCacheInMemImpl{images: make(map[string][]string)}
}

func (impl *ImageCacheInMemImpl) GetAll(id string) ([]string, error) {
	images, ok := impl.images[id]
	if !ok {
		return nil, nil
	}
	return images, nil
}

func (impl *ImageCacheInMemImpl) SaveAll(id string, images []string) (error) {
	impl.images[id] = images
	return nil
}

func (impl *ImageCacheInMemImpl) UpdateAll(id string, images []string) error {
	impl.images[id] = images
	return nil
}

func (impl *ImageCacheInMemImpl) DeleteAll(id string) error {
	delete(impl.images, id)
	return nil
}
