const http = require("http");

const requestHandler = (req, res) => {
  // Simulate some long running processing
  for (let i = 0; i < 1e6; i++) {}
  res.end("Hello, OpenTelemetry eBPF Profiler!");
};

const server = http.createServer(requestHandler);

server.listen(3000, () => {
  console.log("Server is listening on port 3000");
});
