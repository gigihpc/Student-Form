package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Login struct {
	ID       bson.ObjectId `json:"id, omitempty"`
	User     string        `json:"user"`
	Password string        `json:"password"`
}
type LoginResource struct {
	Data Login `json:"data"`
}

type User struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Country  string        `json:"country"`
}
type UserCollection struct {
	Data []User `json:"data"`
}
type UserResource struct {
	Data User `json:"data"`
}
type UserRepo struct {
	coll *mgo.Collection
}

// Mahasiswa Struct
type Mahasiswa struct {
	Id      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string        `json:"name"`
	Address string        `json:"address"`
	Old     string        `json:"old"`
}

type MahasiswaCollection struct {
	Data []Mahasiswa `json:"data"`
}
type MahasiswaResource struct {
	Data Mahasiswa `json:"data"`
}
type MahasiswaRepo struct {
	coll *mgo.Collection
}

//User
func (r *UserRepo) Create(user *User) error {
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
		return errors.New("Data can't save because not complated becaused: " + err)
	}
	id := bson.NewObjectId()
	_, err := r.coll.UpsertId(id, user)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}
func (r *UserRepo) All() (UserCollection, error) {
	res := UserCollection{[]User{}}
	err := r.coll.Find(nil).All(&res.Data)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *UserRepo) Matched(user string, password string) (UserResource, error) {
	res := UserResource{}
	err := r.coll.Find(bson.M{"email": user, "password": password}).One(&res.Data)

	if err != nil {
		return res, err
	}
	return res, nil
}

//Mahasiswa
func (r *MahasiswaRepo) All() (MahasiswaCollection, error) {
	res := MahasiswaCollection{[]Mahasiswa{}}
	err := r.coll.Find(nil).All(&res.Data)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (r *MahasiswaRepo) Find(id string) (MahasiswaResource, error) {
	res := MahasiswaResource{}
	err := r.coll.FindId(bson.ObjectIdHex(id)).One(&res.Data)
	if err != nil {
		return res, err
	}
	return res, nil
}
func (r *MahasiswaRepo) Create(mhsw *Mahasiswa) error {
	id := bson.NewObjectId()
	_, err := r.coll.UpsertId(id, mhsw)
	if err != nil {
		return err
	}
	mhsw.Id = id
	return nil
}
func (r *MahasiswaRepo) Update(mhsw *Mahasiswa) error {
	err := r.coll.UpdateId(mhsw.Id, mhsw)
	if err != nil {
		return err
	}
	return nil
}
func (r *MahasiswaRepo) Delete(id string) error {
	err := r.coll.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}

//Main handler
type appContext struct {
	db *mgo.Database
}

//User
func (c *appContext) createUserHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*UserResource)
	repo := UserRepo{c.db.C("user")}
	// body, checkerr := checkBodyUser(body)
	err := repo.Create(&body.Data)
	log.Println(err)

	if err != nil {
		w.WriteHeader(203)
		json.NewEncoder(w).Encode(body)
		//panic(err)
	} else {
		w.WriteHeader(202)
		json.NewEncoder(w).Encode(body)
	}
}
func (c *appContext) UserHandler(w http.ResponseWriter, r *http.Request) {
	repo := UserRepo{c.db.C("user")}
	user, err := repo.All()
	println(user.Data)
	if err != nil {
		panic(err)
	}
	//w.Header().Set("Accept", "application/vnd.api+json")
	json.NewEncoder(w).Encode(user)
}
func (c *appContext) MatchedHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*LoginResource)
	repo := UserRepo{c.db.C("user")}
	_, err := repo.Matched(body.Data.User, body.Data.Password)
	if err != nil {
		println("Failed Login")
		w.WriteHeader(401)
		w.Write([]byte("Failed"))
		json.NewEncoder(w)
	} else {
		println("Success Login")
		w.WriteHeader(200)
		json.NewEncoder(w)
	}
}

