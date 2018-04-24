func generateToken(secretKey []byte) string {
	
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "APId2js898ilsje6272g726g072gso", 
		IssuedAt: time.Now().Unix(),
	}
	
	//Incorporar informações do usuário ao token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) 
	// token -> string. Only server knows this secret (wsitesb).
	tokenstring, err := token.SignedString(secretKey)          
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}