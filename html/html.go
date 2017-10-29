package html

import (
	"html/template"
	"io"

	"github.com/posener/eztables/table"
)

// Write writes a Table struct to an html page
func Write(w io.Writer, t *table.Table) error {
	return tmplt.Execute(w, t)
}

var tmplt = template.Must(template.New("table").Parse(`
<!doctype html>
<html lang="en">

<head>
	<title>eztables</title>
	<meta charset="utf-8">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css" integrity="sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb" crossorigin="anonymous">
</head>

<body>

<nav class="navbar navbar-light" style="background-color: #e3f2fd;">
	<a class="navbar-brand" href="/">eztables</a>
	<span class="navbar-text">
		<a href="https://github.com/posener/eztables">about</a>
	</span>
</nav>



<div class="container">

	<div class="row">

		<div class="col-3">
			<h2>Chains</h2>
			<ul class="nav flex-column nav-pills" id="chainTabs" role="tablist">
			{{range $i, $chain := .}}
				<li class="nav-item">
					<a class="nav-link{{if eq $i 0}} active{{end}}"
						id="{{$chain.Name}}-tab"
						data-toggle="pill"
						href="#{{$chain.Name}}"
						role="tab"
						aria-controls="{{$chain.Name}}"
						aria-selected="{{if eq $i 0}}true{{else}}false{{end}}"
						onclick="location.hash='#{{$chain.Name}}'">

						{{$chain.Name}}
					</a>
				</li>
			{{end}}
			</ul>
		</div>

		<div class="col-9">
			<div class="tab-content">
			{{range $i, $chain := .}}

				<div class="tab-pane fade{{if eq $i 0}} show active{{end}}"
					id="{{$chain.Name}}"
					role="tabpanel"
					aria-labelledby="{{$chain.Name}}-tab">

					<table class="table table-hover">
						<tr>
							<th>match</th>
							<th>target</th>
							<th>target args</th>
							<th>packets</th>
							<th>bytes</th>
						</tr>
						{{range $chain.Rules}}
						<tr class="{{if .Positive}}table-success{{else if .Negative}}table-danger{{end}}">
							<td>{{.Match}}</td>
							{{if or .Positive .Negative}}
							<td>{{.Target}}</td>
							{{else}}
							<td><a href="#{{.Target}}">{{.Target}}</a></td>
							{{end}}
							<td>{{.TargetArgs}}</td>
							<td>{{.Count.Packets}}</td>
							<td>{{.Count.Bytes}}</td>
						</tr>
						{{end}}
					</table>
				</div>
			{{end}}
			</div>
		</div>
	</div>
</div>

<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js" integrity="sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh" crossorigin="anonymous"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js" integrity="sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ" crossorigin="anonymous"></script>

<script>
	$('#chainTabs a[href="' + location.hash + '"]').tab('show')
	$(window).bind('hashchange', function() {
		console.log(location.hash);
		$('#chainTabs a[href="' + location.hash + '"]').tab('show')
	});
</script>

</body>
</html>
`))
