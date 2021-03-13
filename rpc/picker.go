package rpc

import (
	"math"
	"math/rand"
	"time"
)

func defaultPicker(servers []*ServerForPick) *ServerForPick {
	if len(servers) == 0 {
		return nil
	}
	if len(servers) == 1 {
		if servers[0].Pickable() {
			return servers[0]
		}
		return nil
	}
	start := rand.Intn(len(servers))
	i := start
	first := true
	var normal1, normal2, danger1, danger2, nightmare1, nightmare2 *ServerForPick
	before := time.Now().Add(-time.Millisecond * 100)
	for {
		if !first && i == start {
			break
		}
		first = false
		if servers[i].Pickable() {
			if servers[i].Pickinfo.DServers != 0 &&
				servers[i].Pickinfo.DServerOffline < before.UnixNano() {
				if normal1 == nil {
					normal1 = servers[i]
				} else {
					normal2 = servers[i]
					break
				}
			} else if servers[i].Pickinfo.DServers == 0 {
				if nightmare1 == nil {
					nightmare1 = servers[i]
				} else if nightmare2 == nil {
					nightmare2 = servers[i]
				}
			} else {
				if danger1 == nil {
					danger1 = servers[i]
				} else if danger2 == nil {
					danger2 = servers[i]
				}
			}
		}
		i++
		if i >= len(servers) {
			i = 0
		}
	}
	//check normal
	if normal1 != nil && normal2 == nil {
		return normal1
	} else if normal2 != nil && normal1 == nil {
		return normal2
	} else if normal1 == nil && normal2 == nil {
		//check danger
		if danger1 != nil && danger2 == nil {
			return danger1
		} else if danger2 != nil && danger1 == nil {
			return danger2
		} else if danger1 != nil && danger2 != nil {
			normal1 = danger1
			normal2 = danger2
		} else {
			//check nightmare
			if nightmare1 != nil && nightmare2 == nil {
				return nightmare1
			} else if nightmare2 != nil && nightmare1 == nil {
				return nightmare2
			} else if nightmare1 != nil && nightmare2 != nil {
				normal1 = nightmare1
				normal2 = nightmare2
			} else {
				//all servers are unpickable
				return nil
			}
		}
	}
	//more discoveryservers more safety,so 1 * 2's discoveryserver num
	load1 := float64(normal1.Pickinfo.Activecalls) * math.Log(float64(normal2.Pickinfo.DServers+2))
	if normal1.Pickinfo.Lastfail >= before.UnixNano() {
		//punish
		load1 *= 1.1
	}
	//more discoveryservers more safety,so 2 * 1's discoveryserver num
	load2 := float64(normal2.Pickinfo.Activecalls) * math.Log(float64(normal1.Pickinfo.DServers+2))
	if normal2.Pickinfo.Lastfail >= before.UnixNano() {
		//punish
		load2 *= 1.1
	}
	if load1 > load2 {
		return normal2
	} else if load1 < load2 {
		return normal1
	} else if rand.Intn(2) == 0 {
		return normal1
	} else {
		return normal2
	}
}
