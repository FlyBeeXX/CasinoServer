package util

import (
    "time"
)

func NowUTCTime() (int64){
    return time.Now().UTC().UnixNano() / 1000000;
}
