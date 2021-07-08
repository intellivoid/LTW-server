/*
 * This file is part of LTW-server project (https://github.com/intellivoid/LTW-server).
 * Copyright (c) 2021 WotoTeam.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 2.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package versioning

const (
	// MaxVersion is the maximum version of the game. if the client's version is
	// more than this version, it's not acceptable.
	currentVersion = "1.0.0.0"

	versionHash = "f302bd7ffacbd295194f86620002b8250e8e9be0233a8055bcebc82c8612843ff9e0f09e42015d5e75581cc93d4c29a91388ed411641b543c8fb7b5a26a2a8b8"
)

const (
	//-----------------------------------------------------
	userAgentKey   = "User-Agent" // not optional
	userAgentValue = "ltw-client" // not optional

	//-----------------------------------------------------

	// same domain header field means we are sending the
	// HTTP request to the same domain of our referer domain
	// (and origin).
	// we expect it to send us the respond using the same protocol.
	// The same-origin policy is a critical security mechanism that
	// restricts how a document or script loaded by one origin can interact with
	// a resource from another origin.
	// we should set its value to 1.
	//
	//  > see also: https://developer.mozilla.org/ja/docs/Web/Security/Same-origin_policy
	xSameDomainKey   = "X-Same-Domain" // not optional
	xSameDomainValue = "1"             // not optional

	//-----------------------------------------------------

	ltwVersionKey     = "ltw_version"
	ltwVersionHashKey = "ltw_version_hash"
)
