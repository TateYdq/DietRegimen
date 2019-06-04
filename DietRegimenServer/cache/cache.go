package cache
import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)
var(
	Client *redis.Client
)
func setNx(){

}


func Init(addr string,password string ,db int)(error){
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,  // use default DB
	})
	Client = client
	_, err := client.Ping().Result()
	if err != nil{
		return err
	}
	return err
}



func Set(key string,value interface{},timeout time.Duration)(error){
	err := Client.Set( key, value, timeout).Err()
	if err != nil {
		logrus.WithError(err).Errorf("redis set failed. key:%v,value:%v",key,value)
	}
	return err
}

func Get(key string)(value string,err error){
	value,err = Client.Get(key).Result()
	if err != nil {
		logrus.WithError(err).Errorf("redis get failed. key:%v",key)
	}
	logrus.Infof("redis get key:%v,value:%v",key,value)
	return value,err
}
func GetInt(key string)(value int,err error){
	value,err = Client.Get(key).Int()
	if err != nil {
		logrus.WithError(err).Errorf("redis get failed. key:%v",key)
	}
	logrus.Infof("redis get key:%v,value:%v",key,value)
	return value,err
}