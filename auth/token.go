package auth

import (
	"log"
	"sync"
	"time"

	"cn.a2490/config"
	"cn.a2490/utils"
)

type tokenStore struct {
	reStoreMap map[string]*[]string // loginId -> token
	tokenMap   map[string]string    // token -> loginId
	expTimeMap map[string]time.Time // token -> expTime
	mu         sync.Mutex
}

func (s *tokenStore) getToken() string {
	for {
		token, _ := utils.GetUuid()
		if _, ok := s.tokenMap[token]; !ok {
			return token
		}
	}
}

func (s *tokenStore) cleanInvalidToken() func() {
	tick := time.NewTicker(time.Duration(config.Config.Token.CleanRegular) * time.Second)
	quit := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-tick.C:
				for token, expTime := range s.expTimeMap {
					if expTime.IsZero() {
						log.Printf("clean token: %s\n", token)
						s.cleanToken(token)
						continue
					}
					nowTime := time.Now()
					subS := nowTime.Sub(expTime)
					if config.Config.Token.Timeout < 0 || len(s.tokenMap) > 10000 {
						// 清理长时间不使用的token
						if subS.Seconds() > float64(config.Config.Token.CleanTimeout) {
							log.Printf("clean token: %s\n", token)
							s.cleanToken(token)
						}
					} else {
						// 清理过期token
						if subS.Seconds() > float64(config.Config.Token.Timeout) {
							log.Printf("clean token: %s\n", token)
							s.cleanToken(token)
						}
					}
				}
			case <-quit:
				tick.Stop()
			}
		}
	}()
	return func() {
		close(quit)
	}
}

func (s *tokenStore) cleanToken(token string) {
	loginId := s.tokenMap[token]
	tokens := s.reStoreMap[loginId]
	if tokens != nil && len(*tokens) > 0 {
		i := 0
		for _, t := range *tokens {
			if t != token {
				(*tokens)[i] = t
				i++
			}
		}
		*tokens = (*tokens)[:i]
	}
	delete(s.expTimeMap, token)
	delete(s.tokenMap, token)
}

var store *tokenStore

func InitTokenStore() func() {
	store = &tokenStore{
		reStoreMap: make(map[string]*[]string),
		tokenMap:   make(map[string]string),
		expTimeMap: make(map[string]time.Time),
	}
	return store.cleanInvalidToken()
}

func DoLogin(loginId string) string {
	store.mu.Lock()
	defer store.mu.Unlock()

	tokens := store.reStoreMap[loginId]
	nowTime := time.Now()
	if tokens == nil || len(*tokens) == 0 {
		if tokens == nil {
			tokens = new([]string)
		}
		token := store.getToken()
		*tokens = append(*tokens, token)
		store.reStoreMap[loginId] = tokens
		store.tokenMap[token] = loginId
		store.expTimeMap[token] = nowTime
		return token
	}
	if config.Config.Token.IsConcurrent {
		if config.Config.Token.IsShare {
			token := (*tokens)[0]
			store.expTimeMap[token] = nowTime
			return token
		}
		if config.Config.Token.Timeout >= 0 {
			for _, token := range *tokens {
				loginTime := store.expTimeMap[token]
				if loginTime.IsZero() {
					store.cleanToken(token)
					continue
				}
				subS := nowTime.Sub(loginTime)
				if subS.Seconds() > float64(config.Config.Token.Timeout) {
					store.cleanToken(token)
				}
			}
		}

		token := store.getToken()
		*tokens = append(*tokens, token)
		store.tokenMap[token] = loginId
		store.expTimeMap[token] = nowTime
		return token
	}
	oldToken := (*tokens)[0]
	delete(store.expTimeMap, oldToken)
	delete(store.tokenMap, oldToken)
	newToken := store.getToken()
	(*tokens)[0] = newToken
	store.tokenMap[newToken] = loginId
	store.expTimeMap[newToken] = nowTime
	return newToken
}

func GetLoginId(token string) string {
	if token == "" {
		return ""
	}
	store.mu.Lock()
	defer store.mu.Unlock()

	if config.Config.Token.Timeout >= 0 {
		loginTime := store.expTimeMap[token]
		if loginTime.IsZero() {
			store.cleanToken(token)
			return ""
		}
		nowTime := time.Now()
		subS := nowTime.Sub(loginTime)
		if subS.Seconds() > float64(config.Config.Token.Timeout) {
			store.cleanToken(token)
			return ""
		}
	}
	loginId := store.tokenMap[token]
	if loginId == "" {
		return ""
	}
	nowTime := time.Now()
	store.expTimeMap[token] = nowTime
	return loginId
}

func IsLogin(token string) bool {
	if token == "" {
		return false
	}
	loginId := GetLoginId(token)
	return loginId != ""
}