/*func checkBodyUser(body *UserResource) (*UserResource, error) {
	resbody := UserResource{}
	err := error(nil)
	if body.Data.Name == "" {
		resbody.Data.Name = ""
		err = errors.New("Name empty")
	} else {
		resbody.Data.Name = body.Data.Name
	}
	if body.Data.Paswword == "" {
		resbody.Data.Password = ""
		err = errors.New("Password Empty")
	} else {
		resbody.Data.Password = body.Data.Paswword
	}
	if body.Data.Email == "" {
		resbody.Data.Email = ""
		err = errors.New("Email Empty")
	} else {
		resbody.Data.Email = body.Data.Email
	}
	if body.Data.Country == "" {
		resbody.Data.Country = ""
		err = errors.New("Country Empty")
	} else {
		resbody.Data.Country = body.Data.Country
	}
	return &resbody, err
}*/

//Mahasiswa
func (c *appContext) mhswHandler(w http.ResponseWriter, r *http.Request) {
	repo := MahasiswaRepo{c.db.C("mhsws")}
	mhsw, err := repo.All()
	if err != nil {
		panic(err)
	}
	//w.Header().Set("Accept", "application/vnd.api+json")
	json.NewEncoder(w).Encode(mhsw)
}
func (c *appContext) MahasiswaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := MahasiswaRepo{c.db.C("mhsws")}
	mhsw, err := repo.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	//w.Header().Set("Accept", "application/vnd.api+json")
	// w.WriteHeader(201)
	json.NewEncoder(w).Encode(mhsw)
}
func (c *appContext) createHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*MahasiswaResource)
	repo := MahasiswaRepo{c.db.C("mhsws")}
	err := repo.Create(&body.Data)

	log.Println(err)

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
	body.Data.Id = bson.ObjectIdHex(params.ByName("id"))
	repo := MahasiswaRepo{c.db.C("mhsws")}
	err := repo.Update(&body.Data)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(204)
	w.Write([]byte("\n"))
}
func (c *appContext) deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := MahasiswaRepo{c.db.C("mhsws")}
	err := repo.Delete(params.ByName("id"))
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

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//c := session.DB("myproject").C("mhsws")
	//
	//err = c.Insert(&Mahasiswa{Name: "Michel", Address: "England", Old: "23"})
	//if err != nil {
	//	panic(err)
	//}
	appC := appContext{session.DB("myproject")}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://192.168.1.8:8080"},
		AllowedMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders: []string{"content-type", "Accept", "X-XSRF-TOKEN", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		//ExposedHeaders:   []string{"*", "/"},
		AllowCredentials: true,
	})

	commonHandler := alice.New( /*context.ClearHandler, loggingHandler, recoverHandler, acceptHandler,*/ c.Handler)

	router := NewRouter()
	router.Get("/api/mhsws/:id", commonHandler.ThenFunc(appC.MahasiswaHandler))
	router.Put("/api/mhsws/:id", commonHandler.Append(contentTypeHandler, bodyHandler(MahasiswaResource{})).ThenFunc(appC.updateHandler))
	router.Delete("/api/mhsws/:id", commonHandler.ThenFunc(appC.deleteHandler))
	router.Get("/api/mhsws", commonHandler.ThenFunc(appC.mhswHandler))
	router.Post("/api/mhsws", commonHandler.Append(contentTypeHandler, bodyHandler(MahasiswaResource{})).ThenFunc(appC.createHandler))
	router.Post("/api/user", commonHandler.Append(contentTypeHandler, bodyHandler(UserResource{})).ThenFunc(appC.createUserHandler))
	router.Get("/api/user", commonHandler.ThenFunc(appC.UserHandler))
	router.Post("/api/user_auth", commonHandler.Append(contentTypeHandler, bodyHandler(LoginResource{})).ThenFunc(appC.MatchedHandler))

	port := "8001"
	println("open port: " + port)
	http.ListenAndServe(":"+port, router)
}
