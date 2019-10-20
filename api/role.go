package api

import "github.com/gin-gonic/gin"

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Auth []int  `json:"auth"`
}

func (r *Role) Add(c *gin.Context) {

}

func (r *Role) Delete(c *gin.Context) {

}

func (r *Role) Modify(c *gin.Context) {

}

func (r *Role) Get(c *gin.Context) {

}

func NewRole() *Role {
	return new(Role)
}
