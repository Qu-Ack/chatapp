import { useEffect, useRef, useState } from "react"

export default function App() {

  const [message, setMessage] = useState("")
  const socketRef = useRef(null)
  const [messages, setMessages] = useState([])

  useEffect(() => {
   const socket = new WebSocket("ws://localhost:9000/ws")
    socketRef.current = socket

    socket.onopen = (e) => {
      console.log("socket connected to the server ...")
    }

    socket.onmessage = async (e) => {
     let  data = await parseJSON(e.data) 
     console.log(data)
     if (data.type == 1) {
      setMessages((messages) => [...messages, data.body])
     }
    }

    socket.onerror = (error) => {
      console.log("socket connection close due to ", error)
    }

    socket.onclose = () => {
      console.log("socket connection closed")
    }

    return () => {
      socket.close()
    }

  }, [])


  async function parseJSON(data) {
    const result = await JSON.parse(data)
    return result
  }

  function handleSubmit(e) {
    e.preventDefault()
    if (socketRef.current) {
      socketRef.current.send(message)
    }
    setMessage("")
  }


  return (
    <div className="main-wrapper">
      <h1>Welcome to ChatApp</h1>
      <div className="chat-div">
        <div className="chat-container">
          {
            messages.map(msg => {
              return (
                <div className="message">
                  {msg}
                   </div>
              )
            })
          }
        </div>
      <form onSubmit={handleSubmit}>
        <input type="text" placeholder="your message" className="chat-input" value={message} onChange={(e) => {setMessage(e.target.value)}}></input>
        <button type="submit">Send</button>
      </form>
      </div>
    </div>
  )

}