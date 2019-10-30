package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Joke contains information about a single Joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

// We'll create a list of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

func main() {
	//Usa el router default provisto por gingonic
	router := gin.Default()

	//Sirve los archivos estáticos para el FrontEnd
	router.Use(static.Serve("/", static.LocalFile("./views/js", true)))

	//Setea el grupo de ruta para la API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Mensaje": "Hola mundo",
			})
		})

		api.GET("/jokes", JokeHandler)
		api.POST("/jokes/like/:jokeID", LikeJoke)
	}

	//Arranca el servidor
	router.Run(":3000")
}

//JokeHandler servers dad jokes at an impressive rate
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

//LikeJoke allows the user to Like their lovely jokes
func LikeJoke(c *gin.Context) {

	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes++
			}
		}
		//Devolvemos un puntero a la lista de chistes actualizada
		c.JSON(http.StatusOK, &jokes)
	} else {
		//Joke ID invalido
		c.AbortWithStatus(http.StatusNotFound)
	}
}
