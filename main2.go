	// vamos verificar usando o codigo secreto
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("wsitesb"), nil
	})
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if token.Valid == true {
		routers.Routers()
	}