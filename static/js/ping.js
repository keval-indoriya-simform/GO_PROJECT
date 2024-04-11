function http_ping(fqdn) {

    var NB_ITERATIONS = 4; // number of loop iterations
    var MAX_ITERATIONS = 5; // beware: the number of simultaneous XMLHttpRequest is limited by the browser!
    var TIME_PERIOD = 1000; // 1000 ms between each ping
    var i = 0;
    var over_flag = 0;
    var time_cumul = 0;
    var REQUEST_TIMEOUT = 9000;
    var TIMEOUT_ERROR = 0;

    // document.getElementById('result').innerHTML = "HTTP ping for " + fqdn + "</br>";
    console.log( "HTTP ping for " + fqdn)
    var ping_loop = setInterval(function() {


        // let's change non-existent URL each time to avoid possible side effect with web proxy-cache software on the line
        url = "http://" + fqdn + "?random-no-cache=" + Math.random().toString(36).substring(7);

        if (i < MAX_ITERATIONS) {

            var ping = new XMLHttpRequest();

            i++;
            ping.seq = i;
            over_flag++;

            ping.date1 = Date.now();

            ping.timeout = REQUEST_TIMEOUT; // it could happen that the request takes a very long time

            ping.onreadystatechange = function() { // the request has returned something, let's log it (starting after the first one)

                if (ping.readyState == 4 && TIMEOUT_ERROR == 0) {

                    over_flag--;

                    if (ping.seq > 1) {
                        delta_time = Date.now() - ping.date1;
                        time_cumul += delta_time;
                        // document.getElementById('result').innerHTML += "</br>http_seq=" + (ping.seq-1) + " time=" + delta_time + " ms</br>";
                        console.log("http_seq=" + (ping.seq-1) + " time=" + delta_time + "ms")
                    }
                }
            }


            ping.ontimeout = function() {
                TIMEOUT_ERROR = 1;
            }

            ping.open("GET", url, true);
            ping.send();

        }

        if ((i > NB_ITERATIONS) && (over_flag < 1)) { // all requests are passed and have returned

            clearInterval(ping_loop);
            var avg_time = Math.round(time_cumul / (i - 1));
            // document.getElementById('result').innerHTML += "</br> Average ping latency on " + (i-1) + " iterations: " + avg_time + "ms </br>";
            console.log("Average ping latency on " + (i-1) + " iterations: " + avg_time + "ms ")

        }

        if (TIMEOUT_ERROR == 1) { // timeout: data cannot be accurate

            clearInterval(ping_loop);
            // document.getElementById('result').innerHTML += "<br/> THERE WAS A TIMEOUT ERROR <br/>";
            console.log("THERE WAS A TIMEOUT ERROR ")
            return;

        }

    }, TIME_PERIOD);
}