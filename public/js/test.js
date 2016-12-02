function stripEventData(data) {
  var messageData = {};
  data.replace(/\[(\w+)\]\s:\s\[(.+)\]\s:\s(.+)/i, function(original, command, user, message) {
    messageData = {
      original,
      command,
      user,
      message
    }
  });
  console.log(messageData);
  return messageData;
}
function connect(url) {
  var socket = new WebSocket(url);
  var actions = {
    ready: false,
    send(event, data) {
      if(this.ready) {
        // console.log("sent")
        socket.send(`${data}`);
      } else {
        setTimeout(() => {
          this.send(event, data);
        }, 100)
      }
    },
    on(event, callback) {
      this[event] = callback;
      // console.log(this)
    }
  }
  socket.onopen = function(event) {
    // console.log(event);
    actions.ready = true;
  }
  socket.onmessage = function(event) {
    // console.log(event);
    var messageData = stripEventData(event.data);
    actions[messageData.command](messageData)
  };
  return actions;
}
document.querySelector(".input .submit").addEventListener("submit", sendMessege, false);
document.querySelector(".tools .channel-name").addEventListener("submit", makeConnection, false);
var submitInput = document.querySelector(".input .submit input");
var channelNameInput = document.querySelector(".tools .channel-name input");
var chatElement = document.querySelector(".chat-messages");

var ws;

function sendMessege(e) {
  e.preventDefault();
  ws.send("PRIVMSG", submitInput.value);
  appendMessage({
    user: "piecedigital",
    message: submitInput.value
  });
  submitInput.value = "";
};

function appendMessage(data) {
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
}

function makeConnection(e) {
  e.preventDefault();
  channelNameInput.disabled = true;
  ws = connect("ws://localhost:8080/bot?channel="+channelNameInput.value+"");
  ws.on("PRIVMSG", appendMessage);
}
