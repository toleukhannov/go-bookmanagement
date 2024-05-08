package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/batyrbek/pkg/models"
    // "github.com/batyrbek/pkg/service"
)

// RegisterHandler обрабатывает запросы на регистрацию нового пользователя
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Здесь вы можете выполнить проверки на валидность данных пользователя, например, проверку уникальности имени пользователя

    // После проверок вы можете сохранить пользователя в базе данных или выполнить другие необходимые действия
    // Например:
    // userService.CreateUser(user)

    // Отправляем успешный ответ клиенту
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Пользователь успешно зарегистрирован")
}

// LoginHandler обрабатывает запросы на аутентификацию пользователей
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Здесь вы должны выполнить аутентификацию пользователя, сравнивая его данные с данными из базы данных или другого источника

    // Например, если у вас есть сервис для работы с пользователями, вы можете использовать его для аутентификации:
    // if userService.Authenticate(user.Username, user.Password) {
    //     // Если аутентификация успешна, генерируем JWT токен
    //     token, err := service.GenerateToken(user.Username)
    //     if err != nil {
    //         http.Error(w, err.Error(), http.StatusInternalServerError)
    //         return
    //     }
    //     // Отправляем токен клиенту
    //     json.NewEncoder(w).Encode(token)
    // } else {
    //     http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
    // }

    // В данном примере возвращаем простое сообщение об успешной аутентификации
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Аутентификация успешна")
}



// package handler

// import (
// 	"net/http"

// 	"github.com/batyrbek/pkg/models"
// 	"github.com/gin-gonic/gin"
// )

// func (h *Handler) signUp(c *gin.Context){
// 	var input models.User
// 	if err := c.BindJSON(&input); err != nil{
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	id, err := h.services.Authorization.CreateUser(input)
// 	if err != nil{
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error)
// 		return
// 	}

// 	c.JSON(http.StatusOK, map[string]interface{}){
// 		"id": id,
// 	}
// }

// type SignInInput struct{
// 	Email    string `gorm:""json:"email" binding: "required"`
// 	Password string `gorm:""json:"password" binding: "required"`
// }

// func (h *Handler) signIn(c *gin.Context){
// 	var input models.User
// 	if err := c.BindJSON(&input); err != nil{
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	id, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
// 	if err != nil{
// 		newErrorResponse(c, http.StatusInternalServerError, err.Error)
// 		return
// 	}

// 	c.JSON(http.StatusOK, map[string]interface{}){
// 		"id": id,
// 	}
// } 