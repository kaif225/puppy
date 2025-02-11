/*
First we did go mod init repo path
Then created this puupy.go and wroe the basic code
Now create an another folder and in that create main.go and do go mod init somename and then run
go get github.com/kaif225/puppy@latest (it is becaude we have created a repo and in that we have the code )

-- if we have the module locally then in go.mod where main.go is add

replace module_name => path of module for example it is one folder back then ../module_name

it will get the code and we can use it in call the function in main
*/

package puppy

func Bark() string {
	return "woof"
}

func Barks() string {
	return "woof! woof woof!"
}
