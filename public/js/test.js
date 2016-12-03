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
document.querySelector(".tools .close-connection").addEventListener("submit", closeConnection, false);
var submitInput = document.querySelector(".input .submit input");
var channelNameInput = document.querySelector(".tools .channel-name input");
var channelNameBtn = document.querySelector(".tools .channel-name button");
var closeConnBtn = document.querySelector(".tools .close-connection button");
var chatElement = document.querySelector(".chat-messages");

var ws, lastChild;

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
  console.log(lastChild);
  if(lastChild) {
    chatElement.insertBefore(li, lastChild);
    lastChild = li;
  } else {
    chatElement.appendChild(li);
    lastChild = li;
  }
}

function makeConnection(e) {
  e.preventDefault();
  channelNameInput.disabled = true;
  channelNameBtn.disabled = true;
  closeConnBtn.disabled = false;
  ws = connect("ws://localhost:8080/bot?channel="+channelNameInput.value);
  ws.on("PRIVMSG", appendMessage);
}

function closeConnection(e) {
  e.preventDefault();
  channelNameInput.disabled = false;
  channelNameBtn.disabled = false;
  closeConnBtn.disabled = true;
  ajax({
    url: "http://localhost:8080/stop-bot",
    success: function (data) {
      console.log(data);
    },
    error: function (err) {
      console.error(err);
    }
  });
}

function ajax(optionsObj) {
	optionsObj = optionsObj || {};
	// console.log(optionsObj.data);

	var httpRequest = new XMLHttpRequest();
	if(typeof optionsObj.upload === "function") httpRequest.upload.addEventListener("progress", optionsObj.upload, false);
	httpRequest.onreadystatechange = function(data) {
		if(httpRequest.readyState === 4) {
			if(httpRequest.status < 400) {
				if(typeof optionsObj.success === "function") {
					optionsObj.success(data.target.response);
				} else {
					console.log("no success callback in ajax object");
				}
			} else {
				if(typeof optionsObj.error === "function") {
					optionsObj.error({
						"status": data.target.status,
						"message": data.target.statusText,
						"response": data.target.response
					});
				} else {
					console.log("no error callback in ajax object. logging error below");
					console.error(data.target.status, data.target.statusText);
				}
			}
		}
	};
	var contentTypes = {
		jsonp: "application/javascript; charset=UTF-8",
		json: "application/json; charset=UTF-8",
		text: "text/plain; charset=UTF-8",
		formdata: "multipart/form-data; boundary=---------------------------file0123456789end"
	};

	httpRequest.open(((optionsObj.type || "").toUpperCase() || "GET"), optionsObj.url, optionsObj.multipart || true);
	if(optionsObj.dataType) httpRequest.setRequestHeader("Content-Type", `${contentTypes[(optionsObj.dataType.toLowerCase() || "text")]}`);
	if(typeof optionsObj.beforeSend == "function") {
		optionsObj.beforeSend(httpRequest);
	}
	httpRequest.send(optionsObj.data || null);
};
