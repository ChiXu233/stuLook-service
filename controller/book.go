package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"stuLook-service/model"
	"stuLook-service/service"
	"stuLook-service/utils"

	"net/http"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.CreateBook(book))
}

func DeleteBook(c *gin.Context) {
	if _id, err := primitive.ObjectIDFromHex(c.Query("_id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.DeleteBook(_id))
	}
}

func UpdateBook(c *gin.Context) {
	var book model.BookUpdate
	var _id, err = primitive.ObjectIDFromHex(c.Query("_id"))
	fmt.Println(err)
	if err1 := c.ShouldBind(&book); err1 != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err1.Error()))
		return
	}
	fmt.Println(book)
	book.Id = _id
	c.JSON(http.StatusOK, service.UpdateBook(book))
}

func GetBook(c *gin.Context) {
	// 查询所有图书
	conditions := make(map[string]interface{})
	if method := c.Query("method"); method != "" {
		conditions["method"] = method
	}
	pageSize, err := strconv.ParseInt(c.DefaultQuery("pageSize", "0"), 10, 64)
	//获取 id 参数, 跟 Query 函数的区别是，可以通过第二个参数设置默认值。
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess(err.Error(), nil))
		return
	}
	currPage, err := strconv.ParseInt(c.DefaultQuery("currPage", "1"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, service.GetBook(conditions, pageSize, currPage))
}
func GetById(c *gin.Context) {
	//根据ID查询
	if _id, err := primitive.ObjectIDFromHex(c.Query("_id")); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	} else {
		fmt.Println(_id)
		c.JSON(http.StatusOK, service.GetBookById(_id))
	}
}

func GetByName(c *gin.Context) {
	//根据图书名模糊查找
	var book model.Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.GetByName(book.Name))
	}
}

func GetByAuthor(c *gin.Context) {
	//根据作者模糊查找
	fmt.Println("3333333")
	var bookAuthor model.Book
	if err := c.Bind(&bookAuthor); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数传递错误", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, service.GetByAuthor(bookAuthor.Author))
	}
}
