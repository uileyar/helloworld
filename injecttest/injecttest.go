// injecttest.go
package main

import (
	"fmt"
	//	"github.com/facebookgo/ensure"
	"github.com/facebookgo/inject"
	"github.com/golang/glog"
	"net/http"
	"os"
)

// Our Awesome Application renders a message using two APIs in our fake
// world.
type HomePlanetRenderApp struct {
	// The tags below indicate to the inject library that these fields are
	// eligible for injection. They do not specify any options, and will
	// result in a singleton instance created for each of the APIs.

	NameAPI   *NameAPI   `inject:""`
	PlanetAPI *PlanetAPI `inject:""`
}

func (a *HomePlanetRenderApp) Render(id uint64) string {
	return fmt.Sprintf(
		"%s is from the planet %s.",
		a.NameAPI.Name(id),
		a.PlanetAPI.Planet(id),
	)
}

// Our fake Name API.
type NameAPI struct {
	// Here and below in PlanetAPI we add the tag to an interface value.
	// This value cannot automatically be created (by definition) and
	// hence must be explicitly provided to the graph.

	HTTPTransport http.RoundTripper `inject:""`
}

func (n *NameAPI) Name(id uint64) string {
	// in the real world we would use f.HTTPTransport and fetch the name
	return "Spock"
}

// Our fake Planet API.
type PlanetAPI struct {
	HTTPTransport http.RoundTripper `inject:""`
}

func (p *PlanetAPI) Planet(id uint64) string {
	// in the real world we would use f.HTTPTransport and fetch the planet
	return "Vulcan"
}

type Answerable interface {
	Answer() int
}

type TypeAnswerStruct struct {
	answer  int
	private int
}

func (t *TypeAnswerStruct) Answer() int {
	return t.answer
}

type TypeNestedStruct struct {
	A *TypeAnswerStruct `inject:""`
}

func (t *TypeNestedStruct) Answer() int {
	return t.A.Answer()
}

func RequireTag() {
	var v struct {
		A *TypeAnswerStruct
		B *TypeNestedStruct `inject:""`
	}

	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	if v.A != nil {
		glog.Fatal("v.A is not nil")
	}
	if v.B == nil {
		glog.Fatal("v.B is nil")
	}
}

type TypeWithNonPointerInject struct {
	A *int `inject:""`
}

func ErrorOnNonPointerInject() {
	var a TypeWithNonPointerInject
	err := inject.Populate(&a)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err.Error())
		const msg = "found inject tag on unsupported field A in type *main.TypeWithNonPointerInject"
		if err.Error() != msg {
			fmt.Errorf("expected:\n%s\nactual:\n%s", msg, err.Error())
		}
	}
}

func InjectSimple() {
	var v struct {
		A *TypeAnswerStruct `inject:""`
		B *TypeNestedStruct `inject:""`
	}

	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	fmt.Println(v.A)
	fmt.Println(v.B)
	fmt.Println(v.B.A)
	if v.A == v.B.A {
		fmt.Println("wwwwwwww")
	}
}

func DoesNotOverwrite() {
	a := &TypeAnswerStruct{}
	var v struct {
		A *TypeAnswerStruct `inject:""`
		B *TypeNestedStruct `inject:""`
	}
	fmt.Println(v)
	v.A = a
	a.answer = 1111
	a.private = 22222
	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println(v)
	fmt.Println(v.A)

}

func Private() {
	var v struct {
		A *TypeAnswerStruct `inject:""`
		B *TypeNestedStruct `inject:"private"`
	}

	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", v)
	fmt.Println(v)
	fmt.Println(v.A)
	fmt.Println(v.B)
	fmt.Println(v.B.A == v.A)

}

type TypeWithJustColon struct {
	A *TypeAnswerStruct `inject:"`
}

func TagWithJustColon() {
	var a TypeWithJustColon
	err := inject.Populate(&a)
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err)
	}
}

func ProvideWithFields() {
	var g inject.Graph
	a := &TypeAnswerStruct{}
	err := g.Provide(&inject.Object{Value: a, Fields: map[string]*inject.Object{}})

	fmt.Println(err)
	//ensure.NotNil(glog.Fatal, err)
	//ensure.DeepEqual(fmt, err.Error(), "fields were specified on object *inject_test.TypeAnswerStruct when it was provided")
}

func TestProvideNonPointer() {
	var g inject.Graph
	var i *int
	err := g.Provide(&inject.Object{Value: i})

	fmt.Println(err)

}

