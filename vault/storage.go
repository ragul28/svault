package vault

type VaultData struct {
	CreatedTime int64
	Type        string
	EnctyptData []byte
	Version     int
}
