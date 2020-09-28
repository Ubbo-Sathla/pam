package main

/*
#include <stdlib.h>
#include <security/pam_appl.h>

*/
import "C"
import (
	"log"
	"unsafe"
)

func init() {

}
func main() {

}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	log.Println("pam_sm_authenticate")

	var user string

	var cuser *C.char
	defer C.free(unsafe.Pointer(cuser))
	if getUserName(pamh, &cuser) {
		user = C.GoString(cuser)
	} else {
		return C.PAM_USER_UNKNOWN
	}

	log.Println(user)
	var cPassword *C.char
	defer C.free(unsafe.Pointer(cPassword))
	if !getPassword(pamh, &cPassword) {
		return C.PAM_AUTH_ERR
	}

	log.Println(C.GoString(cPassword))

	return C.PAM_SUCCESS
}

//export pam_sm_acct_mgmt
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	log.Println("pam_sm_acct_mgmt")
	return C.PAM_SUCCESS
}

//export pam_sm_open_session
func pam_sm_open_session(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	log.Println("pam_sm_open_session")

	return C.PAM_SUCCESS
}

//export pam_sm_close_session
func pam_sm_close_session(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	log.Println("pam_sm_close_session")
	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	log.Println("pam_sm_setcred")

	return C.PAM_SUCCESS
}

