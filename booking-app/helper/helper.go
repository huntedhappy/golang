package helper

import "strings"

// 참조 하는 package가 다른 폴더에 있을 경우 ValidateUserInput 여기서 V를 대문자로 변경이 필요 하며 , remainigTicketes 정의를 하지 않았으므로 uint로 정의를 내린다
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainigTicketes uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainigTicketes
	return isValidName, isValidEmail, isValidTicketNumber

}
