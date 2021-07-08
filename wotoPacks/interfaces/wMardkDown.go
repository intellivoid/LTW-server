// Bot.go Project
// Copyright (C) 2021 Sayan Biswas, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package interfaces

type WMarkDown interface {
	Append(md WMarkDown) WMarkDown
	AppendNormal(v string) WMarkDown
	AppendBold(v string) WMarkDown
	AppendItalic(v string) WMarkDown
	AppendMono(v string) WMarkDown
	AppendHyperLink(text, url string) WMarkDown
	AppendMention(text string, id int64) WMarkDown
	ToString() string
}
