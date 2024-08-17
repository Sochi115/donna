package cmd

type Task struct {
	Id      int    `csv:"ID"`
	Desc    string `csv:"Description"`
	Created string `csv:"Created"`
	Done    bool   `csv:"Completed"`
}
