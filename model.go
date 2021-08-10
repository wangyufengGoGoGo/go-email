package go_email

/**
 * @Author wyf
 * @Date 2021/8/9 16:52
 **/

type ClientOption struct {
	Host       string `json:"host"`
	ServerAddr string `json:"serverAddr"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

type Email struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Msg     string   `json:"msg"`
}
