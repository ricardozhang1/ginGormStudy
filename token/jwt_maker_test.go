package token

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateToken(t *testing.T) {
	tokenMaker, err := NewJWTMaker("12345678901234567890123456789012")
	require.NoError(t, err)
	token, err := tokenMaker.CreateToken("zhangsan", "78995", 30*time.Minute)
	require.NoError(t, err)
	fmt.Println(token)
}


