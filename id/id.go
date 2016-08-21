package id

import (
	"errors"
	"time"

	timex "github.com/chanxuehong/time"
)

var (
	ErrInvalidNumber   = errors.New("无效的身份证号码")
	ErrInvalidBirthday = errors.New("无效的出生日期")
)

// ID 表示 中华人民共和国公民身份号码
type ID struct {
	number string // 身份证卡号

	address      string    // 地址码
	birthday     string    // 出生日期
	birthdayTime time.Time // 出生日期的零点时间
	order        string    // 顺序码
	checksum     byte      // 校验码
}

// New 创建一个新的 ID, 不校验 number 的正确性.
func New(number string) (*ID, error) {
	if len(number) != 18 {
		return nil, ErrInvalidNumber
	}
	if number[17] == 'x' {
		number = number[:17] + "X"
	}
	return &ID{
		number:   number,
		address:  number[:6],
		birthday: number[6:14],
		order:    number[14:17],
		checksum: number[17],
	}, nil
}

// NewWithCheck 创建一个新的 ID, 简单校验 number 的正确性.
func NewWithCheck(number string) (*ID, error) {
	if len(number) != 18 {
		return nil, ErrInvalidNumber
	}
	char := number[17]
	switch char {
	case 'X':
	case 'x':
		number = number[:17] + "X"
	default:
		if char < '0' || char > '9' {
			return nil, ErrInvalidNumber
		}
	}
	for i := 0; i < 17; i++ {
		char = number[i]
		if char < '0' || char > '9' {
			return nil, ErrInvalidNumber
		}
	}
	id := ID{
		number:   number,
		address:  number[:6],
		birthday: number[6:14],
		order:    number[14:17],
		checksum: number[17],
	}
	if Checksum(id.number) != id.checksum {
		return nil, ErrInvalidNumber
	}
	if !isAddressCodeValid(id.address) {
		return nil, ErrInvalidNumber
	}
	if str := id.birthday[:2]; str != "19" && str != "20" {
		return nil, ErrInvalidNumber
	}
	t, err := time.ParseInLocation("20060102", id.birthday, timex.BeijingLocation)
	if err != nil {
		return nil, ErrInvalidNumber
	}
	id.birthdayTime = t
	return &id, nil
}

// Number 返回身份证号码
func (id *ID) Number() string {
	return id.number
}

// Address 返回地址码
func (id *ID) Address() string {
	return id.address
}

// Birthday 返回生日码
func (id *ID) Birthday() string {
	return id.birthday
}

// BirthdayTime 返回生日当天的零点时间
func (id *ID) BirthdayTime() (t time.Time, err error) {
	if !id.birthdayTime.IsZero() {
		t = id.birthdayTime
		return
	}
	if id.birthday == "" {
		err = ErrInvalidNumber
		return
	}
	if t, err = time.ParseInLocation("20060102", id.birthday, timex.BeijingLocation); err != nil {
		err = ErrInvalidBirthday
		return
	}
	id.birthdayTime = t
	return
}

// Order 返回顺序码
func (id *ID) Order() string {
	return id.order
}

// Checksum 返回校验码
func (id *ID) Checksum() byte {
	return id.checksum
}
