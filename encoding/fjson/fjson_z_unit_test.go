package fjson_test

import (
	"testing"
	"time"

	"github.com/ZYallers/fine/encoding/fjson"
	"github.com/ZYallers/fine/test/ftest"
	"github.com/ZYallers/fine/text/fstr"
)

func Test_Valid(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		data := []byte(`{"msg":"ok", "code":200}`)
		t.Assert(fjson.Valid(data), true)
	})
}

func Test_Marshal(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		user := User{
			Name: "john",
			Age:  18,
		}
		b, err := fjson.Marshal(user)
		t.Assert(err, nil)
		t.Assert(string(b), `{"name":"john","age":18}`)
	})
}

func Test_Unmarshal(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		data := `{"name":"john","age":18}`
		err := fjson.Unmarshal([]byte(data), &user)
		t.Assert(err, nil)
		t.Assert(user.Name, "john")
		t.Assert(user.Age, 18)
	})
}

func Test_Time_Marshal(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Name string     `json:"name"`
			Age  int        `json:"age"`
			Time *time.Time `json:"time"`
		}
		ts, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-01 12:12:12", time.Local)
		user := User{
			Name: "john",
			Age:  18,
			Time: &ts,
		}
		b, err := fjson.Marshal(user)
		t.AssertNil(err)
		t.Assert(fstr.Contains(string(b), `"time":"2020-01-01 12:12:12"`), true)
	})
}

func Test_Time_Format(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Name string     `json:"name"`
			Age  int        `json:"age"`
			Time *time.Time `json:"time" time_format:"2006-01-02"`
		}
		ts, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-01 12:12:12", time.Local)
		user := User{
			Name: "john",
			Age:  18,
			Time: &ts,
		}
		b, err := fjson.Marshal(user)
		t.AssertNil(err)
		t.Assert(fstr.Contains(string(b), `"time":"2020-01-01"`), true)
	})
}

func Test_Time_Location(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		type User struct {
			Name string     `json:"name"`
			Age  int        `json:"age"`
			Time *time.Time `json:"time" time_location:"Japan"`
		}
		ts, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-01 12:12:12", time.Local)
		user := User{
			Name: "john",
			Age:  18,
			Time: &ts,
		}
		b, err := fjson.Marshal(user)
		t.AssertNil(err)
		t.Assert(fstr.Contains(string(b), `"time":"2020-01-01 13:12:12"`), true)
	})
}
