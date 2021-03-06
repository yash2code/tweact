package generator

import "text/template"

type header struct {
	Source  string
	Comment string
}

var headerTemplate = template.Must(template.New("header").Parse(`// generated by protoc-gen-tweact. DO NOT EDIT.
// source: {{.Source}}

{{- .Comment}}

import { withTwirp, TwirpService } from "@department/twirp-component";
`))
