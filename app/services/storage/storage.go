package storage

import "go-observatory/app/models/pinger"

var RawBucket []string
var Bucket []pinger.Pinger

func Place(data []string) {
	RawBucket = append(RawBucket, data...)
}

func Transform() {
	// for _, piace := range RawBucket {
	// 	// map data
	// }
}
