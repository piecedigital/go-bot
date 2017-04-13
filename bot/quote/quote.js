process.on("uncaughtException", function (err) {
  console.error(err.stack || err);
});
// var Firebase = require("firebase");
var express = require("express");
var admin = require("firebase-admin");

var app = express();

var serviceAccount = (require(__dirname + "/strange-stuff-pds-firebase-adminsdk-uy91z-a4cfec4a22.json"));
var fire = admin;

fire.initializeApp({
  credential: admin.credential.cert(serviceAccount),
  databaseURL: "https://strange-stuff-pds.firebaseio.com"
});

var ref = {
  root: fire.database().ref(),
  quotesRef: fire.database().ref("quotes"),
};

app
.post("/", function (req,res,next) {
  var body = "";
  req.on("data", function (chunk) {
    // console.log("chunk", chunk);
    body += chunk;
  });
  req.on("end", function (chunk) {
    // console.log("the end");
    // console.log(body, typeof body);
    req.body = JSON.parse(body);
    // var clearedBody = body.replace(/(\+)/g, " ");
    // // console.log("cleared", clearedBody);
    next();
  });
}, function (req, res) {
  console.log(req.body);
//   res.status(200).send("ok");
// return;
  var operation = req.body.op;
  var value = req.body.val.replace(/[#]/g, "");
  var quoter = req.body.quoter;

  switch (operation) {
    case "get":
    case "who":
      if(value.match(/^[0-9]+$/g)) {
        ref.quotesRef
        .orderByChild("quoteNumber")
        .equalTo(parseInt(value) || 0)
        .once("value")
        .then(function (snap) {
          var quotes = snap.val();
          if(quotes) {
            var key = Object.keys(quotes)[0];
            var val = quotes[key];

            console.log(val.quote);
            if(operation === "who") {
              res.status(200).send("Quote #" + val.quoteNumber + " was created by @" + val.quoter);
            } else {
              res.status(200).send(val.quote);
            }
          } else {
            res.status(400).send("Couldn't find that quote.");
          }
        }, function (err) {
          if(err) {
            console.error(err.stack || err);
            res.status(400).send("There was a problem getting the quote");
            return;
          }
          console.log("done");
        });
      } else {
        console.log("err");
        res.status(400).send("Value must be a number");
      }
    break;
    case "add":
    case "set":
      ref.quotesRef
      .orderByKey()
      .limitToLast(1)
      .once("value")
      .then(function (snap) {
        var quotes = snap.val();
        var key, val;
        if(quotes) {
          key = Object.keys(quotes)[0];
          val = quotes[key];
          // console.log("got quote", val);
        }

        if(value.match(/[a-z]/gi)) {
          var nextNumber = val ? val.quoteNumber + 1 : 0;
          try {
            ref.quotesRef
            .push()
            .set({
              quoteNumber: nextNumber,
              quote: value,
              quoter: quoter,
              date: Date.now()
            }, function (err) {
              console.log("done");
              res.status(200).send("quote #" + nextNumber + " has been set");
            });
          } catch (err) {
            console.error(err.stack || err);
            res.status(400).send("There was a problem setting the quote");
          }
        } else {
          console.log("err");
        }
      })
      .catch(function (e) {
        console.error(e);
      })
    break;
    case "count":
      ref.quotesRef
      .once("value")
      .then(function (snap) {
        var quotes = snap.val();
        var count = 0;
        if(quotes) {
          count = Object.keys(quotes).length;
        }
        res.status(200).send("There are " + count + " quotes.");
      });
    break;
    default:
      var ops = ["get #", "who #", "set x"];
      console.log("Must call !quote with either " + ops.map(function (str) {
        return "'" + str + "'";
      }));
      res.status(400).send("Must call !quote with either " + ops.map(function (str) {
        return "'" + str + "'";
      }));
  }
});

app.listen(8090, function () {
  console.log("listening on", 8090);
})
