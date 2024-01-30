package server

import (
	"log"

	"swc-api-back.com/cmd/swc-api-back/handlers/web/publication"
	publicationRepo "swc-api-back.com/internal/repository/publication"
	publicationSrv "swc-api-back.com/internal/service/publication"
)

func resolvePublicationHandler() *publication.Handler {
	h, err := publication.New(
		resolvePublicationService(),
	)
	if err != nil {
		panicHandler(err)
	}
	return h

}

func resolvePublicationService() publicationSrv.Publication {
	s, err := publicationSrv.New(
		resolvePublicationRepository(),
	)
	if err != nil {
		panicHandler(err)
	}
	return s
}

func resolvePublicationRepository() publicationRepo.Publication {
	r, err := publicationRepo.New()
	if err != nil {
		panicHandler(err)
	}
	return r
}

func panicHandler(err error) {
	log.Default().Panicf("(panic) error on initialize app dependencias: %v", err)
}
