package falcata

import "testing"

type cutFieldsTest struct {
    arg1 string
	expected string
	expectedLength int
}

var target_1 = `I am a student.`
var target_2 = `I am a   student.`
var target_3 = `I am 
a  
student.`
var target_4 = ` I am   
a
student.`
var target_5 = `  I am 
a 
student.`
var target_6 = `    I am 
a 
student.`
var target_7 = `I am 
a 
student. `
var target_8 = `I am 
a 
student.  `
var target_9 = `I am 
a 
student.    `
var target_10 = `I am 
a student.  `
var expected = `I am a student.`
var expectedLength = 15

var target2_1 = ""
var target2_2 = " "
var target2_3 = "  "
var target2_4 = "    "
var target2_5 = `
`
var target2_6 = `
 `
var expected2 = ""
var expectedLength2 = 0

var cutFieldsTests = []cutFieldsTest{
    cutFieldsTest{target_1, expected, expectedLength},
    cutFieldsTest{target_2, expected, expectedLength},
    cutFieldsTest{target_3, expected, expectedLength},
    cutFieldsTest{target_4, expected, expectedLength},
    cutFieldsTest{target_5, expected, expectedLength},
    cutFieldsTest{target_6, expected, expectedLength},
    cutFieldsTest{target_7, expected, expectedLength},
    cutFieldsTest{target_8, expected, expectedLength},
    cutFieldsTest{target_9, expected, expectedLength},
    cutFieldsTest{target_10, expected, expectedLength},
    cutFieldsTest{target2_1, expected2, expectedLength2},
    cutFieldsTest{target2_2, expected2, expectedLength2},
    cutFieldsTest{target2_3, expected2, expectedLength2},
    cutFieldsTest{target2_4, expected2, expectedLength2},
    cutFieldsTest{target2_5, expected2, expectedLength2},
    cutFieldsTest{target2_6, expected2, expectedLength2},
}

func TestCutFields(t *testing.T){

    for _, test := range cutFieldsTests{
		var copied = test.arg1
		CutFields(&test.arg1)
        if (test.arg1 != test.expected || len(test.arg1) != test.expectedLength) {
            t.Errorf("Input %q, Output %q not equal to expected %q or Output length %q not equal to expectedLength %q", copied, test.arg1, test.expected, len(test.arg1), test.expectedLength)
        }
    }
}
