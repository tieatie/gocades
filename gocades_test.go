package gocades

import (
	"testing"
)

// commented tests require certificates installed and can be used as examples
var Cert = "MIINsjCCDJqgAwIBAgIRAPq8ife/MxCUCgAAAAEl/TIwDQYJKoZIhvcNAQELBQAwRjELMAkGA1UEBhMCVVMxIjAgBgNVBAoTGUdvb2dsZSBUcnVzdCBTZXJ2aWNlcyBMTEMxEzARBgNVBAMTCkdUUyBDQSAxQzMwHhcNMjExMTI5MDIyMjMzWhcNMjIwMjIxMDIyMjMyWjAXMRUwEwYDVQQDDAwqLmdvb2dsZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAShwtJ0zDJohmgYDI9a4Sxu+2c8JyYLtfnS/wdyRoIXUchfFuyrWO+bwp1BW6Fkauoqu0LeDXO8oysHN8gba4Vdo4ILkzCCC48wDgYDVR0PAQH/BAQDAgeAMBMGA1UdJQQMMAoGCCsGAQUFBwMBMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFAQOcJ1gEZcBGw5+crGZ5Sj3KeByMB8GA1UdIwQYMBaAFIp0f6+Fze6VzT2c0OJGFPNxNR0nMGoGCCsGAQUFBwEBBF4wXDAnBggrBgEFBQcwAYYbaHR0cDovL29jc3AucGtpLmdvb2cvZ3RzMWMzMDEGCCsGAQUFBzAChiVodHRwOi8vcGtpLmdvb2cvcmVwby9jZXJ0cy9ndHMxYzMuZGVyMIIJQgYDVR0RBIIJOTCCCTWCDCouZ29vZ2xlLmNvbYIWKi5hcHBlbmdpbmUuZ29vZ2xlLmNvbYIJKi5iZG4uZGV2ghIqLmNsb3VkLmdvb2dsZS5jb22CGCouY3Jvd2Rzb3VyY2UuZ29vZ2xlLmNvbYIYKi5kYXRhY29tcHV0ZS5nb29nbGUuY29tggsqLmdvb2dsZS5jYYILKi5nb29nbGUuY2yCDiouZ29vZ2xlLmNvLmlugg4qLmdvb2dsZS5jby5qcIIOKi5nb29nbGUuY28udWuCDyouZ29vZ2xlLmNvbS5hcoIPKi5nb29nbGUuY29tLmF1gg8qLmdvb2dsZS5jb20uYnKCDyouZ29vZ2xlLmNvbS5jb4IPKi5nb29nbGUuY29tLm14gg8qLmdvb2dsZS5jb20udHKCDyouZ29vZ2xlLmNvbS52boILKi5nb29nbGUuZGWCCyouZ29vZ2xlLmVzggsqLmdvb2dsZS5mcoILKi5nb29nbGUuaHWCCyouZ29vZ2xlLml0ggsqLmdvb2dsZS5ubIILKi5nb29nbGUucGyCCyouZ29vZ2xlLnB0ghIqLmdvb2dsZWFkYXBpcy5jb22CDyouZ29vZ2xlYXBpcy5jboIRKi5nb29nbGV2aWRlby5jb22CDCouZ3N0YXRpYy5jboIQKi5nc3RhdGljLWNuLmNvbYIPZ29vZ2xlY25hcHBzLmNughEqLmdvb2dsZWNuYXBwcy5jboIRZ29vZ2xlYXBwcy1jbi5jb22CEyouZ29vZ2xlYXBwcy1jbi5jb22CDGdrZWNuYXBwcy5jboIOKi5na2VjbmFwcHMuY26CEmdvb2dsZWRvd25sb2Fkcy5jboIUKi5nb29nbGVkb3dubG9hZHMuY26CEHJlY2FwdGNoYS5uZXQuY26CEioucmVjYXB0Y2hhLm5ldC5jboILd2lkZXZpbmUuY26CDSoud2lkZXZpbmUuY26CEWFtcHByb2plY3Qub3JnLmNughMqLmFtcHByb2plY3Qub3JnLmNughFhbXBwcm9qZWN0Lm5ldC5jboITKi5hbXBwcm9qZWN0Lm5ldC5jboIXZ29vZ2xlLWFuYWx5dGljcy1jbi5jb22CGSouZ29vZ2xlLWFuYWx5dGljcy1jbi5jb22CF2dvb2dsZWFkc2VydmljZXMtY24uY29tghkqLmdvb2dsZWFkc2VydmljZXMtY24uY29tghFnb29nbGV2YWRzLWNuLmNvbYITKi5nb29nbGV2YWRzLWNuLmNvbYIRZ29vZ2xlYXBpcy1jbi5jb22CEyouZ29vZ2xlYXBpcy1jbi5jb22CFWdvb2dsZW9wdGltaXplLWNuLmNvbYIXKi5nb29nbGVvcHRpbWl6ZS1jbi5jb22CEmRvdWJsZWNsaWNrLWNuLm5ldIIUKi5kb3VibGVjbGljay1jbi5uZXSCGCouZmxzLmRvdWJsZWNsaWNrLWNuLm5ldIIWKi5nLmRvdWJsZWNsaWNrLWNuLm5ldIIOZG91YmxlY2xpY2suY26CECouZG91YmxlY2xpY2suY26CFCouZmxzLmRvdWJsZWNsaWNrLmNughIqLmcuZG91YmxlY2xpY2suY26CEWRhcnRzZWFyY2gtY24ubmV0ghMqLmRhcnRzZWFyY2gtY24ubmV0gh1nb29nbGV0cmF2ZWxhZHNlcnZpY2VzLWNuLmNvbYIfKi5nb29nbGV0cmF2ZWxhZHNlcnZpY2VzLWNuLmNvbYIYZ29vZ2xldGFnc2VydmljZXMtY24uY29tghoqLmdvb2dsZXRhZ3NlcnZpY2VzLWNuLmNvbYIXZ29vZ2xldGFnbWFuYWdlci1jbi5jb22CGSouZ29vZ2xldGFnbWFuYWdlci1jbi5jb22CGGdvb2dsZXN5bmRpY2F0aW9uLWNuLmNvbYIaKi5nb29nbGVzeW5kaWNhdGlvbi1jbi5jb22CJCouc2FmZWZyYW1lLmdvb2dsZXN5bmRpY2F0aW9uLWNuLmNvbYIWYXBwLW1lYXN1cmVtZW50LWNuLmNvbYIYKi5hcHAtbWVhc3VyZW1lbnQtY24uY29tggtndnQxLWNuLmNvbYINKi5ndnQxLWNuLmNvbYILZ3Z0Mi1jbi5jb22CDSouZ3Z0Mi1jbi5jb22CCzJtZG4tY24ubmV0gg0qLjJtZG4tY24ubmV0ghRnb29nbGVmbGlnaHRzLWNuLm5ldIIWKi5nb29nbGVmbGlnaHRzLWNuLm5ldIIMYWRtb2ItY24uY29tgg4qLmFkbW9iLWNuLmNvbYINKi5nc3RhdGljLmNvbYIUKi5tZXRyaWMuZ3N0YXRpYy5jb22CCiouZ3Z0MS5jb22CESouZ2NwY2RuLmd2dDEuY29tggoqLmd2dDIuY29tgg4qLmdjcC5ndnQyLmNvbYIQKi51cmwuZ29vZ2xlLmNvbYIWKi55b3V0dWJlLW5vY29va2llLmNvbYILKi55dGltZy5jb22CC2FuZHJvaWQuY29tgg0qLmFuZHJvaWQuY29tghMqLmZsYXNoLmFuZHJvaWQuY29tggRnLmNuggYqLmcuY26CBGcuY2+CBiouZy5jb4IGZ29vLmdsggp3d3cuZ29vLmdsghRnb29nbGUtYW5hbHl0aWNzLmNvbYIWKi5nb29nbGUtYW5hbHl0aWNzLmNvbYIKZ29vZ2xlLmNvbYISZ29vZ2xlY29tbWVyY2UuY29tghQqLmdvb2dsZWNvbW1lcmNlLmNvbYIIZ2dwaHQuY26CCiouZ2dwaHQuY26CCnVyY2hpbi5jb22CDCoudXJjaGluLmNvbYIIeW91dHUuYmWCC3lvdXR1YmUuY29tgg0qLnlvdXR1YmUuY29tghR5b3V0dWJlZWR1Y2F0aW9uLmNvbYIWKi55b3V0dWJlZWR1Y2F0aW9uLmNvbYIPeW91dHViZWtpZHMuY29tghEqLnlvdXR1YmVraWRzLmNvbYIFeXQuYmWCByoueXQuYmWCGmFuZHJvaWQuY2xpZW50cy5nb29nbGUuY29tghtkZXZlbG9wZXIuYW5kcm9pZC5nb29nbGUuY26CHGRldmVsb3BlcnMuYW5kcm9pZC5nb29nbGUuY26CGHNvdXJjZS5hbmRyb2lkLmdvb2dsZS5jbjAhBgNVHSAEGjAYMAgGBmeBDAECATAMBgorBgEEAdZ5AgUDMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly9jcmxzLnBraS5nb29nL2d0czFjMy9RcUZ4Ymk5TTQ4Yy5jcmwwggEFBgorBgEEAdZ5AgQCBIH2BIHzAPEAdwApeb7wnjk5IfBWc59jpXflvld9nGAK+PlNXSZcJV3HhAAAAX1pt0b0AAAEAwBIMEYCIQC1x9lAp2IZtthi0mxfjtSDXCR714tElKslM61c4kJv/gIhAJPjgGFY5XBnQSaV/jHFfoRcbksMU+lruflT+sfvLLqcAHYAQcjKsd8iRkoQxqE6CUKHXk4xixsD6+tLx2jwkGKWBvYAAAF9abdHmAAABAMARzBFAiEAhzUCqPYG77z0wZUE3y2DmZhTaStBB9BMVBfYSQmYTogCIHbA8hBzj6kjpaw57ZidtP5Dz5Zli2xOJQB+W4s15tI8MA0GCSqGSIb3DQEBCwUAA4IBAQB2aorYRKwUQJI280q1y1Q2Z8c6pem1MWxRX/Ptapmsp1ucrsmu+VaC1YMCSlW1exVieyMhOIvcKDBy/8L0FUEAmsL8fCTroRQ4DnZVLelFk9vmVsiSQtfYHOf6rwEbOrv+94kk964iQCLQFYN5klRqI0hWa3wYe6tnXm/2PvPbAwqsnAq3q+Iek+3pGm6YTshJyA7P9L176psddm6slAYpHOryFcrvXzu1lHSylCAFNT/OYcH1GLTf0qJXuN7YnX9swoYu2oCDkIyAHss2DDp7f8qf0VgDNNxZB8drZ9ID85YA3qgeIbHHAB8UIj8qkKXkmybfMxVL0lz3iY73asSm"
var Crl = "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"

