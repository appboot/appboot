
export function decodeParams(params) {
    var result = []

    var stringParams = params.string
    var intParams = params.int

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

    return result
}