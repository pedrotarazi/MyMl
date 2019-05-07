package myml

import (
	"github.com/gin-gonic/gin"
	myml2 "github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/services/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"net/http"
	"strconv"
	"sync"
)

const (
	paramUserID = "userID"
)

func GetMyML(context *gin.Context) {
	userID := context.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)

	if err != nil {
		apiErr := &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		context.JSON(apiErr.Status, apiErr)
		return
	}

	var mymlChannel myml2.ChannelError

	// use waitgroup for execute all go routines
	var wg sync.WaitGroup
	ch := make(chan myml2.ChannelError)

	user, apiErr := myml.GetUserFromAPI(id)
	if apiErr != nil {
		context.JSON(apiErr.Status, apiErr)
		return
	}
	mymlChannel.MyMl.User = *user

	wg.Add(2)
	go myml.GetSiteFromAPI(ch, user.SiteID, &wg)
	go myml.GetCountryFromAPI(ch, user.CountryID, &wg)

	for i := 0; i < 2; i++ {
		mymlChannelData := <-ch
		if mymlChannelData.Error.Status != 500 {
			if mymlChannelData.MyMl.Site.ID != "" {
				mymlChannel.MyMl.Site = mymlChannelData.MyMl.Site
			}
			if mymlChannelData.MyMl.Country.ID != "" {
				mymlChannel.MyMl.Country = mymlChannelData.MyMl.Country
			}
		}
	}
	wg.Wait()
	context.JSON(200, mymlChannel.MyMl)
}
