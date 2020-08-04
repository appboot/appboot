
export function decodeParams(params) {
    var result = []

    var stringParams = params.string
    var intParams = params.int
    var floatParams = params.float
    var selectParams = params.select

    if (stringParams != null) {
        for (let i = 0; i < stringParams.length; i++) {
            var stringParam = stringParams[i]
            result.push({
                type: "string",
                key: stringParam.key,
                value: stringParam.default
            });
        }
    }

    if (intParams != null) {
        for (let i = 0; i < intParams.length; i++) {
            var intParam = intParams[i]
            result.push({
                type: "int",
                key: intParam.key,
                value: intParam.default,
                min: intParam.min,
                max: intParam.max
            });
        }
    }

    if (floatParams != null) {
        for (let i = 0; i < floatParams.length; i++) {
            var floatParam = floatParams[i]
            result.push({
                type: "float",
                key: floatParam.key,
                value: floatParam.default,
                min: floatParam.min,
                max: floatParam.max
            });
        }
    }

    if (selectParams != null) {
        for (let i = 0; i < selectParams.length; i++) {
            var selectParam = selectParams[i]
            result.push({
                type: "select",
                key: selectParam.key,
                options: selectParam.options,
            });
        }
    }

    return result
}

export function encodeParams(params) {
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