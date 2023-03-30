package security

import "testing"

func TestAESDecrypt(t *testing.T) {
	deStr, err := AESDecrypt("VEmHXzPPPKzYZtY2w0tYHM3OGMA7jPZn50BermJ37Jc=", "KEY_SiteOL_Stone")
	t.Log(deStr)
	t.Log(err)
}

func TestAESEncrypt(t *testing.T) {
	enStr, err := AESEncrypt("AdminAdminAdmin1", "KEY_SiteOL_Stone")
	t.Log(enStr)
	t.Log(err)
}
