package util

//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"golang-boilerplate/ent"
//	"golang-boilerplate/ent/group"
//	"time"
//
//	"github.com/google/uuid"
//	"github.com/pkg/errors"
//)
//
//const (
//	AccessTokenDuration  = 30 * time.Minute
//	RefreshTokenDuration = 30 * 24 * time.Hour // 7 days
//
//)
//
//func GenerateAccessToken(user ent.User) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//	claims["id"] = user.ID
//	claims["username"] = user.Username
//	claims["exp"] = time.Now().Add(AccessTokenDuration).Unix()
//	claims["role"] = user.Edges.Group.Type.String()
//	tokenString, err := token.SignedString([]byte("your-secret-key"))
//	if err != nil {
//		return "", err
//	}
//	//err = sendMailWithTemplate()
//	return tokenString, nil
//}
//
//func GenerateRefreshToken(user ent.User) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := token.Claims.(jwt.MapClaims)
//	claims["id"] = user.ID
//	claims["username"] = user.Username
//	claims["exp"] = time.Now().Add(RefreshTokenDuration).Unix()
//	claims["role"] = user.Edges.Group.Type.String()
//
//	tokenString, err := token.SignedString([]byte("your-secret-key"))
//	if err != nil {
//		return "", err
//	}
//	return tokenString, nil
//}
//
//func ValidateToken(tokenString string) (*ent.User, error) {
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.Errorf("Unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte("your-secret-key"), nil
//	})
//
//	if err != nil {
//		return nil, err
//	}
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		id, err := uuid.Parse(claims["id"].(string))
//		username, usernameOk := claims["username"].(string)
//		expirationTime, expOk := claims["exp"].(float64)
//		role, roleOk := claims["role"].(string)
//
//		if !expOk {
//			return nil, errors.Errorf("Token does not contain an expiration time")
//		}
//
//		// Convert the expiration time to a Go time.Time.
//		expTime := time.Unix(int64(expirationTime), 0)
//
//		// Check if the token has expired.
//		if time.Now().After(expTime) {
//			return nil, errors.Errorf("Token has expired")
//		}
//		if err == nil && usernameOk && roleOk {
//			group := &ent.Group{
//				Type: group.Type(role),
//			}
//			user := &ent.User{
//				ID:       id,
//				Username: username,
//				Edges:    ent.UserEdges{Group: group},
//			}
//			return user, nil
//		}
//	}
//
//	return nil, errors.Errorf("Token validation failed")
//}
//
////func getUserFromToken(ctx context.Context) *ent.User {
////	// Extract the token from the context.
////	token := ctx.Value("token")
////	if token == nil {
////		return nil
////	}
////
////	// Convert the token to a string.
////	tokenString, ok := token.(string)
////	if !ok {
////		return nil
////	}
////
////	// Validate and decode the token.
////	user, err := ValidateToken(tokenString)
////	if err != nil {
////		// Token is invalid or expired.
////		return nil
////	}
////
////	return user
////}
