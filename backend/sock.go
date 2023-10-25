package main

import(
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

var newline = []byte{'\n'}

func (c *Client) readPump(){
	for{
		_, p, err := c.conn.ReadMessage();
		_ = p;
		if err!=nil{
			//log.Println(err)
			return
		}else{
			//fmt.Println("msg got : ", messageType, " _ ", string(p));
			c.send <- []byte("omg hi!!")
		}
	}
}

func (c *Client) writePump(){
	for{
		select{
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil{
				return
			}
			w.Write(message)
			
			n := len(c.send)
			for i := 0; i<n;i++{
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err!=nil{
				return
			}
		}
	}
}

func check_header(r *http.Request) bool {
	//fmt.Println(r)
	return true;
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

func serveWs( hub *ConnManager, w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func(r *http.Request) bool {return check_header(r)}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(err)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte, 256)}

	//fmt.Println("client connected")

	hub.register <- client
	
	go client.readPump()
	go client.writePump()
}
