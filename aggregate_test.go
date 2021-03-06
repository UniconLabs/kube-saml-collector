package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregator(t *testing.T) {
	asrt := assert.New(t)

	a := aggregator{}
	asrt.NoError(a.add(strings.NewReader(testMD)))
	asrt.Len(a.Entities, 3)
	asrt.Equal("EntityDescriptor", a.Entities["https://portal.astuart.co/uPortal"].Tag)

	asrt.Len(a.xmlns, 4)

	d := a.Doc()

	ct := 0
	for _, at := range d.Root().Attr {
		asrt.True(at.Key == "xmlns" || at.Space == "xmlns")
		ct++
	}
	asrt.Equal(4, ct)
}

const testMD = `
<EntitiesDescriptor xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:shibmd="urn:mace:shibboleth:metadata:1.0"
                    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                    xmlns="urn:oasis:names:tc:SAML:2.0:metadata" ID="INC20150622T184912"
                    Name="urn:mace:incommon" validUntil="2017-01-01T00:00:00Z"
                    xsi:schemaLocation="urn:oasis:names:tc:SAML:2.0:metadata sstc-saml-schema-metadata-2.0.xsd urn:mace:shibboleth:metadata:1.0 shibboleth-metadata-1.0.xsd http://www.w3.org/2000/09/xmldsig# xmldsig-core-schema.xsd">
    <!-- SSP uPortal dev env -->
    <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" ID="https://portal.astuart.co/uPortal"
                         entityID="https://portal.astuart.co/uPortal">
        <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false"
                            protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
            <md:KeyDescriptor use="signing">
                <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
                    <ds:X509Data>
                        <ds:X509Certificate>MIIDUjCCAjqgAwIBAgIEUOLIQTANBgkqhkiG9w0BAQUFADBrMQswCQYDVQQGEwJGSTEQMA4GA1UE
                            CBMHVXVzaW1hYTERMA8GA1UEBxMISGVsc2lua2kxGDAWBgNVBAoTD1JNNSBTb2Z0d2FyZSBPeTEM
                            MAoGA1UECwwDUiZEMQ8wDQYDVQQDEwZhcG9sbG8wHhcNMTMwMTAxMTEyODAxWhcNMjIxMjMwMTEy
                            ODAxWjBrMQswCQYDVQQGEwJGSTEQMA4GA1UECBMHVXVzaW1hYTERMA8GA1UEBxMISGVsc2lua2kx
                            GDAWBgNVBAoTD1JNNSBTb2Z0d2FyZSBPeTEMMAoGA1UECwwDUiZEMQ8wDQYDVQQDEwZhcG9sbG8w
                            ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCXqP0wqL2Ai1haeTj0alwsLafhrDtUt00E
                            5xc7kdD7PISRA270ZmpYMB4W24Uk2QkuwaBp6dI/yRdUvPfOT45YZrqIxMe2451PAQWtEKWF5Z13
                            F0J4/lB71TtrzyH94RnqSHXFfvRN8EY/rzuEzrpZrHdtNs9LRyLqcRTXMMO4z7QghBuxh3K5gu7K
                            qxpHx6No83WNZj4B3gvWLRWv05nbXh/F9YMeQClTX1iBNAhLQxWhwXMKB4u1iPQ/KSaal3R26pON
                            UUmu1qVtU1quQozSTPD8HvsDqGG19v2+/N3uf5dRYtvEPfwXN3wIY+/R93vBA6lnl5nTctZIRsyg
                            0Gv5AgMBAAEwDQYJKoZIhvcNAQEFBQADggEBAFQwAAYUjso1VwjDc2kypK/RRcB8bMAUUIG0hLGL
                            82IvnKouGixGqAcULwQKIvTs6uGmlgbSG6Gn5ROb2mlBztXqQ49zRvi5qWNRttir6eyqwRFGOM6A
                            8rxj3Jhxi2Vb/MJn7XzeVHHLzA1sV5hwl/2PLnaL2h9WyG9QwBbwtmkMEqUt/dgixKb1Rvby/tBu
                            RogWgPONNSACiW+Z5o8UdAOqNMZQozD/i1gOjBXoF0F5OksjQN7xoQZLj9xXefxCFQ69FPcFDeEW
                            bHwSoBy5hLPNALaEUoa5zPDwlixwRjFQTc5XXaRpgIjy/2gsL8+Y5QRhyXnLqgO67BlLYW/GuHE=
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </md:KeyDescriptor>
            <md:KeyDescriptor use="encryption">
                <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
                    <ds:X509Data>
                        <ds:X509Certificate>MIIDUjCCAjqgAwIBAgIEUOLIQTANBgkqhkiG9w0BAQUFADBrMQswCQYDVQQGEwJGSTEQMA4GA1UE
                            CBMHVXVzaW1hYTERMA8GA1UEBxMISGVsc2lua2kxGDAWBgNVBAoTD1JNNSBTb2Z0d2FyZSBPeTEM
                            MAoGA1UECwwDUiZEMQ8wDQYDVQQDEwZhcG9sbG8wHhcNMTMwMTAxMTEyODAxWhcNMjIxMjMwMTEy
                            ODAxWjBrMQswCQYDVQQGEwJGSTEQMA4GA1UECBMHVXVzaW1hYTERMA8GA1UEBxMISGVsc2lua2kx
                            GDAWBgNVBAoTD1JNNSBTb2Z0d2FyZSBPeTEMMAoGA1UECwwDUiZEMQ8wDQYDVQQDEwZhcG9sbG8w
                            ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCXqP0wqL2Ai1haeTj0alwsLafhrDtUt00E
                            5xc7kdD7PISRA270ZmpYMB4W24Uk2QkuwaBp6dI/yRdUvPfOT45YZrqIxMe2451PAQWtEKWF5Z13
                            F0J4/lB71TtrzyH94RnqSHXFfvRN8EY/rzuEzrpZrHdtNs9LRyLqcRTXMMO4z7QghBuxh3K5gu7K
                            qxpHx6No83WNZj4B3gvWLRWv05nbXh/F9YMeQClTX1iBNAhLQxWhwXMKB4u1iPQ/KSaal3R26pON
                            UUmu1qVtU1quQozSTPD8HvsDqGG19v2+/N3uf5dRYtvEPfwXN3wIY+/R93vBA6lnl5nTctZIRsyg
                            0Gv5AgMBAAEwDQYJKoZIhvcNAQEFBQADggEBAFQwAAYUjso1VwjDc2kypK/RRcB8bMAUUIG0hLGL
                            82IvnKouGixGqAcULwQKIvTs6uGmlgbSG6Gn5ROb2mlBztXqQ49zRvi5qWNRttir6eyqwRFGOM6A
                            8rxj3Jhxi2Vb/MJn7XzeVHHLzA1sV5hwl/2PLnaL2h9WyG9QwBbwtmkMEqUt/dgixKb1Rvby/tBu
                            RogWgPONNSACiW+Z5o8UdAOqNMZQozD/i1gOjBXoF0F5OksjQN7xoQZLj9xXefxCFQ69FPcFDeEW
                            bHwSoBy5hLPNALaEUoa5zPDwlixwRjFQTc5XXaRpgIjy/2gsL8+Y5QRhyXnLqgO67BlLYW/GuHE=
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </md:KeyDescriptor>
            <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
                                    Location="https://portal.ccctcportal.org/uPortal/saml/SingleLogout"/>
            <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
                                    Location="https://portal.ccctcportal.org/uPortal/saml/SingleLogout"/>
            <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
            <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:transient</md:NameIDFormat>
            <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:persistent</md:NameIDFormat>
            <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
            <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName</md:NameIDFormat>
            <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
                                         Location="https://portal.astuart.co/uPortal/saml/SSO" index="0" isDefault="true"/>
        </md:SPSSODescriptor>
    </md:EntityDescriptor>

    <md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" ID="https://saml2.test.astuart.co/sso/saml2"
      entityID="https://saml2.test.astuart.co/sso/saml2">
      <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="true" WantResponseSigned="true"
        protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
        <md:KeyDescriptor use="signing">
          <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
            <ds:X509Data>
              <ds:X509Certificate>
                MIIDWTCCAkGgAwIBAgIUFs23hkwDtx84sFzaWnFbD6Yb4K8wDQYJKoZIhvcNAQEL
                BQAwHjEcMBoGA1UEAxMTY2EudmF1bHQuYXN0dWFydC5jbzAeFw0xNjA3MDcxODQ5
                MjNaFw0xNzA3MDcxODQ5NTNaMCQxIjAgBgNVBAMTGXNhbWwyLmVuYy1zaWduLmFz
                dHVhcnQuY28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDKKUlqAvwc
                3uUviYKBa+T6J4OV0Dvk1keqxc/iRuFACoAUnsU54seHg2skHUGaWnDvreLEKtTk
                vMWzp00NGcoKQfMJ1+m5XV9OCOOfcxGYKmPpq9Q35Ni3tv3TznNYXQYkl1NYTQtC
                fVtelYEQEuPYa4awm1P8X7+WHxnNIeswkgtlU+YhzxjghoFJzChYzxkl5A8wUCUl
                MBIkNhs6PftCCXlPgblvQ45t2KYebi3x2S3W6ScpDplne7HWsSr1gHsfSMq8qaea
                VWzpVlRnFbVuoYYB4A4tQwjEWPIOqNbSNw4ZPQ7zzI6KHf7CEa5Tckt22RxCOL4Q
                KSJsAblv2hnXAgMBAAGjgYgwgYUwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUF
                BwMCMB0GA1UdDgQWBBQxGrnl0BWf3UKdafCgeWRk+j9hBzAfBgNVHSMEGDAWgBQT
                3WXwhXbSp+X7oWSVF2+K/GCaHjAkBgNVHREEHTAbghlzYW1sMi5lbmMtc2lnbi5h
                c3R1YXJ0LmNvMA0GCSqGSIb3DQEBCwUAA4IBAQCj+/MJI+nOHKSCLHpme9OdpJmE
                Ne6WDr8OvBKQ+dUonaTcV6fmJ5EMtTGvcK+NDAz+/i3MHBdWKUApXB4Voao+QtDx
                CTkZKfePfX/gThXw0v92tdgTcAchFLQ+SBa29iArNXR4K9t37P+xrNQiFf2J0IXP
                /qk+iwjMYmGij3JQfRoL1oeH+uyoOAcmZOBexxE+f7V07N8iaqNJDqc56VCLLqhy
                AUcJ0JMnJFR+LX7LIpwvQ4muESkgSEVllG71Q8Rxj6GAMc7UkgWyrMp9bPlJfTvl
                WI2JSTdNlpoTamOP1VjRwk0SUfCyXoGnMRPuLShkHt2tq7vWM1ff/UuCn5QI
              </ds:X509Certificate>
            </ds:X509Data>
          </ds:KeyInfo>
        </md:KeyDescriptor>
        <md:KeyDescriptor use="encryption">
          <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
            <ds:X509Data>
              <ds:X509Certificate>
                MIIDWTCCAkGgAwIBAgIUFs23hkwDtx84sFzaWnFbD6Yb4K8wDQYJKoZIhvcNAQEL
                BQAwHjEcMBoGA1UEAxMTY2EudmF1bHQuYXN0dWFydC5jbzAeFw0xNjA3MDcxODQ5
                MjNaFw0xNzA3MDcxODQ5NTNaMCQxIjAgBgNVBAMTGXNhbWwyLmVuYy1zaWduLmFz
                dHVhcnQuY28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDKKUlqAvwc
                3uUviYKBa+T6J4OV0Dvk1keqxc/iRuFACoAUnsU54seHg2skHUGaWnDvreLEKtTk
                vMWzp00NGcoKQfMJ1+m5XV9OCOOfcxGYKmPpq9Q35Ni3tv3TznNYXQYkl1NYTQtC
                fVtelYEQEuPYa4awm1P8X7+WHxnNIeswkgtlU+YhzxjghoFJzChYzxkl5A8wUCUl
                MBIkNhs6PftCCXlPgblvQ45t2KYebi3x2S3W6ScpDplne7HWsSr1gHsfSMq8qaea
                VWzpVlRnFbVuoYYB4A4tQwjEWPIOqNbSNw4ZPQ7zzI6KHf7CEa5Tckt22RxCOL4Q
                KSJsAblv2hnXAgMBAAGjgYgwgYUwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUF
                BwMCMB0GA1UdDgQWBBQxGrnl0BWf3UKdafCgeWRk+j9hBzAfBgNVHSMEGDAWgBQT
                3WXwhXbSp+X7oWSVF2+K/GCaHjAkBgNVHREEHTAbghlzYW1sMi5lbmMtc2lnbi5h
                c3R1YXJ0LmNvMA0GCSqGSIb3DQEBCwUAA4IBAQCj+/MJI+nOHKSCLHpme9OdpJmE
                Ne6WDr8OvBKQ+dUonaTcV6fmJ5EMtTGvcK+NDAz+/i3MHBdWKUApXB4Voao+QtDx
                CTkZKfePfX/gThXw0v92tdgTcAchFLQ+SBa29iArNXR4K9t37P+xrNQiFf2J0IXP
                /qk+iwjMYmGij3JQfRoL1oeH+uyoOAcmZOBexxE+f7V07N8iaqNJDqc56VCLLqhy
                AUcJ0JMnJFR+LX7LIpwvQ4muESkgSEVllG71Q8Rxj6GAMc7UkgWyrMp9bPlJfTvl
                WI2JSTdNlpoTamOP1VjRwk0SUfCyXoGnMRPuLShkHt2tq7vWM1ff/UuCn5QI
              </ds:X509Certificate>
            </ds:X509Data>
          </ds:KeyInfo>
        </md:KeyDescriptor>
        <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
          Location="https://portal.ccctcportal.org/uPortal/saml/SingleLogout"/>
        <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
          Location="https://portal.ccctcportal.org/uPortal/saml/SingleLogout"/>
        <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
        <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:transient</md:NameIDFormat>
        <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:persistent</md:NameIDFormat>
        <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
        <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName</md:NameIDFormat>
        <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
          Location="https://saml2.test.astuart.co/sso/saml2" index="0" isDefault="true"/>
      </md:SPSSODescriptor>
    </md:EntityDescriptor>

    <!-- ========================================================================================= -->
    <!-- Locally-run College IDP -->
    <!-- ========================================================================================= -->
    <EntityDescriptor xmlns:ds="http://www.w3.org/2000/09/xmldsig#" xmlns:shibmd="urn:mace:shibboleth:metadata:1.0"
                      xmlns:mdui="urn:oasis:names:tc:SAML:metadata:ui" xmlns="urn:oasis:names:tc:SAML:2.0:metadata"
                      entityID="https://idp.astuart.co/idp/shibboleth">
        <IDPSSODescriptor
                protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol urn:oasis:names:tc:SAML:1.1:protocol urn:mace:shibboleth:1.0">
            <Extensions>
                <!-- <shibmd:Scope regexp="false">example.org</shibmd:Scope> -->
                <!-- <shibmd:Scope regexp="false">test</shibmd:Scope> -->
                <!--
                                    Fill in the details for your IdP here

                                            <mdui:UIInfo>
                                                <mdui:DisplayName xml:lang="en">A Name for the IdP at college.ccctca.edu</mdui:DisplayName>
                                                <mdui:Description xml:lang="en">Enter a description of your IdP at college.ccctca.edu</mdui:Description>
                                                <mdui:Logo height="HeightInPixels" width="WidthInPixels">https://college.ccctca.edu/Path/To/Logo.png</mdui:Logo>
                                            </mdui:UIInfo>

                -->
            </Extensions>
            <KeyDescriptor use="signing">
                <ds:KeyInfo>
                    <ds:X509Data>
                        <ds:X509Certificate>
                            MIIDODCCAiCgAwIBAgIUDPz+OwougAXSuQmKDyAEL46KlPgwDQYJKoZIhvcNAQEL
                            BQAwHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1MB4XDTE1MDYwNDIyMTA0
                            NFoXDTM1MDYwNDIyMTA0NFowHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1
                            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApa0K3OtkHwOnBKSJ7PxT
                            7zry+p8kpu20d+whJs9mHW8w+DikLQ2orLPDZA34Xor0QdR6Y6+gqezIJqqpvuaj
                            YTneQQtXD3neCGD9pPemyF4efEnl21YHryt6Juy6VXIcB6ytHGhmaWg41btdxweD
                            li0b6M7Z6KAW5FjJUoqA+GqFY8rvdm0HZQN+ko4KRK7zTft6ZaPOSbQd7vMtU8bj
                            Msh2XGLWx9G10jvCOFDUbsCNQ3xeFkV30rlUgjb6p2eRUSDWcVPs2Q/FG3t8TVfJ
                            dDtRYps7QW0GDaCPM5hYnlSm+gXwkS8V0j8bGPjv7TfxxK3VMx6okIVsKga7swuZ
                            4QIDAQABo3AwbjAdBgNVHQ4EFgQUT56D4cLSoNxs17FBY+evwXvL2jowTQYDVR0R
                            BEYwRIISY29sbGVnZS5jY2N0Y2EuZWR1hi5odHRwczovL2NvbGxlZ2UuY2NjdGNh
                            LmVkdTo4NDQzL2lkcC9zaGliYm9sZXRoMA0GCSqGSIb3DQEBCwUAA4IBAQBiG0Nw
                            KxslU74tcgjK8CBVahTs5s9Nh2s/En9lP6iWqS2wOHotZ19qqp+AJoIG0pJJpQ6o
                            fRSHdWD2uHmF0F7Uzu1XBxxbV3oG8DmbhzUw2TAOsn0Czt8V30Tfn9U+auNW2XSb
                            z27FACHplll7/T+pycCW6vUcw+boDJIG92TxqIMzlBQOzDGGOTGf/OaKXLb48rWT
                            kEfMv//2Kh735TytX0bJsPmmCLlI9kLcrBNKgHGPNB7oeQNGnYOu+ALxSIugZ7MW
                            LRx2jHND7RSVTetgfEEkkSzsebCxNKMdhIL62Z8VZgYUGD07EeV/3RZ0eV0q5Yf8
                            BhBA6Owk2P264O4R
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </KeyDescriptor>
            <KeyDescriptor use="signing">
                <ds:KeyInfo>
                    <ds:X509Data>
                        <ds:X509Certificate>
                            MIIDODCCAiCgAwIBAgIUQH54kyyeacU69J2iwz9bzeLmMaswDQYJKoZIhvcNAQEL
                            BQAwHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1MB4XDTE1MDYwNDIyMTAz
                            MVoXDTM1MDYwNDIyMTAzMVowHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1
                            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlJhN20ng2VN/cTrWtqUI
                            NaUsrHCkYXbm2y1PTN4b6fJI5hbvcv+LWCuLkLi3+iPGlBpcHHfrdJcyhmBHRHQ9
                            Sos3RIH5Lsn1IgjWe3hxQQmVeEi5xVxnw2YZGHaeX4YnI1TEBJwhtJmyitk74LHy
                            bPGEqOJdApUnLz54L7I+252G/cOfEqUHMbxxtmHSc/9chF8bBxQ8OzIbJsByHnqi
                            awQHwtsttre7n328gVqmf1VHE27cfAYiSjuK5pCsx/1kuJMBN+kg/3Gg9oi6aR50
                            WX1VUF3IBcnTDeiAXRz3PgsT8FlVZou6Ik9NT/Y5IHOZVGk64SRDaG8FuGxLexXr
                            swIDAQABo3AwbjAdBgNVHQ4EFgQUjQwaAoY3u/iToIE3ADeNEW+Uu34wTQYDVR0R
                            BEYwRIISY29sbGVnZS5jY2N0Y2EuZWR1hi5odHRwczovL2NvbGxlZ2UuY2NjdGNh
                            LmVkdTo4NDQzL2lkcC9zaGliYm9sZXRoMA0GCSqGSIb3DQEBCwUAA4IBAQB26rdx
                            phN1YKad3yDhLg6Y1ZwbmAjc+l4QB1KSL+cLqhDn5iMy4VdWh8HpSKRqCwofLtlw
                            3qOwospj+mJaguXRMpjYODRQaKRkTrCGxJhuNrQxDXL/b6FOEIJnUYenbPevuNgR
                            Jc1VnREhWUUXT44KN5YUz9FEiG0BsBK8ecCPKBzTQ/hwaczhpqw6uqVMqxJaTGcn
                            lCUHJAhVHiA8lWJ7vaNPsJ86xBFs/F76EwyFXIKQaruvcvChU7GNNSYdNJBa6HO9
                            9QWdGbr5aNQ4diunnBQdrdjgbQIwyhKTfbFWa2l5vbqEKDc0dwuPa6c25l8ruqxq
                            CQ1CF8ZDDJ0XV6Ab
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </KeyDescriptor>
            <KeyDescriptor use="encryption">
                <ds:KeyInfo>
                    <ds:X509Data>
                        <ds:X509Certificate>
                            MIIDODCCAiCgAwIBAgIUc25IlWI0hNfpkIMwMUn3fpi+GjIwDQYJKoZIhvcNAQEL
                            BQAwHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1MB4XDTE1MDYwNDIyMTAz
                            MloXDTM1MDYwNDIyMTAzMlowHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1
                            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvdr57JQoOxb/UJDuj2oZ
                            jHdbCkIRMf6P/9U9oYSnptWaGOxV7/zuwc78i1JJgSu7ZptPQMn85mSNqCrkhdcX
                            pMkrF98RnT+beB/BK5Y3WqP9z7VEnOkbEXr1BL1toQdU3bedwEcubI9B1cbhqoDW
                            TgWJo4c3dnOWGDBSjnkv6ZXvbPI43FmSXBPNvRfLNLLdPnXoovVkgc0jKEt+PyCc
                            DKEkLNuAD/JoA/RWiM8On8x6K5pPpcEe+yQD6LQPRGaOerBJrh320XTj40ggpDo9
                            UQ5/ywyJmtrg6LVCYs0IXLN5sSScJ34Qc7S92UbSFWByYfABRiAqzcqgVnogVh3Z
                            awIDAQABo3AwbjAdBgNVHQ4EFgQUr3No1lQ65FvD8Serw0cykAQamPwwTQYDVR0R
                            BEYwRIISY29sbGVnZS5jY2N0Y2EuZWR1hi5odHRwczovL2NvbGxlZ2UuY2NjdGNh
                            LmVkdTo4NDQzL2lkcC9zaGliYm9sZXRoMA0GCSqGSIb3DQEBCwUAA4IBAQAE4em2
                            lTVfBrxV9iv8SZC1eVwIWi9KSrRLyX48c/MEDTHUdQL1oZvoy1VXF8HY4mZ1O48W
                            hWAzFQe1gDp6G0Fw48r9olbhlsEj74xH5UztEF0u3SH2i45QGqzx4rOPeSMiT8JE
                            tvravFrqMOYK3cubc/ugIn0AuOIsKWLFkV5aOAbOyv3INbp1FfR3WMEnmOz7kvXo
                            2Paeitjqd2DxjGbrFnpUj/2acLucZ3bJ48vKK5GK22rb3a2aBGG/HRaW4XJdC+NE
                            qnpUhcM0xrtS7rqYaZDvd2CRbZHJaWPi1W+JUIf5gBJ8McvK5JJbKCpa1ujU1Glw
                            lbOTf5UpBgHl9RG0
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </KeyDescriptor>
            <ArtifactResolutionService Binding="urn:oasis:names:tc:SAML:1.0:bindings:SOAP-binding"
                                       Location="https://idp.astuart.co/idp/profile/SAML1/SOAP/ArtifactResolution"
                                       index="1"/>
            <ArtifactResolutionService Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"
                                       Location="https://idp.astuart.co/idp/profile/SAML2/SOAP/ArtifactResolution"
                                       index="2"/>
            <NameIDFormat>urn:mace:shibboleth:1.0:nameIdentifier</NameIDFormat>
            <NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:transient</NameIDFormat>
            <SingleSignOnService Binding="urn:mace:shibboleth:1.0:profiles:AuthnRequest"
                                 Location="https://https://idp.astuart.co/idp/profile/Shibboleth/SSO"/>
            <SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST"
                                 Location="https://ccc-idp-preview-lb-1334668992.us-west-2.elb.amazonaws.com:8443/idp/profile/SAML2/POST/SSO"/>
            <SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST-SimpleSign"
                                 Location="https://https://idp.astuart.co/idp/profile/SAML2/POST-SimpleSign/SSO"/>
            <SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect"
                                 Location="https://https://idp.astuart.co/idp/profile/SAML2/Redirect/SSO"/>
        </IDPSSODescriptor>
        <AttributeAuthorityDescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:1.1:protocol">
            <Extensions>
                <shibmd:Scope regexp="false">example.org</shibmd:Scope>
            </Extensions>
            <KeyDescriptor use="signing">
                <ds:KeyInfo>
                    <ds:X509Data>
                        <ds:X509Certificate>
                            MIIDODCCAiCgAwIBAgIUDPz+OwougAXSuQmKDyAEL46KlPgwDQYJKoZIhvcNAQEL
                            BQAwHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1MB4XDTE1MDYwNDIyMTA0
                            NFoXDTM1MDYwNDIyMTA0NFowHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1
                            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApa0K3OtkHwOnBKSJ7PxT
                            7zry+p8kpu20d+whJs9mHW8w+DikLQ2orLPDZA34Xor0QdR6Y6+gqezIJqqpvuaj
                            YTneQQtXD3neCGD9pPemyF4efEnl21YHryt6Juy6VXIcB6ytHGhmaWg41btdxweD
                            li0b6M7Z6KAW5FjJUoqA+GqFY8rvdm0HZQN+ko4KRK7zTft6ZaPOSbQd7vMtU8bj
                            Msh2XGLWx9G10jvCOFDUbsCNQ3xeFkV30rlUgjb6p2eRUSDWcVPs2Q/FG3t8TVfJ
                            dDtRYps7QW0GDaCPM5hYnlSm+gXwkS8V0j8bGPjv7TfxxK3VMx6okIVsKga7swuZ
                            4QIDAQABo3AwbjAdBgNVHQ4EFgQUT56D4cLSoNxs17FBY+evwXvL2jowTQYDVR0R
                            BEYwRIISY29sbGVnZS5jY2N0Y2EuZWR1hi5odHRwczovL2NvbGxlZ2UuY2NjdGNh
                            LmVkdTo4NDQzL2lkcC9zaGliYm9sZXRoMA0GCSqGSIb3DQEBCwUAA4IBAQBiG0Nw
                            KxslU74tcgjK8CBVahTs5s9Nh2s/En9lP6iWqS2wOHotZ19qqp+AJoIG0pJJpQ6o
                            fRSHdWD2uHmF0F7Uzu1XBxxbV3oG8DmbhzUw2TAOsn0Czt8V30Tfn9U+auNW2XSb
                            z27FACHplll7/T+pycCW6vUcw+boDJIG92TxqIMzlBQOzDGGOTGf/OaKXLb48rWT
                            kEfMv//2Kh735TytX0bJsPmmCLlI9kLcrBNKgHGPNB7oeQNGnYOu+ALxSIugZ7MW
                            LRx2jHND7RSVTetgfEEkkSzsebCxNKMdhIL62Z8VZgYUGD07EeV/3RZ0eV0q5Yf8
                            BhBA6Owk2P264O4R
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </KeyDescriptor>
            <KeyDescriptor use="signing">
                <ds:KeyInfo>
                    <ds:X509Data>
                        <ds:X509Certificate>
                            MIIDODCCAiCgAwIBAgIUQH54kyyeacU69J2iwz9bzeLmMaswDQYJKoZIhvcNAQEL
                            BQAwHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1MB4XDTE1MDYwNDIyMTAz
                            MVoXDTM1MDYwNDIyMTAzMVowHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1
                            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlJhN20ng2VN/cTrWtqUI
                            NaUsrHCkYXbm2y1PTN4b6fJI5hbvcv+LWCuLkLi3+iPGlBpcHHfrdJcyhmBHRHQ9
                            Sos3RIH5Lsn1IgjWe3hxQQmVeEi5xVxnw2YZGHaeX4YnI1TEBJwhtJmyitk74LHy
                            bPGEqOJdApUnLz54L7I+252G/cOfEqUHMbxxtmHSc/9chF8bBxQ8OzIbJsByHnqi
                            awQHwtsttre7n328gVqmf1VHE27cfAYiSjuK5pCsx/1kuJMBN+kg/3Gg9oi6aR50
                            WX1VUF3IBcnTDeiAXRz3PgsT8FlVZou6Ik9NT/Y5IHOZVGk64SRDaG8FuGxLexXr
                            swIDAQABo3AwbjAdBgNVHQ4EFgQUjQwaAoY3u/iToIE3ADeNEW+Uu34wTQYDVR0R
                            BEYwRIISY29sbGVnZS5jY2N0Y2EuZWR1hi5odHRwczovL2NvbGxlZ2UuY2NjdGNh
                            LmVkdTo4NDQzL2lkcC9zaGliYm9sZXRoMA0GCSqGSIb3DQEBCwUAA4IBAQB26rdx
                            phN1YKad3yDhLg6Y1ZwbmAjc+l4QB1KSL+cLqhDn5iMy4VdWh8HpSKRqCwofLtlw
                            3qOwospj+mJaguXRMpjYODRQaKRkTrCGxJhuNrQxDXL/b6FOEIJnUYenbPevuNgR
                            Jc1VnREhWUUXT44KN5YUz9FEiG0BsBK8ecCPKBzTQ/hwaczhpqw6uqVMqxJaTGcn
                            lCUHJAhVHiA8lWJ7vaNPsJ86xBFs/F76EwyFXIKQaruvcvChU7GNNSYdNJBa6HO9
                            9QWdGbr5aNQ4diunnBQdrdjgbQIwyhKTfbFWa2l5vbqEKDc0dwuPa6c25l8ruqxq
                            CQ1CF8ZDDJ0XV6Ab
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </KeyDescriptor>
            <KeyDescriptor use="encryption">
                <ds:KeyInfo>
                    <ds:X509Data>
                        <ds:X509Certificate>
                            MIIDODCCAiCgAwIBAgIUc25IlWI0hNfpkIMwMUn3fpi+GjIwDQYJKoZIhvcNAQEL
                            BQAwHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1MB4XDTE1MDYwNDIyMTAz
                            MloXDTM1MDYwNDIyMTAzMlowHTEbMBkGA1UEAwwSY29sbGVnZS5jY2N0Y2EuZWR1
                            MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvdr57JQoOxb/UJDuj2oZ
                            jHdbCkIRMf6P/9U9oYSnptWaGOxV7/zuwc78i1JJgSu7ZptPQMn85mSNqCrkhdcX
                            pMkrF98RnT+beB/BK5Y3WqP9z7VEnOkbEXr1BL1toQdU3bedwEcubI9B1cbhqoDW
                            TgWJo4c3dnOWGDBSjnkv6ZXvbPI43FmSXBPNvRfLNLLdPnXoovVkgc0jKEt+PyCc
                            DKEkLNuAD/JoA/RWiM8On8x6K5pPpcEe+yQD6LQPRGaOerBJrh320XTj40ggpDo9
                            UQ5/ywyJmtrg6LVCYs0IXLN5sSScJ34Qc7S92UbSFWByYfABRiAqzcqgVnogVh3Z
                            awIDAQABo3AwbjAdBgNVHQ4EFgQUr3No1lQ65FvD8Serw0cykAQamPwwTQYDVR0R
                            BEYwRIISY29sbGVnZS5jY2N0Y2EuZWR1hi5odHRwczovL2NvbGxlZ2UuY2NjdGNh
                            LmVkdTo4NDQzL2lkcC9zaGliYm9sZXRoMA0GCSqGSIb3DQEBCwUAA4IBAQAE4em2
                            lTVfBrxV9iv8SZC1eVwIWi9KSrRLyX48c/MEDTHUdQL1oZvoy1VXF8HY4mZ1O48W
                            hWAzFQe1gDp6G0Fw48r9olbhlsEj74xH5UztEF0u3SH2i45QGqzx4rOPeSMiT8JE
                            tvravFrqMOYK3cubc/ugIn0AuOIsKWLFkV5aOAbOyv3INbp1FfR3WMEnmOz7kvXo
                            2Paeitjqd2DxjGbrFnpUj/2acLucZ3bJ48vKK5GK22rb3a2aBGG/HRaW4XJdC+NE
                            qnpUhcM0xrtS7rqYaZDvd2CRbZHJaWPi1W+JUIf5gBJ8McvK5JJbKCpa1ujU1Glw
                            lbOTf5UpBgHl9RG0
                        </ds:X509Certificate>
                    </ds:X509Data>
                </ds:KeyInfo>
            </KeyDescriptor>
            <AttributeService Binding="urn:oasis:names:tc:SAML:1.0:bindings:SOAP-binding"
                              Location="https://idp.astuart.co/idp/profile/SAML1/SOAP/AttributeQuery"/>
            <AttributeService Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP"
                              Location="https://idp.astuart.co/idp/profile/SAML2/SOAP/AttributeQuery"/>
        </AttributeAuthorityDescriptor>
    </EntityDescriptor>
</EntitiesDescriptor>
`
