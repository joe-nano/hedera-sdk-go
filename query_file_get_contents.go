package hedera

// #include "hedera.h"
import "C"

type QueryFileGetContents struct {
	query
}

func newQueryFileGetContents(client Client, id FileID) QueryFileGetContents {
	return QueryFileGetContents{
		query{C.hedera_query__file_get_contents__new(client.inner, cFileID(id))},
	}
}

//func (query QueryFileGetContents) GetContents() ([]byte, error) {
//	var cContents C.HederaArray
//	res := C.hedera_query__file_get_contents_get(query.inner, &cContents)
//	if res != 0 {
//		return nil, hederaLastError()
//	}
//	return []byte(cContents.ptr), nil
//}
