package main

import (
	"fmt"
	"os"
	"time"
)

type stdoutActionsEmailSenderStruct struct{}

func newStdoutActionsEmailSender() *stdoutActionsEmailSenderStruct {
	stdoutFrontendActionsEmailSender := &stdoutActionsEmailSenderStruct{}
	return stdoutFrontendActionsEmailSender
}

func (*stdoutActionsEmailSenderStruct) SendSignupEmailAddressVerificationCode(emailAddress string, emailAddressVerificationCode string) error {
	message := fmt.Sprintf("Your email address verification code is %s.", emailAddressVerificationCode)
	fmt.Fprintf(os.Stdout, "[EMAIL] To %s: %s\n", emailAddress, message)
	return nil
}

func (*stdoutActionsEmailSenderStruct) SendUserEmailAddressUpdateEmailVerificationCode(emailAddress string, _ string, emailAddressVerificationCode string) error {
	message := fmt.Sprintf("Hi, your email address verification code is %s.", emailAddressVerificationCode)
	fmt.Fprintf(os.Stdout, "[EMAIL] To %s: %s\n", emailAddress, message)
	return nil
}

func (*stdoutActionsEmailSenderStruct) SendUserPasswordResetTemporaryPassword(emailAddress string, _ string, temporaryPassword string) error {
	message := fmt.Sprintf("Hi, your password reset temporary password is %s.", temporaryPassword)
	fmt.Fprintf(os.Stdout, "[EMAIL] To %s: %s\n", emailAddress, message)
	return nil
}

func (*stdoutActionsEmailSenderStruct) SendUserSignedInNotification(emailAddress string, _ string, _ time.Time) error {
	message := "Hi, we detected a sign-in to your account."
	fmt.Fprintf(os.Stdout, "[EMAIL] To %s: %s\n", emailAddress, message)
	return nil
}

func (*stdoutActionsEmailSenderStruct) SendUserPasswordUpdatedNotification(emailAddress string, _ string, _ time.Time) error {
	message := "Hi, your account password was updated."
	fmt.Fprintf(os.Stdout, "[EMAIL] To %s: %s\n", emailAddress, message)
	return nil
}

func (*stdoutActionsEmailSenderStruct) SendUserEmailAddressUpdatedNotification(emailAddress string, _ string, newEmailAddress string, _ time.Time) error {
	message := fmt.Sprintf("Hi, your account email address was updated to %s.", newEmailAddress)
	fmt.Fprintf(os.Stdout, "[EMAIL] To %s: %s\n", emailAddress, message)
	return nil
}
