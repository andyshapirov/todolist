package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andyshapirov/todolist/internal/database"
	"github.com/golang-jwt/jwt"
)

const Layout = "20060102"

func CreateJWT(pass, secret string) (string, error) {
	hash := sha256.Sum256([]byte(pass))
	hexHash := hex.EncodeToString(hash[:])

	claims := jwt.MapClaims{
		"password_hash": hexHash,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Auth(pass, secret string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(pass) > 0 {
			var jwt string

			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			var valid bool
			if token, err := CreateJWT(pass, secret); token == jwt && err == nil {
				valid = true
			}

			if !valid {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func NextDate(now time.Time, date string, repeat string) (string, error) {
	now, _ = time.Parse(Layout, now.Format(Layout))

	startDate, err := time.Parse(Layout, date)
	if err != nil {
		return "", err
	}

	if len(repeat) == 0 {
		return "", errors.New("empty repeat")
	}

	r := strings.Split(repeat, " ")

	switch true {
	case r[0] == "d" && len(r) == 2:
		d, err := strconv.Atoi(r[1])
		if err != nil {
			return "", errors.New("invalid repeat format")
		}
		if d < 1 || d > 400 {
			return "", errors.New("invalid repeat format")
		}

		diff := int(now.Sub(startDate).Hours()) / 24
		n := 1
		if diff > 0 {
			n += diff / d
		}

		return startDate.AddDate(0, 0, d*n).Format(Layout), nil
	case r[0] == "y" && len(r) == 1:
		if now.Before(startDate) || now.Equal(startDate) {
			return startDate.AddDate(1, 0, 0).Format(Layout), nil
		}

		diff := now.Year() - startDate.Year()
		if now.After(startDate.AddDate(diff, 0, 0)) {
			diff++
		}

		return startDate.AddDate(diff, 0, 0).Format(Layout), nil
	case r[0] == "w" && len(r) == 2:
		wd := strings.Split(r[1], ",")
		if len(wd) > 7 {
			return "", errors.New("invalid repeat format")
		}

		if now.After(startDate) {
			startDate = now
		}

		startWd := int(startDate.Weekday())
		if startWd == 0 {
			startWd = 7
		}
		diff := 7

		var iDiff int
		for _, v := range wd {
			iWd, err := strconv.Atoi(v)
			if err != nil {
				return "", errors.New("invalid repeat format")
			}
			if iWd < 1 || iWd > 7 {
				return "", errors.New("invalid repeat format")
			}

			iDiff = iWd - startWd

			if iDiff <= 0 {
				iDiff += 7
			}

			if diff > iDiff {
				diff = iDiff
			}
		}

		return startDate.AddDate(0, 0, diff).Format(Layout), nil
	case r[0] == "m" && len(r) == 2:
		md := strings.Split(r[1], ",")
		if len(md) > 33 {
			return "", errors.New("invalid repeat format")
		}

		if now.After(startDate) {
			startDate = now
		}

		startMd := startDate.Day()
		startMm := int(startDate.Month())
		diff := 61

		var iDiff int
		for _, sMd := range md {
			iMd, err := strconv.Atoi(sMd)
			if err != nil {
				return "", errors.New("invalid repeat format")
			}
			if iMd < -2 || iMd > 31 || iMd == 0 {
				return "", errors.New("invalid repeat format")
			}

			if iMd < 0 {
				iMd++

				iDiff = startDate.AddDate(0, 1, iMd-startMd).Day() - startMd

				if iDiff <= 0 {
					iDiff = int(startDate.AddDate(0, 2, iMd-startMd).Sub(startDate).Hours()) / 24
				}
			} else {
				iDiff = iMd - startMd

				if iDiff <= 0 || iMd > startDate.AddDate(0, 1, -startMd).Day() {
					if int(startDate.AddDate(0, 1, iMd-startMd).Month())-startMm == 1 {
						iDiff = int(startDate.AddDate(0, 1, iMd-startMd).Sub(startDate).Hours()) / 24
					} else {
						iDiff = int(startDate.AddDate(0, 2, iMd-startMd).Sub(startDate).Hours()) / 24
					}
				}
			}

			if diff > iDiff {
				diff = iDiff
			}

		}

		return startDate.AddDate(0, 0, diff).Format(Layout), nil
	case r[0] == "m" && len(r) == 3:
		md := strings.Split(r[1], ",")
		if len(md) > 33 {
			return "", errors.New("invalid repeat format")
		}

		mm := strings.Split(r[2], ",")
		if len(mm) > 12 {
			return "", errors.New("invalid repeat format")
		}

		if now.After(startDate) {
			startDate = now
		}

		startMd := startDate.Day()
		startMm := int(startDate.Month())
		diff := 367

		var iDiff int
		for _, sMm := range mm {
			iMm, err := strconv.Atoi(sMm)
			if err != nil {
				return "", errors.New("invalid repeat format")
			}
			if iMm < 1 || iMm > 12 {
				return "", errors.New("invalid repeat format")
			}

			for _, sMd := range md {
				iMd, err := strconv.Atoi(sMd)
				if err != nil {
					return "", errors.New("invalid repeat format")
				}
				if iMd < -2 || iMd > 31 || iMd == 0 {
					return "", errors.New("invalid repeat format")
				}

				if iMd < 0 {
					iMd++
					iMm++

					iDiff = int(startDate.AddDate(0, iMm-startMm, iMd-startMd).Sub(startDate).Hours()) / 24

					if iDiff <= 0 {
						iDiff = int(startDate.AddDate(1, iMm-startMm, iMd-startMd).Sub(startDate).Hours()) / 24
					}
				} else {
					if int(startDate.AddDate(0, iMm-startMm, iMd-startMd).Month()) != iMm {
						continue
					}

					iDiff = int(startDate.AddDate(0, iMm-startMm, iMd-startMd).Sub(startDate).Hours()) / 24

					if iDiff <= 0 {
						iDiff = int(startDate.AddDate(1, iMm-startMm, iMd-startMd).Sub(startDate).Hours()) / 24
					}
				}

				if diff > iDiff {
					diff = iDiff
				}
			}
		}

		if diff < 367 {
			return startDate.AddDate(0, 0, diff).Format(Layout), nil
		}
	}

	return "", errors.New("invalid repeat format")
}

func prepareTask(task *database.Task) error {
	if len(task.Title) == 0 {
		return errors.New("empty title")
	}

	now, _ := time.Parse(Layout, time.Now().UTC().Format(Layout))

	var nextDate string
	var err error
	if len(task.Repeat) > 0 {
		nextDate, err = NextDate(now, now.Format(Layout), task.Repeat)
		if err != nil {
			return err
		}
	}

	if len(task.Date) == 0 {
		task.Date = now.Format(Layout)
	}

	date, err := time.Parse(Layout, task.Date)
	if err != nil {
		return errors.New("invalid date format")
	}

	if date.Before(now) {
		if len(task.Repeat) > 0 {
			task.Date = nextDate
			return nil
		}

		task.Date = now.Format(Layout)
	}

	return nil
}
