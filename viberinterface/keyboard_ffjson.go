// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: keyboard.go
// DO NOT EDIT!

package viberinterface

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Button) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Button) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if mj.Columns != 0 {
		buf.WriteString(`"Columns":`)
		fflib.FormatBits2(buf, uint64(mj.Columns), 10, mj.Columns < 0)
		buf.WriteByte(',')
	}
	if mj.Rows != 0 {
		buf.WriteString(`"Rows":`)
		fflib.FormatBits2(buf, uint64(mj.Rows), 10, mj.Rows < 0)
		buf.WriteByte(',')
	}
	if len(mj.BgColor) != 0 {
		buf.WriteString(`"BgColor":`)
		fflib.WriteJsonString(buf, string(mj.BgColor))
		buf.WriteByte(',')
	}
	if len(mj.BgMediaType) != 0 {
		buf.WriteString(`"BgMediaType":`)
		fflib.WriteJsonString(buf, string(mj.BgMediaType))
		buf.WriteByte(',')
	}
	if len(mj.BgMedia) != 0 {
		buf.WriteString(`"BgMedia":`)
		fflib.WriteJsonString(buf, string(mj.BgMedia))
		buf.WriteByte(',')
	}
	if mj.BgLoop != false {
		if mj.BgLoop {
			buf.WriteString(`"BgLoop":true`)
		} else {
			buf.WriteString(`"BgLoop":false`)
		}
		buf.WriteByte(',')
	}
	if len(mj.ActionType) != 0 {
		buf.WriteString(`"ActionType":`)
		fflib.WriteJsonString(buf, string(mj.ActionType))
		buf.WriteByte(',')
	}
	if len(mj.ActionBody) != 0 {
		buf.WriteString(`"ActionBody":`)
		fflib.WriteJsonString(buf, string(mj.ActionBody))
		buf.WriteByte(',')
	}
	if len(mj.Image) != 0 {
		buf.WriteString(`"Image":`)
		fflib.WriteJsonString(buf, string(mj.Image))
		buf.WriteByte(',')
	}
	if len(mj.Text) != 0 {
		buf.WriteString(`"Text":`)
		fflib.WriteJsonString(buf, string(mj.Text))
		buf.WriteByte(',')
	}
	if len(mj.TextVAlign) != 0 {
		buf.WriteString(`"TextVAlign":`)
		fflib.WriteJsonString(buf, string(mj.TextVAlign))
		buf.WriteByte(',')
	}
	if len(mj.TextHAlign) != 0 {
		buf.WriteString(`"TextHAlign":`)
		fflib.WriteJsonString(buf, string(mj.TextHAlign))
		buf.WriteByte(',')
	}
	if mj.TextOpacity != 0 {
		buf.WriteString(`"TextOpacity":`)
		fflib.FormatBits2(buf, uint64(mj.TextOpacity), 10, mj.TextOpacity < 0)
		buf.WriteByte(',')
	}
	if len(mj.TextSize) != 0 {
		buf.WriteString(`"TextSize":`)
		fflib.WriteJsonString(buf, string(mj.TextSize))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Buttonbase = iota
	ffj_t_Buttonno_such_key

	ffj_t_Button_Columns

	ffj_t_Button_Rows

	ffj_t_Button_BgColor

	ffj_t_Button_BgMediaType

	ffj_t_Button_BgMedia

	ffj_t_Button_BgLoop

	ffj_t_Button_ActionType

	ffj_t_Button_ActionBody

	ffj_t_Button_Image

	ffj_t_Button_Text

	ffj_t_Button_TextVAlign

	ffj_t_Button_TextHAlign

	ffj_t_Button_TextOpacity

	ffj_t_Button_TextSize
)

var ffj_key_Button_Columns = []byte("Columns")

var ffj_key_Button_Rows = []byte("Rows")

var ffj_key_Button_BgColor = []byte("BgColor")

var ffj_key_Button_BgMediaType = []byte("BgMediaType")

var ffj_key_Button_BgMedia = []byte("BgMedia")

var ffj_key_Button_BgLoop = []byte("BgLoop")

var ffj_key_Button_ActionType = []byte("ActionType")

var ffj_key_Button_ActionBody = []byte("ActionBody")

var ffj_key_Button_Image = []byte("Image")

