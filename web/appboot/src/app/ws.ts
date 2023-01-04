import EventEmitter from "eventemitter3";
import { API_URL } from "./config";

export enum SokcetEvent {
    open = 'onopen',
    close = 'onclose',
    message = 'onmessage'
}

export enum SokcetCMD {
    getTemplates = "getTemplates",
    getTemplateConfig = "getTemplateConfig",
    updateTemplates = "updateTemplates",
    createApp = "createApp"
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
            that.emit(SokcetEvent.message, e.data)
        }
    }

    getTemplates() {
        this.send({
            cmd: SokcetCMD.getTemplates
        })
    }

    getConfigs(template: string) {
        this.send({
            cmd: SokcetCMD.getTemplateConfig,
            data: {
                template,
            }
        })
    }

    updateTemplates() {
        this.send({
            cmd: SokcetCMD.updateTemplates
        })
    }

    createApp(name: string, template: string, params: string, skipBeforeScripts: string, skipAfterScripts:string) {
        this.send({
            cmd: SokcetCMD.createApp,
            data: {
                name,
                template,
                params,
                skipBeforeScripts,
                skipAfterScripts
            }
        })
    }

    send(obj: object) {
        const message = JSON.stringify(obj)
        this.sock.send(message)
    }
}

var socket = new Sokcet()

export default socket
