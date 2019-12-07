package main

var templateHTML = `<html>

<head>
    <meta charset="UTF-8">
    <meta name="robots" content="noindex,nofollow,noarchive,nosnippet,noodp,noydir">
    <title>{{.Title}}</title>
    <link href="https://ajax.googleapis.com/ajax/static/modules/gviz/1.0/core/tooltip.css" rel="stylesheet" type="text/css">
</head>

<body>

    <script type="text/javascript" src="https://www.google.com/jsapi"></script>
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
    <script type="text/javascript">
        $(window).resize(function() {
            drawChart2o();
        });

        // Load the Visualization API and the chart package.
        google.load('visualization', '1.0', {
            'packages': ['corechart']
        });

        // Set a callback to run when the Google Visualization API is loaded.
        google.setOnLoadCallback(drawChart2o);

        // Callback that creates and populates a data table,
        // instantiates the chart, passes in the data and
        // draws it.
        function drawChart2o() {
            // Create the data table.
            var data = google.visualization.arrayToDataTable({{.Data}});

            // Set chart options
            var options = {
                title: '{{.Title}}',
                curveType: 'function',
                legend: {
                    position: 'bottom'
                }
            };

            // Instantiate and draw our chart, passing in some options.
            var chart = new google.visualization.LineChart(document.getElementById('chart_div_2o'));
            chart.draw(data, options);
        }
    </script>
    <script src="https://www.google.com/uds/?file=visualization&amp;v=1.0&amp;packages=corechart" type="text/javascript"></script>
    <link href="https://www.google.com/uds/api/visualization/1.0/40ff64b1d9d6b3213524485974f36cc0/ui+en.css" type="text/css" rel="stylesheet">
    <script src="https://www.google.com/uds/api/visualization/1.0/40ff64b1d9d6b3213524485974f36cc0/format+en,default+en,ui+en,corechart+en.I.js" type="text/javascript"></script>

    <div id="chart_div_2o" style="height:600px;width:100%;"></div>

</body>

</html>`
