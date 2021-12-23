mapboxgl.accessToken = 'pk.eyJ1IjoiamluZ2ZlbmdsaSIsImEiOiJja3hieWlycGE0MzVzMnNvNnR0cnd6eHZoIn0.LH-XpbouONSC1PaztPB19w';
const map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v11',
    center: [116.39111557492419, 39.90584776852543],
    zoom: 13
});

map.addControl(new MapboxLanguage({
    defaultLanguage: 'zh-Hans'
}));

const marker1 = new mapboxgl.Marker({ color: 'green' })
const marker2 = new mapboxgl.Marker({ color: 'orange' })

const host = "http://localhost:9000"
const location_server = host + "/locations/"
const direction_server = host + "/direction/"

map.on('load', () => {
    map.addSource('route', {
        'type': 'geojson',
        'data': {
            'type': 'Feature',
        }
    });
    map.addLayer({
        'id': 'route',
        'type': 'line',
        'source': 'route',
        'layout': {
            'line-join': 'round',
            'line-cap': 'round'
        },
        'paint': {
            'line-color': '#1E90FF',
            'line-width': 4
        }
    });
});


function Navigate() {
    sloc = document.getElementById("sloc").value
    dloc = document.getElementById("dloc").value

    var slat, slon, dlat, dlon

    $.when($.get(location_server, {
        name: sloc,
    }), $.get(location_server, {
        name: dloc,
    })).done(function (res1, res2) {
        if (res1[0].hasOwnProperty("error")) {
            alert("出发地不存在！")
            return
        }
        if (res2[0].hasOwnProperty("error")) {
            alert("目的地不存在！")
            return
        }
        slat = res1[0].lat
        slon = res1[0].lon
        dlat = res2[0].lat
        dlon = res2[0].lon
        $.get(direction_server, {
            slat: slat,
            slon: slon,
            dlat: dlat,
            dlon: dlon,
        }, function (data) {
            // Create a default Marker and add it to the map.
            marker1.setLngLat([slon, slat])
                .addTo(map);
            marker2.setLngLat([dlon, dlat])
                .addTo(map);

            if (data.nodes != null) {
                map.getSource('route').setData({
                    'type': 'Feature',
                    'properties': {},
                    'geometry': {
                        'type': 'LineString',
                        'coordinates': data.nodes
                    }
                });
            }
            document.getElementById("direction").innerHTML = "<div class='direction-text'>" + data.text + "</div>";
        });
    })
}

function Search(ele) {
    var keyword = ele.value
    if (keyword != '') {
        $.get(location_server, {
            prefix: ele.value
        }, function (data) {
            $('#node_list').empty()
            $.each(data.names, function (index, item) {
                if (ele.value != item) {
                    $('#node_list').append('<option value="' + item + '"/>')
                }
            })
        });
    }
}
