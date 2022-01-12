package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"irisStudy/common"
	"irisStudy/rabbitmq"
	"irisStudy/token"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var tokenMaker token.Maker

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func init() {
	var err error
	tokenMaker, err = token.NewJWTMaker("12345678901234567890123456789012")
	if err != nil {
		log.Fatalf("failed to init tokenMaker, err: %v", err)
	}
}

// 用来存放控制信息
type AccessControl struct {
	// 用来存放用户想要存放的信息
	sourceArray map[string]interface{}
	sync.RWMutex
}

// 创建全局变量
var accessControl = &AccessControl{sourceArray: make(map[string]interface{})}

// 获取制定数据
func (a *AccessControl) GetNewRecord(uid string) interface{} {
	a.RLock()
	defer a.RUnlock()

	data := a.sourceArray[uid]
	return data
}

// 设置记录
func (a *AccessControl) SetNewRecord(uid string) {
	a.RLock()
	defer a.RUnlock()
	a.sourceArray[uid] = "hello imocc"
}

func (a *AccessControl) GetDistributedRight(request *http.Request) bool {
	// 获取用户uid
	queryData, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		log.Println(err)
		return false
	}

	if len(queryData) < 1 || len(queryData["uid"]) < 1 {
		err = errors.New("query params is not found")
		log.Println(err)
		return false
	}
	uid := queryData["uid"][0]

	// 采用一致性hash算法，根据用户ID，判断获取具体机器
	hostRequest, err := hashConsistent.Get(uid)
	if err != nil {
		return false
	}

	// 判断是否为本机
	if hostRequest == localHost {
		// 执行本机数据读取和校验
		return a.GetDataFromMap(uid)
	} else {
		// 不是本机充当代理访问数据返回结果
		return GetDataFromOtherMap(hostRequest, request)
	}
}

// 获取本机map，并且处理业务逻辑，返回的结果类型为bool类型
func (a *AccessControl) GetDataFromMap(uid string) (isOk bool) {
	//data := a.GetNewRecord(uid)
	//if data != nil {
	//	return false
	//}
	return true
}

// GetDataFromOtherMap 获取其它节点处理结果
func GetDataFromOtherMap(host string, request *http.Request) bool {
	hostUrl := fmt.Sprintf("http://%s:%s/checkRight", host, port)
	response, body, err := GetCurl(hostUrl, request)
	if err != nil {
		return false
	}

	// 判断状态
	if response.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}

// 模拟请求
func GetCurl(hostUrl string, request *http.Request) (response *http.Response, body []byte, err error) {
	// 获取uid
	queryData, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		log.Println(err)
		return
	}
	if len(queryData) < 1 || len(queryData["uid"]) < 1 {
		err = errors.New("query params is not found")
		log.Println(err)
		return
	}
	uid := queryData["uid"][0]

	// 获取token值
	authorizationHeader := request.Header.Get(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		err = errors.New("authorization header is not provided")
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err = errors.New("invalid authorization header format")
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err = fmt.Errorf("unsupported authorization type %s", authorizationType)
		return
	}

	accessToken := fields[1]

	hostUrl = fmt.Sprintf("%s?uid=%s", hostUrl, uid)

	// 模拟接口访问
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, hostUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add(authorizationHeaderKey, accessToken)

	// 获取返回结果
	response, err = client.Do(req)
	defer response.Body.Close()

	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(response.Body)
	return
}


// Auth 统一验证拦截器，每个接口都需要提前验证
func Auth(response http.ResponseWriter, request *http.Request) error {
	log.Println("执行验证操作")
	// 添加权限验证
	err := CheckUserInfo(request)
	if err != nil {
		return err
	}
	return nil
}

