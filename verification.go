package main

import (
	"net"
	"strings"
)

const SPFPrefix string = "v=spf1"
const DMARCPrefix string = "v=DMARC1"
const DMARCDomainPrefix string = "_dmarc."

type DomainInfo struct {
	HasMX       bool   `json:"hasMX"`
	HasSPF      bool   `json:"hasSPF"`
	HasDMARC    bool   `json:"hasDMARC"`
	SpfRecord   string `json:"SPFRecord"`
	DMARCRecord string `json:"DMARCRecord"`
}

func checkMX(domain string) (bool, error) {
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		return false, err
	}

	return len(mxRecords) > 0, nil
}

func checkSPF(domain string) (bool, string, error) {
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		return false, "", err
	}

	var hasSPF bool
	var SPFRecord string

	for _, record := range txtRecords {
		if strings.HasPrefix(record, SPFPrefix) {
			hasSPF = true
			SPFRecord = record
			break
		}
	}

	return hasSPF, SPFRecord, nil
}

func checkDMARC(domain string) (bool, string, error) {
	dmarcRecords, err := net.LookupTXT(DMARCDomainPrefix + domain)

	if err != nil {
		return false, "", err
	}

	var hasDMARC bool
	var DMARCRecord string

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, DMARCPrefix) {
			hasDMARC = true
			DMARCRecord = record
			break
		}
	}

	return hasDMARC, DMARCRecord, nil
}

func CheckDomain(domain string) (*DomainInfo, error) {
	hasMX, err := checkMX(domain)

	if err != nil {
		return &DomainInfo{}, err
	}

	hasSPF, SPFRecord, err := checkSPF(domain)

	if err != nil {
		return &DomainInfo{}, err
	}

	hasDMARC, DMARCRecord, err := checkDMARC(domain)

	if err != nil {
		return &DomainInfo{}, err
	}

	return &DomainInfo{
		HasMX:       hasMX,
		HasSPF:      hasSPF,
		HasDMARC:    hasDMARC,
		SpfRecord:   SPFRecord,
		DMARCRecord: DMARCRecord,
	}, nil
}
