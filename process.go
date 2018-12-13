package main

type (
	Data struct {
		id int
		name string
	}
)

func (data Data)CountDataLength() interface{}{
	return len(data.name)
}
