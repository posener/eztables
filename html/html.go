package html

import (
	"html/template"
	"io"

	"github.com/posener/eztables/table"
)

type data struct {
	Table  table.Table
	Others []table.Table
}

// Write writes a Table struct to an html page
func Write(w io.Writer, t table.Table, tables []table.Table) error {
	return tmplt.Execute(w, data{Table: t, Others: tables})
}

var tmplt = template.Must(template.New("table").Parse(`
<!doctype html>
<html lang="en">

<head>
	<title>eztables</title>
	<meta charset="utf-8">
	<link rel="stylesheet" href="/static/bootstrap-4.0.0-beta.2.min.css">
</head>

<body>

<nav class="navbar navbar-expand-lg navbar-light bg-light">
  <a class="navbar-brand" href="/">eztables</a>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>

  <div class="collapse navbar-collapse" id="navbarSupportedContent">
    <ul class="navbar-nav mr-auto">
      <li class="nav-item dropdown" >
        <a
			class="nav-link dropdown-toggle"
			href="#"
			id="navbarDropdown"
			role="button"
			data-toggle="dropdown"
			aria-haspopup="true"
			aria-expanded="false">
		Table: {{.Table.Name}}
        </a>
        <div class="dropdown-menu" aria-labelledby="navbarDropdown">
		{{range $_, $table := .Others}}
          <a class="dropdown-item" href="/tables/{{$table.Name}}" {{$table.ToolTipAttributes}}>{{$table.Name}}</a>
		{{end}}
        </div>
      </li>
    </ul>
	<span class="navbar-text">
		<a href="https://github.com/posener/eztables">about</a>
	</span>
  </div>
</nav>

<div class="container">

	<div class="row">

		<div class="col-3">
			<ul class="nav flex-column nav-pills" id="chainTabs" role="tablist">
			{{range $i, $chain := .Table.Chains}}
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
			{{range $i, $chain := .Table.Chains}}

				<div class="tab-pane fade{{if eq $i 0}} show active{{end}}"
					id="{{$chain.Name}}"
					role="tabpanel"
					aria-labelledby="{{$chain.Name}}-tab"
					style="overflow-y:scroll;height:90vh"
					>

					<table class="table table-hover">
						<thead>
						<tr>
							<th>match</th>
							<th>target</th>
							<th>packets</th>
							<th>bytes</th>
						</tr>
						</thead>
						<tbody>
						{{range $chain.Rules}}
						<tr class="{{if .Positive}}table-success{{else if .Negative}}table-danger{{end}}">
							<td>
								{{if .Match}}
									<ul class="list-inline">
									{{range $_, $arg := .Match}}
										<li class="list-inline-item">
											<span class="badge badge-info" {{$arg.ToolTipAttributes}}>
												{{$arg}}
											</span>
										</li>
									{{end}}
									</ul>
								{{end}}
							</td>
							<td>
								{{if or .TargetIsChain}}
								<a href="#{{.Target}}">{{.Target}}</a>
								{{else}}
								{{.Target}}
								{{end}}
								{{if .TargetArgs}}
									<br>
									{{range $_, $arg := .TargetArgs}}
										<li class="list-inline-item">
											<span class="badge badge-info" {{$arg.ToolTipAttributes}}>
												{{$arg}}
											</span>
										</li>
									{{end}}
								{{end}}
							</td>
							<td>{{.Count.Packets}}</td>
							<td>{{.Count.Bytes}}</td>
						</tr>
						{{end}}
						</tbody>
					</table>
				</div>
			{{end}}
			</div>
		</div>
	</div>
</div>

<script src="/static/jquery-3.2.1.slim.min.js"></script>
<script src="/static/popper-1.12.3.min.js"></script>
<script src="/static/bootstrap-4.0.0-beta.2.min.js"></script>

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
