package request

type Email struct {
	From    string // From source email
	To      string // To destination email(s)
	Subject string // Subject text to send
	Text    string // Text is the text body representation
	HTML    string // HTMLBody is the HTML body representation
	ReplyTo string // Reply-To email(s)
}
