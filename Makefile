build: 
	rsrc.exe --manifest ".\.windows\windows.manifest" -ico ".\.windows\icon.ico" -o ".\cmd\rsrc.syso"
	go build -o "DCS Updater.exe" -ldflags -H=windowsgui .\cmd\
