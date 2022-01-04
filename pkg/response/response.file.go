package response

import (
	"io"
	"net/http"
)

type FileWriter interface {
	Write(w io.Writer) error
}

type FileResponse struct {
	BasicResponse
	file FileWriter
	filename string
}
// type JSONBody struct {
// 	*Error
// 	Message string      `json:"message,omitempty"`
// 	Data    interface{} `json:"data,omitempty"`
// }

func NewFileResponse() *FileResponse {
	return &FileResponse{
		BasicResponse: BasicResponse{
			ContentType: FileContentType,
			StatusCode:  http.StatusOK,
		},
	}
}

func (r *FileResponse) SetFileWriter(data FileWriter, filename string) *FileResponse {
	r.file = data
	r.filename = filename
	return r
}


func (r *FileResponse) WriteResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.ContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+ r.filename)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	r.file.Write(w)
}
