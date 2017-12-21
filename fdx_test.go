//
// fdx is a package encoding/decoding fdx formatted XML files.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// BSD 2-Clause License
//
// Copyright (c) 2017, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package fdx

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

var (
	expectedDocs map[string][]byte
)

func fdxFile(t *testing.T, fname string) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "skipping %s, %s", fname, err)
		return
	}
	fdx := new(FinalDraft)
	if err := xml.Unmarshal(src, &fdx); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	} else {
		os.RemoveAll(path.Join("testout", path.Base(fname)))
		if src2, err := xml.MarshalIndent(fdx, " ", "    "); err != nil {
			t.Errorf("%s", err)
		} else {
			if err := ioutil.WriteFile(path.Join("testout", path.Base(fname)), src2, 0666); err != nil {
				t.Errorf("%s", err)
			}
		}
	}
}

func TestConversion(t *testing.T) {
	fdxFile(t, "testdata/testplay-01a.fdx")
	fdxFile(t, "testdata/testplay-01b.fdx")
	fdxFile(t, "testdata/Big Fish.fdx")
}

func TestMain(m *testing.M) {
	// Setup everything, process flags, etc.
	os.Exit(m.Run())
}
