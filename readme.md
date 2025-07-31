# Execute all tests
go test ./...

# Execute Demo
go run ./cmd

# UML
+---------------------------------------------------------+
|                    Creational Patterns                  |
+---------------------------------------------------------+
| Singleton:                                              |
|   + GetInstance()                                       |
|                                                         |
| Factory:                                                |
|   ShapeFactory → Shape                                  |
|         ↙       ↘                                       |
|     Circle      Square                                  |
|                                                         |
| Abstract Factory:                                       |
|   GUIFactory → Button, Checkbox                         |
|        ↙           ↘                                    |
|   WinFactory   MacFactory                               |
|     ↙   ↘       ↙   ↘                                   |
| WinBtn WinChk MacBtn MacChk                             |
|                                                         |
| Builder:                                                |
|   HouseBuilder → House                                  |
|                                                         |
| Prototype:                                              |
|   Person.Clone() → Person                               |
+---------------------------------------------------------+


+---------------------------------------------------------+
|                   Structural Patterns                   |
+---------------------------------------------------------+
| Adapter:                                                |
|   NewPrinter → Adapter → OldPrinter                     |
|                                                         |
| Decorator:                                              |
|   Notifier → SlackDecorator → EmailNotifier             |
+---------------------------------------------------------+

+---------------------------------------------------------+
|                   Behavioral Patterns                   |
+---------------------------------------------------------+
| Strategy:                                               |
|   Context → Strategy                                    |
|            ↙        ↘                                   |
|         Add        Multiply                             |
|                                                         |
| Observer:                                               |
|   Subject → Observer[]                                  |
|                                                         |
| Command:                                                |
|   Invoker → Command → Receiver (Light)                  |
+---------------------------------------------------------+