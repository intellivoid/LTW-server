// Bot.go Project
// Copyright (C) 2021 Sayan Biswas, ALiwoto
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of the source code.

package wotoValues

const (
	Yes    = "yes"
	CapYes = "Yes"
	No     = "no"
	CapNo  = "No"
)

// Type of chat, can be either
// "private", "group", "supergroup" or "channel"
const (
	Private    = "private"    // a private chat (a user)
	Group      = "group"      // a group (a normal group)
	Supergroup = "supergroup" // a super group
	Channel    = "channel"    // a channel
)

// The member's status in the chat.
// Can be "creator", "administrator", "member",
// "restricted", "left" or "kicked"
const (
	Creator       = "creator"
	Administrator = "administrator"
	Member        = "member"
	Restricted    = "restricted"
	Left          = "left"
	Kicked        = "kicked"
)

// Type of the entity.
const (
	MentionEntity       = "mention"       // "mention" (@username)
	HashtagEntity       = "hashtag"       // "hashtag" (#hashtag)
	CashtagEntity       = "cashtag"       // "cashtag" ($USD)
	BotCommandEntity    = "bot_command"   // "bot_command" (/start@jobs_bot)
	UrlEntity           = "url"           // "url" (https://telegram.org)
	EmailEntity         = "email"         // "email" (do-not-reply@telegram.org)
	PhoneNumberEntity   = "phone_number"  // "phone_number" (+1-212-555-0123)
	BoldEntity          = "bold"          // "bold" (bold text)
	ItalicEntity        = "italic"        // "italic" (italic text)
	UnderlineEntity     = "underline"     // "underline" (underlined text)
	StrikethroughEntity = "strikethrough" // strikethrough text
	CodeEntity          = "code"          // "code" (monowidth string)
	PreEntity           = "pre"           // "pre" (monowidth block)
	TextLinkEntity      = "text_link"     // "text_link" (for clickable text URLs)
	TextMentionEntity   = "text_mention"  // "text_mention"
)

const (
	MarkDownV2 = "markdownv2"
)

const (
	DebugChat int64 = -1001159890779
)
