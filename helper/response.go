package helper

import (
	"net/http"

	"google.golang.org/grpc/codes"
	grpcCode "google.golang.org/grpc/codes"
)

const (
	// Success status
	Success codes.Code = 200
	//SuccessCreated status
	SuccessCreated codes.Code = 201
	// SuccessNoContent status
	SuccessNoContent codes.Code = 204
	// InvalidArgument status
	InvalidArgument codes.Code = 400
	// Unauthorized status
	Unauthorized codes.Code = 401
	// Forbidden status
	Forbidden codes.Code = 403
	// NotFound status
	NotFound codes.Code = 404
	// Cancelled status
	Cancelled codes.Code = 405
	// RequestTimeout status
	RequestTimeout codes.Code = 408

	// InactiveAccount status
	InactiveAccount codes.Code = 410
	// InvalidToken status
	InvalidToken codes.Code = 411
	// InvalidAPIKey status
	InvalidAPIKey codes.Code = 412
	// InvalidSession status
	InvalidSession codes.Code = 413
	// ResourceExhausted status
	ResourceExhausted codes.Code = 414

	// InvalidTransaction status
	InvalidTransaction codes.Code = 430
	// DuplicateTransaction status
	DuplicateTransaction codes.Code = 431

	// InternalError status
	InternalError codes.Code = 500
	// ProcessingError status
	ProcessingError codes.Code = 502
)

// StatusMessage represent string message for code
var StatusMessage = map[codes.Code]string{
	Success:              "Berhasil",
	SuccessCreated:       "Berhasil, Data Tersimpan",
	SuccessNoContent:     "Berhasil, Tanpa Konten",
	InvalidArgument:      "Parameter tidak valid",
	Unauthorized:         "Password kamu salah",
	Forbidden:            "Akses tidak dibolehkan atau kamu tidak memiliki akses",
	NotFound:             "Data tidak ditemukan",
	Cancelled:            "Permintaan dibatalkan",
	RequestTimeout:       "Permintaan melebihi batas waktu",
	InactiveAccount:      "Akun tidak aktif",
	InvalidToken:         "Akses tidak valid karena token tidak cocok atau login sudah kadaluarsa",
	InvalidAPIKey:        "API key tidak valid",
	InvalidSession:       "Sesi tidak valid atau sudah berakhir",
	ResourceExhausted:    "Sudah mencapai batas limit",
	InvalidTransaction:   "Data transaksi tidak valid atau tidak sesuai",
	DuplicateTransaction: "Data transaksi duplikat",
	InternalError:        "Error dari server",
}

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

// Response struct
type Response struct {
	Code       grpcCode.Code          `json:"code"`
	Message    string                 `json:"message,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
	Pagination *Pagination            `json:"pagination,omitempty"`
	Errors     []string               `json:"errors,omitempty"`
	TraceID    string                 `json:"trace_id,omitempty"`
	Header     map[string]interface{} `json:"-"`
	HttpCode   int32                  `json:"-"`
}

// Pagination struct
type Pagination struct {
	CurrentPage int32  `json:"current_page"`
	PageSize    int32  `json:"page_size"`
	TotalPage   int32  `json:"total_page"`
	TotalResult int32  `json:"total_result"`
	Next        string `json:"next,omitempty"`
	Prev        string `json:"prev,omitempty"`
}
