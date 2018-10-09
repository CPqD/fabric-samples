var app = require('./config/server');
const PORT = 8080;
app.listen(PORT, function(){
	console.log('Web application is up and running on port ' + PORT);
});