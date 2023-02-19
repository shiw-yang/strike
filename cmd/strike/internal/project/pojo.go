package project

// Project is a project template
type Project struct {
	Name string
	Path string
}

// New a project from remote repo.
func (p *Project) New(dir, layout string) error {
	panic("not impl")
}
