package codes

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

// HTTPStatusFromCode return HTTP Status for each code
func HTTPStatusFromCode(c codes.Code) int {
	switch c {
	case Success:
		return http.StatusOK
	case SuccessCreated:
		return http.StatusCreated
	case SuccessNoContent:
		return http.StatusOK
	case InvalidArgument:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case Cancelled:
		return http.StatusRequestTimeout
	case RequestTimeout:
		return http.StatusRequestTimeout
	case InactiveAccount:
		return http.StatusUnauthorized
	case InvalidToken:
		return http.StatusUnauthorized
	case InvalidAPIKey:
		return http.StatusUnauthorized
	case InvalidSession:
		return http.StatusUnauthorized
	case ResourceExhausted:
		return http.StatusTooManyRequests
	case InvalidSubdomain:
		return http.StatusNotFound
	case InactiveSubdomain:
		return http.StatusNotFound
	case SuspendedSubdomain:
		return http.StatusForbidden
	case InvalidTransaction:
		return http.StatusBadRequest
	case DuplicateTransaction:
		return http.StatusConflict
	case ProcessingError:
		return http.StatusInternalServerError
	case InternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
