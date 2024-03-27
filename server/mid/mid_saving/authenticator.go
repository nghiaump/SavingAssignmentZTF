package main

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
	"time"
)

const SecretKey = `NguyenDaiNghia`

// Claims struct để chứa thông tin JWT payload
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateTokenString(req *pb.LoginRequest) (string, error) {
	// Tạo Claims chứa thông tin payload cho JWT
	log.Printf("CreateTokenString() %v", req)
	claims := Claims{
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token hết hạn sau 1 ngày
		},
	}

	// Tạo JWT từ Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký JWT bằng key bí mật và chuyển thành chuỗi
	tokenString, err := token.SignedString([]byte(SecretKey))
	log.Printf("Token string: %v", tokenString)
	return tokenString, err
}

func GetTokenFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	token := md.Get("authorization")
	if len(token) == 0 {
		return ""
	}
	return strings.TrimPrefix(token[0], "Bearer ")
}

func ValidateJWTToken(tokenString string) (*Claims, error) {
	// Parse JWT token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil // Key bí mật để giải mã token
	})
	if err != nil {
		return nil, err
	}

	// Kiểm tra token có hợp lệ không và lấy thông tin payload
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Invalid JWT token")
	}
}
