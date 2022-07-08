{{template "base" .}}

{{define "content"}}
    <div class="container">
        <div class="row">
            <div class="col">
                <h1>This is the About Page</h1>
                <p>Some text...</p>
                <p>This came from the template: {{index .StringMap "test"}}</p>

                <p>
                    {{if ne (index .StringMap "remote_ip") ""}}
                        Your remote IP is: {{index .StringMap "remote_ip"}}
                    {{else}}
                        I don't know your IP address yet. Visit the <a href="/">Home</a> page so I can set it.
                    {{end}}
                </p>
            </div>
        </div>
    </div>
{{end}}