func TestProvideTwoOfTheSame() {
	var g inject.Graph
	a := TypeAnswerStruct{}
	err := g.Provide(&inject.Object{Value: &a})
	fmt.Println(err)

	err = g.Provide(&inject.Object{Value: &a})
	fmt.Println(err)

}

func TestProvideTwoOfTheSameWithPopulate() {
	a := TypeAnswerStruct{}
	err := inject.Populate(&a, &a)
	fmt.Println(err)

}

func TestProvideTwoWithTheSameName() {
	var g inject.Graph
	const name = "foo"
	a := TypeAnswerStruct{}
	err := g.Provide(&inject.Object{Value: &a, Name: name})
	fmt.Println(err)

	err = g.Provide(&inject.Object{Value: &a, Name: name})
	fmt.Println(err)

}

func TestNamedInstanceWithDependencies() {
	var g inject.Graph
	a := &TypeNestedStruct{}
	if err := g.Provide(&inject.Object{Value: a, Name: "foo"}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	var c struct {
		A *TypeNestedStruct `inject:"foo"`
	}
	if err := g.Provide(&inject.Object{Value: &c}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.A.A)
	if c.A.A == nil {
		fmt.Println("c.A.A was not injected")
	}
}

func TestTwoNamedInstances() {
	var g inject.Graph
	a := &TypeAnswerStruct{}
	b := &TypeAnswerStruct{}
	if err := g.Provide(&inject.Object{Value: a, Name: "foo"}); err != nil {
		fmt.Println(err)
	}
	a.answer = 1
	a.private = 1
	b.answer = 2
	b.private = 2

	if err := g.Provide(&inject.Object{Value: b, Name: "bar"}); err != nil {
		fmt.Println(err)
	}

	var c struct {
		A *TypeAnswerStruct `inject:"foo"`
		B *TypeAnswerStruct `inject:"bar"`
	}
	if err := g.Provide(&inject.Object{Value: &c}); err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(c.A.answer)
	fmt.Println(c.B)
	/*
		if c.A != a {
			fmt.Println("did not find expected c.A")
		}
		if c.B != b {
			fmt.Println("did not find expected c.B")
		}
	*/
}

func TestCompleteProvides() {
	var g inject.Graph
	var v struct {
		A *TypeAnswerStruct `inject:""`
		i int
	}

	if err := g.Provide(&inject.Object{Value: &v, Complete: true, Name: "foo"}); err != nil {
		fmt.Println(err)
	}

	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(v)

}

type TypeInjectInterfaceMissing struct {
	Answerable Answerable `inject:""`
}

func TestInjectInterfaceMissing() {
	var v TypeInjectInterfaceMissing
	err := inject.Populate(&v)
	fmt.Println(err)

}

type TypeInjectInterface struct {
	Answerable Answerable        `inject:""`
	A          *TypeNestedStruct `inject:""`
}

func TestInjectInterface() {
	var v TypeInjectInterface
	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Answerable)
	fmt.Println(v.A)
	if v.Answerable == nil || v.Answerable != v.A {
		fmt.Println(
			v.Answerable,
			v.A,
		)
	}
}

type TypeWithInvalidNamedType struct {
	A *TypeNestedStruct `inject:"foo"`
}

func TestInvalidNamedInstanceType() {
	var g inject.Graph
	a := &TypeAnswerStruct{}
	if err := g.Provide(&inject.Object{Value: a, Name: "foo"}); err != nil {
		fmt.Println(err)
	}

	var c TypeWithInvalidNamedType
	if err := g.Provide(&inject.Object{Value: &c}); err != nil {
		fmt.Println(err)
	}

	err := g.Populate()

	fmt.Println(err)

}

type TypeWithInjectOnPrivateInterfaceField struct {
	a Answerable `inject:""`
}

func TestInjectOnPrivateInterfaceField() {
	var a TypeWithInjectOnPrivateInterfaceField
	err := inject.Populate(&a)
	fmt.Println(err)
}

type TypeWithInjectOnPrivateField struct {
	a *TypeAnswerStruct `inject:""` //注意大小写！！！
}

func TestInjectOnPrivateField() {
	var a TypeWithInjectOnPrivateField
	err := inject.Populate(&a)
	fmt.Println(err)
}

type TypeInjectPrivateInterface struct {
	Answerable Answerable        `inject:""`
	B          *TypeAnswerStruct `inject:"private"`
}

