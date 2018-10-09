var express = require('express');
var consign = require('consign');

var app = express();


//Sets view engine and view files path.
app.set('view engine', 'ejs');
app.set('views', './app/views');

//Sets static content path
app.use(express.static('./app/public'));

consign()
	.include('app/routes')		
	.then('app/controllers')
	.into(app);

module.exports = app;