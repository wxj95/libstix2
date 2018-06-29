// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
//
// Types
//
// ----------------------------------------------------------------------

/*
AliasesProperty - A property used by one or more STIX objects.
*/
type AliasesProperty struct {
	Aliases []string `json:"aliases,omitempty"`
}

// ----------------------------------------------------------------------
//
// Public Methods - AliasesProperty
//
// ----------------------------------------------------------------------

/*
AddAlias - This method takes in a takes in a string value that represents an
alias for something in STIX and adds it to the property.
*/
func (p *AliasesProperty) AddAlias(s string) error {
	p.Aliases = append(p.Aliases, s)
	return nil
}
