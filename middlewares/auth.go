package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/spf13/viper"

	"github.com/your-org/medical-records/utils"
)

// Authenticate middleware checks if the request contains a valid JWT token
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid authorization header format")
			return
		}

		token := bearerToken[1]

		claims, err := utils.ValidateToken(token, viper.GetString("jwt.secret"))
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims.Subject)
		r = r.WithContext(ctx)

		// Call the next middleware/handler in the chain
		next(w, r)
	}
}

// Authorize middleware checks if the request is authorized by checking if the peer
// certificate contains the required attributes for the specific role
func Authorize(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stub, ok := r.Context().Value("stub").(*shim.ChaincodeStub)
		if !ok {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve chaincode stub from context")
			return
		}

		creatorBytes, err := stub.GetCreator()
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve creator from transaction context")
			return
		}

		creator, err := utils.ParseX509Certificate(creatorBytes)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to parse creator certificate")
			return
		}

		if !creator.HasAttribute("role", role) {
			utils.RespondWithError(w, http.StatusUnauthorized, "Unauthorized request")
			return
		}

		// Call the next middleware/handler in the chain
		next(w, r)
	}
}
