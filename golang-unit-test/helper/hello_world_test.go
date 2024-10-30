package helper

import (
	"testing"

	"fmt"
	"runtime"


	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//* Benchmark

//* go test -v -bench=. (if want run all benchmark)
//? go test -v -run=NotMathUnitTest -bench=. (if want run all benchmark without test)
//* go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorld (if want run a spesific benchmark without test)
//? go test -v -run=NotMathUnitTest -bench=. ./... (if want run all benchmark all package without test)

//* Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name:    "Faris",
			request: "Faris",
		},
		{
			name:    "Masud",
			request: "Masud",
		},
		{
			name:    "FarisMasud",
			request: "FarisMasud",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}


//! Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Faris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Faris")
		}
	})
	b.Run("Masud", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Masud")
		}
	})
}

func BenchmarkHelloWorldFaris(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Faris")
	}
}

func BenchmarkHelloWorldMasud(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Masud")
	}
}

//? Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Masud",
			request:  "Masud",
			expected: "Hello Masud",
		},
		{
			name:     "Faris",
			request:  "Faris",
			expected: "Hello Faris",
		},
		{
			name:     "FarisMasud",
			request:  "FarisMasud",
			expected: "Hello FarisMasud",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

//? Sub Test
// go test -v -run=TestSubTest (if want run all subtest)
// go test -v -run=TestSubTest/Faris (if want run specific subtest)
// go test -v -run=/Masud (if want run all subtest with name Masud)
func TestSubTest(t *testing.T) {
	t.Run("Faris", func(t *testing.T) {
		result := HelloWorld("Faris")
		require.Equal(t, "Hello Faris", result, "Result must be 'Hello Faris'")
		fmt.Println("TestHelloFaris with SubTest Done")
})
	t.Run("Masud", func(t *testing.T) {
		result := HelloWorld("Masud")
		require.Equal(t, "Hello Masud", result, "Result must be 'Hello Masud'")
		fmt.Println("TestHelloMasud with SubTest Done")
	})
	t.Run("FarisMasud", func(t *testing.T) {
		result := HelloWorld("FarisMasud")
		require.Equal(t, "Hello FarisMasud", result, "Result must be 'Hello FarisMasud'")
		fmt.Println("TestHelloFarisMasud with SubTest Done")
	})
}

//? Test Main
func TestMain(m *testing.M) {
	fmt.Println("Before Testing")

	//Before
	m.Run()

	//After
	fmt.Println("After Testing") 
}

//? Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Faris")
	require.Equal(t, "Hello Faris", result, "Result must be 'Hello Faris'")
}

//* Assertions
// Require(same like FailNow())
func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Faris")
	require.Equal(t, "Hello Faris", result, "Result must be 'Hello Faris'")
	fmt.Println("TestHelloWorld with Require Done")
}

// Assert(same like fail())
func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Faris")
	assert.Equal(t, "Hello Faris", result, "Result must be 'Hello Faris'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldFaris(t *testing.T){
	result := HelloWorld("Faris")

	if result != "Hello Faris" {
		// Error
		//panic("Result must be 'Hello Faris'")
		//t.Fail()
		t.Error("Result must be 'Hello Faris'")
	}
	fmt.Println("TestHelloWorldFaris Done")
}
// Test Spesific function go test -v -run FunctionName
func TestHelloWorldMasud(t *testing.T){
	result := HelloWorld("Masud")
	
	if result != "Hello Masud" {
		// Error
		//panic("Result must be 'Hello Masud'")
		//t.FailNow()
		t.Fatal()
	}
	fmt.Println("TestHelloWorldMas'ud Done")
}

func TestHelloWorldFarisMasud(t *testing.T){
	result := HelloWorld("FarisMasud")

	if result != "Hello Faris" {
		// Error
		t.Fatal("Result must be 'Hello FarisMasud'")
	}
	fmt.Println("TestHelloWorldFarisMasud Done")
}