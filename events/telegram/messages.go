package telegram

const msgHelp = `I can save and keep your page. Also i can offer yo them to read.

In order to save the page, just send me all link to it.

In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list!`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand = "UnknownCommand"
	msgNoSavedPages   = "You have not saved pages"
	msgSaved          = "Saved!"
	msgAlreadyExists  = "You have already have this page in your list"
)
