package cmd

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// func TestCmdDelivery_TaxiFares(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockTaxiFaresUc := mockUC.NewMockITaxeFares(ctrl)

// 	tests := []struct {
// 		name  string
// 		input string
// 		d     *CmdDelivery
// 		mock  func()
// 	}{
// 		{
// 			name:  "",
// 			input: "10:00:00.000 2.5\n10:05:00.000 3.1\n10:10:00.000 4.0\n\n",
// 			d: &CmdDelivery{
// 				taxiFaresUc: mockTaxiFaresUc,
// 			},
// 			mock: func() {
// 				mockTaxiFaresUc.EXPECT().CalculateFare(gomock.Any()).Return(0.0)
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()

// 			// Redirect os.Stdin to a pipe
// 			oldStdin := os.Stdin
// 			defer func() { os.Stdin = oldStdin }()

// 			r, w, _ := os.Pipe()
// 			os.Stdin = r

// 			// Write the test input to the pipe
// 			_, _ = w.Write([]byte(tt.input))
// 			_ = w.Close()

// 			tt.d.TaxiFares()
// 		})
// 	}
// }

func TestCmdDelivery_TaxiFares(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	msg := "Your name please? Press the Enter key when done"
	fmt.Fprintln(os.Stdout, msg)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	name := scanner.Text()
	if len(name) == 0 {
		t.Log("empty input")
	}
	t.Logf("You entered: %s\n", name)
}
