# Visitor Design Pattern in golang

Wikipedia says :

> In object-oriented programming and software engineering, the visitor design pattern is a way of separating an algorithm from an object structure on which it operates. A practical result of this separation is the ability to add new operations to existing object structures without modifying the structures. It is one way to follow the open/closed principle.

An example being better than 100 words, lets take an example.
So we have our Employee Developer and Directors that is a structure that works well and are critical.

```golang
type Employee interface {
	FullName()
}

type Developer struct {
	FirstName string
	LastName  string
	Income    float32
	Age 	  int
}
func (d  Developer) FullName() {
	fmt.Println("Developer ",d.FirstName," ",d.LastName)
}

type Director struct {
	FirstName string
	LastName  string
	Income    float32
	Age       int
}
func (d  Director) FullName() {
	fmt.Println("Director ",d.FirstName," ",d.LastName)
}
```

An for some raison we want to *extend* that structure. But as these objects are highly critical we don't want to change it (not too much).
So We will use for that the Visitor Pattern, in order to "inject" new operations to existing object structures without modifying the structures.

```golang
type Visitor interface {
	VisitDeveloper(d Developer)
	VisitDirector(d Director)
}
```
 
 Now each new "additional Operation" or Behavior can be coded in a separate structure.

 ```golang
 type CalculIncome struct {
	bonusRate int
}

func (c CalculIncome) VisitDeveloper(d Developer) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}

func (c CalculIncome) VisitDirector(d Director) {
	fmt.Println(d.Income + d.Income*c.bonusRate/100)
}
 ```

 And I still need to inject it on our Employees, for that we need to add one modification to our Employee Interface, adding the Accept Method so Employee becomes :

 ```golang
type Employee interface {
	FullName()
	Accept(Visitor)
}
 ```

 And our objects 

 ```golang
func (d  Developer) Accept(v Visitor) {
	v.VisitDeveloper(d)
}

func (d  Director) Accept(v Visitor) {
	v.VisitDirector(d)
}
 ```

There we go, we have finished to implement our Visitor Pattern, let's test it.

```golang
backend := Developer{"Bob", "Bilbo", 1000, 32}
boss := Director{"Bob", "Baggins", 2000, 40}

backend.FullName()
backend.Accept(CalculIncome{20})

boss.FullName()
boss.Accept(CalculIncome{10})
```

output :
```bash
Developer  Bob   Bilbo
1200
Director  Bob   Baggins
2200
```

In the same wy, if we want to add more behaviors, we only need to implement visitor interface :

```golang
type AddingCaptainAge struct {
	captainAge int
}

func (c AddingCaptainAge) VisitDeveloper(d Developer) {
	fmt.Println(d.Age + c.captainAge)
}

func (c AddingCaptainAge) VisitDirector(d Director) {
	fmt.Println(d.Age + c.captainAge)
}
```

and using it the same way as previsouly :
```golang
backend := Developer{"Bob", "Bilbo", 1000, 32}
boss := Director{"Bob", "Baggins", 2000, 40}

backend.FullName()
backend.Accept(CalculIncome{20})
backend.Accept(AddingCaptainAge{42})

boss.FullName()
boss.Accept(CalculIncome{10})
boss.Accept(AddingCaptainAge{42})
```

output :

```bash
Developer  Bob   Bilbo
1200
74

Director  Bob   Baggins
2200
82
```

Execute the code : https://play.golang.org/p/ddouZJj2z2S
Read on medium : https://medium.com/@felipedutratine/visitor-design-pattern-in-golang-3c142a12945a

