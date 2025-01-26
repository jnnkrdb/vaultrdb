package storedb

import (
	"testing"
)

// ########################################################################################### calculateBucketsFromPath

func Test_calculateBucketsFromPath(t *testing.T) {

	var tests = []struct {
		name      string
		input     string
		estimated []string
	}{

		{"no level bucket", "", nil},
		{"no level bucket, with space", " ", nil},
		{"no level bucket, with .", ".", nil},
		{"no level bucket, with . and spaces", " .  ", nil},

		{"single level bucket", "singlelevelpath", []string{"singlelevelpath"}},
		{"single level bucket, with prefix .", ".singlelevelpath", []string{"singlelevelpath"}},
		{"single level bucket, with suffix .", "singlelevelpath.", []string{"singlelevelpath"}},
		{"single level bucket, with prefix and suffix .", ".singlelevelpath.", []string{"singlelevelpath"}},

		{"multi level bucket", "multi.level.path", []string{"multi", "level", "path"}},
		{"multi level bucket, with prefix .", ".multi.level.path", []string{"multi", "level", "path"}},
		{"multi level bucket, with suffix .", "multi.level.path.", []string{"multi", "level", "path"}},
		{"multi level bucket, with prefix and suffix .", ".multi.level.path.", []string{"multi", "level", "path"}},

		{"single level bucket, with whitespaces before bucketname", " singlelevelpath", []string{"singlelevelpath"}},
		{"single level bucket, with whitespaces in bucketname", "single levelpath", []string{"singlelevelpath"}},
		{"single level bucket, with whitespaces after bucketname", "singlelevelpath ", []string{"singlelevelpath"}},
		{"single level bucket, with whitespaces everywhere", " single  level  path ", []string{"singlelevelpath"}},

		{"multi level bucket, with whitespaces before .", "multi .level .path", []string{"multi", "level", "path"}},
		{"multi level bucket, with whitespaces around .", "multi . level . path", []string{"multi", "level", "path"}},
		{"multi level bucket, with whitespaces after .", "multi. level. path", []string{"multi", "level", "path"}},
		{"multi level bucket, with whitespaces before first .", " .multi.level.path", []string{"multi", "level", "path"}},
		{"multi level bucket, with whitespaces after last .", ".multi.level.path. ", []string{"multi", "level", "path"}},
		{"multi level bucket, with whitespaces before first and after last .", " .multi.level.path. ", []string{"multi", "level", "path"}},
	}

	// running the tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testResult := calculateBucketsFromPath(tt.input)

			t.Logf("path: [%s] estimatedResult: %v - testResult: %v", tt.input, tt.estimated, testResult)

			// check if the length of the arrays are the same
			if len(testResult) != len(tt.estimated) {
				t.Fatalf("lengths of the array are not the same: estimatedResult[%d] != testResult[%d]", len(tt.estimated), len(testResult))
			}

			// parse through the array and check the corresponding value
			for index := range testResult {
				if testResult[index] != tt.estimated[index] {
					t.Fatalf("testResult[%d] is not estimatedResult[%d] -> %s != %s", index, index, testResult[index], tt.estimated[index])
				}
			}
		})
	}
}
