/*
Copyright 2021 The Pixiu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
 @Version : 1.0
 @Author  : steven.wang
 @Email   : 'wangxk1991@gamil.com'
 @Time    : 2022/2022/14 14/32/15
 @Desc    :
*/

package cipher

import "testing"

var (
	kb1 = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM2VENDQWRHZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRJek1EZ3pNREF6TkRRek9Gb1lEekl4TWpNd09EQTJNRE0wTkRNNFdqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCnVjUkVEK3pTVTdUNmtRYjc2aWNuVWhLM3BTbUEyMVl5VWRlUUtnb3dFMHZIb3RIU1hBTnJoRnp0ZmdjdzIxT08KZGdoQ2FOYTdYc1MrY2tVY0p0Qjl5VmZqUUZ0ei9RbnlQc0RDdVVaZ0EyTDdNeldkK3NxbVRhQUovRllsVlZzLwpGeS9tRWg5eEdNRzdSOWltNkpDVUZnNkhvMUtQdkRzSGdNUEdlc1dhSEl1YVJHMjZDUE9EMlhjMUxaTU5LRTQ1CkZBUTJpRDIydkFDbUlrMEpvd2NjWnFad3JJMFRMc2kwZVBXeGR4anhHNjZVajM4L3Q2dGJCZy9HSUt6SGlvTUUKcnJKUmxsemdmZ1YzbHpKbE03cXV0QldtemJuR3I4RGZFVzR6RFVwTlIweUN1YjZhZUNzbFllbWxZMXIzWmdQZQp0N1h6RDZMRVVPaGNuY1drSE1USWxRSURBUUFCbzBJd1FEQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVqeFRvbTRDUThReTY1V1JWY2lFV1c2MlNnSzh3RFFZSktvWkkKaHZjTkFRRUxCUUFEZ2dFQkFKV3Q4VTMvN1VpQWFPNDF0UHdDOVJTaHIzc3FSaHVvYjlEdERoaHZsbng2SE80VwpkOEJWQXJLK1dPTlZVcWkwTDRhZWUzQjNvV2g5andoYVl1TkRKWXhXNnlyUWtGR25POGVCT1pXSjMwbWEyajYzCm80MU9LbXp6VVRXbTgwTUtqODl2dGVlamdsQUNMRDROWUw0Nk9UL1dDSjhsUXpOVUtKQ2FWS2o0bXlzcGl6OVgKMUdkSitvWXhUSDQyNTJBZkJGUzhDcWdkYnpyY2w5c2s4ckh0NVVjN3NFYmgvcHJYTmdKVzRUOElTb0x2Ujk4eAo0T1FMeEZLNWFuaURkTVJkUlQ3WnQzRm5YaFd3bXl6ZGVKbmtoazdDemNoVm90WmhSaFNXa20zUjEySDVudkRqCk43R3dhQ2tBbkVkakd5K1czUkpWdTduNGdhc3pyQXpFNEY1RGk4TT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: https://120.48.100.85:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lJTHpoR1N5NjZaS3d3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWdGdzB5TXpBNE16QXdNelEwTXpoYUdBOHlNVEl6TURnd05qQXpOVEF3TWxvdwpOREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhHVEFYQmdOVkJBTVRFR3QxWW1WeWJtVjBaWE10CllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRRHFVMWQ0SFByZzlOeEIKcUpzT1JOeGlGTnBXamNlMW1sRTR2TjgxTjJnMldoTzlQYmVYaVJYcUhZbktCVmJocGhodWxQL3hhKzFsUWRCYQpPU1lmM3ZjQ1NXazFlcUlJa0g2VEFuYVZSQm1xTzU5ZUZqYUV2bVJqYVZUUi9JM3BZTlhtOGliVXBDd2hrd0Z2Cm42dDB3WGJ5Z1MvVERJbFdQZUNic3R1WThKSTlMTXF3aFRydVlKVG1CZkZQeXVrQmNZWHRnQXNOYUJiSWNGaFQKdVR5TGE1Q0hlZDFJNWRwaWFTRWFZUm1WVEljOEVvTGF4bC8reXB3YzFPNWhYVDhqdHcyTngxSzlMSGFpU3pnZwpNN0E3WGxwODhpMkhodERseGMvTEFrd01zWmpsZVFHM2RMNkNDNWd3R1QvZG4zODh5T2FERXZzNHAxUUhQeVFaCmZBa284eXkzQWdNQkFBR2pTREJHTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQWZCZ05WSFNNRUdEQVdnQlNQRk9pYmdKRHhETHJsWkZWeUlSWmJyWktBcnpBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFRejJ3NlNLZTdvSXlDQmdRWmo3M3podHZKNVZuUnd1QkVGYW9zS0orWGpsZ1paNml2SGU5CmRLM3dqK0xHR2xpTnozRUM2czcyb0FHSmY4NUpwcHdkR1lpVGFIcll3VjNaclVSaVEwUW9ZcHlJZE90dmJ4ZEEKNC9YTjlhUElXQThtbXVQVnA4R1NzbGZaN1JYRW12SjVWZWxOMkNLWEQxbFFIeElzSjRiMm5SUDE3NHhnUE5QeQpQT2xrWWhJUWprcTFybEk4TXEwYmNUbjRhTjBCZVFFU0RUb3RkenhMamNEbkc1S0I5bERmdVRBZ2FIMHRUYVN4ClVDUi9lcjhnSitNaXR4RWhkUmZwcHVkWHhHT0M4MGZQRDNNOWhiNllsdWtmREJNOHd3NWdmcmVaZENhNUlrbTEKd2RQS2tsSUk0Yk5KVnlQRUV5emdIN0tCWHlSckhyR0dBdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBNmxOWGVCejY0UFRjUWFpYkRrVGNZaFRhVm8zSHRacFJPTHpmTlRkb05sb1R2VDIzCmw0a1Y2aDJKeWdWVzRhWVlicFQvOFd2dFpVSFFXamttSDk3M0FrbHBOWHFpQ0pCK2t3SjJsVVFacWp1ZlhoWTIKaEw1a1kybFUwZnlONldEVjV2SW0xS1FzSVpNQmI1K3JkTUYyOG9FdjB3eUpWajNnbTdMYm1QQ1NQU3pLc0lVNgo3bUNVNWdYeFQ4cnBBWEdGN1lBTERXZ1d5SEJZVTdrOGkydVFoM25kU09YYVlta2hHbUVabFV5SFBCS0Myc1pmCi9zcWNITlR1WVYwL0k3Y05qY2RTdlN4Mm9rczRJRE93TzE1YWZQSXRoNGJRNWNYUHl3Sk1ETEdZNVhrQnQzUysKZ2d1WU1Cay8zWjkvUE1qbWd4TDdPS2RVQno4a0dYd0pLUE1zdHdJREFRQUJBb0lCQUVVdWgvT1Jjc3g1M3BHUQpDdjlpZTBLRjc1U0FWWWNiWVlHRDJjNm5aaUF0NWdnSU1hME9yeWFvdythZVB2RnVEV2FRblNRVS9ralg3b0RFCndLbkdOVjA2ZVZKdHUyc1Y1aGpYTUZXUGZwaFE5Qk5lTVlodGdjd0w3QnYzR1pyZW1TaVBNOEd5R3NscTBwQzMKTG5MUkw2SVhYQnhPWmpuTkdRUS9BVENkVGorNjZmbENIL0k2c3o1VU41cVk0OE50NytiY1ErSjNBVHQ1bHk3WApBaEpWR2U5VGU5dmExYkc2bVFrVkVCWmF4ZUJGZllJS1VxWmZyYzhSTndhd1o3cksyQ2tLUHZJQVJ2M0R3MTRuCno4L0pDOTMrNENJdFFjWjV1Nks2Nk8zc2hMdGhLMjdIVDYvanlaNUVsc1BHWnZ2dC9sNXNXdFZ6cGRGUStoUlQKTlp4SGR1a0NnWUVBLzA5V3I3RnhVQmwwNE9BS0E2UXNhcDlaaHc2SnFOQjhGZUloUW5uaEhMZ0k1bHNFSWJZVAplWlZyajQrMW4wcTlOZDJJU0JqbVlSOVQrK0s0M2RrWkxCaW5LdzZSQ0JsNEpDRkMydXJ6ejB2M2lVRFdXcWc5ClNTK0UxcUdvRUludDZNd2tCVG1rVlFhYU5JRXA1N3JyeGlZZ1BZU2J5b0tGZnZ4YWdYYm1YSFVDZ1lFQTZ2VjcKb3VCa0VBK3hZemtrNEJGZW44eTlPcnBQSG54VVIybXFmTmxjZFlaUFZHMEJoVWNoSWJ3MWhUYXRWRFZtYkpPMwprWEJZbEJOY1JVZTZxd3krR1JTck1qay9qdVZuRm5VMURRd0hYdzJrdFBXVWlGZnI5MGxwdjk2Rmo3Y1dheXpGClVjTjBsNVlRZGt5OU5tcVQveUtvVjNXWENhVUdJRXpxa2gxanJ2c0NnWUJZYkhnUGZadExMS1hvUkcvemR1YlQKUHNGWkgxWGFHK3IrQVVXWHFuS202YTViYWVaeWJvY3NNdTMvMkx3RW9YUUt0ZUUvWnpCVEx2QjlhcmF3VDBLQwpuVDg0a0dEMnR5dC9CKzRKdUJvZEQ2aDU1dlRXalRJOHFMS1BTNG5Ud0pHZmNTU0Frc2d4a29uUjZYSkVDZDhpCno0MlhBTmljT0x4WGtMc2YzN2hGQ1FLQmdRQ3V5RTVncTJ2RVJ1ZytWNVFHMnVUNUd1VHBldzdBTHR6REF1ZkEKV0RGZmU3b25vNjh0NWV2UG5LSnpwbWpwNjQ5MTFhMElrL3BRck5aRzgwVDVXVkZRZHYrakxOTUxDYnZXQ1BxZAo2Z1RSd09SMWVQbzFNckptRWthSVpKQVdpS3cwcmVsYkdYcnlDTC9XUU4vQVN5Z3M4eXNuWkdJbTJvdmtHMU9qCnZoa1pZUUtCZ1FDb21oZkNsWEtZUXMxRE1LZkFJOEVLaG83c1lUTjVaNWhXWkZMQXZVM1d2OG1QcnRXb2l2cXIKditGci9DaTQrOGJJOHBreWIrbGhZVzVHMldlUDBMdWRNY2wwT3Z0Ti8vRDNlWXFBOEdTZkU5OTBMY1VKQ25rawptTHJOMmlNMzdLZkNRU0tuVEl6YS9qM3FSKzVReWliNWtuYWYxa1lCT3pIOXRGL2ozcHRWR0E9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=`
	kb2 = `bbbbbbb`
	kb3 = `ccccccc`
)

func TestEncrypt(t *testing.T) {
	cases := []struct {
		Name string
		text []byte
	}{
		{"a", []byte(kb1)},
		{"b", []byte(kb2)},
		{"c", []byte(kb3)},
		{"d", []byte("p4aNIOGsMUTlRKGWrwdefrnKBSujVF")},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans, err := Encrypt(c.text, "uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl"); err != nil {
				t.Fatalf("encrypt text %s failed: %+v",
					c.text, err)
			} else {
				t.Logf("encrypt text: { %s }", ans)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	cases := []struct {
		Name string
		text string
	}{
		{"a", `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`},
		{"b", "RR3jDsJrut4Lz1ffcSYMzw=="},
		{"c", "ho1s5qTr+11w9t/9W/d6yg=="},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans, err := Decrypt(c.text, "uOvKLmVfztaXGpNYd4Z0I1SiT7MweJhl"); err != nil {
				t.Fatalf("decrypt text %s failed: %+v",
					c.text, err)
			} else {
				t.Logf("decrypt is %s", ans)
			}
		})
	}
}