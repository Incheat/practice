package main

import (
	"fmt"

	// Creational
	"example.com/go-design-pattern/creational/singleton"
	"example.com/go-design-pattern/creational/factory"
	"go-design-patterns/creational/abstractfactory"
	"example.com/go-design-pattern/creational/builder"
	"example.com/go-design-pattern/creational/prototype"

	// Structural
	"example.com/go-design-pattern/structural/adapter"
	"example.com/go-design-pattern/structural/decorator"

	// Behavioral
	"example.com/go-design-pattern/behavioral/strategy"
	"example.com/go-design-pattern/behavioral/observer"
	"example.com/go-design-pattern/behavioral/command"
)

func main() {
	// === Creational ===
	fmt.Println("=== Creational Patterns ===")

	// Singleton
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	fmt.Println("Singleton same instance:", s1 == s2)

	// Factory
	shape := factory.NewShape("circle")
	fmt.Println("Factory Shape:", shape.Draw())

	// Abstract Factory
	winFactory := abstractfactory.GetFactory("windows")
	winBtn := winFactory.CreateButton()
	winChk := winFactory.CreateCheckbox()
	fmt.Println("Abstract Factory Windows:", winBtn.Paint(), "|", winChk.Paint())

	macFactory := abstractfactory.GetFactory("mac")
	macBtn := macFactory.CreateButton()
	macChk := macFactory.CreateCheckbox()
	fmt.Println("Abstract Factory Mac:", macBtn.Paint(), "|", macChk.Paint())

	// Builder
	b := &builder.HouseBuilder{}
	house := b.BuildWalls("Brick").BuildRoof("Tile").GetResult()
	fmt.Println("Builder House:", house.Walls, house.Roof)

	// Prototype
	p1 := &prototype.Person{Name: "John"}
	p2 := p1.Clone().(*prototype.Person)
	fmt.Println("Prototype Clone:", p1.Name, "==", p2.Name, "Different instance:", p1 != p2)

	// === Structural ===
	fmt.Println("\n=== Structural Patterns ===")

	// Adapter
	oldPrinter := adapter.LegacyPrinter{}
	newPrinter := adapter.Adapter{OldPrinter: oldPrinter}
	newPrinter.PrintNew("Hello Adapter")

	// Decorator
	email := decorator.EmailNotifier{}
	slack := decorator.SlackDecorator{Notifier: email}
	slack.Send("Hello Decorator")

	// === Behavioral ===
	fmt.Println("\n=== Behavioral Patterns ===")

	// Strategy
	ctx := strategy.Context{Strategy: strategy.Add{}}
	fmt.Println("Strategy Add:", ctx.ExecuteStrategy(3, 5))
	ctx = strategy.Context{Strategy: strategy.Multiply{}}
	fmt.Println("Strategy Multiply:", ctx.ExecuteStrategy(3, 5))

	// Observer
	subject := &observer.Subject{}
	obs := &observer.ConcreteObserver{}
	subject.Add(obs)
	subject.Notify("Hello Observer")
	fmt.Println("Observer received:", obs.LastMsg)

	// Command
	light := &command.Light{}
	onCmd := command.LightOnCommand{Light: light}
	offCmd := command.LightOffCommand{Light: light}
	onCmd.Execute()
	fmt.Println("Command Light On:", light.OnState)
	offCmd.Execute()
	fmt.Println("Command Light Off:", light.OnState)
}
