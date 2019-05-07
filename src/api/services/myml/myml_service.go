package myml

import (
	"fmt"
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"sync"
)

func GetUserFromAPI(id int64) (*myml.User, *apierrors.ApiError) {
	user := &myml.User{
		ID: id,
	}
	//fmt.Println("Id en Services:", user.ID)
	if apiErr := user.Get(); apiErr != nil {
		fmt.Println(apiErr)
		return nil, apiErr
	}
	return user, nil
}

func GetSiteFromAPI(ch chan myml.ChannelError, siteID string, wg *sync.WaitGroup) {
	site := &myml.Site{
		ID: siteID,
	}
	apiErr := site.Get()
	if apiErr != nil {
		fmt.Println(apiErr)
		ch <- myml.ChannelError{
			MyMl: myml.MyMl{},
			Error: apierrors.ApiError{
				Message: "Error en Site",
				Status:  500,
			},
		}
		return
	}
	mymlSite := myml.MyMl{
		Site: *site,
	}
	ch <- myml.ChannelError{
		MyMl:  mymlSite,
		Error: apierrors.ApiError{},
	}
	wg.Done()
}

func GetCountryFromAPI(ch chan myml.ChannelError, countryID string, wg *sync.WaitGroup) {
	country := &myml.Country{
		ID: countryID,
	}
	apiErr := country.Get()
	if apiErr != nil {
		fmt.Println(apiErr)
		ch <- myml.ChannelError{
			MyMl: myml.MyMl{},
			Error: apierrors.ApiError{
				Message: "Error en Country",
				Status:  500,
			},
		}
		return
	}
	mymlCountry := myml.MyMl{
		Country: *country,
	}
	ch <- myml.ChannelError{
		MyMl:  mymlCountry,
		Error: apierrors.ApiError{},
	}
	wg.Done()
}