func TestInjectPrivateInterface() {
	var v TypeInjectPrivateInterface
	err := inject.Populate(&v)
	fmt.Println(err)
}

type TypeInjectTwoSatisfyInterface struct {
	Answerable Answerable        `inject:""`
	A          *TypeAnswerStruct `inject:""`
	B          *TypeNestedStruct `inject:""`
}

func TestInjectTwoSatisfyInterface() {
	var v TypeInjectTwoSatisfyInterface
	err := inject.Populate(&v)
	fmt.Println(err)
}

type TypeInjectNamedTwoSatisfyInterface struct {
	Answerable Answerable        `inject:""`
	A          *TypeAnswerStruct `inject:""`
	B          *TypeNestedStruct `inject:""`
}

func TestInjectNamedTwoSatisfyInterface() {
	var g inject.Graph
	var v TypeInjectNamedTwoSatisfyInterface
	if err := g.Provide(&inject.Object{Name: "foo", Value: &v}); err != nil {
		fmt.Println(err)
	}

	err := g.Populate()
	fmt.Println(err)
}

type TypeWithNonPointerNamedInject struct {
	A int `inject:"foo"`
	B int `inject:"foo1"`
}

func TestErrorOnNonPointerNamedInject() {
	var g inject.Graph
	if err := g.Provide(&inject.Object{Name: "foo", Value: 10}); err != nil {
		fmt.Println(err)
	}

	if err := g.Provide(&inject.Object{Name: "foo1", Value: 20}); err != nil {
		fmt.Println(err)
	}

	var v TypeWithNonPointerNamedInject
	if err := g.Provide(&inject.Object{Value: &v}); err != nil {
		fmt.Println(err)
	}

	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

}

func TestInjectInline() {
	var v struct {
		Inline struct {
			A *TypeAnswerStruct `inject:""`
			B *TypeNestedStruct `inject:""`
		} `inject:"inline"`
	}

	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}

	v.Inline.A.answer = 1

	fmt.Println(v)
	fmt.Println(v.Inline)
	fmt.Println(v.Inline.A)
	fmt.Println(v.Inline.B)
	fmt.Println(v.Inline.B.A)
	if v.Inline.A != v.Inline.B.A {
		fmt.Println("got different instances of A")
	}
}

func TestInjectInlineOnPointer() {
	var v struct {
		Inline *struct {
			A *TypeAnswerStruct `inject:""`
			B *TypeNestedStruct `inject:""`
		} `inject:""`
	}

	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	if v.Inline.A == nil {
		fmt.Println("v.Inline.A is nil")
	}
	if v.Inline.B == nil {
		fmt.Println("v.Inline.B is nil")
	}
	if v.Inline.B.A == nil {
		fmt.Println("v.Inline.B.A is nil")
	}
	if v.Inline.A != v.Inline.B.A {
		fmt.Println("got different instances of A")
	}
}

type TypeWithInlineStructWithPrivate struct {
	Inline struct {
		A *TypeAnswerStruct `inject:"private"`
		B *TypeNestedStruct `inject:"private"`
	} `inject:"inline"`
}

func TestInjectInlinePrivate() {
	var v TypeWithInlineStructWithPrivate
	err := inject.Populate(&v)
	fmt.Println(err)
}

type TypeWithStructValue struct {
	Inline TypeNestedStruct `inject:"inline"`
}

func TestInjectWithStructValue() {
	var v TypeWithStructValue
	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}

type TypeWithNonpointerStructValue struct {
	Inline TypeNestedStruct `inject:"inline"`
}

func TestInjectWithNonpointerStructValue() {
	var v TypeWithNonpointerStructValue
	var g inject.Graph
	if err := g.Provide(&inject.Object{Value: &v}); err != nil {
		fmt.Println(err)
	}
	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}
	if v.Inline.A == nil {
		fmt.Println("v.Inline.A is nil")
	}
	//n := len(g.Objects())

	tek := g.Objects()
	for n, value := range tek {

		fmt.Printf("n = %d, value = %#v\n", n, value)
	}

}

func TestPrivateIsFollowed() {
	var i int
	var v struct {
		//A *TypeNestedStruct `inject:"private"`
		B *int `inject:"foo"`
		//c int `inject:"private"`
	}
	i = 100
	//v.B = &i
	var g inject.Graph
	if err := g.Provide(&inject.Object{Name: "foo", Value: &i}); err != nil {
		fmt.Println(err)
	}
	/*
		if err := g.Provide(&inject.Object{Value: &v}); err != nil {
			fmt.Println(err)
		}
	*/
	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
}

