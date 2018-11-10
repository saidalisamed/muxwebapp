import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func HoursMinutes(seconds int) string {
	hr := seconds / 3600
	min := (seconds % 3600) / 60
	//sec := seconds % 60;
	//return fmt.Sprintf("%02d:%02d:%02d", hr, min, sec);
	// Adapting for mysql time_zone
	var prefix string
	if hr > 0 {
		prefix = "+"
	}
	return fmt.Sprintf("%s%02d:%02d", prefix, hr, min)
}

func TimeZoneOffset(location string) int {
	t := time.Now()
	loc, _ := time.LoadLocation(location)
	_, offset := t.In(loc).Zone()
	return offset
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
