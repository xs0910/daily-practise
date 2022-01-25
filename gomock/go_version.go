package gomock

import "daily-practise/gomock/spider"

func GetGoVersion(s spider.Spider) string {
	body := s.GetBody()
	return body
}
