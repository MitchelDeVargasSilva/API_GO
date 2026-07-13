package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tarefa struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
}

var ListaTarefas = []Tarefa{
	{Id: 1, Titulo: "Estudar GO"},
	{Id: 2, Titulo: "Criar projetos"},
}

func Rotas(router *gin.Engine) {

	//GET inicial de teste
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Primeira Api com Go rodando",
		})
	})

	//GET para retornar todas as tarefas
	router.GET("/tarefas", func(c *gin.Context) {
		rows, err := DB.Query("SELECT id, titulo FROM tarefas")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		defer rows.Close()

		var tasks []Tarefa

		for rows.Next() {
			var task Tarefa
			if err := rows.Scan(&task.Id, &task.Titulo); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			tasks = append(tasks, task)
		}

		c.JSON(http.StatusOK, tasks)
	})

	//POST para adicionar novas tarefas
	router.POST("/tarefas", func(c *gin.Context) {
		var novaTarefa Tarefa

		if err := c.BindJSON(&novaTarefa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro": err.Error(),
			})
			return
		}

		result, err := DB.Exec("Insert into tarefas (titulo) values (?)", novaTarefa.Titulo)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro": err.Error(),
			})
			return
		}

		id, err := result.LastInsertId()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro": err.Error(),
			})
			return
		}

		novaTarefa.Id = int(id)

		c.JSON(http.StatusOK, novaTarefa)
	})

	//GET buscando tarefas por id
	router.GET("/tarefas/:id", func(c *gin.Context) {
		id := c.Param("id")

		var tarefas Tarefa
		row := DB.QueryRow("SELECT id, titulo FROM tarefas WHERE id = ?", id)

		if err := row.Scan(&tarefas.Id, &tarefas.Titulo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tarefas)

	})
	//DELETE apagando tarefa pelo id
	router.DELETE("/tarefas/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := DB.Exec("DELETE FROM tarefas WHERE id = ?", id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Tarefa deletada com sucesso!"})
	})

	//Atualizando registros
	router.PUT("/tarefas/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var updatedTarefa Tarefa

		if err := c.BindJSON(&updatedTarefa); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Tarefa não encontrada!"})
			return
		}

		_, err := DB.Exec("UPDATE tarefas SET titulo = ? WHERE id = ?", updatedTarefa.Titulo, id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedTarefa.Id = id
		c.JSON(http.StatusOK, updatedTarefa)
	})
}
