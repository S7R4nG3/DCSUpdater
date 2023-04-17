build: 
	rsrc.exe --manifest ".\.windows\exec.manifest" -o ".\cmd\rsrc.syso"
	go build -o "DCS Updater.exe" -ldflags -H=windowsgui .\cmd\
