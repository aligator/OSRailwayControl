export default class Websocket {
    /**
     * @type Websocket
     */
    socket = null
    host = ""

    actions = new Map()

    isOpen = false

    constructor(host) {
        this.host = host
    }

    init() {
        return new Promise((resolve, reject) => {
            this.socket = new WebSocket(this.host)
            this.socket.onmessage = (e) => this.onMessage(e)
            this.socket.onopen = (_) => {
                this.isOpen = true
                resolve()
            }
            this.socket.onerror = (_) => {
                this.isOpen = false
                reject()
            }
            this.socket.onclose = (_) => {
                this.isOpen = false
            }
        })
    }

    onMessage(e) {
        const message = JSON.parse(e.data)
        if (typeof message.key !== "string" && typeof message.value !== "string") {
            console.error("received message has wrong format")
        }

        if (this.actions.has(message.key)) {
            this.actions.get(message.key)(message.value)
        }
    }

    send(key, value) {
        if (typeof key !== "string" && (value === undefined || typeof value !== "string")) {
            console.error("message to send has wrong format")
        }

        if (!value) {
            value = ""
        }

        this.socket.send(JSON.stringify({
            key, value
        }))
    }

    register(key, action) {
        this.actions.set(key, action)
    }
}