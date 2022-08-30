// Modules to control application life and create native browser window
const { app, BrowserWindow } = require('electron');
const path = require('path');
const { exec, spawn } = require('node:child_process');

let backgroundProcess = undefined;

function createWindow() {
  // Create the browser window.
  const mainWindow = new BrowserWindow({
    resizable: false,
    width: 1600 * 0.8,
    height: 900 * 0.8,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js')
    },
  });

  // and load the index.html of the app.
  if (process.env.NODE_ENV === 'development') {
    mainWindow.menuBarVisible = true;

    backgroundProcess = spawn("go", ["run", `${path.join(__dirname, '../server/main.go')}`], {
      cwd: path.join(__dirname, "../server"),
      detached: true,
    });
    backgroundProcess.unref();

    backgroundProcess.on('exit', (code) => {
      console.log(`Command go run finished, code[${code}]`);
      backgroundProcess = undefined;
    });

    mainWindow.loadURL('http://127.0.0.1:8000');
  } else {
    mainWindow.menuBarVisible = false;

    backgroundProcess = spawn(path.join(__dirname, '../../../backstage/main.exe'), {
      cwd: path.join(__dirname, "../../../backstage"),
      detached: true,
    });

    let localPath = __dirname.substring(0, __dirname.length - 8);
    mainWindow.loadFile(`${localPath}/dist/index.html`);
  }
  // mainWindow.loadFile('index.html');
  // Open the DevTools.
  // mainWindow.webContents.openDevTools()
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.whenReady().then(() => {
  createWindow();

  app.on('activate', function () {
    // On macOS it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

// Quit when all windows are closed, except on macOS. There, it's common
// for applications and their menu bar to stay active until the user quits
// explicitly with Cmd + Q.
app.on('window-all-closed', function () {
  if (backgroundProcess !== undefined) {
    backgroundProcess.kill()
    backgroundProcess = undefined
  }

  if (process.platform !== 'darwin') app.quit();
});

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.
