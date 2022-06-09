package controllers

/* import (
	"github.com/golang-jwt/jwt"
) */

/* func Middleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				SECRET_KEY := os.Getenv("DRIVERS_API_SECRET_JWT_KEY")

				return []byte(SECRET_KEY), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "userId", claims["userId"])
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	}
} */

/* type CustomClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func GetToken(user model.User) (*string, error) {
	claims := CustomClaims{
		user.Id,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := os.Getenv("DRIVERS_API_SECRET_JWT_KEY")
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		fmt.Printf("error signing token: %v", err)
		return nil, err
	}

	return &signedToken, nil
} */
