package student

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

type StudentController struct {
	Dao *sql.DB
}
type Student struct {
	Roll int    `json:"roll"`
	Name string `json:"name"`
}
type ServerErr struct {
	Code int
	Msg  string
}

func NewController() *StudentController {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/raj")
	if err != nil {
		panic("not able to get the db connection" + err.Error())
	}
	v := StudentController{Dao: db}
	return &v
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func PutStudent() {
	runtime.GC()
	s := studentPool.Get().(*Student)
	defer studentPool.Put(s)
	json.NewDecoder(strings.NewReader(`{"roll":8,"name":"raj"}`)).Decode(&s)
	//fmt.Println(err)
}

// go:noinline
func getStudent() *Student {
	return &Student{}
}
func PutStudent1() {
	s := Student{}
	json.NewDecoder(strings.NewReader(`{"roll":8,"name":"raj"}`)).Decode(&s)
	//fmt.Println(err)
}

func (a StudentController) PutStudent(w http.ResponseWriter, t *http.Request) {
	s := studentPool.Get().(Student)
	defer studentPool.Put(s)
	err := json.NewDecoder(t.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ServerErr{400, "input is not proper"})
		return
	}
	a.AddStudent(s)

}
func (a StudentController) AddStudent(s Student) error {
	ab := "insert into student values (?,?)"
	st, err := a.Dao.Prepare(ab)
	if err != nil {
		return err
	}
	_, err = st.Exec(s.Roll, s.Name)
	return err

}
func (a StudentController) GetStudent(w http.ResponseWriter, t *http.Request) {
	return
}
