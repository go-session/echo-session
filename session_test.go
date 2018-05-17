package echosession

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"gopkg.in/session.v2"
)

func TestSession(t *testing.T) {
	cookieName := "test_echo_session"

	e := echo.New()
	e.Use(New(
		session.SetCookieName(cookieName),
		session.SetSign([]byte("sign")),
	))

	e.GET("/", func(ctx echo.Context) error {
		store := FromContext(ctx)
		if ctx.QueryParam("login") == "1" {
			foo, ok := store.Get("foo")
			fmt.Fprintf(ctx.Response(), "%s:%v", foo, ok)
			return nil
		}

		store.Set("foo", "bar")
		err := store.Save()
		if err != nil {
			t.Error(err)
			return nil
		}
		fmt.Fprint(ctx.Response(), "ok")
		return nil
	})

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
		return
	}
	e.ServeHTTP(w, req)

	res := w.Result()
	cookie := res.Cookies()[0]
	if cookie.Name != cookieName {
		t.Error("Not expected value:", cookie.Name)
		return
	}

	buf, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if string(buf) != "ok" {
		t.Error("Not expected value:", string(buf))
		return
	}

	req, err = http.NewRequest("GET", "/?login=1", nil)
	if err != nil {
		t.Error(err)
		return
	}
	req.AddCookie(cookie)

	w = httptest.NewRecorder()
	e.ServeHTTP(w, req)

	res = w.Result()
	buf, _ = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if string(buf) != "bar:true" {
		t.Error("Not expected value:", string(buf))
		return
	}
}