func TestDoesNotOverwriteInterface() {
	i := 100
	a := &TypeAnswerStruct{}
	var v struct {
		K *int              `inject:""`
		A Answerable        `inject:""`
		B *TypeNestedStruct `inject:""`
	}
	fmt.Println(v)
	v.A = a
	v.K = &i
	/*
		var g inject.Graph
		if err := g.Provide(&inject.Object{Name: "inttest", Value: &i}); err != nil {
			fmt.Println(err)
		}
	*/
	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	a.answer = 1
	i = 99
	fmt.Println(v.A)
	fmt.Println(v.B)
	fmt.Println(*v.K)
	if v.A != a {
		fmt.Println("original A was lost")
	}
	if v.B == nil {
		fmt.Println("v.B is nil")
	}
}

func TestInterfaceIncludingPrivate() {
	var v struct {
		A Answerable        `inject:""`
		B *TypeNestedStruct `inject:"private"`
		C *TypeAnswerStruct `inject:""`
	}
	if err := inject.Populate(&v); err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	if v.A == nil {
		fmt.Println("v.A is nil")
	}
	if v.B == nil {
		fmt.Println("v.B is nil")
	}
	if v.C == nil {
		fmt.Println("v.C is nil")
	}
	if v.A != v.C {
		fmt.Println("v.A != v.C")
	}
	if v.A == v.B {
		fmt.Println("v.A == v.B")
	}
}

type TypeInjectWithMapWithoutPrivate struct {
	A map[string]int `inject:"private"`
}

func TestInjectMapWithoutPrivate() {
	var v TypeInjectWithMapWithoutPrivate
	err := inject.Populate(&v)
	fmt.Println(err)
}

type TypeForObjectString struct {
	A *TypeNestedStruct `inject:"foo"`
	B *TypeNestedStruct `inject:""`
}

func TestObjectString() {
	var g inject.Graph

	a := &TypeNestedStruct{}
	if err := g.Provide(&inject.Object{Value: a, Name: "foo"}); err != nil {
		fmt.Println(err)
	}

	var c TypeForObjectString
	if err := g.Provide(&inject.Object{Value: &c}); err != nil {
		fmt.Println(err)
	}

	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}

	a.A.answer = 100
	fmt.Printf("c.A.A = %#v, c.B.A = %#v\n", c.A.A, c.B.A)

	var actual []string
	for _, o := range g.Objects() {

		fmt.Printf("%#v\n", o)
		actual = append(actual, fmt.Sprint(o))
	}

	//fmt.Printf("%#v \n", actual)

}

type TypeForGraphObjects struct {
	//TypeNestedStruct `inject:"inline"`
	//A                *TypeNestedStruct `inject:"foo"`
	E struct {
		B *TypeNestedStruct `inject:""`
	} `inject:"inline"`
}

func TestGraphObjects() {
	var g inject.Graph
	err := g.Provide(
		//&inject.Object{Value: &TypeNestedStruct{}, Name: "foo"},
		&inject.Object{Value: &TypeForGraphObjects{}},
	)
	fmt.Println(err)

	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}

	var actual []string
	for _, o := range g.Objects() {
		actual = append(actual, fmt.Sprint(o))
	}
	fmt.Printf("%#v\n", actual)
}

type logger struct {
	Expected []string
	next     int
}

func (l *logger) Debugf(f string, v ...interface{}) {
	actual := fmt.Sprintf(f, v...)

	fmt.Println(actual)

	if l.next == len(l.Expected) {
		fmt.Printf(`unexpected log "%s"`, actual)
	}
	expected := l.Expected[l.next]
	if actual != expected {
		fmt.Printf(`expected log "%s" got "%s"`, expected, actual)
	}
	l.next++
}

type TypeForLoggingInterface interface {
	Foo()
}

type TypeForLoggingCreated struct{}

func (t TypeForLoggingCreated) Foo() {}

type TypeForLoggingEmbedded struct {
	TypeForLoggingCreated      *TypeForLoggingCreated  `inject:""`
	TypeForLoggingInterface    TypeForLoggingInterface `inject:""`
	TypeForLoggingCreatedNamed *TypeForLoggingCreated  `inject:"name_for_logging"`
	Map                        map[string]string       `inject:"private"`
}

