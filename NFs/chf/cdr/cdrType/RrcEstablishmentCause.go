package cdrType

import "github.com/free5gc/chf/cdr/asn"

// Need to import "gofree5gc/lib/aper" if it uses "aper"

type RrcEstablishmentCause struct {
	Value asn.OctetString
}
