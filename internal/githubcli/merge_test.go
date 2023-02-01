package githubcli

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGhCli_Exec(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mock := NewMockGhCliExecutor(ctrl)
	mock.
		EXPECT().
		Exec("pr", "merge", "5", "--merge", "--subject", "my subject", "--body", "some body").
		Return(bytes.Buffer{}, bytes.Buffer{}, nil).
		Times(1)

	_ = Merge(mock, 5, "my subject", "some body", "merge")
}
