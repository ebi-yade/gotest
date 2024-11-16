# gotest

Unit Testing Utilities for Go

## Installation

```sh
go get github.com/ebi-yade/gotest/cases
```

## Usage

### Error Check

Works as a wrapper for [testify/require](https://pkg.go.dev/github.com/stretchr/testify/require) in table-driven test cases.

```go
package foo_test

import (
	"testing"

	"github.com/ebi-yade/gotest/cases"
)

func TestFoo(t *testing.T) {
	tests := []struct {
		name     string
		input    foo.Input
		errCheck cases.ErrorCheck // func(*testing.T, error)
	}{
		{
			name: "case1",
			input: foo.Input{
				Amout: 1,
			},
			errCheck: cases.NoError,
		},
		{
			name: "case2",
			input: foo.Input{
				Amout: 999,
			},
			errCheck: cases.ErrorIs(ErrTooLargeAmount),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := foo.DoSomething(tt.input)
			tt.errCheck(t, err)
		})
	}
}
```

<details>
<summary>Memo</summary>

### ここまで書いて気づいたこと

- assert.NoError は bool を返すが、 require.NoError は何も返さない
- t.Run するテーブル駆動テストであれば require にわたる t は無名関数のスコープ
- エラーチェックをする場合は基本的に　require でいいはず（暗黙直和）
- ErrorIs などはシグニチャが異なるのでラッパーの有り難みがある

</details>
