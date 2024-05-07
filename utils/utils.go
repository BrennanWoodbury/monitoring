package utils

import (
	"github.com/gosnmp/gosnmp"
)

func TranslatePDU(pdu gosnmp.SnmpPDU) interface{} {
	switch pdu.Type {
	case gosnmp.OctetString:
		value := string(pdu.Value.([]byte))
		return value

	case gosnmp.Integer:
		value := pdu.Value.(int)
		return value

	case gosnmp.Counter32, gosnmp.Counter64:
		return pdu.Value.(string)

	case gosnmp.OpaqueFloat:
		value := pdu.Value.(float32)
		return value

	case gosnmp.OpaqueDouble:
		value := pdu.Value.(float64)
		return value

	default:
		return pdu.Value
	}

}
