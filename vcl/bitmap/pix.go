//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build !darwin

package bitmap

//var (
//	// B G R A
//	// Little Endian
//	pixIndex = [4]int{2, 1, 0, 3}
//)

// Little Endian
type bgra struct {
	B, G, R, A uint8
}
