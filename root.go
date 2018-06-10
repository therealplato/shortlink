package main

var root = []byte(`<!DOCTYPE html>
<html>
<head><title>short.link</title>
<style>
	body, input{
		max-width:80em;
		margin: auto;
		display: block;
	}
	input[type="text"] {
		width: 100%;
		margin: 5em auto 1em auto;
	}
	input[type="submit"] {
		width: 20%;
    height: 4em;
    font-size: 1.5em;
	}
</style>
</head>
<body>
	<form action="/create/">
		<input type="text" name="destination" placeholder="https://github.com/therealplato/shortlink/issues">
		<input type="submit" value="Shorten">
	</form>
</body>
</html>`)