func TestAbout(t *testing.T) {
	about, _ := About()
	about.GetVersion()
	about.GetMajorVersion()
	about.GetMinorVersion()
	about.GetBuildVersion()
	about.GetCSPVersion("", 80)
	about.GetCSPName(80)
	about.MediaFilter(MEDIA_TYPE_SCARD)
	about.ReaderFilter(ENABLE_ANY_CARRIER_TYPE, DISABLE_EVERY_CARRIER_OPERATION, ".*rutoken.*")
}

func TestVersion(t *testing.T) {
	about, _ := About()
	version, _ := about.GetCSPVersion("", 80)
	version.ToString()
	version.GetMajorVersion()
	version.GetMinorVersion()
	version.GetBuildVersion()
}

func TestAttribute(t *testing.T) {
	attribute, _ := Attribute()
	attribute.PutValue("ABCD")
	attribute.GetValue()
	attribute.PutName(2)
	attribute.GetName()
	attribute.PutValueEncoding(1)
	attribute.GetValueEncoding()
	attribute.PutOID("1.2.3.3.5")
	attribute.GetOID()
}

func TestOID(t *testing.T) {
	attribute, _ := Attribute()
	oid, _ := attribute.GetOID()
	oid.SetFriendlyName("ГОСТ Р 34.10-2001")
	oid.GetFriendlyName()
	oid.GetValue()
	oid.SetValue("1.2.643.7.1.1.2.2")
	oid.GetFriendlyName()
	oid.SetName(1)
	oid.GetName()
	oid.GetValue()
}

