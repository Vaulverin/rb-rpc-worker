package main

import (
	"net/url"
)

func DefaultQueryParams(u *url.URL, msisdn string) {
	q := u.Query()
	q.Set("loyaltyProgramMember.characteristic.name", "AccountIdentifierTypeCode,AccountIdentifierCode")
	q.Set("loyaltyProgramMember.characteristic.value", "MSISDN," + msisdn)
	u.RawQuery = q.Encode()
}
