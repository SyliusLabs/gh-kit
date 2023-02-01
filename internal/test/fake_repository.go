package test

type FakeRepository struct {
}

func (f *FakeRepository) Owner() string {
	return "JohnnyWalker"
}

func (f *FakeRepository) Name() string {
	return "BlackLabel"
}

func (f *FakeRepository) Host() string {
	return "github.com"
}
