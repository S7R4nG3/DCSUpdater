## Forgive my random piecewise code, I'm having to add pieces from work and from home
## My office's firewall policies are a pain :P

# We're gonna start from a clean slate
cls

# First, we need to check and be sure that your PowerShell window is running as an administrator
## I have this code at home
## I have this code at home
## I have this code at home

# This piece now looks to see if it finds an Eagle Dynamics folder on your C: drive
$directory = Get-ChildItem c:\Test -Recurse | Where-Object {$_.PSIsContainer -eq $true -and $_.Name -match "Eagle Dynamics"}

# If Eagle Dynamics folder isn't found, prompt you to enter the Drive Letter for your install directory
if($directory -eq $null){
    write-host "Looks like your install directory isn't on C:\"
    $installdrive = Read-Host "Please provide the drive letter for your install (Ex. E)"
    $installdrive = $installdrive + ":\Eagle Dynamics"
    }
else{}

