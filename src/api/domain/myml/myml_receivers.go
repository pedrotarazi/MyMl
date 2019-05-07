package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

var (
	siteID = ""
)

func (user *User) Get() *apierrors.ApiError {
	if user.ID == 0 {
		return &apierrors.ApiError{
			Message: "ID is Empty",
			Status:  http.StatusInternalServerError,
		}
	}

	url := fmt.Sprintf("%s%d", "https://api.mercadolibre.com/users/", user.ID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return nil
}

func (site *Site) Get() *apierrors.ApiError {
	if site.ID == "" {
		return &apierrors.ApiError{
			Message: "ID is not valid",
			Status:  http.StatusInternalServerError,
		}
	}
	siteID = site.ID
	url := fmt.Sprintf("%s%s", "https://api.mercadolibre.com/sites/", site.ID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	err = json.Unmarshal([]byte(data), &site)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}

func (country *Country) Get() *apierrors.ApiError {
	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "ID is not valid",
			Status:  http.StatusInternalServerError,
		}
	}
	url := fmt.Sprintf("%s%s", "https://api.mercadolibre.com/countries/", country.ID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	err = json.Unmarshal([]byte(data), &country)
	if err != nil {
		fmt.Println(err)
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil
}
