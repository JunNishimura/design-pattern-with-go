package animal

type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) BuildCat() {
	d.builder.setName("cat")
	d.builder.setVoice("meow")
}

func (d *Director) BuildDog() {
	d.builder.setName("dog")
	d.builder.setVoice("wow")
}

type Builder interface {
	setName(string)
	setVoice(string)
}

type Cat struct {
	name  string
	voice string
}

type CatBuilder struct {
	cat *Cat
}

func NewCatBuilder() *CatBuilder {
	return &CatBuilder{
		cat: &Cat{},
	}
}

func (c *CatBuilder) setName(name string) {
	c.cat.name = name
}

func (c *CatBuilder) setVoice(voice string) {
	c.cat.voice = voice
}

func (c *CatBuilder) Get() *Cat {
	return c.cat
}

type Dog struct {
	name  string
	voice string
}

type DogBuilder struct {
	dog *Dog
}

func NewDogBuilder() *DogBuilder {
	return &DogBuilder{
		dog: &Dog{},
	}
}

func (d *DogBuilder) setName(name string) {
	d.dog.name = name
}

func (d *DogBuilder) setVoice(voice string) {
	d.dog.voice = voice
}

func (d *DogBuilder) Get() *Dog {
	return d.dog
}
