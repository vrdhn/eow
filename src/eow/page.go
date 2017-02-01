package main

import (
	"html/template"
	"net/http"
	"os"
)

const testpage = `
InputText: [{{.InputText}}]
{{ range  $ind, $ele  := .TTSEngines }}
<option value="{{$ele.External}}">{{$ele.Internal}}</option>
{{end}}

langVariants = {
{{ range  $tts, $langs  := .TTSLangs }}
"{{$tts}}" : 
'{{ range  $idx, $ele  := $langs }}<option value="{{$ele.External}}">{{$ele.Internal}}</option>{{end}}'
{{end}}
}

`

type pair struct {
	Internal, External string
}

type PageData struct {
	ID          int
	InputText   string
	TTSEngines  []pair
	TTSLangs    map[string][]pair
	TTSVersion  string
	LangVariant string
}

// []installation ===> [ [ display_name, ident] , ... ]

// just a dumb O(n*n) for now
// [ ins,[]lang],.... ===>      [ lang, [ins]] ,...
func make_html_lang(tts []installation) map[string][]pair {
	ret := make(map[string][]pair)
	for _, ele := range tts {
		this := make([]pair, 0)
		for _, lang := range ele.Langs {
			p := pair{lang.Display, lang.Option}
			this = append(this, p)
		}
		ret[ele.Ident] = this
	}
	return ret
}

func make_html_tts(tts []installation) []pair {
	ret := make([]pair, 0)
	for _, ele := range tts {
		p := pair{ele.Name, ele.Ident}
		ret = append(ret, p)
	}
	return ret
}

func send_newpage(w http.ResponseWriter, tts []installation, inv invocation) {
	var tpl = template.Must(template.ParseFiles("templates/index.html"))

	things := PageData{
		ID:          inv.Id,
		InputText:   inv.TextInput,
		TTSEngines:  make_html_tts(tts),
		TTSLangs:    make_html_lang(tts),
		TTSVersion:  inv.TTSVer,
		LangVariant: inv.Language,
	}
	tpl.Execute(w, things)
}

func send_testpage(tts []installation) {

	things := PageData{
		ID:          -1,
		InputText:   "HwlloWorld",
		TTSEngines:  make_html_tts(tts),
		TTSLangs:    make_html_lang(tts),
		TTSVersion:  "",
		LangVariant: "",
	}
	t := template.Must(template.New("Page").Parse(testpage))
	t.Execute(os.Stdout, things)
}
