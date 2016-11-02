const fs = require('fs');
const http = require('http');

const PORT = 3000;

function handleRequest(request, response){
  fs.readFile('cache.html', 'utf-8', (err, content) => {
    if (err) {
      console.log(err);
    }
    response.end(content);
  })
}

const server = http.createServer(handleRequest);

server.listen(PORT, function(){
    console.log("Node server listening on %s", PORT);
});
