package main

import (
	"path/filepath"
	//"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

// tts packages installed ... only espeak supported for now
//

var engines []installation

// ByAge implements sort.Interface for []Person based on
// the Age field.
type byIdent []installation

func (a byIdent) Len() int           { return len(a) }
func (a byIdent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIdent) Less(i, j int) bool { return a[i].Name < a[j].Name }

// scan installation directory for tts instalations...
// Directory Naming Convention:
// tts/<engine>-<version>-(Name)/
func tts_init() {
	tts_espeak_init()
	sort.Sort(byIdent(engines))
}

func tts_engines() []installation {
	return engines
}

func tts_espeak_init() {
	// load espeakng's
	paths, err := filepath.Glob("inst/espeakng/*/bin/espeak-ng")
	if err != nil || len(paths) == 0 {
		panic(" Can not find any TTS : inst/espeakng/*/bin/espeak-ng")
	}
	for _, ele := range paths {
		ver_date := strings.Split(strings.Split(ele, "/")[2], "@")
		ver := ver_date[0]
		date := ver_date[1]
		ins := installation{
			Tts:   "espeakng",
			Ident: ver,
			Date:  date,
			Path:  ele,
			Name:  "Espeak-NG " + ver + " (" + date + ")",
			Langs: tts_espeak_langs(ele),
		}
		engines = append(engines, ins)

	}
}

func tts_espeak_langs(exe string) []language {
	langs := make([]language, 0)

	stdout, err := exec.Command("./"+exe, "--voices").Output()
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(stdout), "\n")
	for idx, ele := range lines {
		//Pty Language       Age/Gender VoiceName          File                 Other Languages
		// 5  en-gb-x-gbcwmd  --/M      english_wmids      gmw/en-GB-x-gbcwmd   (en-gb 9)(en 9)
		if idx > 0 {
			words := strings.Fields(ele)
			if len(words) > 4 {
				lang := language{
					Display: words[3],
					Option:  words[4],
				}
				langs = append(langs, lang)
			}
		}
	}
	return langs
}

// updates WavePath...
func tts_find(inv *invocation) (*installation, *language) {
	var ins *installation
	var lang *language
	for _, ele := range engines {
		if ele.Ident == inv.TTSVer {
			ins = &ele
			for _, ln := range ele.Langs {
				if ln.Option == inv.Language {
					lang = &ln
					break
				}
			}
			break
		}

	}
	return ins, lang
}

func tts_run(inv *invocation) {
	engine, lang := tts_find(inv)
	if engine == nil || lang == nil {
		inv.Wavepath = ""
		inv.Analysis = ""
		return
	}
	// only speak for now...
	wave := "waves/" + strconv.Itoa(inv.Id) + ".wav"
	mp3 := "waves/" + strconv.Itoa(inv.Id) + ".mp3"
	os.Remove(wave)
	os.Remove(mp3)
	stdout, err := exec.Command(
		"./"+engine.Path,
		"-v", lang.Option,
		"-w", wave,
		"-X",
		"--", inv.TextInput).Output()
	if err != nil {
		inv.Analysis = ""
		inv.Wavepath = ""
		return
	}
	_, err = exec.Command("lame", wave, mp3).Output()
	if err != nil {
		inv.Analysis = string(stdout)
		inv.Wavepath = mp3
	} else {
		inv.Analysis = string(stdout)
		inv.Wavepath = mp3
	}
}
