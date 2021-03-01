package main

import "fmt"

//Email contains the information of a mail service
type Email interface {
	Write(string, string, string)
	Send() string
	Read()
}

//Test contains the information of a mail
type Test struct {
	To      string
	From    string
	Subject string
	Message string
}

func (t Test) Write(to string, from string, subject string) {
	fmt.Println(to, from, subject)
}

//Send return the recipients address
func (t Test) Send() string {
	//fmt.Println(t.To, "email sent")
	return t.To
}

func (t Test) Read() {
	fmt.Println(t.From, "received from")
}

func main() {
	var tst Test
	tst.To = "mnh.nahid35@gmail.com"
	tst.From = "test@mail.com"
	tst.Subject = "Test Email"
	tst.Message = "Hello there, this is a test email"

	tst.Write(tst.To, tst.From, tst.Subject)
}
