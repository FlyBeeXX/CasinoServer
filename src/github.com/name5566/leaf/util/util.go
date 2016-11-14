package util

import (
    "time"
    "crypto/md5"
    "encoding/hex"
    "io"
    "crypto/rand"
    "encoding/base64"
)

func NowUTCTime() (int64){
    return time.Now().UTC().UnixNano() / 1000000;
}

func GetMd5(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

func GetGuid() string {
    b := make([]byte, 48)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return GetMd5(base64.URLEncoding.EncodeToString(b))
}
