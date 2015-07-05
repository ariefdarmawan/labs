package main

import (
	"fmt"
	"strings"
)

type IPerson interface {
	Say()
}

type Person struct {
	FullName string
}

type Programmer struct {
	Person
}

type Analyst struct {
	Person
}

func (p *Programmer) Say() {
	fmt.Printf("Hi my name is %s and I am a programmer\n", p.FullName)
}

func (p *Analyst) Say() {
	fmt.Printf("Hi my name is %s and I am an analyst\n", p.FullName)
}

type IProject interface {
	Stat()
	AddMember(IPerson)
}

type ProjectBase struct {
	Teams []IPerson
	Name  string
}

func (p *ProjectBase) Stat() {
	fmt.Println("Development Project: ", p.Name)
	for _, t := range p.Teams {
		t.Say()
	}
	fmt.Println("=========================")
	fmt.Println("")
}

func (p *ProjectBase) AddMember(person IPerson) {
	p.Teams = append(p.Teams, person)
}

type ProjectDevelopment struct {
	ProjectBase
}

type ProjectImplementation struct {
	ProjectBase
	Cost int
}

func (p *ProjectImplementation) Stat() {
	fmt.Printf("Implementation Project: %s cost: %d\n", p.Name, p.Cost)
	for _, t := range p.Teams {
		t.Say()
	}
	fmt.Println("=========================")
	fmt.Println("")
}

func NewProject(projectType string, name string, persons ...IPerson) IProject {
	p := func(projectType string, name string) IProject {
		projectType = strings.ToLower(projectType)
		if projectType == "development" {
			p := new(ProjectDevelopment)
			return p
		} else {
			p := new(ProjectImplementation)
			return p
		}
	}(projectType, name)
	for _, v := range persons {
		p.AddMember(v)
	}
	return p
}

func NewProgrammer(fullname string) *Programmer {
	o := new(Programmer)
	o.FullName = fullname
	return o
}

func NewAnalyst(fullname string) *Analyst {
	o := new(Analyst)
	o.FullName = fullname
	return o
}

func main() {
	_ = "breakpoint"
	p := NewProject("development", "Project Dev 1", NewProgrammer("Prog 1"), NewAnalyst("Analyst 1"))
	p.Stat()

	i := NewProject("implementation", "New Implementation", NewProgrammer("Prog 2"), NewAnalyst("Analyst 2"))
	i.(*ProjectImplementation).Cost = 500
	i.Stat()
}
