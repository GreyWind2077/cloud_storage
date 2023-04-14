package controller

import (
	"cloud_storage/config"
	"cloud_storage/dao/redis"
	"cloud_storage/service"
	"cloud_storage/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// 检查是否登录
func CheckLogin(c *gin.Context) {
	token, err := c.Cookie("Token")
	if err != nil {
		log.Println("cookie", err.Error())
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	openId, err := redis.GetKey(token)
	if err != nil {
		log.Println("Get Redis Err:", err.Error())
		c.Redirect(http.StatusFound, "/")
		c.Abort()
	}

	user := service.FindUser(openId)

	if user.Id == 0 {
		//校验失败 返回登录页面
		c.Redirect(http.StatusFound, "/")
	} else {
		//校验成功 继续执行
		c.Set("openId", openId)
		c.Next()
	}
}

// QQ登录成功 处理登录
func LoginSucceed(userInfo, openId string, c *gin.Context) {
	var qUserInfo QUserInfo
	//将数据转为结构体
	if err := json.Unmarshal([]byte(userInfo), &qUserInfo); err != nil {
		fmt.Println("转换json失败", err.Error())
		return
	}

	//创建一个token
	hashToken := utils.Md5Crypt("token" + string(time.Now().Unix()) + openId)
	//存入redis
	if err := redis.SetKey(hashToken, openId, 24*3600); err != nil {
		fmt.Println("Redis Set Err:", err.Error())
		return
	}
	//设置cookie
	c.SetCookie("Token", hashToken, 3600*24, "/", "pyxgo.cn", false, true)

	if ok := service.FindUser(openId); ok.Id != 0 { //用户存在直接登录
		//登录成功重定向到首页
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	} else {
		service.CreateUser(openId, qUserInfo.Nickname, qUserInfo.FigureUrlQQ)
		//登录成功重定向到首页
		c.Redirect(http.StatusMovedPermanently, "/cloud/index")
	}
}

// 退出登录
func Logout(c *gin.Context) {
	token, err := c.Cookie("Token")
	if err != nil {
		log.Println("cookie", err.Error())
	}

	if err := redis.DelKey(token); err != nil {
		log.Println("Del Redis Err:", err.Error())
	}

	c.SetCookie("Token", "", 0, "/", "pyxgo.cn", false, false)
	c.Redirect(http.StatusFound, "/")
}

type PrivateInfo struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
}

// 登录成功获取的QQ用户信息
type QUserInfo struct {
	Nickname    string
	FigureUrlQQ string `json:"figureurl_qq"`
}

// 获取access_token
func QQLogin(c *gin.Context) {
	code := c.Query("code")

	loginUrl := "https://graph.qq.com/oauth2.0/token?grant_type=authorization_code&client_id=" + config.Cfg.QQ.App_id + "&client_secret=" + config.Cfg.QQ.App_key + "&redirect_uri=" + config.Cfg.QQ.Redirect_url + "&code=" + code

	//访问qq的api
	response, err := http.Get(loginUrl)
	if err != nil {
		log.Println("请求错误", err.Error())
		return
	}
	defer response.Body.Close()

	bs, _ := ioutil.ReadAll(response.Body)
	body := string(bs)
	resultMap := ConvertToMap(body)

	info := &PrivateInfo{}
	info.AccessToken = resultMap["access_token"]
	info.RefreshToken = resultMap["refresh_token"]
	info.ExpiresIn = resultMap["expires_in"]

	GetOpenId(info, c)
}

// 获取QQ openId
func GetOpenId(info *PrivateInfo, c *gin.Context) {
	resp, err := http.Get(fmt.Sprintf("%s?access_token=%s", "https://graph.qq.com/oauth2.0/me", info.AccessToken))
	if err != nil {
		fmt.Println("GetOpenId Err", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	body := string(bs)
	info.OpenId = body[45:77]

	GetUserInfo(info, c)
}

// 获取QQ用户信息
func GetUserInfo(info *PrivateInfo, c *gin.Context) {
	params := url.Values{}
	params.Add("access_token", info.AccessToken)
	params.Add("openid", info.OpenId)
	params.Add("oauth_consumer_key", config.Cfg.QQ.App_id)

	uri := fmt.Sprintf("https://graph.qq.com/user/get_user_info?%s", params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("GetUserInfo Err:", err.Error())
		return
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)

	LoginSucceed(string(bs), info.OpenId, c)
}

// 装换类型
func ConvertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}
