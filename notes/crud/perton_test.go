package crud

import (
	"fmt"
	"testing"
)

func TestPersonCrud(t *testing.T) {
	pp := &Person{
		Id:       24234,
		Username: "zhangsan",
		Gender:   "male",
		Email:    "23432",
	}
	repository := NewPersonRepository()
	err, id := repository.SaveOrUpdate(*pp)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
	fmt.Println(pp.say())
}
