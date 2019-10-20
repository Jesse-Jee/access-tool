package api

import "github.com/gin-gonic/gin"

type Rule struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	TeamId int    `json:"team_id"`
	Role   []int  `json:"role"`
	Custom int    `json:"custom"`
}

func (r *Rule) Add(c *gin.Context) {

}

func (r *Rule) Delete(c *gin.Context) {

}

func (r *Rule) Modify(c *gin.Context) {

}

func (r *Rule) Get(c *gin.Context) {

}

func NewRule() *Rule {
	return new(Rule)
}
