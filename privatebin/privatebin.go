package privatebin

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gearnode/base58"
	"golang.org/x/crypto/pbkdf2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	URL      url.URL
	Username string
	Password string
}

type CreatePasteRequest struct {
	V     int                    `json:"v"`
	AData []interface{}          `json:"adata"`
	Meta  CreatePasteRequestMeta `json:"meta"`
	CT    string                 `json:"ct"`
}

type CreatePasteRequestMeta struct {
	Expire string `json:"expire"`
}

type CreatePasteResponse struct {
	ID          string `json:"id"`
	Status      int    `json:"status"`
	URL         string `json:"url"`
	DeleteToken string `json:"deletetoken"`
}

type PasteSpec struct {
	IV          string
	Salt        string
	Iterations  int
	KeySize     int
	TagSize     int
	Algorithm   string
	Mode        string
	Compression string
}

func (ps *PasteSpec) SpecArray() []interface{} {
	return []interface{}{
		ps.IV,
		ps.Salt,
		ps.Iterations,
		ps.KeySize,
		ps.TagSize,
		ps.Algorithm,
		ps.Mode,
		ps.Compression,
	}
}

func NewClient(uri *url.URL, username, password string) *Client {
	return &Client{URL: *uri, Username: username, Password: password}
}

type PasteContent struct {
	Paste string `json:"paste"`
}

func (c *Client) CreatePaste(message string) (*CreatePasteResponse, error) {
	masterKey, err := generateRandomBytes(32)
	if err != nil {
		return nil, fmt.Errorf("cannot generate random bytes: %w", err)
	}

	pasteContent, err := json.Marshal(&PasteContent{Paste: message})
	if err != nil {
		return nil, fmt.Errorf("cannot marshal paste content: %w", err)
	}

	pasteData, err := encrypt(masterKey, pasteContent)
	if err != nil {
		return nil, fmt.Errorf("cannot encrypt data: %w", err)
	}

	createPasteReq := &CreatePasteRequest{
		V:     2,
		AData: pasteData.adata(),
		Meta:  CreatePasteRequestMeta{Expire: "1week"},
		CT:    base64.RawStdEncoding.EncodeToString(pasteData.Data),
	}

	body, err := json.Marshal(createPasteReq)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal paste request: %w", err)
	}

	req, err := http.NewRequest("POST", c.URL.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header.Set("X-Requested-With", "JSONHttpRequest")
	req.SetBasicAuth(c.Username, c.Password)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot execute http request: %w", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read response body: %w", err)
	}

	pasteResponse := CreatePasteResponse{}
	err = json.Unmarshal(resBody, &pasteResponse)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal response: %w", err)
	}

	pasteId, err := url.Parse(pasteResponse.URL)
	if err != nil {
		return nil, fmt.Errorf("cannot parse paste url: %w", err)
	}

	var uri url.URL
	uri.Scheme = c.URL.Scheme
	uri.Host = c.URL.Host
	uri.RawQuery = pasteId.RawQuery
	uri.Fragment = base58.Encode(masterKey)

	pasteResponse.URL = uri.String()

	return &pasteResponse, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

type PasteData struct {
	*PasteSpec
	Data []byte
}

func (p *PasteData) adata() []interface{} {
	return []interface{}{
		p.SpecArray(),
		"plaintext",
		0,
		0,
	}
}

func encrypt(masterKey []byte, message []byte) (*PasteData, error) {
	iv, err := generateRandomBytes(12)
	if err != nil {
		return nil, err
	}

	salt, err := generateRandomBytes(8)
	if err != nil {
		return nil, err
	}

	paste := &PasteData{
		PasteSpec: &PasteSpec{
			IV:          base64.RawStdEncoding.EncodeToString(iv),
			Salt:        base64.RawStdEncoding.EncodeToString(salt),
			Iterations:  100000,
			KeySize:     256,
			TagSize:     128,
			Algorithm:   "aes",
			Mode:        "gcm",
			Compression: "none",
		},
	}

	key := pbkdf2.Key(masterKey, salt, paste.Iterations, 32, sha256.New)

	adata, err := json.Marshal(paste.adata())
	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	data := gcm.Seal(nil, iv, message, adata)

	paste.Data = data

	return paste, nil
}
