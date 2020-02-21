package service_products

import (
	"time"
	"log"
	"net/http"
	"errors"
	
	"github.com/sony/gobreaker"
	"github.com/nahwinrajan/resilient-go/shops-service/internal/config"
)
// ServiceProducts structure holding methods for available datas from Products-Service
type ServiceProducts struct {
	client *http.Client
	conf config.ServiceProducts
	cb *gobreaker.CircuitBreaker
}


// New return new instance of ServiceProducts
func New(conf config.ServiceProducts, hclient *http.Client) *ServiceProducts {
	cbSetting := gobreaker.Settings{
		Name:        conf.CircuitBreaker.Name,
		MaxRequests: uint32(conf.CircuitBreaker.MaxRequest),
		Interval:    time.Duration(conf.CircuitBreaker.Interval) * time.Second,
		Timeout:     time.Duration(conf.CircuitBreaker.TimeOut) * time.Second,
		OnStateChange: func(name string, from, to gobreaker.State) {
			if to == gobreaker.StateOpen && from == gobreaker.StateClosed {
				// in real word sample we will do slack notification here
				log.Printf("[circuit-breaker][cb-%s] is open, database is unavailable", name)
			}

			if from == gobreaker.StateHalfOpen && to == gobreaker.StateClosed {
				// in real word sample we will do slack notification here
				log.Printf("[circuit-breaker][cb-%s] is closed, database back to normal", name)
			}
		},
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > uint32(conf.CircuitBreaker.ConsecutiveFailures)
		},
	}

	cbCli := gobreaker.NewCircuitBreaker(cbSetting)
	
	return &ServiceProducts {
		client: hclient,
		conf:	conf,
		cb: cbCli,
	}
}

// do centralized function 
func (sp *ServiceProducts) do(req *http.Request) (*http.Response, error) {
	respInterface, err := sp.cb.Execute(func() (interface{}, error){
		return sp.client.Do(req)
	})

	if (err == ErrOpenState || err == ErrTooManyRequests) && respInterface == nil {
		// send metrics
		return nil, errors.New("service is unavailable")
	}

	if respInterface == nil {
		return nil, err
	}

	return respInterface.(*http.Response), nil
}