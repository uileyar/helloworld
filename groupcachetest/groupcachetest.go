// groupcachetest.go
package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/golang/glog"

	"github.com/golang/groupcache"
)

func main() {

	flag.Set("logtostderr", "true")
	flag.Parse()

	me := ":8080"
	peers := groupcache.NewHTTPPool("http://localhost" + me)
	peers.Set("http://localhost:8081", "http://localhost:8082", "http://localhost:8083")

	helloworld := groupcache.NewGroup("helloworld", 1024*1024*1024*16, groupcache.GetterFunc(
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			glog.Infof("%v, key = %v", me, key)
			dest.SetString(key)
			return nil
		}))

	glog.Infof("GroupName: %v", helloworld.Name())
	http.HandleFunc("/xbox/",
		func(w http.ResponseWriter, r *http.Request) {
			parts := strings.SplitN(r.URL.Path[len("/xbox/"):], "/", 1)
			glog.Infof("parts: %v", parts)
			if len(parts) != 1 {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			var data []byte
			helloworld.Get(nil, parts[0], groupcache.AllocatingByteSliceSink(&data))
			w.Write(data)

			glog.Infof("data: %s", data)
			glog.Infof("Stats: %#v", helloworld.Stats)
			//glog.Infof("Gets: %v", helloworld.Stats.Gets.String())
			//glog.Infof("Load: %v", helloworld.Stats.Loads.String())
			//glog.Infof("LocalLoad: %v", helloworld.Stats.LocalLoads.String())
			//glog.Infof("PeerError: %v", helloworld.Stats.PeerErrors.String())
			//glog.Infof("PeerLoad: %v", helloworld.Stats.PeerLoads.String())
		})

	http.ListenAndServe(me, nil)
}
