console.log("started\r\n");
return;
// var Firebase = require("firebase");
var admin = require("firebase-admin");

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

var operation = process.argv[2];
var value = process.argv[3];

console.log("operation", operation);
console.log("value", value);

switch (operation) {
  case "get":
    if(value.match(/[0-9]/g)) {
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
        }
      }, function () {
        console.log("done get");
        console.log("done");
      });
    } else {
      console.log("err");
    }
  break;
  case "set":
    ref.quotesRef
    .orderByKey()
    .limitToLast(1)
    .once("value")
    .then(function (snap) {
      var quotes = snap.val();
      if(quotes) {
        var key = Object.keys(quotes)[0];
        var val = quotes[key];

        console.log("got quote", key, val);

        if(value.match(/[a-z]/gi)) {
          ref.quotesRef
          .push()
          .set({
            quoteNumber: val ? val.quoteNumber + 1 : 0,
            quote: value
          }, function () {
            console.log("done set");
            console.log("done");
          });
        } else {
          console.log("err");
        }
      }
    });
  break;
}
