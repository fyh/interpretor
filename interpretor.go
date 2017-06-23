package interpreter

import (
    "io/ioutil"
    "fmt"
    "github.com/bitly/go-simplejson"
    "net/url"
    "net/http"
    "bytes"
    "errors"
)

const apiUrl string = "https://translate.googleapis.com/translate_a/single"

// returns the current implementation version
func Version() string {
    return "0.0.1"
}

type Interpreter struct {
    client string // "gtx"
    sourceLang string
    targetLang string
    targetType string
    inputEncoding string  // constant: "utf-8"
    outputEncoding string  // constant: "utf-8"
    destType string
    text string
}

func (i *Interpreter) Link() string {
    formatter := "?client=%s&ie=%s&oe=%s&sl=%s&tl=%s&dt=%s&q=%s"
    return apiUrl + fmt.Sprintf(
        formatter,
        i.client,
        i.inputEncoding,
        i.outputEncoding,
        i.sourceLang,
        i.targetLang,
        i.destType,
        url.QueryEscape(i.text),
    )
}

// New returns a pointer to a new, empty `Interpreter` object
func New() *Interpreter {
    return &Interpreter{
        client: "gtx",
        inputEncoding: "utf-8",
        outputEncoding: "utf-8",
        sourceLang: "auto",
        targetLang: "en",
        destType: "t",
        text: "",
    }
}

// New returns a pointer to a new, empty `Interpreter` object
func NewText(t string, tl string) *Interpreter {
    i := New()
    i.targetLang = tl
    i.text = t
    return i
}

// Append append text `t` to the interpretor
func (i *Interpreter) Append(t string) {
    i.text += t
}

// Interpret interpret source text `t`
func (i *Interpreter) Interpret() (string, error) {
    resp, err := http.Get(i.Link())

    defer resp.Body.Close()

    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    json, err := simplejson.NewJson(body)

    if err != nil {
        return "", err
    }

    a0 := json.MustArray()[0]
    a1 := a0.([]interface{})
    var bf bytes.Buffer
    for _, a2 := range a1 {
        s := a2.([]interface{})
        r, ok := s[0].(string)
        if !ok {
            return "", errors.New(fmt.Sprintf("type assertion failed: %p (%T) -> string\n", s[0], s[0]))
        }
        bf.WriteString(r)
    }

    return bf.String(), nil
}

// Translate for simple translation like `interpretor.Translate("hello, world", interpretor.zh_CN)`
// source language is automaticly detected with `auto` mark
func Translate(t string, tl string) (string, error) {
    return NewText(t, tl).Interpret()
}