func TestLicense(t *testing.T) {
	license, _ := License()
	license.GetCompanyName(0)
	license.GetFirstInstallDate(0)
	license.GetSerialNumber(0)
	license.GetType(0)
	license.GetValidTo(0)
}

func TestCRL(t *testing.T) {
	crl, _ := CRL()
	crl.Import(Crl)
	crl.GetIssuerName()
	crl.GetThisUpdate()
	crl.GetNextUpdate()
	crl.GetThumbprint()
	crl.GetAuthKeyID()
	crl.Export(CADESCOM_ENCODE_BASE64)
}

func TestCertificate(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	certificate.Export(CADESCOM_ENCODE_BASE64)
	certificate.GetBasicConstraints()
	certificate.GetExtendedKeyUsage()
	certificate.GetKeyUsage()
	certificate.GetPublicKey()
	certificate.GetPrivateKey()
	certificate.IsValid()
	certificate.GetInfo(CAPICOM_CERT_INFO_SUBJECT_SIMPLE_NAME)
	certificate.HasPrivateKey()
	certificate.FindPrivateKey()
	certificate.GetIssuerName()
	certificate.GetSerialNumber()
	certificate.GetSubjectName()
	certificate.GetThumbprint()
	certificate.GetValidToDate()
	certificate.GetValidFromDate()
	certificate.GetVersion()
}

