package main

const html = `<html lang="en">
	<head>
	  <meta charset="utf-8">
	  <title>Coyote Tester - Results</title>
	</head>
        <body>
	  <style type="text/css">
	    .tg  {border-collapse:collapse;border-spacing:0;border-color:#aaa;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:0px;overflow:hidden;word-break:normal;border-color:#aaa;color:#333;background-color:#fff;border-top-width:1px;border-bottom-width:1px;}
            .tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:0px;overflow:hidden;word-break:normal;border-color:#aaa;color:#fff;background-color:#f38630;border-top-width:1px;border-bottom-width:1px;}
            .tg .tg-j2zy{background-color:#FCFBE3;vertical-align:top}
            .tg .tg-zczf{background-color:#FCFBE3;text-align:left;vertical-align:top}
            .tg .tg-lqy6{text-align:left;vertical-align:top}
            .tg .tg-9hbo{font-weight:bold;vertical-align:top}
            .tg .tg-b5xm{background-color:#fe0000;text-align:center;vertical-align:top}
            .tg .tg-yw4l{vertical-align:top}
            .tg .tg-y0xi{background-color:#32cb00;text-align:center;vertical-align:top}
            .hideContent {overflow:hidden;line-height:1em;height:2em;}
            .showContent {line-height:1em;height:auto;}
          </style>
          <table class="tg">
            <tr>
              <th class="tg-9hbo">Test</th>
              <th class="tg-9hbo">Status</th>
              <th class="tg-9hbo">Time (sec)</th>
              <th class="tg-9hbo">Error</th>
              <th class="tg-9hbo">Command</th>
              <th class="tg-9hbo">StdOut</th>
              <th class="tg-9hbo">StdErr</th>
            </tr>
            {{ range $k, $v := . }}
	    <tr>
              {{ if isEven $k }}
	      <td class="tg-j2zy">{{ $v.Name }}</td>
	      <td class="{{ if eq $v.Status "ok"}}tg-y0xi{{ else }}tg-b5xm{{ end }}">{{ $v.Status }}</td>
	      <td class="tg-zczf">{{ $v.Time | printf "%.3f" }}</td>
	      <td class="tg-zczf">{{ $v.Error }}</td>
	      <td class="tg-zczf">{{ $len := len $v.Command }}{{ if gt $len 0 }}{{ $x := splitString $v.Command }}
                {{ if showmore $x }}
                <div class="content hideContent">{{ range $x }}{{ . }}</br>{{ end }}</div>
	        <div class="show-more"><a href="#">Show more</a></div>
                {{ else }}
                {{ range $x }}{{ . }}</br>{{ end }}
                {{ end }}
                {{ end }}
              </td>
	      <td class="tg-zczf">{{ if showmore $v.Stdout }}
                <div class="content hideContent">{{ range $v.Stdout }}{{ . }}</br>{{ end }}</div><div class="show-more"><a href="#">Show more</a></div>
                {{ else }}
                {{ range $v.Stdout }}{{ . }}</br>{{ end }}
                {{ end }}
              </td>
	      <td class="tg-zczf">{{ if showmore $v.Stderr }}
                <div class="content hideContent">{{ range $v.Stderr }}{{ . }}</br>{{ end }}</div><div class="show-more"><a href="#">Show more</a></div>
                {{ else }}
                {{ range $v.Stderr }}{{ . }}</br>{{ end }}
                {{ end }}
              </td>
              {{ else }}
              <td class="tg-yw4l">{{ $v.Name }}</td>
              <td class="{{ if eq $v.Status "ok"}}tg-y0xi{{ else }}tg-b5xm{{ end }}">{{ $v.Status }}</td>
	      <td class="tg-lqy6">{{ $v.Time | printf "%1.3f" }}</td>
	      <td class="tg-lqy6">{{ $v.Error }}</td>
	      <td class="tg-lqy6">{{ $len := len $v.Command }}{{ if gt $len 0 }}{{ $x := splitString $v.Command }}
                {{ if showmore $x }}
                <div class="content hideContent">{{ range $x }}{{ . }}</br>{{ end }}</div>
	        <div class="show-more"><a href="#">Show more</a></div>
                {{ else }}
                {{ range $x }}{{ . }}</br>{{ end }}
                {{ end }}
                {{ end }}
              </td>
	      <td class="tg-lqy6">{{ if showmore $v.Stdout }}
                <div class="content hideContent">{{ range $v.Stdout }}{{ . }}</br>{{ end }}</div><div class="show-more"><a href="#">Show more</a></div>
                {{ else }}
                {{ range $v.Stdout }}{{ . }}</br>{{ end }}
                {{ end }}
              </td>
	      <td class="tg-lqy6">{{ if showmore $v.Stderr }}
                <div class="content hideContent">{{ range $v.Stderr }}{{ . }}</br>{{ end }}</div><div class="show-more"><a href="#">Show more</a></div>
                {{ else }}
                {{ range $v.Stderr }}{{ . }}</br>{{ end }}
                {{ end }}
              </td>
              {{ end }}
	    </tr>
            {{ end }}
          </table>
          <script src="https://code.jquery.com/jquery-1.12.2.min.js"></script>
          <script src="https://code.jquery.com/ui/1.11.4/jquery-ui.min.js"></script>
          <!-- http://stackoverflow.com/questions/12307666/how-to-create-a-show-more-button-and-specify-how-many-lines-of-text-can-be-ini -->
          <script>
            $(".show-more a").on("click", function() {
            var $this = $(this);
            var $content = $this.parent().prev("div.content");
            var linkText = $this.text().toUpperCase();

            if(linkText === "SHOW MORE"){
            linkText = "Show less";
            $content.switchClass("hideContent", "showContent", 400);
            } else {
            linkText = "Show more";
            $content.switchClass("showContent", "hideContent", 400);
            };

            $this.text(linkText);
            });
          </script>
        </body>
</html>`
