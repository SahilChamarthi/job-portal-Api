package services

import (
	"errors"
	"fmt"
	"math/rand"
	"net/smtp"

	"project/internal/model"
	redispack "project/internal/redisPack"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *Services) UserSignup(nu model.UserSignup) (model.User, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Msg("error occured in hashing password")
		return model.User{}, errors.New("hashing password failed")
	}

	user := model.User{UserName: nu.UserName, Email: nu.Email, PasswordHash: string(hashedPass)}
	cu, err := s.r.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.User{}, errors.New("user creation failed")
	}

	return cu, nil

}
func (s *Services) UserLogin(l model.UserLogin) (jwt.RegisteredClaims, error) {
	fu, err := s.r.FetchUserByEmail(l.Email)
	if err != nil {
		log.Error().Err(err).Msg("couldnot find user")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	fmt.Println(fu)
	err = bcrypt.CompareHashAndPassword([]byte(fu.PasswordHash), []byte(l.Password))
	if err != nil {
		log.Error().Err(err).Msg("password of user incorrect")
		return jwt.RegisteredClaims{}, errors.New("user login failed")
	}
	c := jwt.RegisteredClaims{
		Issuer:    "service project",
		Subject:   strconv.FormatUint(uint64(fu.ID), 10),
		Audience:  jwt.ClaimStrings{"users"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	fmt.Println(c)

	return c, nil

}
func (s *Services) GenerateOtp(m string) (bool, error) {

	ok, err := s.r.FindEmail(m)

	if err != nil {
		log.Error().Err(err)
		return false, err
	}

	if ok {

		otp := generateOTP()
		otpStr := strconv.Itoa(otp)

		from := "sahilchamarthi06@gmail.com"
		password := "kzfr utuj wbzt xkuy"

		// Recipient's email address
		to := m

		// SMTP server and port
		smtpServer := "smtp.gmail.com"
		smtpPort := 587

		// Message
		subject := "opt regarding job portal"
		body := fmt.Sprintf("one time password is:%s", otpStr)

		// Set up authentication information
		auth := smtp.PlainAuth("", from, password, smtpServer)

		// Compose the email
		message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

		// Connect to the SMTP server
		err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
		if err != nil {
			fmt.Println("Error sending email:", err)
			log.Err(err)
			return false, err
		}

		redis := redispack.NewRedisClient()

		err = redis.Set(m, otpStr, 5*time.Minute).Err()

		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func generateOTP() int {

	return rand.Intn(999999) + 100000
}

func (s *Services) NewPasswordVerify(rp model.PasswordReset) error {

	redis := redispack.NewRedisClient()

	pass, err := redis.Get(rp.Email).Result()

	if err != nil {
		return err
	}

	if rp.Otp == pass {
		_, err = s.r.UpdatePassword(rp.Email, rp.NewPassword)
		if err != nil {
			return err
		}

		from := "sahilchamarthi06@gmail.com"
		password := "kzfr utuj wbzt xkuy"

		// Recipient's email address
		to := rp.Email

		// SMTP server and port
		smtpServer := "smtp.gmail.com"
		smtpPort := 587
		infoMesge := "password has been successfully reset"
		// Message
		subject := "Reset password"
		body := fmt.Sprintf(":%s", infoMesge)

		// Set up authentication information
		auth := smtp.PlainAuth("", from, password, smtpServer)

		// Compose the email
		message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

		// Connect to the SMTP server
		err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
		if err != nil {
			fmt.Println("Error sending email:", err)
			return err
		}
	}
	return nil
}
