package api

import (
	"access-tool/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Team struct {
	Id       int    `json:"id"`
	TeamName string `json:"team_name"`
}

func (t *Team) Add(c *gin.Context) {
	teamName := c.PostForm("team_name")
	if len(teamName) == 0{
		c.String(http.StatusBadRequest,"team_name can't null")
		return
	}
	if QueryTeamIsExist(teamName) {
		c.String(http.StatusBadRequest,"%s is exist!",teamName)
		return
	}
	repo:= repo.AccessRepo{}
	repo.TeamName = teamName


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
