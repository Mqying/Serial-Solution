export function formatTime(arg) {
    let mouth = arg.substr(5, 2);
    let day = arg.substr(8, 2);
    let year = arg.substr(0, 4);
    let hour = arg.substr(11, 2);
    let minute = arg.substr(14, 2);

    return `${mouth}/${day}/${year}  ${hour}:${minute}`;
}

export function formatUnit(arg) {
    arg = arg.trim();

    if (arg.includes(" ")) {
        return arg.slice(0, arg.indexOf(" "));
    } else {
        return arg;
    }
}