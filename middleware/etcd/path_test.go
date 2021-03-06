package etcd

import "testing"

func TestPath(t *testing.T) {
	for _, path := range []string{"mydns", "skydns"} {
		e := Etcd{PathPrefix: path}
		result := e.Path("service.staging.skydns.local.")
		if result != "/"+path+"/local/skydns/staging/service" {
			t.Errorf("Failure to get domain's path with prefix: %s", result)
		}
	}
}
