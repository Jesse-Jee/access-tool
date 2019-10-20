package api

import "github.com/gin-gonic/gin"

type Team struct {
	Id       int    `json:"id"`
	TeamName string `json:"team_name"`
}

func (t *Team) Add(c *gin.Context) {

}

func (t *Team) Delete(c *gin.Context) {

}

func (t *Team) Modify(c *gin.Context) {

}

func (t *Team) Get(c *gin.Context) {

}

func NewTeam() *Team {
	return new(Team)
}
