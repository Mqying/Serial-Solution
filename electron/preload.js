const { contextBridge, shell } = require('electron')

contextBridge.exposeInMainWorld('shell', {
    open: (href) => shell.openExternal(href),
})