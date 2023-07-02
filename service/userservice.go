package service

import (
	"Gin+WebSocket/models"
	"Gin+WebSocket/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"node",""message}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "成功获取用户列表",
		"data":    data,
	})
}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param phone query string false "手机号"
// @param email query string false "电子邮箱"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code",""message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	password := c.Query("password")
	repassword := c.Query("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	user1 := models.FindUserByName(user.Name)
	if user1.Name != "" {
		c.JSON(-1, gin.H{
			"message": "用户名已注册！",
		})
		return
	}

	user2 := models.FindUserByPhone(user.Phone)
	if user2.Phone != "" {
		c.JSON(-1, gin.H{
			"message": "手机号已注册！",
		})
		return
	}

	user3 := models.FindUserByEmail(user.Email)
	if user3.Email != "" {
		c.JSON(-1, gin.H{
			"message": "电子邮箱已注册！",
		})
		return
	}

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt

	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "新增用户成功！！！",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// @Summary 用户登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"node",""message}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}

	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, // 0 成功   -1失败
			"message": "该用户不存在！！！",
			"data":    data,
		})
		return
	}

	flag := utils.VaildPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, // 0 成功   -1失败
			"message": "密码不正确",
			"data":    data,
		})
		return
	}

	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(http.StatusOK, gin.H{
		"code":    0, // 0 成功   -1失败
		"message": "登录成功",
		"data":    data,
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code",""message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": "删除用户成功！！！",
		"data":    user,
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code",""message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改参数不匹配！",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "修改用户成功！！！",
			"data":    user,
		})
	}

}
