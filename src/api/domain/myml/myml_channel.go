package myml

import "github.com/mercadolibre/myml/src/api/utils/apierrors"

type ChannelError struct {
	MyMl  MyMl
	Error apierrors.ApiError
}
