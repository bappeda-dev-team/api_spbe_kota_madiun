package service

import (
	"api_spbe_kota_madiun/model/web"
	"context"
)

type DomainSPBEService interface {
	Insert(ctx context.Context, request web.DomainSPBECreateRequest) web.DomainSPBEResponse
	Update(ctx context.Context, request web.DomainSPBEUpdateRequest) web.DomainSPBEResponse
	Delete(ctx context.Context, domainSPBEId int)
	FindAll(ctx context.Context) []web.DomainSPBEResponse
	FindById(ctx context.Context, domainSPBEId int) web.DomainSPBEResponse
}
