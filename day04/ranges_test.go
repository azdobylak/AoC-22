package main

import "testing"

func TestRangeContains(t *testing.T) {
	r1 := &SectionsRange{min: 2, max: 6}
	r2 := &SectionsRange{min: 7, max: 9}

	if r1.Contains(r2) {
		t.Errorf("%#v does not contain %#v", r1, r2)
	}
	if r2.Contains(r1) {
		t.Errorf("%#v does not contain %#v", r2, r1)
	}

	r3 := &SectionsRange{min: 3, max: 6}

	if !r1.Contains(r3) {
		t.Errorf("%#v should contain %#v", r1, r3)
	}
}

func TestRangeOverlaps(t *testing.T) {
	r1 := &SectionsRange{min: 2, max: 6}
	r2 := &SectionsRange{min: 7, max: 9}
	r3 := &SectionsRange{min: 9, max: 10}
	r4 := &SectionsRange{min: 3, max: 4}

	if r1.Overlaps(r2) {
		t.Errorf("%#v and %#v should not have overlap", r1, r2)
	}
	if !r2.Overlaps(r3) {
		t.Errorf("%#v and %#v should have overlap", r2, r3)
	}
	if !r1.Overlaps(r4) {
		t.Errorf("%#v and %#v should have overlap", r1, r4)
	}
	if !r4.Overlaps(r1) {
		t.Errorf("%#v and %#v should have overlap", r1, r4)
	}
}
