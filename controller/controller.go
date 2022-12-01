package src

import (
	"fmt"
	"gin_app/maria"
	"gin_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaevor/go-nanoid"
)

var (
	db        = maria.DB
	user_list []models.User
	user      models.User
)

// @Summary 测试
// @Schemes
// @Description 行内以下描述
// @Tags 测试
// @Accept json
// @Produce json
// @Router /test [get]
func FindTest(ctx *gin.Context) {

	// 查全部
	db.Find(&user_list)

	// 查单条
	// db.First(&u)

	// fmt.Printf("%v", u)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user_list,
	})
}

// @Summary 用户注册
// @Schemes
// @Description 行内以下描述
// @Tags 用户相关
// @param user body models.User true "提交的参数"
// @Accept json
// @Produce json
// @Router /regis [post]
func RegisterUser(ctx *gin.Context) {
	if err := ctx.ShouldBind(&user); err != nil {
		return
	}
	fmt.Println(user)
	db.Select("name", "pswd", "stateMsg", "portrait").Create(&user)
}

// @Summary 添加用户
// @Schemes
// @Description 行内以下描述
// @Tags 用户相关
// @param name formData string true "名称"
// @param pswd formData string true "密码"
// @Accept mpfd
// @Produce mpfd
// @Router /add [post]
func AddUser(ctx *gin.Context) {
	name := ctx.PostForm("name")
	pswd := ctx.PostForm("pswd")

	// 自动生成消息id
	nanoID, err := nanoid.Standard(8)
	if err != nil {
		return
	}

	user.Mid = nanoID()
	user.Name = name
	user.Password = pswd
	db.Omit("id", "stateMsg", "portrait").Create(&user)
}

// @Summary 删除用户
// @Schemes
// @Description 行内以下描述
// @Tags 用户相关
// @param name path string true "提交的参数"
// @Accept json
// @Produce json
// @Router /del/{name} [delete]
func DeleteByName(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
	db.Where("name = ?", name).Delete(&user)
}

// @Summary 更新用户
// @Schemes
// @Description 根据name更新用户信息
// @Tags 用户相关
// @param name path string true "被更改的名字"
// @param user body models.User true "更改的参数"
// @Accept json
// @Produce json
// @Router /upd/{name} [patch]
func UpdateByName(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.ShouldBind(&user)
	fmt.Println(name)
	db.Model(&user).Where("name=?", name).Updates(
		models.User{
			Name:     user.Name,
			Password: user.Password,
			StateMsg: user.StateMsg,
			Portrait: user.Portrait,
		},
	)
}

// @Summary 查询用户
// @Schemes
// @Description 根据name查询用户信息
// @Tags 用户相关
// @param name path string true "要查询的名字"
// @Accept json
// @Produce json
// @Router /find/{name} [get]
func FindUserByName(ctx *gin.Context) {
	name := ctx.Param("name")
	db.Where("name = ?", name).Find(&user_list)

	var statucCode int
	if len(user_list) == 0 {
		statucCode = http.StatusNotFound
	} else {
		statucCode = http.StatusOK
	}

	ctx.JSON(statucCode, gin.H{
		"data": user_list,
	})
}
