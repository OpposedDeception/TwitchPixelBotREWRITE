package pixeldata

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

const canvasWidth = 1000
const canvasHeight = 1000
const apiBaseURL = "https://api.twitch.tv"
const oauthTokenURL = "https://id.twitch.tv/oauth2/token"
const placeURL = "https://twitch.tv/place"

var (
    httpClient *http.Client
)

func init() {
    httpClient = &http.Client{}
}

type wrongTokenError struct {
    message string
}

func (e *wrongTokenError) Error() string {
    return e.message
}

type pixelNotPlacedError struct {
    message string
}

func (e *pixelNotPlacedError) Error() string {
    return e.message
}
