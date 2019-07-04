## go-streaming-httpserver dir tree

```
.
├── README.md
├── controllers             
│   ├── chat                ＃聊天室
│   │   ├── chatroom.go
│   │   └── client.go
│   └── webpage             ＃頁面route
│       └── getweb.go
├── main.go
├── routers
│   └── routers.go          #route管理
├── static
│   ├── RtmpStreamer.as     
│   ├── RtmpStreamer.swf    #rtmp h264 flash推拉流
│   ├── main.js             #推拉流頁面js
│   ├── require.js
│   └── rtmp-streamer.min.js
└── view
    ├── mplay_flv.html      #flv player
    ├── mplay_m3u8.html     #hls player
    ├── webplay.html        #rtmp player
    └── webpush.html        #rtmp push
```