// Check 执行正常业务逻辑
func Check(response http.ResponseWriter, request *http.Request) {
	log.Println("执行check 正常的业务逻辑...")
	// 将传递过来的数据放到rabbitMQ中
	// 获取到uid
	queryData, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		_, _ = response.Write([]byte("query params wrong"))
	}
	if len(queryData) < 1 || len(queryData["uid"]) < 1 {
		err = errors.New("query params is not found")
		_, _ = response.Write([]byte(err.Error()))
	}
	uid := queryData["uid"][0]
	fmt.Println("获取到uid值: ", uid)

	// 获取到productId
	productId := queryData["productId"][0]
	fmt.Println("获取到productId值: ", productId)

	// 1.分布式权限验证
	right := accessControl.GetDistributedRight(request)
	if right == false {
		_, _ = response.Write([]byte("false"))
		return
	}



	// 2.获取数量控制权限，防止出现超卖现象
	//hostUrl := fmt.Sprintf("http://%s:%s/getOne", "127.0.0.1", port)
	//responseValidate, validateBody, err := GetCurl(hostUrl, request)
	//if err != nil {
	//	_, _ = response.Write([]byte("false"))
	//	return
	//}

	// 判断数量控制请求
	//if responseValidate.StatusCode == 200 {
	//	if string(validateBody) == "true" {
	//		// 整合下单
	//		// 1.获取商品ID
	//		productID, err := strconv.ParseInt(productId, 10, 64)
	//		if err != nil {
	//			_, _ = response.Write([]byte("false"))
	//			return
	//		}
	//
	//		// 2.获取用户ID
	//		userID, err := strconv.ParseInt(uid, 10, 64)
	//		if err != nil {
	//			_, _ = response.Write([]byte("false"))
	//			return
	//		}
	//
	//		// 3.创建消息体
	//		message := datamodels.NewMessage(userID, productID)
	//		// 类型转换
	//		byteMessage, err := json.Marshal(message)
	//		if err != nil {
	//			_, _ = response.Write([]byte("false"))
	//			return
	//		}
	//
	//		// 4.生产消息
	//		err = rabbitMqValidate.PublishSimple(string(byteMessage))
	//		if err != nil {
	//			_, _ = response.Write([]byte("false"))
	//			return
	//		}
	//
	//		_, _ = response.Write([]byte("ture"))
	//		return
	//
	//	}
	//}
	//_, _ = response.Write([]byte("false"))
	return
}

// 用户身份验证函数
func CheckUserInfo(request *http.Request) error {
	fmt.Println("CheckUserInfo: ", request.URL)
	// 获取到由url传入的参数uid
	queryData, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		return err
	}
	if len(queryData) < 1 || len(queryData["uid"]) < 1 {
		err = errors.New("query params is not found")
		return err
	}
	uid := queryData["uid"][0]

	// 获取用户加密串
	authorizationHeader := request.Header.Get(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		err = errors.New("authorization header is not provided")
		return err
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err = errors.New("invalid authorization header format")
		return err
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != authorizationTypeBearer {
		err = fmt.Errorf("unsupported authorization type %s", authorizationType)
		return err
	}

	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)

	// 判断是否有解析都token值，若无则是token过期
	if payload == nil {
		return errors.New("auth failed, login again")
	}

	if uid == payload.UserId {
		return nil
	}
	// 非当前用户登录时
	return errors.New("auth failed")
}

var hostArray = []string{"127.0.0.1", "127.0.0.1"}

var localHost = "127.0.0.1"

var port = "8083"

var hashConsistent *common.Consistent

//rabbitmq
var rabbitMqValidate *rabbitmq.RabbitMQ

func main() {
	// 负载均衡器设置
	// 采用一致性hash算法
	hashConsistent = common.NewConsistent()
	// 采用一致性hash算法，添加节点
	for _, v := range hostArray {
		fmt.Println("add: ", v)
		hashConsistent.Add(v)
	}

	//localIp, err := common.GetIntranceIp()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//localHost = localIp
	//fmt.Println(localHost)

	rabbitMqValidate =rabbitmq.NewRabbitMQSimple("imoocProduct")
	defer rabbitMqValidate.Destory()


	// 过滤器
	filter := common.NewFilter()
	// 注册拦截器
	filter.RegisterFilterUri("/check", Auth)
	// 启动服务
	http.HandleFunc("/check", filter.Handle(Check))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal("failed to start http server")
	}
}
