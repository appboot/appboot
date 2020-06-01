import { method } from "./const";
import { WSS_URL } from "./config";

export let websocket = ''

export function createSocket() {
    window.console.log('creating websocket: '+WSS_URL)
    websocket = new WebSocket(WSS_URL)
}

export function sendGetTemplates() {
    const msg = JSON.stringify({method: method.GetTemplates})
    websocket.send(msg);
}

export function sendUpdateAllTemplates() {
    const msg = JSON.stringify({method: method.UpdateAllTemplates})
    websocket.send(msg);
}

export function sendGetConfig(template) {
    const msg = JSON.stringify({
        method: method.GetConfig,
        application: {
            template: template,
        }})
    websocket.send(msg);
}

export function sendCreateApp(name, template, params) {
    const msg = JSON.stringify({
        method: method.CreateApp, 
        application: {
            name: name,
            template: template,
            params: params
        }})
    websocket.send(msg);
}

export function jsonParams(params) {
    var obj = {};
    for (var j = 0, len = params.length; j < len; j++) {
        const param = params[j];
        var value = param.value
        if (param.type != "string") {
            value = value.toString()
        }
        obj[param.key] = value;
    }
    return JSON.stringify(obj);
}