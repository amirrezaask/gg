$("#generate").on('click', function() {
    var template_name = document.getElementById('template').value;
    console.log('https://raw.githubusercontent.com/amirrezaask/gg/master/'+template_name);
    fetch('https://raw.githubusercontent.com/amirrezaask/gg/master/'+template_name).then(response=> {
        response.text().then(template => {
            var text = Mustache.render(template, JSON.parse(document.getElementById("args").value));
            document.getElementById('output').innerText = text;
        });
    });
});

