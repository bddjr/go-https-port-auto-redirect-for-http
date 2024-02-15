// node main.js

const os = require("os");
const ostype = os.type();
const fs = require("fs");
const path = require("path");

console.log(ostype);

var goPath = "";
switch (ostype) {
  case "Windows_NT":
    goPath = "C:\\Program Files\\Go";
    break;
  case "Linux":
  case "Darwin":
    goPath = "/usr/local/go";
    break;
  default:
    throw "unknow system";
}
goPath = path.join(goPath, "src", "net", "http", "server.go");

console.log(goPath);

const serverGo = fs.readFileSync(goPath).toString();
var newServerGo = "";

const onOffStr = {
  off: `"HTTP/1.0 400 Bad Request\\r\\n\\r\\nClient sent an HTTP request to an HTTPS server.\\n"`,
  on: `"HTTP/1.0 400 Bad Request\\r\\nContent-Type: text/html\\r\\n\\r\\n<!-- Client sent an HTTP request to an HTTPS server. -->\\n<script> location.protocol = 'https:' </script>\\n"`,
  olds: [
    `"HTTP/1.0 400 Bad Request\\r\\nContent-Type: text/html\\r\\n\\r\\n<!-- Client sent an HTTP request to an HTTPS server. -->\\n<!-- https://github.com/bddjr/go-https-port-auto-redirect-for-http -->\\n<html><head><script>location.protocol='https:'</script></head><body></body></html>\\n"`,
  ],
};

var argv = process.argv[2]?.toLowerCase();

/**
 * @param {string} oldStr
 * @param {string} newStr
 */
function serverGoReplace(oldStr, newStr) {
  newServerGo = serverGo.replace(oldStr, newStr);
  if (newServerGo !== serverGo) return;
  for (const i of onOffStr.olds) {
    newServerGo = serverGo.replace(i, newStr);
    if (newServerGo !== serverGo) return;
  }
}

switch (argv) {
  case "on":
    serverGoReplace(onOffStr.off, onOffStr.on);
    break;
  case "off":
    serverGoReplace(onOffStr.on, onOffStr.off);
    break;
  default:
    throw "unknow argv";
}

if (newServerGo !== serverGo) {
  console.log(`Writing ${argv}`);
  fs.writeFileSync(goPath, newServerGo);
} else {
  console.log("No changes needed");
}
