package main

import (
	"github.com/gorilla/mux"
)

func apiRouterInit(router *mux.Router) error {
	router.HandleFunc("/api/owner/new", ownerNewHandler).Methods("POST")
	router.HandleFunc("/api/owner/confirm-hex", ownerConfirmHexHandler).Methods("GET")
	router.HandleFunc("/api/owner/login", ownerLoginHandler).Methods("POST")
	router.HandleFunc("/api/owner/send-reset-hex", ownerSendResetHexHandler).Methods("POST")
	router.HandleFunc("/api/owner/reset-password", ownerResetPasswordHandler).Methods("POST")
	router.HandleFunc("/api/owner/self", ownerSelfHandler).Methods("POST")

	router.HandleFunc("/api/domain/new", domainNewHandler).Methods("POST")
	router.HandleFunc("/api/domain/delete", domainDeleteHandler).Methods("POST")
	router.HandleFunc("/api/domain/list", domainListHandler).Methods("POST")
	router.HandleFunc("/api/domain/update", domainUpdateHandler).Methods("POST")
	router.HandleFunc("/api/domain/moderator/new", domainModeratorNewHandler).Methods("POST")
	router.HandleFunc("/api/domain/moderator/delete", domainModeratorDeleteHandler).Methods("POST")
	router.HandleFunc("/api/domain/statistics", domainStatisticsHandler).Methods("POST")
	router.HandleFunc("/api/domain/import/disqus", domainImportDisqusHandler).Methods("POST")

	router.HandleFunc("/api/commenter/token/new", commenterTokenNewHandler).Methods("GET")
	router.HandleFunc("/api/commenter/new", commenterNewHandler).Methods("POST")
	router.HandleFunc("/api/commenter/login", commenterLoginHandler).Methods("POST")
	router.HandleFunc("/api/commenter/self", commenterSelfHandler).Methods("POST")

	router.HandleFunc("/api/oauth/google/redirect", googleRedirectHandler).Methods("GET")
	router.HandleFunc("/api/oauth/google/callback", googleCallbackHandler).Methods("GET")

	router.HandleFunc("/api/comment/new", commentNewHandler).Methods("POST")
	router.HandleFunc("/api/comment/list", commentListHandler).Methods("POST")
	router.HandleFunc("/api/comment/vote", commentVoteHandler).Methods("POST")
	router.HandleFunc("/api/comment/approve", commentApproveHandler).Methods("POST")
	router.HandleFunc("/api/comment/delete", commentDeleteHandler).Methods("POST")

	return nil
}
