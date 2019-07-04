require.config({
    paths: {
        "rtmp-streamer": "/static/rtmp-streamer.min"
    }
});

require(["rtmp-streamer"], function (RtmpStreamer) {

    var getUrl = function () {
        return document.getElementById('url').value;
    };

    var getName = function () {
        return document.getElementById('stream-name').value;
    };

    var streamer = new RtmpStreamer(document.getElementById('rtmp-streamer'));
    var player = new RtmpStreamer(document.getElementById('rtmp-player'));

    if (document.getElementById("play") != null)
        document.getElementById("play").addEventListener("click", function () {
            player.play(getUrl(), getName());
        });
    if (document.getElementById("publish") != null)
        document.getElementById("publish").addEventListener("click", function () {
            streamer.publish(getUrl(), getName());
        });
    if (document.getElementById("streamer-disconnect") != null)
        document.getElementById("streamer-disconnect").addEventListener("click", function () {
            streamer.disconnect();
        });
    if (document.getElementById("player-disconnect") != null)
        document.getElementById("player-disconnect").addEventListener("click", function () {
            player.disconnect();
        });

});