var ffj_key_Button_Text = []byte("Text")

var ffj_key_Button_TextVAlign = []byte("TextVAlign")

var ffj_key_Button_TextHAlign = []byte("TextHAlign")

var ffj_key_Button_TextOpacity = []byte("TextOpacity")

var ffj_key_Button_TextSize = []byte("TextSize")

func (uj *Button) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Button) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Buttonbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Buttonno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'A':

					if bytes.Equal(ffj_key_Button_ActionType, kn) {
						currentKey = ffj_t_Button_ActionType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_ActionBody, kn) {
						currentKey = ffj_t_Button_ActionBody
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'B':

					if bytes.Equal(ffj_key_Button_BgColor, kn) {
						currentKey = ffj_t_Button_BgColor
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_BgMediaType, kn) {
						currentKey = ffj_t_Button_BgMediaType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_BgMedia, kn) {
						currentKey = ffj_t_Button_BgMedia
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_BgLoop, kn) {
						currentKey = ffj_t_Button_BgLoop
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'C':

					if bytes.Equal(ffj_key_Button_Columns, kn) {
						currentKey = ffj_t_Button_Columns
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffj_key_Button_Image, kn) {
						currentKey = ffj_t_Button_Image
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'R':

					if bytes.Equal(ffj_key_Button_Rows, kn) {
						currentKey = ffj_t_Button_Rows
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'T':

					if bytes.Equal(ffj_key_Button_Text, kn) {
						currentKey = ffj_t_Button_Text
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_TextVAlign, kn) {
						currentKey = ffj_t_Button_TextVAlign
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_TextHAlign, kn) {
						currentKey = ffj_t_Button_TextHAlign
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_TextOpacity, kn) {
						currentKey = ffj_t_Button_TextOpacity
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Button_TextSize, kn) {
						currentKey = ffj_t_Button_TextSize
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_Button_TextSize, kn) {
					currentKey = ffj_t_Button_TextSize
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_TextOpacity, kn) {
					currentKey = ffj_t_Button_TextOpacity
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_TextHAlign, kn) {
					currentKey = ffj_t_Button_TextHAlign
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_TextVAlign, kn) {
					currentKey = ffj_t_Button_TextVAlign
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_Text, kn) {
					currentKey = ffj_t_Button_Text
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_Image, kn) {
					currentKey = ffj_t_Button_Image
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_ActionBody, kn) {
					currentKey = ffj_t_Button_ActionBody
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_ActionType, kn) {
					currentKey = ffj_t_Button_ActionType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_BgLoop, kn) {
					currentKey = ffj_t_Button_BgLoop
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_BgMedia, kn) {
					currentKey = ffj_t_Button_BgMedia
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_BgMediaType, kn) {
					currentKey = ffj_t_Button_BgMediaType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Button_BgColor, kn) {
					currentKey = ffj_t_Button_BgColor
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Button_Rows, kn) {
					currentKey = ffj_t_Button_Rows
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Button_Columns, kn) {
					currentKey = ffj_t_Button_Columns
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Buttonno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Button_Columns:
					goto handle_Columns

				case ffj_t_Button_Rows:
					goto handle_Rows

				case ffj_t_Button_BgColor:
					goto handle_BgColor

				case ffj_t_Button_BgMediaType:
					goto handle_BgMediaType

				case ffj_t_Button_BgMedia:
					goto handle_BgMedia

				case ffj_t_Button_BgLoop:
					goto handle_BgLoop

				case ffj_t_Button_ActionType:
					goto handle_ActionType

				case ffj_t_Button_ActionBody:
					goto handle_ActionBody

				case ffj_t_Button_Image:
					goto handle_Image

				case ffj_t_Button_Text:
					goto handle_Text

				case ffj_t_Button_TextVAlign:
					goto handle_TextVAlign

				case ffj_t_Button_TextHAlign:
					goto handle_TextHAlign

				case ffj_t_Button_TextOpacity:
					goto handle_TextOpacity

				case ffj_t_Button_TextSize:
					goto handle_TextSize

				case ffj_t_Buttonno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Columns:

	/* handler: uj.Columns type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Columns = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Rows:

	/* handler: uj.Rows type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Rows = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BgColor:

	/* handler: uj.BgColor type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BgColor = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BgMediaType:

	/* handler: uj.BgMediaType type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BgMediaType = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BgMedia:

	/* handler: uj.BgMedia type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BgMedia = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BgLoop:

	/* handler: uj.BgLoop type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.BgLoop = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.BgLoop = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ActionType:

	/* handler: uj.ActionType type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ActionType = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ActionBody:

	/* handler: uj.ActionBody type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.ActionBody = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Image:

	/* handler: uj.Image type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Image = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Text:

	/* handler: uj.Text type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Text = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TextVAlign:

	/* handler: uj.TextVAlign type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.TextVAlign = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TextHAlign:

	/* handler: uj.TextHAlign type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.TextHAlign = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TextOpacity:

	/* handler: uj.TextOpacity type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.TextOpacity = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TextSize:

	/* handler: uj.TextSize type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.TextSize = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *Keyboard) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Keyboard) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "Type":`)
	fflib.WriteJsonString(buf, string(mj.Type))
	buf.WriteString(`,"Buttons":`)
	if mj.Buttons != nil {
		buf.WriteString(`[`)
		for i, v := range mj.Buttons {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				err = v.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if len(mj.BgColor) != 0 {
		buf.WriteString(`"BgColor":`)
		fflib.WriteJsonString(buf, string(mj.BgColor))
		buf.WriteByte(',')
	}
	if mj.DefaultHeight != false {
		if mj.DefaultHeight {
			buf.WriteString(`"DefaultHeight":true`)
		} else {
			buf.WriteString(`"DefaultHeight":false`)
		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Keyboardbase = iota
	ffj_t_Keyboardno_such_key

	ffj_t_Keyboard_Type

	ffj_t_Keyboard_Buttons

	ffj_t_Keyboard_BgColor

	ffj_t_Keyboard_DefaultHeight
)

var ffj_key_Keyboard_Type = []byte("Type")

var ffj_key_Keyboard_Buttons = []byte("Buttons")

var ffj_key_Keyboard_BgColor = []byte("BgColor")

var ffj_key_Keyboard_DefaultHeight = []byte("DefaultHeight")

func (uj *Keyboard) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Keyboard) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Keyboardbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Keyboardno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'B':

					if bytes.Equal(ffj_key_Keyboard_Buttons, kn) {
						currentKey = ffj_t_Keyboard_Buttons
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Keyboard_BgColor, kn) {
						currentKey = ffj_t_Keyboard_BgColor
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'D':

					if bytes.Equal(ffj_key_Keyboard_DefaultHeight, kn) {
						currentKey = ffj_t_Keyboard_DefaultHeight
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'T':

					if bytes.Equal(ffj_key_Keyboard_Type, kn) {
						currentKey = ffj_t_Keyboard_Type
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffj_key_Keyboard_DefaultHeight, kn) {
					currentKey = ffj_t_Keyboard_DefaultHeight
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Keyboard_BgColor, kn) {
					currentKey = ffj_t_Keyboard_BgColor
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Keyboard_Buttons, kn) {
					currentKey = ffj_t_Keyboard_Buttons
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Keyboard_Type, kn) {
					currentKey = ffj_t_Keyboard_Type
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Keyboardno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Keyboard_Type:
					goto handle_Type

				case ffj_t_Keyboard_Buttons:
					goto handle_Buttons

				case ffj_t_Keyboard_BgColor:
					goto handle_BgColor

				case ffj_t_Keyboard_DefaultHeight:
					goto handle_DefaultHeight

				case ffj_t_Keyboardno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Type:

	/* handler: uj.Type type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Type = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Buttons:

	/* handler: uj.Buttons type=[]viberinterface.Button kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Buttons = nil
		} else {

			uj.Buttons = []Button{}

			wantVal := true

			for {

				var tmp_uj__Buttons Button

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmp_uj__Buttons type=viberinterface.Button kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

						state = fflib.FFParse_after_value
						goto mainparse
					}

					err = tmp_uj__Buttons.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
					if err != nil {
						return err
					}
					state = fflib.FFParse_after_value
				}

				uj.Buttons = append(uj.Buttons, tmp_uj__Buttons)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BgColor:

	/* handler: uj.BgColor type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.BgColor = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_DefaultHeight:

	/* handler: uj.DefaultHeight type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.DefaultHeight = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.DefaultHeight = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
