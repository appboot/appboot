const WSS_URL = process.env.WS_URL || "ws://127.0.0.1:8888/ws"

export let websocket = ''

export function createSocket() {
    if (!websocket) {
        window.console.log('creating websocket: '+WSS_URL)
        websocket = new WebSocket(WSS_URL)
    } else {
        window.console.log('websocket was created')
    }
}

export function sendGetTemplates() {
    const msg = JSON.stringify({method: 'GetTemplates'})
    websocket.send(msg);
}

export function sendCreateApp(name, template, params) {
    const msg = JSON.stringify({
        method: 'CreateApp', 
        application: {
            name: name,
            template: template,
            params: params,
        }})
    websocket.send(msg);
}

export function jsonParams(params) {
    var obj = {};
    for (var j = 0, len = params.length; j < len; j++) {
      const param = params[j];
      obj[param.key] = param.value;
    }
    return JSON.stringify(obj);
}