package tests

import (
	"backend/controllers"
	"backend/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup router khusus category testing
func setupCategoryRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/categories", controllers.GetCategories)
	r.POST("/categories", controllers.CreateCategory)

	return r
}


func TestCategoryCRUD(t *testing.T) {
	r := setupCategoryRouter()

	// --- Create Category ---
	body := `{"name":"Test Category","color":"red"}`
	req := httptest.NewRequest("POST", "/categories", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var categoryResp struct {
		Data models.Category `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &categoryResp)
	assert.Nil(t, err)
	category := categoryResp.Data
	assert.Equal(t, "Test Category", category.Name)
	assert.Equal(t, "red", category.Color)

	// --- Get All Categories ---
	req2 := httptest.NewRequest("GET", "/categories", nil)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)

	var allCategoriesResp struct {
		Data []models.Category `json:"data"`
	}
	err = json.Unmarshal(w2.Body.Bytes(), &allCategoriesResp)
	assert.Nil(t, err)
	assert.Len(t, allCategoriesResp.Data, 1)
	assert.Equal(t, "Test Category", allCategoriesResp.Data[0].Name)
}
