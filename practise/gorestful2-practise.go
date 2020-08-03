package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id, Name string
}

type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

func (u UserResource) Register(container *restful.Container) {
	// 创建新的WebService
	ws := new(restful.WebService)

	// 设定WebService对应的路径("/users")和支持的MIME类型(restful.MIME_XML/ restful.MIME_JSON)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	// 添加路由： GET /{user-id} --> u.findUser
	ws.Route(ws.GET("/{user-id}").To(u.findUser))

	// 添加路由： POST / --> u.updateUser
	ws.Route(ws.POST("").To(u.updateUser))

	// 添加路由： PUT /{user-id} --> u.createUser
	ws.Route(ws.PUT("/{user-id}").To(u.createUser))

	// 添加路由： DELETE /{user-id} --> u.removeUser
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))

	// 将初始化好的WebService添加到Container中
	container.Add(ws)
}

// GET http://localhost:8080/users/1
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	fmt.Println("-------", id)
	usr := u.users[id]
	if len(usr.Id) == 0 {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

// POST http://localhost:8080/users
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = *usr
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// PUT http://localhost:8080/users/1
// <User><Id>1</Id><Name>Melissa</Name></User>
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	fmt.Println(id)
	usr := User{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = usr
		response.WriteHeader(http.StatusCreated)
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8080/users/1
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	fmt.Println(id)
	delete(u.users, id)
}

func main() {
	// 创建一个空的Container
	wsContainer := restful.NewContainer()

	// 设定路由为CurlyRouter(快速路由)
	wsContainer.Router(restful.CurlyRouter{})

	// 创建自定义的Resource Handle(此处为UserResource)
	u := UserResource{map[string]User{}}

	// 创建WebService，并将WebService加入到Container中
	u.Register(wsContainer)

	fmt.Println(wsContainer)
	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}

	// 启动服务
	log.Fatal(server.ListenAndServe())
}
