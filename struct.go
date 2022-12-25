package main

import "fmt"

type person struct {
	name string
	age  int
}

type person2 struct {
	name string
	age  int
}

// Todo構造体を定義する
type Todo struct {
	ID   string
	Text string
	Done bool
}

// Todo構造体のポインタ型を定義する
type TodoPtr []*Todo

// 既存の型に新しい名前をつけているのでtypeを使用
// person2のポインタが入ったスライス
type person3 []*person2

type person4 []person2

type person5 person2

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func renamePerson(p []*person2) {
	// [0xc000010030]
	fmt.Println(p)

	fmt.Println(p[0])

	// 構造体の要素にアクセスするときはポインタであっても*をつける必要がない
	p[0].age = 100
	fmt.Println(p[0])

}

func renamePersonWithoutPtr(p []person2) {
	fmt.Println(p)

	p[0].age = 200

}

func renamePersonSingle(p person5) person5 {

	p.age = 1000
	return p

}

// TodoPtr型に対するDoneフィールドの値を変更するメソッド
func (t TodoPtr) SetDone(done bool) {
	for _, todo := range t {
		//スライスの値にアクセスして値を変えるときはポインタである必要はない
		todo.Done = done
	}
}

func struct_main() {
	// Todoの一覧を作成する
	//TodoPtrはTodo構造体のポインタ型なので中身もポインタでなければならない。
	//だから&をつけてポインタにしている
	todos := TodoPtr{
		&Todo{ID: "1", Text: "買い物をする", Done: false},
		&Todo{ID: "2", Text: "宿題をする", Done: false},
		&Todo{ID: "3", Text: "運動する", Done: false},
	}

	// 全てのTodoのDoneフィールドの値をtrueにする
	todos.SetDone(true)
	// {Bob 20}
	// fmt.Println(person{"Bob", 20})

	// &{Ann 40}
	// fmt.Println(&person{name: "Ann", age: 40})

	//person2のポインタであってname, ageのポインタではないのでポインタは1つしかない
	personPtr := person3{&person2{name: "Bob", age: 40}}

	// [0xc000010030]
	fmt.Println(&personPtr)

	// // &{Bob 100}
	// renamePerson(personPtr)

	// // {Bob 100}
	// fmt.Println(*personPtr[0])

	personWithoutPtr := person4{person2{name: "Joy", age: 55}}

	renamePersonWithoutPtr(personWithoutPtr)

	fmt.Println(personWithoutPtr)

	peronSingle := person5{name: "Bob", age: 3}
	fmt.Println(peronSingle)
	var val person5
	val = renamePersonSingle(peronSingle)

	fmt.Println(val)

	// fmt.Println(&person2{name: "Ann", age: 40})

	person1 := &person{name: "Ann", age: 40}
	fmt.Println(person1)

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
