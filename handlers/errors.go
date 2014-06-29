package handlers

// type APIResponder interface {
// 	JSON()
// 	Log()
// }
//
// type APIErr struct {
// }
//
// func (apiErr *APIErr) Log() {
// 	log.Println("Error: Some loggedError")
// }
// func (apiErr *APIErr) JSON() {
// 	writeJSON("Something happened!")
// }
//
// type APIResponse struct {
// }
//
// func (apiRes *APIResponse) Log() {
// 	log.Println("Status: 200 Type:Get")
// }
// func (apiRes *APIResponse) JSON() {
// 	writeJSON("{Something: somethingElse}")
// }
//
// func (a *APIHandler) GetAllPeople() *APIResponser {
// 	people, err := db.People.Get(a["ID"])
//
// 	if err != nil {
// 		return &APIErr{Error: err}
// 	}
//
// 	return &APIResponse{"people": people}
// }
