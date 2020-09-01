// Copyright 2020 ConsenSys AG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by goff (v0.3.2) DO NOT EDIT

package bls381

import (
	"golang.org/x/sys/cpu"
)

// supportAdx will be set only on amd64 that has MULX and ADDX instructions
var supportAdx = cpu.X86.HasADX && cpu.X86.HasBMI2

// q (modulus)
var qE2 = [6]uint64{
	13402431016077863595,
	2210141511517208575,
	7435674573564081700,
	7239337960414712511,
	5412103778470702295,
	1873798617647539866,
}

// q'[0], see montgommery multiplication algorithm
var qE2Inv0 uint64 = 9940570264628428797

//go:noescape
func addE2(res, x, y *E2)

//go:noescape
func subE2(res, x, y *E2)

//go:noescape
func doubleE2(res, x *E2)

//go:noescape
func negE2(res, x *E2)