type TypeForLogging struct {
	TypeForLoggingEmbedded `inject:"inline"`
	TypeForLoggingCreated  *TypeForLoggingCreated `inject:""`
}

func TestInjectLogging() {
	g := inject.Graph{
		Logger: &logger{
			Expected: []string{
				"provided *main.TypeForLoggingCreated named name_for_logging",
				"provided *main.TypeForLogging",
				"provided embedded *main.TypeForLoggingEmbedded",
				"created *main.TypeForLoggingCreated",
				"assigned newly created *main.TypeForLoggingCreated to field TypeForLoggingCreated in *main.TypeForLogging",
				"assigned existing *main.TypeForLoggingCreated to field TypeForLoggingCreated in *main.TypeForLoggingEmbedded",
				"assigned *main.TypeForLoggingCreated named name_for_logging to field TypeForLoggingCreatedNamed in *main.TypeForLoggingEmbedded",
				"made map for field Map in *main.TypeForLoggingEmbedded",
				"assigned existing *main.TypeForLoggingCreated to interface field TypeForLoggingInterface in *main.TypeForLoggingEmbedded",
			},
		},
	}
	var v TypeForLogging

	err := g.Provide(
		&inject.Object{Value: &TypeForLoggingCreated{}, Name: "name_for_logging"},
		&inject.Object{Value: &v},
	)
	if err != nil {
		fmt.Println(err)
	}
	if err := g.Populate(); err != nil {
		fmt.Println(err)
	}
}

func main() {

	//TestInjectLogging()
	//TestInjectOnPrivateField()//注意大小写
	//TestInjectOnPrivateInterfaceField()
	//TestInjectPrivateInterface()
	//TestInjectNamedTwoSatisfyInterface()
	//TestErrorOnNonPointerNamedInject()
	//TestInjectInline()
	//TestInjectInlineOnPointer()
	//TestInjectInlinePrivate()
	//TestInjectWithStructValue()
	//TestInjectWithNonpointerStructValue()
	//TestPrivateIsFollowed()
	//TestDoesNotOverwriteInterface()
	//TestInterfaceIncludingPrivate()
	//TestInjectMapWithoutPrivate()
	//TestObjectString()
	TestGraphObjects()
	//RequireTag()
	//ErrorOnNonPointerInject()
	//InjectSimple()
	//DoesNotOverwrite()
	//Private()
	//TagWithJustColon()
	//ProvideWithFields()////////////////////
	//TestProvideNonPointer()
	//TestProvideTwoOfTheSame()
	//TestProvideTwoOfTheSameWithPopulate()
	//TestProvideTwoWithTheSameName()
	//TestNamedInstanceWithDependencies()
	//TestTwoNamedInstances()
	//TestCompleteProvides()

	/* 这几个一并考虑
	//TestInjectInterfaceMissing()//////////////
	//TestInjectInterface() ///////////////////
	//TestInjectTwoSatisfyInterface()
	*/
	//TestInvalidNamedInstanceType() //类型不一致

	//injectTest()

}

func injectTest() {
	// Typically an application will have exactly one object graph, and
	// you will create it and use it within a main function:
	var g inject.Graph

	// We provide our graph two "seed" objects, one our empty
	// HomePlanetRenderApp instance which we're hoping to get filled out,
	// and second our DefaultTransport to satisfiy our HTTPTransport
	// dependency. We have to provide the DefaultTransport because the
	// dependency is defined in terms of the http.RoundTripper interface,
	// and since it is an interface the library cannot create an instance
	// for it. Instead it will use the given DefaultTransport to satisfy
	// the dependency since it implements the interface:
	var a HomePlanetRenderApp
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: http.DefaultTransport},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Here the Populate call is creating instances of NameAPI &
	// PlanetAPI, and setting the HTTPTransport on both to the
	// http.DefaultTransport provided above:
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("a = %#v\n", a)
	fmt.Printf("http.DefaultTransport = %#v\n", http.DefaultTransport)

	for _, o := range g.Objects() {
		fmt.Printf(" = %#v\n", o)

	}

	// There is a shorthand API for the simple case which combines the
	// three calls above is available as inject.Populate:
	//
	//   inject.Populate(&a, http.DefaultTransport)
	//
	// The above API shows the underlying API which also allows the use of
	// named instances for more complex scenarios.

	fmt.Println(a.Render(42))
	// Output: Spock is from the planet Vulcan.
}