func TestCertificateStatus(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	certstatus, _ := certificate.IsValid()
	certstatus.GetResult()
}

func TestBasicConstraints(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	result, _ := certificate.GetBasicConstraints()
	result.GetIsCertificateAuthority()
	result.GetIsCritical()
	result.GetIsPathLenConstraintPresent()
	result.GetIsPresent()
	result.GetPathLenConstraint()
}

func TestExtendedKeyUsage(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	extkeyusage, _ := certificate.GetExtendedKeyUsage()
	extkeyusage.GetIsCritical()
	extkeyusage.GetIsPresent()
	extkeyusage.GetEKUs()
}

func TestEKUs(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	extkeyusage, _ := certificate.GetExtendedKeyUsage()
	ekus, _ := extkeyusage.GetEKUs()
	ekus.GetCount()
	ekus.GetItem(1)
}

func TestEKU(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	extkeyusage, _ := certificate.GetExtendedKeyUsage()
	ekus, _ := extkeyusage.GetEKUs()
	eku, _ := ekus.GetItem(1)
	eku.GetName()
	eku.GetOID()
	eku.SetName(2)
	eku.GetName()
	eku.GetOID()
	eku.PutOID("1.2.8.4")
	eku.GetOID()
}

func TestKeyUsage(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	keyusage, _ := certificate.GetKeyUsage()
	keyusage.GetIsCRLSignEnabled()
	keyusage.GetIsCritical()
	keyusage.GetIsDataEnciphermentEnabled()
	keyusage.GetIsDecipherOnlyEnabled()
	keyusage.GetIsDigitalSignatureEnabled()
	keyusage.GetIsEncipherOnlyEnabled()
	keyusage.GetIsKeyAgreementEnabled()
	keyusage.GetIsKeyCertSignEnabled()
	keyusage.GetIsKeyEnciphermentEnabled()
	keyusage.GetIsNonRepudationEnabled()
	keyusage.GetIsPresent()
}

