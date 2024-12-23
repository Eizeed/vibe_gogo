package handlers

import (
	"net/http"

	"github.com/Eizeed/vibe_gogo/forms"
	"github.com/Eizeed/vibe_gogo/models"
	"github.com/Eizeed/vibe_gogo/services"
	"github.com/Eizeed/vibe_gogo/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {};

var userService = services.UserService {};

func (h *UserHandler) GetAll(c *gin.Context) {
    res, err := userService.GetAll();
    if err != nil {
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "Got all users", "users": res})
}

func (h *UserHandler) Register(c *gin.Context) {
    var userForm forms.UserForm;
    var registerForm forms.RegisterForm;

    if err := c.ShouldBindJSON(&registerForm); err != nil {
        message := userForm.Register(err);
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message})
        return
    }

    res, err := userService.Create(registerForm);
    if err != nil {
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *UserHandler) Login(c *gin.Context) {
    var userForm forms.UserForm;
    var loginForm forms.LoginForm;

    if err := c.ShouldBindJSON(&loginForm); err != nil {
        message := userForm.Login(err);
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message})
        return
    }

    res, token, err := userService.Login(loginForm);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.SetCookie("jwt_token", token, 3600, "/api", "localhost", false, true);
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in", "user": res, "token": token})
}

func (h *UserHandler) Logout(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, gin.H{"message": "Logout user"})
}

func (h *UserHandler) Update(c *gin.Context) {
    var v utils.Validate;
    var updateForm forms.UpdateForm;
    token, exists := c.Get("token");
    if !exists {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claims, ok := token.(models.Claims);
    if !ok {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to get claims from token"})
        return
    }

    if err := c.ShouldBindJSON(&updateForm); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request??"})
        return
    }


    if updateForm.Email != "" {
        if !v.Email(updateForm.Email) {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
            return
        }
    }
    if updateForm.Username != "" {
        if !v.Username(updateForm.Username) {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
            return
        }
    }
    if updateForm.Fullname != "" {
        if !v.Fullname(updateForm.Fullname) {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid name"})
            return
        }
    }
    if updateForm.Password != "" {
        if len(updateForm.Password) < 3 {
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid password. Must be at least 3 characters long"})
            return
        }
    }

    res, err := userService.Update(updateForm, claims.UserUUID);
    if err != nil {
        return
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": res})
}

func (h *UserHandler) Delete(c *gin.Context) {
    token, exists := c.Get("token");
    if !exists {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    claims, ok := token.(*models.Claims);
    if !ok {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unable to get claims from token"})
        return
    }

    uuid, err := uuid.Parse(c.Param("uuid"));
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user uuid"})
        return
    }

    if uuid != claims.UserUUID {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "uuids don't match"})
        return
    }

    res, err := userService.Delete(uuid);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"message": "User delete successfully", "user": res})
}














