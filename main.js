// import
const os = require("os");
const fs = require("fs");
const path = require("path");

// Matching operating systems for path server.go
const ostype = os.type();
console.log(ostype);
var goPath = "";
switch (ostype) {
  case "Windows_NT":
    // Windows
    goPath = "C:\\Program Files\\Go";
    break;
  case "Linux":
  case "Darwin":
    // Linux or MacOS
    goPath = "/usr/local/go";
    break;
  default:
    // ?
    throw "unknow system";
}

// path server.go
goPath = path.join(goPath, "src", "net", "http", "server.go");
console.log(goPath);

// Reading server.go
const serverGo = fs.readFileSync(goPath).toString();

// Configure response string
const onOffStr = {
  off: JSON.stringify(
    "HTTP/1.0 400 Bad Request\r\n" +
      "\r\n" +
      "Client sent an HTTP request to an HTTPS server.\n"
  ),
  on: JSON.stringify(
    "HTTP/1.1 400 Bad Request\r\n" +
      "Content-Type: text/html\r\n" +
      "Connection: close\r\n"+
      "\r\n" +
      "<!-- Client sent an HTTP request to an HTTPS server. -->\n" +
      "<script> location.protocol = 'https:' </script>\n"
  ),
  old_versions: [
    JSON.stringify(
      "HTTP/1.0 400 Bad Request\r\n" +
        "Content-Type: text/html\r\n" +
        "\r\n" +
        "<!-- Client sent an HTTP request to an HTTPS server. -->\n" +
        "<script> location.protocol = 'https:' </script>\n"
    ),
    JSON.stringify(
      "HTTP/1.0 400 Bad Request\r\n" +
        "Content-Type: text/html\r\n" +
        "\r\n" +
        "<!-- Client sent an HTTP request to an HTTPS server. -->\n" +
        "<!-- https://github.com/bddjr/go-https-port-auto-redirect-for-http -->\n" +
        "<html><head><script>location.protocol='https:'</script></head><body></body></html>\n"
    ),
  ],
};

// new server.go
var newServerGo = "";

/**
 * replace new server.go
 * @param {string} oldStr
 * @param {string} newStr
 */
function replaceNewServerGo(oldStr, newStr) {
  // old to new
  newServerGo = serverGo.replace(oldStr, newStr);
  if (newServerGo !== serverGo) return; // if changed
  // Searching for old versions
  for (const i of onOffStr.old_versions) {
    newServerGo = serverGo.replace(i, newStr);
    if (newServerGo !== serverGo) return; // if changed
  }
}

// Matching arguments for new server.go
const argv = process.argv[2]?.toLowerCase();
switch (argv) {
  case "on":
    replaceNewServerGo(onOffStr.off, onOffStr.on);
    break;
  case "off":
    replaceNewServerGo(onOffStr.on, onOffStr.off);
    break;
  default:
    throw "unknow argv";
}

// Writing new server.go
if (newServerGo !== serverGo) {
  console.log(`Writing ${argv}`);
  fs.writeFileSync(goPath, newServerGo);
} else {
  console.log("No changes needed");
}
