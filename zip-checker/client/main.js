$(document).ready(function() {
    var $container = $('#container');
    var $cityNameInput = document.getElementById('city-name');

    var zipForm = $('#zip-form');
    var cityName, url;

    var output;

    zipForm.submit(function(e) {
        e.preventDefault();
        $container.empty();

        cityName = $cityNameInput.value;
        url = '/zips/' + cityName;
        $.ajax({
            url: url
        })
            .done(function(data) {
                if (Array.isArray(data)) {
                    output = '<table><tbody><tr><th>Zip Code</th><th>City</th><th>State</th></tr>';
                    for (var i = 0; i < data.length; i++) {
                        output += '<tr>';
                        output += '<td>' + data[i].code + '</td>';
                        output += '<td>' + data[i].city + '</td>';
                        output += '<td>' + data[i].state + '</td>';
                        output += '</tr>';
                    }
                    output += '</tbody></table>';
                } else {
                    output = data;
                }
                $container.append(output);
            })
            .fail(function(error) {
                console.log(error);
            });
    });
});
