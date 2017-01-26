# Espeakng On Web

## 1. Goals

EoW is a website for colloboratively developing langauges for eSpeakNG.
EoW will let user play with dict and phoneme settings for a language, and 
listen to generated audio in real time. User will also be able to share
the changes for otheres to approve, and generated changes to eSpeakNG's
source

The site should be completely screen reader accesible.


## 2. Specs

Create a website which allows:
  * Paste Arbitrary text and listen to generated audio
  * Show dict & phoneme break up of provieded text
  * Allow to change dict / phoneme and hear audio
  * Create/Share link with text, dict changes, and phoneme changes.
  * Download espeakng data files diff for all the changes
  
In first version, the emscripten interface will not be used.
The server will compile language files and generate audio.

Initially the site will be built with plan html, rather than
SPA+REST. This will be revisited later.

## 3. Info. Arch

The information to be entered by  user is
  * Espeak version to be used
  * Text to be played out
  * language / variant to be used 
  * changes to phoneems
  * changes to dict - rules
  * changes to dict - list
  
Information to be displayed to user is
  * phonemes of the text
  * the dict rules/list 
  * the inbuild phone/dict/rules
  * 

The server will need to parse and send the list
of phonemes, dict-rules, and dict-list to page
so that it can be displayed to user.
  
All of these should sit behind a url, which can be shared.

## 4. Interaction Design

There will be single screen, with all the input/edit fields.
and buttons 'Generate' and 'Play'. Generate will save the values in server db

## 5. Layouts

Nothing special, just series of controls, grouped by functionality.

## 6. Visual Design


## 7. URLs, payloads
The randomly generated _slug_ as query string: http://.../q=_slug_
'Generate' will always submit the full form,
and return a new html page, with path of wave file embededed.

## 8. Frontend Design
No fancy, just HTML.

## 9. Frontend Impl


## 10. Server Design
A SQLite database for storing url's data.
A python function to do everything,
A bottle.py to wrap and call the function.

## 11. Server Impl


## 12. Deployment

Estimate: 100 MB per espeak vesion
          20  MB per url ( data compiled, wave file )
		  
10 GB Space can  support 5 espeak version and 500 urls.

Cheap 1$ VPS, with 20 GB Space should be sufficient for now.



