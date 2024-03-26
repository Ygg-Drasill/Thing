package pages

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SubmitHandler(context *gin.Context) {
	session := sessions.Default(context)
	person := context.PostForm("person")
	fine := context.PostForm("fine")
	session.Set("fine_"+person, fine)
	session.Set("selectedPerson", person)
	session.Save()

	FinesPage(context)
}
