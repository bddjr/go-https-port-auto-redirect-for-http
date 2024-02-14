// node main.js

const os = require("os");
const ostype = os.type();
const fs = require("fs");

console.log(ostype);

var goPath = "";
if (ostype === "Windows_NT") {
  goPath = "C:\\Program Files\\Go\\src\\net\\http\\server.go";
} else {
  throw "unknow system";
}

console.log(goPath);

const serverGo = fs.readFileSync(goPath).toString();
var newServerGo = "";
const onOffStr = {
  off: `"HTTP/1.0 400 Bad Request\\r\\n\\r\\nClient sent an HTTP request to an HTTPS server.\\n"`,
  on: `"HTTP/1.0 400 Bad Request\\r\\nContent-Type: text/html\\r\\n\\r\\n<!-- Client sent an HTTP request to an HTTPS server. -->\\n<!-- https://github.com/bddjr/go-https-port-auto-redirect-for-http -->\\n<html><head><script>location.protocol='https:'</script></head><body></body></html>\\n"`,
};

var argv = process.argv[2]?.toLowerCase();

switch (argv) {
  case "on":
    newServerGo = serverGo.replace(onOffStr.off, onOffStr.on);
    break;
  case "off":
    newServerGo = serverGo.replace(onOffStr.on, onOffStr.off);
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
