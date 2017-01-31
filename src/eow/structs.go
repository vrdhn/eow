package main


// a language name, and the command line option to select it
type language struct {
	Display  string		// option to display to user
	Option   string		// option to pass to TTS engine
}
// an TTS installation , and supported languages
type installation struct {
	Tts    string		`json:"-"` // espeakng for now.
	Ident  string		// for html usage
	Date   string		// release date.
	Path   string		`json:"-"` // the installation path
	Name   string		// user visible name
	Langs []language	// supported languages
}


// a single web-invocation 
type invocation struct {	
	Id          int64	// An id under which this is saved.
	TextInput   string	// UTF8 text to play out
	Language    string      // selected language, or empty
	Wavepath    string	// Path where wav is stored
	Analysis    string	// phoneme analysis of text
}
