function connect(url) {
  var socket = new WebSocket(url);
  var actions = {
    ready: false,
    send(event, data) {
      if(this.ready) {
        console.log("sent")
        socket.send(`[${event}] : ${data}`);
      } else {
        setTimeout(() => {
          this.send(event, data);
        }, 100)
      }
    },
    on(event, callback) {
      this[event] = callback;
      console.log(this)
    }
  }
  socket.onopen = function(event) {
    // console.log(event);
    actions.ready = true;
  }
  socket.onmessage = function(event) {
    var messageData = {};
    event.data.replace(/\[(\w+)\]\s:\s\[(.+)\]\s:\s(.+)/i, function(original, command, user, message) {
      messageData = {
        original,
        command,
        user,
        message
      }
    });
    actions[messageData.command](messageData)
  };
  return actions;
}

var chatElement = document.querySelector(".chat-messages")
var ws = connect("ws://localhost:8080/bot");
ws.on("PRIVMSG", function(data) {
  var li = document.createElement("li");
  li.setAttribute("data-username", data.user)
  var user = document.createElement("span");
  user.innerText = data.user;
  var message = document.createElement("span");
  message.innerText = data.message;

  li.innerHTML += "[";
  li.appendChild(user);
  li.innerHTML += "]: ";
  li.appendChild(message);
  chatElement.appendChild(li);
});
