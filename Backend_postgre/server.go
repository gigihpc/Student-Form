package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	// "gopkg.in/mgo.v2"
	"database/sql"

	_ "github.com/lib/pq"
)

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
type LoginResource struct {
	Data Login `json:"data"`
}

var mySigningKey = []byte("Scret")

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
}
type UserCollection struct {
	Data []User `json:"data"`
}
type UserResource struct {
	Data User `json:"data"`
}

// Mahasiswa Struct
type Mahasiswa struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Old     string `json:"old"`
}

type MahasiswaCollection struct {
	Data []Mahasiswa `json:"data"`
}
type MahasiswaResource struct {
	Data Mahasiswa `json:"data"`
}

//User
func (c *appContext) Create(user *User) (int, error) {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Country == "" {
		err := string("")
		if user.Name == "" {
			err = err + "Name Empty, "
		}
		if user.Email == "" {
			err = err + "Email Empty "
		}
		if user.Password == "" {
			err = err + "Password Empty "
		}
		if user.Country == "" {
			err = err + "Country Empty "
		}
		return 0, errors.New("Data can't save because not complated becaused: " + err)
	}
	id := 0
	sqlStatement := `insert into users (name, email, password, country) values($1, $2, $3, $4) RETURNING id`
	err := c.db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, user.Country).Scan(&id)
	user.ID = strconv.Itoa(id)
	return id, err
}

// func (r *UserRepo) All() (UserCollection, error) {
// 	res := UserCollection{[]User{}}
// 	err := r.coll.Find(nil).All(&res.Data)
// 	if err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

func (c *appContext) Matched(user string, password string) (UserResource, error) {
	res := UserResource{}
	var id int
	query := `select id from users where email=$1 and password=$2`
	err := c.db.QueryRow(query, user, password).Scan(&id)

	if err != nil {
		return res, err
	}
	return res, nil
}

//Mahasiswa
func (c *appContext) All() (MahasiswaCollection, error) {
	res := MahasiswaCollection{[]Mahasiswa{}}
	//For Limit
	// query := `select id, name, address, old from mhsw LIMIT $1`
	// rows, err := c.db.Query(query, 10)

	query := `select id, name, address, old from mhsw`
	rows, err := c.db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var _id, _name, _address, _old string
	for rows.Next() {
		err = rows.Scan(&_id, &_name, &_address, &_old)
		if err != nil {
			panic(err)
		}
		mhsw := Mahasiswa{Id: _id, Name: _name, Address: _address, Old: _old}
		res.Data = append(res.Data, mhsw)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	if err != nil {
		return res, err
	}
	return res, nil
}
func (c *appContext) Find(id string) (MahasiswaResource, error) {
	res := MahasiswaResource{}
	query := `select id, name, address, old from mhsw where id=$1`
	row := c.db.QueryRow(query, id)
	err := row.Scan(&res.Data.Id, &res.Data.Name, &res.Data.Address, &res.Data.Old)

	if err != nil {
		panic(err)
	}

	if err != nil {
		return res, err
	}
	return res, nil
}
func (c *appContext) CreateMhsw(mhsw *Mahasiswa) (int, error) {
	id := 0
	query := `insert into mhsw (name, address, old) values($1,$2,$3) RETURNING id`
	err := c.db.QueryRow(query, mhsw.Name, mhsw.Address, mhsw.Old).Scan(&id)
	mhsw.Id = strconv.Itoa(id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
func (c *appContext) Update(mhsw Mahasiswa) error {
	query := `update mhsw set name=$1, address=$2, old=$3 where id=$4`
	_, err := c.db.Query(query, mhsw.Name, mhsw.Address, mhsw.Old, mhsw.Id)

	if err != nil {
		return err
	}
	return nil
}
func (c *appContext) Delete(id string) error {
	query := "delete from mhsw where id=$1"
	_, err := c.db.Query(query, id)
	// _, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

//Main handler
type appContext struct {
	db *sql.DB
}

//User
func (c *appContext) createUserHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*UserResource)
	id, err := c.Create(&body.Data)
	if err != nil {
		println("erreor is: ", err.Error)
		// panic(err)
	}
	query := `select id, name, email, password, country from users where id=$1`
	row := c.db.QueryRow(query, id)

	err = row.Scan(&body.Data.ID, &body.Data.Name, &body.Data.Email, &body.Data.Password, &body.Data.Country)
	if err != nil {
		w.WriteHeader(203)
		json.NewEncoder(w).Encode(body)
		//panic(err)
	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(body)
	}
}

// func (c *appContext) UserHandler(w http.ResponseWriter, r *http.Request) {
// 	repo := UserRepo{c.db.C("user")}
// 	user, err := repo.All()
// 	println(user.Data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//w.Header().Set("Accept", "application/vnd.api+json")
// 	json.NewEncoder(w).Encode(user)
// }
func (c *appContext) MatchedHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*LoginResource)
	// repo := UserRepo{c.db.C("user")}
	_, err := c.Matched(body.Data.User, body.Data.Password)
	if err != nil {
		println("Failed Login")
		w.WriteHeader(401)
		w.Write([]byte("Failed"))
		json.NewEncoder(w)
	} else {
		println("Success Login")
		w.WriteHeader(200)
		w.Write([]byte(createToken()))
		json.NewEncoder(w)
	}
}

func createToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "Gigih"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	tokenString, _ := token.SignedString(mySigningKey)

	// if clm, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	// 	println("test claim: ", clm["admin"], clm["name"])
	// }

	return tokenString
}

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

//Mahasiswa
func (c *appContext) mhswHandler(w http.ResponseWriter, r *http.Request) {
	mhsw, err := c.All()
	if err != nil {
		panic(err)
	}
	//w.Header().Set("Accept", "application/vnd.api+json")
	json.NewEncoder(w).Encode(mhsw)
}
func (c *appContext) MahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	mhsw, err := c.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	//w.Header().Set("Accept", "application/vnd.api+json")
	// w.WriteHeader(201)
	json.NewEncoder(w).Encode(mhsw)
}
func (c *appContext) createHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*MahasiswaResource)
	id, err := c.CreateMhsw(&body.Data)
	if err != nil {
		panic(err)
	}
	query := `select id, name, address, old from mhsw where id=$1`
	row := c.db.QueryRow(query, id)
	err = row.Scan(&body.Data.Id, &body.Data.Name, &body.Data.Address, &body.Data.Old)

	if err != nil {
		panic(err)
	}
	//w.Header().Set("Accept", "application/vnd.api+json")
	w.WriteHeader(201) // response fill new resource being create
	json.NewEncoder(w).Encode(body)
}
func (c *appContext) updateHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	body := context.Get(r, "body").(*MahasiswaResource)
	body.Data.Id = params.ByName("id") //params[0].Value
	err := c.Update(body.Data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(204)
	w.Write([]byte("\n"))
}
func (c *appContext) deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	err := c.Delete(params.ByName("id"))
	if err != nil {
		panic(err)
	}

	w.WriteHeader(204) //response body update
	w.Write([]byte("\n"))
}

// Errors

type Errors struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func WriteError(w http.ResponseWriter, err *Error) {
	//w.Header().Set("Accept", "application/vnd.api+json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}

var (
	ErrBadRequest           = &Error{"bad_request", 400, "Bad request", "Request body is not well-formed. It must be JSON."}
	ErrNotAcceptable        = &Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/vnd.api+json'."}
	ErrUnsupportedMediaType = &Error{"unsupported_media_type", 415, "Unsupported Media Type", "Content-Type header must be set to: 'application/vnd.api+json'."}
	ErrInternalServer       = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
)

// Middlewares

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				WriteError(w, ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func acceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//if r.Header.Get("Accept") != "application/vnd.api+json" {
		//	WriteError(w, ErrNotAcceptable)
		//	return
		//}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func contentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//if r.Header.Get("Accept") != "application/vnd.api+json" {
		//	WriteError(w, ErrUnsupportedMediaType)
		//	return
		//}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func bodyHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)

	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)

			println("body: ", r.Body)

			if err != nil {
				WriteError(w, ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", val)
				next.ServeHTTP(w, r)
			}
		}

		return http.HandlerFunc(fn)
	}

	return m
}

// Router

type router struct {
	*httprouter.Router
}

func (r *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "content-type, Accept, X-XSRF-TOKEN,Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	r.Router.ServeHTTP(rw, req)
}

func (r *router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func (r *router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

func (r *router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

func (r *router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

func NewRouter() *router {
	return &router{httprouter.New()}
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

const (
	host     = "192.168.1.8"
	port     = 5432
	user     = "postgres"
	password = "gigih"
	dbname   = "myproject"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	appC := appContext{db}

	commonHandler := alice.New()

	router := NewRouter()
	router.Get("/api/mhsws/:id", commonHandler.ThenFunc(appC.MahasiswaHandler))
	router.Put("/api/mhsws/:id", commonHandler.Append(contentTypeHandler, bodyHandler(MahasiswaResource{})).ThenFunc(appC.updateHandler))
	router.Delete("/api/mhsws/:id", commonHandler.ThenFunc(appC.deleteHandler))
	router.Get("/api/mhsws", jwtMiddleware.Handler(commonHandler.ThenFunc(appC.mhswHandler)))
	router.Post("/api/mhsws", jwtMiddleware.Handler(commonHandler.Append(contentTypeHandler, bodyHandler(MahasiswaResource{})).ThenFunc(appC.createHandler)))
	router.Post("/api/user", commonHandler.Append(contentTypeHandler, bodyHandler(UserResource{})).ThenFunc(appC.createUserHandler))
	// router.Get("/api/user", commonHandler.ThenFunc(appC.UserHandler))
	router.Post("/api/user_auth", commonHandler.Append(contentTypeHandler, bodyHandler(LoginResource{})).ThenFunc(appC.MatchedHandler))
	port := "8002"
	println("open port: " + port)
	http.ListenAndServe(":"+port, router)
}
