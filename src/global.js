const global = {
    frozen: false,
    lastQueryTime: undefined,
    first: true,

    setFrozen: () => {
        global.frozen = true;
        global.lastQueryTime = new Date();
    },

    clearFrozen: () => {
        global.frozen = false;
    },

    isFrozen: () => {
        if (global.frozen == false) {
            return false;
        }

        let limitSecond = 5
        let nowTime = new Date()
        if (((nowTime.getTime() - new Date(global.lastQueryTime).getTime()) / 1000) >= limitSecond) {
            global.frozen = false;
            return false;
        }

        return true;
    },

    isFirst: () => {
        if (global.first == true) {
            return true;
        } else {
            return false;
        }
    },

    setFirst: () => {
        global.first = false;
    }
};

export {
    global
}