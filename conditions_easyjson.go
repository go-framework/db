// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package db

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson53ae5e32DecodeGithubComGoFrameworkDb(in *jlexer.Lexer, out *Conditions) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "limit":
			out.Limit = int(in.Int())
		case "offset":
			out.Offset = int(in.Int())
		case "cursor":
			if m, ok := out.Cursor.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Cursor.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Cursor = in.Interface()
			}
		case "fields":
			if in.IsNull() {
				in.Skip()
				out.Fields = nil
			} else {
				in.Delim('[')
				if out.Fields == nil {
					if !in.IsDelim(']') {
						out.Fields = make([]string, 0, 4)
					} else {
						out.Fields = []string{}
					}
				} else {
					out.Fields = (out.Fields)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Fields = append(out.Fields, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "order":
			if in.IsNull() {
				in.Skip()
				out.Order = nil
			} else {
				in.Delim('[')
				if out.Order == nil {
					if !in.IsDelim(']') {
						out.Order = make([]string, 0, 4)
					} else {
						out.Order = []string{}
					}
				} else {
					out.Order = (out.Order)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Order = append(out.Order, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "group":
			if in.IsNull() {
				in.Skip()
				out.Group = nil
			} else {
				in.Delim('[')
				if out.Group == nil {
					if !in.IsDelim(']') {
						out.Group = make([]string, 0, 4)
					} else {
						out.Group = []string{}
					}
				} else {
					out.Group = (out.Group)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.Group = append(out.Group, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "having":
			if in.IsNull() {
				in.Skip()
				out.Having = nil
			} else {
				in.Delim('[')
				if out.Having == nil {
					if !in.IsDelim(']') {
						out.Having = make([]interface{}, 0, 4)
					} else {
						out.Having = []interface{}{}
					}
				} else {
					out.Having = (out.Having)[:0]
				}
				for !in.IsDelim(']') {
					var v4 interface{}
					if m, ok := v4.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v4.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v4 = in.Interface()
					}
					out.Having = append(out.Having, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "and":
			if in.IsNull() {
				in.Skip()
				out.And = nil
			} else {
				in.Delim('[')
				if out.And == nil {
					if !in.IsDelim(']') {
						out.And = make([]interface{}, 0, 4)
					} else {
						out.And = []interface{}{}
					}
				} else {
					out.And = (out.And)[:0]
				}
				for !in.IsDelim(']') {
					var v5 interface{}
					if m, ok := v5.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v5.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v5 = in.Interface()
					}
					out.And = append(out.And, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "or":
			if in.IsNull() {
				in.Skip()
				out.Or = nil
			} else {
				in.Delim('[')
				if out.Or == nil {
					if !in.IsDelim(']') {
						out.Or = make([]interface{}, 0, 4)
					} else {
						out.Or = []interface{}{}
					}
				} else {
					out.Or = (out.Or)[:0]
				}
				for !in.IsDelim(']') {
					var v6 interface{}
					if m, ok := v6.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v6.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v6 = in.Interface()
					}
					out.Or = append(out.Or, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "not":
			if in.IsNull() {
				in.Skip()
				out.Not = nil
			} else {
				in.Delim('[')
				if out.Not == nil {
					if !in.IsDelim(']') {
						out.Not = make([]interface{}, 0, 4)
					} else {
						out.Not = []interface{}{}
					}
				} else {
					out.Not = (out.Not)[:0]
				}
				for !in.IsDelim(']') {
					var v7 interface{}
					if m, ok := v7.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v7.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v7 = in.Interface()
					}
					out.Not = append(out.Not, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "between":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v8 := 0
				for !in.IsDelim(']') {
					if v8 < 2 {
						if m, ok := (out.Between)[v8].(easyjson.Unmarshaler); ok {
							m.UnmarshalEasyJSON(in)
						} else if m, ok := (out.Between)[v8].(json.Unmarshaler); ok {
							_ = m.UnmarshalJSON(in.Raw())
						} else {
							(out.Between)[v8] = in.Interface()
						}
						v8++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson53ae5e32EncodeGithubComGoFrameworkDb(out *jwriter.Writer, in Conditions) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Limit != 0 {
		const prefix string = ",\"limit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Limit))
	}
	if in.Offset != 0 {
		const prefix string = ",\"offset\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Offset))
	}
	if in.Cursor != nil {
		const prefix string = ",\"cursor\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if m, ok := in.Cursor.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Cursor.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Cursor))
		}
	}
	if len(in.Fields) != 0 {
		const prefix string = ",\"fields\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Fields {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	if len(in.Order) != 0 {
		const prefix string = ",\"order\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Order {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if len(in.Group) != 0 {
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v13, v14 := range in.Group {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	if len(in.Having) != 0 {
		const prefix string = ",\"having\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v15, v16 := range in.Having {
				if v15 > 0 {
					out.RawByte(',')
				}
				if m, ok := v16.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v16.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v16))
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.And) != 0 {
		const prefix string = ",\"and\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v17, v18 := range in.And {
				if v17 > 0 {
					out.RawByte(',')
				}
				if m, ok := v18.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v18.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v18))
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Or) != 0 {
		const prefix string = ",\"or\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v19, v20 := range in.Or {
				if v19 > 0 {
					out.RawByte(',')
				}
				if m, ok := v20.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v20.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v20))
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Not) != 0 {
		const prefix string = ",\"not\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v21, v22 := range in.Not {
				if v21 > 0 {
					out.RawByte(',')
				}
				if m, ok := v22.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v22.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v22))
				}
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"between\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawByte('[')
		for v23 := range in.Between {
			if v23 > 0 {
				out.RawByte(',')
			}
			if m, ok := (in.Between)[v23].(easyjson.Marshaler); ok {
				m.MarshalEasyJSON(out)
			} else if m, ok := (in.Between)[v23].(json.Marshaler); ok {
				out.Raw(m.MarshalJSON())
			} else {
				out.Raw(json.Marshal((in.Between)[v23]))
			}
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Conditions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson53ae5e32EncodeGithubComGoFrameworkDb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Conditions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson53ae5e32EncodeGithubComGoFrameworkDb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Conditions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson53ae5e32DecodeGithubComGoFrameworkDb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Conditions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson53ae5e32DecodeGithubComGoFrameworkDb(l, v)
}