func TestPublicKey(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	publickey, _ := certificate.GetPublicKey()
	publickey.GetAlgorithm()
	publickey.GetEncodedKey()
	publickey.GetEncodedParameters()
	publickey.GetLength()
}

func TestEncodedData(t *testing.T) {
	certificate, _ := Certificate()
	certificate.Import(Cert)
	publickey, _ := certificate.GetPublicKey()
	encodeddata, _ := publickey.GetEncodedKey()
	encodeddata.Format(true)
	encodeddata.GetValue(CADESCOM_ENCODE_BASE64)
}

func TestStore(t *testing.T) {
	store, _ := Store()
	store.Open(CADESCOM_MEMORY_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	defer store.Close()
	store.GetCertificates()
	store.GetLocation()
	store.GetName()
	cert, _ := Certificate()
	cert.Import(Cert)
	store.Add(*cert)
	crl, _ := CRL()
	crl.Import(Crl)
	store.AddCRL(*crl)
	store.Close()
}

func TestCertificates(t *testing.T) {
	store, _ := Store()
	store.Open(CADESCOM_CURRENT_USER_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	certificates, _ := store.GetCertificates()
	certificates.GetCount()
}

func TestHashedData(t *testing.T) {
	hashed, _ := HashedData()
	hashed.PutAlgorithm(CADESCOM_HASH_ALGORITHM_CP_GOST_3411_2012_256)
	hashed.GetAlgorithm()
	hashed.PutDataEncoding(1)
	hashed.GetDataEncoding()
	hashed.Hash("data to hash")
	hashed.GetValue()
	hashed.SetHashValue("DEADBEEFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEE")
	hashed.GetValue()
	hashed.PutAlgorithm(CADESCOM_HASH_ALGORITHM_CP_GOST_3411_2012_256_HMAC)
	hashed.PutKey("DEADBEEF6ED7D38DF2D4356807BABA09F91325C6C705AC01F27274B952FFEABEEBFEAAAA")
	hashed.GetKey()
}

func TestSigner(t *testing.T) {
	signer, _ := Signer()
	signer.GetAuthenticatedAttributes2()
	signer.GetCertificate()
	signer.PutOptions(CAPICOM_CERTIFICATE_INCLUDE_CHAIN_EXCEPT_ROOT)
	signer.GetOptions()
	signer.GetCRLs()
	signer.GetOCSPResponses()
	signer.PutTSAAddress("http://www.cryptopro.ru/tsp/tsp.srf")
	signer.GetTSAAddress()
	signer.GetUnauthenticatedAttributes()
	signer.GetSignatureTimestampTime()
	signer.GetSigningTime()
	signer.GetSignatureStatus()
	signer.PutCheckCertificate(true)
	signer.GetCheckCertificate()
}

func TestBlobs(t *testing.T) {
	signer, _ := Signer()
	crls, _ := signer.GetCRLs()
	crls.GetCount()
}

func TestSignatureStatus(t *testing.T) {
	signer, _ := Signer()
	status, _ := signer.GetSignatureStatus()
	status.IsValid()
}

func TestEnvelopedData(t *testing.T) {
	enveloped, _ := EnvelopedData()
	enveloped.PutContent("content to encrypt")
	enveloped.PutContentEncoding(CADESCOM_STRING_TO_UCS2LE)
	enveloped.GetContentEncoding()
	enveloped.GetRecipients()

	//store, _ := Store()
	//store.Open(CADESCOM_CURRENT_USER_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	//defer store.Close()
	//certificates, _ := store.GetCertificates()
	//cert, _ := certificates.GetItem(1)
	//recipients.Add(*cert)

	//encrypted, _ := enveloped.Encrypt(CADESCOM_ENCODE_BASE64)
	//encrypted

	//enveloped2, _ := EnvelopedData()
	//enveloped2.Decrypt(encrypted)
	//enveloped2.GetContent()
	//enveloped2.GetAlgorithm()

	//enveloped3, _ := EnvelopedData()
	//recipients3, _ := enveloped3.GetRecipients()
	//recipients3.Add(*cert)
	//enc1, _ := enveloped3.StreamEncrypt("AAAA", false)
	//enc2, _ := enveloped3.StreamEncrypt("ABBABAAB", false)
	//enc3, _ := enveloped3.StreamEncrypt("BASDBAAA", true)
	//enc1, enc2, enc3

	//enveloped4, _ := EnvelopedData()
	//recipients4, _ := enveloped4.GetRecipients()
	//recipients4.Add(*cert)
	//enveloped4.StreamDecrypt(enc1, false)
	//enveloped4.StreamDecrypt(enc2, false)
	//enveloped4.StreamDecrypt(enc3, true)
}

func TestRecipients(t *testing.T) {
	//store, _ := Store()
	//store.Open(CADESCOM_CURRENT_USER_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	//defer store.Close()
	//certificates, _ := store.GetCertificates()
	//cert, _ := certificates.GetItem(1)
	//enveloped, _ := EnvelopedData()
	//recipients, _ := enveloped.GetRecipients()
	//recipients.Add(*cert)
	//recipients.GetCount()
	//recipients.GetItem(1)
}

func TestAlgorithm(t *testing.T) {
	//store, _ := Store()
	//store.Open(CADESCOM_CURRENT_USER_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	//defer store.Close()
	//certificates, _ := store.GetCertificates()
	//cert, _ := certificates.GetItem(1)
	//enveloped, _ := EnvelopedData()
	//recipients, _ := enveloped.GetRecipients()
	//recipients.Add(*cert)
	//enveloped.PutContent("content to encrypt")
	//encrypted, _ := enveloped.Encrypt(CADESCOM_ENCODE_BASE64)

	//enveloped2, _ := EnvelopedData()
	//enveloped2.Decrypt(encrypted)
	//algorithm, _ := enveloped2.GetAlgorithm()
	//algorithm.GetKeyLength()
	//algorithm.GetName()
}

func TestSignedData(t *testing.T) {
	//store, _ := Store()
	//store.Open(CADESCOM_CURRENT_USER_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	//defer store.Close()
	//certificates, _ := store.GetCertificates()
	//cert, _ := certificates.GetItem(2)
	//signer, _ := Signer()
	//signer.PutCertificate(*cert)

	//signeddata, _ := SignedData()
	//signeddata.PutContent("content to sign")
	//signature, err := signeddata.SignCades(*signer, CADESCOM_CADES_BES, false, CADESCOM_ENCODE_BASE64)
	//signature, err

	//signeddata2, _ := SignedData()
	//signeddata2.VerifyCades(signature, CADESCOM_CADES_BES, false)
}

func TestRawSignature(t *testing.T) {
	//store, _ := Store()
	//store.Open(CADESCOM_CURRENT_USER_STORE, "My", CAPICOM_STORE_OPEN_MAXIMUM_ALLOWED)
	//defer store.Close()
	//certificates, _ := store.GetCertificates()
	//cert, _ := certificates.GetItem(2)

	//hashed, _ := HashedData()
	//hashed.PutAlgorithm(CADESCOM_HASH_ALGORITHM_CP_GOST_3411)
	//hashed.SetHashValue("DEADBEEFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABEE")
	//rs, _ := RawSignature()
	//hash, err := rs.SignHash(*hashed, *cert)

	//rs2, _ := RawSignature()
	//rs2.VerifyHash(*hashed, *cert, hash)
}

func TestSignedXML(t *testing.T) {
	//signedxml, _ := SignedXML()
	//signedxml.GetSigners()
	//signedxml.PutContent("test")
	//signedxml.GetContent()
}

func TestSymmetricAlgorithm(t *testing.T) {
	symalg, _ := SymmetricAlgorithm()
	symalg.PutIV("12AB34CD56EFBCE5")
	symalg.GetIV()
	symalg.PutDiversData("D29429ADB833913892938412233339520135829A9A4839B9DDDB38992128AAAA28841838295001BB")

}
