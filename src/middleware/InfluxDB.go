package middleware

import (
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/Injector"
	"git.sxxfuture.net/filfi/letsfil/letsfil-job/src/configs"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxDB struct {
}

func NewInfluxDB() *InfluxDB {
	return &InfluxDB{}
}

func (this *InfluxDB) InitInfluxDB() influxdb2.Client {
	client := influxdb2.NewClient(configs.GetValue(configs.InfluxdbUrl, configs.Get().Influxdb.Url), configs.GetValue(configs.InfluxdbToken, configs.Get().Influxdb.Token))
	//get query client
	queryApi := client.QueryAPI(configs.GetValue(configs.InfluxdbOrg, configs.Get().Influxdb.Org))
	writeApi := client.WriteAPI(configs.GetValue(configs.InfluxdbOrg, configs.Get().Influxdb.Org), configs.Get().Influxdb.Bucket)
	Injector.BeanFactory.Set(queryApi)
	Injector.BeanFactory.Set(writeApi)
	client.Options()
	return client
}
