package crud

import (
	"fmt"
	"reflect"
)

type AbstractRepository[T any] interface {
	// SaveOrUpdate 新增或者更新一条数据，返回id
	SaveOrUpdate(bean T) (error, int64)

	// Query 条件查询
	Query(map[string]interface{}) (error, []T)

	// Get 根据ID获取
	Get(id int64) T

	// Delete 根据ID删除操作返回是否删除成功
	Delete(id int64) bool
}

type PersonRepository struct {
	AbstractRepository[Person]
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{}
}

func (a *PersonRepository) SaveOrUpdate(bean Person) (error, int64) {
	fmt.Println(reflect.TypeOf(bean))
	return nil, bean.Id
}

func (a *PersonRepository) Query(map[string]interface{}) (error, []Person) {
	return nil, []Person{}
}

func (a *PersonRepository) Get(id int64) Person {
	return Person{}
}

func (a *PersonRepository) Delete(id int64) bool {
	return false
}

type Person struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}

func (p *Person) say() string {
	return fmt.Sprintf("hello %s\n", p.Username)
}
