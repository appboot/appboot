import EventEmitter from "eventemitter3";
import { API_URL } from "./config";

export enum SokcetEvent {
    open = 'onopen',
    close = 'onclose',
    message = 'onmessage'
}

export enum SokcetCMD {
    getTemplates = "getTemplates"
}

class Sokcet extends EventEmitter{
    sock: WebSocket = new WebSocket(API_URL)

    constructor() {
        super()

        var that = this;
        this.sock.onopen = function() {
            console.log("connected to " + API_URL);
            that.emit(SokcetEvent.open)
        }

        this.sock.onclose = function(e) {
            console.log("connection closed (" + e.code + ")");
            that.emit(SokcetEvent.close)
        }

        this.sock.onmessage = function(e) {
            console.log("message received: " + e.data);
            that.emit(SokcetEvent.message, e.data)
        }
    }

    getTemplates() {
        const message = `{"cmd": "${SokcetCMD.getTemplates}"}`
        this.sock.send(message)
    }
}

var socket = new Sokcet()

export default socket
