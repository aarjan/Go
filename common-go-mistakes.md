Common go mistakes 

====================================================
# Mistake No.1	- Not accepting interfaces

## Stop Doing This
func (page *Page) saveSourceAs(path string) {
	b := new(bytes.Buffer)
	b.Write(page.Source.Content)
	page.saveSource(b.Bytes(),path)
}

func (page *Page) saveSource(by []bytes, inpath string) {
	writeToDisk(inpath, bytes.NewReader(by))
}

## Instead
func (page *Page) saveSourceAs(path string) {
	b := new(bytes.Buffer)
	b.Write(page.Source.Content)
	page.saveSource(b,path)
}

				// We have used an interface instead of the specific data type which is alot more efficient
func (page *Page) saveSource(b io.Reader, inpath string) {
	writeToDisk(inpath, b)
}


====================================================
Mistake No.2 - Not using io.Reader & io.Writer 

These interfaces supports reading/writing from various sources including std. i/o in different formats.
Rather than using "bytes.Buffer" which can be very large and inefficient, use these interfaces.


====================================================
# Mistake No.3 - Using Broad Interfaces

	- Functions should only accept interfaces that require the methods they need
	- Functions should not accept a broad interface when a narrow one would work

Eg. of broad interface
	type File interface {
		io.Closer
		io.Reader
		io.ReaderAt
		io.Seeker
		io.Writer
		io.WriterAt
	}

	// Using this interface to read a file 
	(Instead, we could have used 'io.Reader' interface only) 
	
	func ReadIn(f File) {
		b := []byte{}
		n,err := f.Read(b)
		...
	}


=====================================
# Mistake No.4 - Methods vs Functions

## Function
	- Operations performed in N1 inputs that results in N2 outputs.
	- The same inputs will always result in the same outputs
	- Functions should not depend upon state

## Method
	- Defines the behaviour of a type
	- A function that operates against a value
	- Shoud use state
	- Logically connected


================================
# Mistake No.5 - Pointers vs Values

	- It's not a question of performance(generally),
	but one of shared access.
	- If you want to share the value with a function or method, then use a pointer.
	- If you don't want to share it, then use a value(copy).

## Pointer Reciever 
	- If you want to share a value with it's method, use a pointer reciever.
	- Since methods commonly manage state, this is a common usage.
	- Not safe for concurrent usage.
	
	Eg:
		type InMemoryFile struct {
			at int64
			name string
			data []byte
			closed bool
		}

		// Here we need to alter the state of InMemoryFile

		func (f *InMemoryFile) Close() error { 	
			atomic.StoreInt64(&f.at, 0)
			f.closed = true
			return nil
		}

## Value Reciever 
	- If you want the value copied (not shared), use values
	- If the type is an empty struct(stateless, just behaviour), then just use value
	- Safe for concurrent usage

	Eg:
		type Time struct {
			sec int64
			nsec uintptr
		}

		// Since time always changes

		func (t Time) isZero() bool {
			return t.sec == 0 && t.nsec == 0
		}


===================================================
# Thinking of Errors as Strings

## Standard Errors
	- errors.New("error here") is usually sufficient
	- Exported Error Variables can be easily checked

// To find the exact type of error that occurred , checking error variables is alot easier than checking strings/substrings in the error

## Custom Errors
	- Can provide context to guarantee consistent feedback
	- Provide a type which can be different from the error value
	- Can provide dynamic values (based on internal error state)


===================================================









