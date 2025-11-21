package tests

import (
	"backend/controllers"
	"backend/initializers"
	"backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup router khusus testing
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos/:id", controllers.GetTodoById)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)
	r.PATCH("/todos/:id/complete", controllers.ToggleComplete)

	return r
}



func TestTodoCRUD(t *testing.T) {
	r := setupRouter()

	// --- Buat category dulu ---
	category := models.Category{Name: "Test Category", Color: "blue"}
	err := initializers.DB.Create(&category).Error
	if err != nil {
		t.Fatal(err)
	}

	// --- Create Todo --- (pakai category_id valid)
	body := fmt.Sprintf(`{"title":"Test Todo","description":"desc","priority":"high","category_id":%d}`, category.ID)
	req := httptest.NewRequest("POST", "/todos", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var todoResp struct {
		Data models.Todo `json:"data"`
	}
	err = json.Unmarshal(w.Body.Bytes(), &todoResp)
	assert.Nil(t, err)
	todo := todoResp.Data
	assert.Equal(t, "Test Todo", todo.Title)
	assert.False(t, todo.Completed)

	// --- Get Todo By ID ---
	req2 := httptest.NewRequest("GET", "/todos/"+strconv.Itoa(int(todo.ID)), nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)

	var todoByIDResp struct {
		Data models.Todo `json:"data"`
	}
	err = json.Unmarshal(w2.Body.Bytes(), &todoByIDResp)
	assert.Nil(t, err)
	assert.Equal(t, todo.ID, todoByIDResp.Data.ID)

	// --- Update Todo ---
	updateBody := fmt.Sprintf(`{"title":"Updated Todo","description":"desc update","priority":"medium","category_id":%d}`, category.ID)
	req3 := httptest.NewRequest("PUT", "/todos/"+strconv.Itoa(int(todo.ID)), strings.NewReader(updateBody))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()
	r.ServeHTTP(w3, req3)
	assert.Equal(t, http.StatusOK, w3.Code)

	var updatedTodoResp struct {
		Data models.Todo `json:"data"`
	}
	err = json.Unmarshal(w3.Body.Bytes(), &updatedTodoResp)
	assert.Nil(t, err)
	assert.Equal(t, "Updated Todo", updatedTodoResp.Data.Title)
	assert.Equal(t, "medium", updatedTodoResp.Data.Priority)

	// --- Toggle Complete ---
	req4 := httptest.NewRequest("PATCH", "/todos/"+strconv.Itoa(int(todo.ID))+"/complete", nil)
	w4 := httptest.NewRecorder()
	r.ServeHTTP(w4, req4)
	assert.Equal(t, http.StatusOK, w4.Code)

	var toggledResp struct {
		Data models.Todo `json:"data"`
	}
	err = json.Unmarshal(w4.Body.Bytes(), &toggledResp)
	assert.Nil(t, err)
	assert.True(t, toggledResp.Data.Completed)

	// --- Test filter by category (before delete) ---
	reqCat := httptest.NewRequest("GET", fmt.Sprintf("/todos?category_id=%d", category.ID), nil)
	wCat := httptest.NewRecorder()
	r.ServeHTTP(wCat, reqCat)
	assert.Equal(t, http.StatusOK, wCat.Code)

	var filteredTodosResp struct {
		Data []models.Todo `json:"data"`
	}
	err = json.Unmarshal(wCat.Body.Bytes(), &filteredTodosResp)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(filteredTodosResp.Data))
	assert.Equal(t, category.ID, filteredTodosResp.Data[0].CategoryID)

	// --- Delete Todo ---
	req5 := httptest.NewRequest("DELETE", "/todos/"+strconv.Itoa(int(todo.ID)), nil)
	w5 := httptest.NewRecorder()
	r.ServeHTTP(w5, req5)
	assert.Equal(t, http.StatusOK, w5.Code)

	// --- Get All should be empty ---
	req6 := httptest.NewRequest("GET", "/todos", nil)
	w6 := httptest.NewRecorder()
	r.ServeHTTP(w6, req6)
	assert.Equal(t, http.StatusOK, w6.Code)

	var allTodosResp struct {
		Data []models.Todo `json:"data"`
	}
	err = json.Unmarshal(w6.Body.Bytes(), &allTodosResp)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(allTodosResp.Data))
}
