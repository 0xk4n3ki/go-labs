package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (s *GinServer) AddUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData UserParam
		if err := ctx.BindJSON(&userData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user := User{
			Name: userData.Name,
			Email: userData.Email,
			Age: userData.Age,
		}

		if err := s.db.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, user)
	}
}

func (s *GinServer) ListUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var users []User
		if err := s.db.Find(&users).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		ctx.JSON(http.StatusOK, users)
	}
}

func (s *GinServer) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userData UserParam
		if err := ctx.BindJSON(&userData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userId := ctx.Param("id")
		var user User
		if err := s.db.First(&user, userId).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		user.Name = userData.Name
		user.Email = userData.Email
		user.Age = userData.Age

		if err := s.db.Save(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

func (s *GinServer) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		var user User
		if err := s.db.First(&user, userId).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err := s.db.Delete(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
	}
}