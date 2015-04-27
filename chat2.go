package main

func readFromSocket(conn net.Conn, rMessage chan string) {
	var buff [512]byte
	for {
		n, err := conn.Read(buff[0:])
		if err != nil {
			return
		}
	}
	stringBuf := string(buff[:n])
	rMessage <- stringBuf

}
func writeToSocket(conn net.Conn, wMessage chan string) {
	for message := range wMessage {
		_, err2 := conn.Write()
		if err2 != nil {
			return
		}
	}
}
func messageMenager(rMessage chan string, wMessage chan string) {
	for message := range rMessage {
		wMessage <- message
	}
}
