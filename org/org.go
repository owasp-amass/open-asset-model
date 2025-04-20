// Copyright © by Jeff Foley 2017-2025. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package org

import (
	"encoding/json"

	model "github.com/owasp-amass/open-asset-model"
)

// Organization represents an organization.
// Should support relationships for the following:
// - Principal place of business
// - Parent organizations
// - Subsidiaries
// - Sister companies
// - DUNS number
// - Tax identification number
// - Trader identification number
// - ARIN organization identifier
// - Decentralized identifier (DID)
// - Ticker symbol
// - Global Location Number (GLN)
// - ISIC, NAICS, SIC, BIC, and ISO 6523 code
// - Legal Entity Identifier (LEI) ISO 17442 code
// - Registration number
// - Website
// - Social media profiles
// - Contact information
// - Founder, sponsorships, and funding sources
type Organization struct {
	ID             string   `json:"unique_id"`
	Name           string   `json:"name"`
	LegalName      string   `json:"legal_name,omitempty"`
	FoundingDate   string   `json:"founding_date,omitempty"`
	Jurisdiction   string   `json:"jurisdiction,omitempty"`
	RegistrationID string   `json:"registration_id,omitempty"`
	Industry       string   `json:"industry,omitempty"`
	TargetMarkets  []string `json:"target_markets,omitempty"`
	Active         bool     `json:"active,omitempty"`
	NonProfit      bool     `json:"non_profit,omitempty"`
	Headcount      int      `json:"headcount,omitempty"`
}

// Key implements the Asset interface.
func (o Organization) Key() string {
	return o.ID
}

// AssetType implements the Asset interface.
func (o Organization) AssetType() model.AssetType {
	return model.Organization
}

// JSON implements the Asset interface.
func (o Organization) JSON() ([]byte, error) {
	return json.Marshal(o)
}

// GICSIndustryMap captures the sectors, industry groups, and industries
// of the Global Industry Classification Standard (GICS) taxonomy.
// It is a standardized classification system for categorizing companies
// into sectors and industries based on their primary business activities.
// GICS was developed by MSCI and Standard & Poor's (S&P) to provide a
// consistent framework for analyzing and comparing companies across
// different sectors and industries. The GICS taxonomy is widely used by
// investors, analysts, and financial professionals to assess the
// performance of various sectors and industries in the global economy.
// The GICS taxonomy consists of 11 sectors, each of which is further
// divided into 24 industry groups, 69 industries, and 158 sub-industries.
var GICSIndustryMap = map[string]map[string][]string{
	"Communication Services": {
		"Telecommunication Services": {"Diversified Telecommunication Services", "Wireless Telecommunication Services"},
		"Media & Entertainment":      {"Media", "Entertainment", "Interactive Media & Services"},
	},
	"Consumer Discretionary": {
		"Automobiles & Components": {"Automobiles", "Automobile Components"},
		"Consumer Durable & Apparel": {
			"Household Durables",
			"Leisure Products",
			"Textiles, Apparel & Luxury Goods",
		},
		"Consumer Services":                            {"Hotels, Restaurants & Leisure", "Diversified Consumer Services"},
		"Consumer Discretionary Distribution & Retail": {"Distributors", "Broadline Retail", "Specialty Retail"},
	},
	"Consumer Staples": {
		"Consumer Staples Distribtuion & Retail": {"Consumer Staples Distribution & Retail"},
	},
	"Energy": {
		"Energy": {"Energy Equipment & Services", "Oil, Gas & Consumable Fuels"},
	},
	"Financials": {
		"Banks": {"Banks"},
		"Financial Services": {
			"Financial Services",
			"Consumer Finance",
			"Capital Markets",
			"Mortgage Real Estate Investment Trusts",
		},
		"Insurance": {"Insurance"},
	},
	"Health Care": {
		"Health Care Equipment & Services":               {"Health Care Equipment & Supplies", "Health Care Providers & Services"},
		"Pharmaceuticals, Biotechnology & Life Sciences": {"Pharmaceuticals", "Biotechnology", "Life Sciences Tools & Services"},
	},
	"Industrials": {
		"Capital Goods": {
			"Aerospace & Defense",
			"Building Products",
			"Construction & Engineering",
			"Electrical Equipment",
			"Industrial Conglomerates",
			"Machinery",
			"Trading Companies & Distributors",
		},
		"Commercial & Professional Services": {"Commercial Services & Supplies", "Professional Services"},
		"Transportation": {
			"Air Freight & Logistics",
			"Passenger Airlines",
			"Marine Transportation",
			"Ground Transportation",
			"Transportation Infrastructure",
		},
	},
	"Information Technology": {
		"Software & Services": {"Software", "IT Services"},
		"Technology Hardware & Equipment": {
			"Technology Hardware, Storage & Peripherals",
			"Electronic Equipment, Instruments & Components",
			"Communications Equipment",
		},
		"Semiconductors & Semiconductor Equipment": {"Semiconductors & Semiconductor Equipment"},
	},
	"Materials": {
		"Materials": {
			"Chemicals",
			"Construction Materials",
			"Containers & Packaging",
			"Metals & Mining",
			"Paper & Forest Products",
		},
	},
	"Real Estate": {
		"Equity Real Estate Investment Trusts": {
			"Diversified REITs",
			"Industrial REITs",
			"Hotel & Resort REITs",
			"Office REITs",
			"Health Care REITs",
			"Residential REITs",
			"Retail REITs",
			"Specialized REITs",
		},
		"Real Estate Management & Development": {"Real Estate Management & Development"},
	},
	"Utilities": {
		"Utilities": {
			"Electric Utilities",
			"Gas Utilities",
			"Multi-Utilities",
			"Water Utilities",
			"Independent Power and Renewable Electricity Producers",
		},
	},
}
