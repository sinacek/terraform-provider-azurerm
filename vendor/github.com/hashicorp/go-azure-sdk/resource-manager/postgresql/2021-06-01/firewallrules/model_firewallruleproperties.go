package firewallrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallRuleProperties struct {
	EndIpAddress   string `json:"endIpAddress"`
	StartIpAddress string `json:"startIpAddress"`
}
