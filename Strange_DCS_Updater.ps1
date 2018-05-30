## Forgive my random piecewise code, I'm having to add pieces from work and from home
## My office's firewall policies are a pain :P

# We're gonna start from a clean slate
cls

# First, we need to check and be sure that your PowerShell window is running as an administrator
# If it isn't, open an Admin PowerShell window
$currentPrincipal = New-Object Security.Principal.WindowsPrincipal([Security.Principal.WindowsIdentity]::GetCurrent())
$admin = $currentPrincipal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)

if($admin -eq $false){
If (-NOT ([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator"))

    {   $arguments = "& '" + $myinvocation.mycommand.definition + "'"
        Start-Process powershell -Verb runAs -ArgumentList $arguments
        Break
    }

}

# Next do a quick sanity check just asking if you'd like to continue
write-host "Welcome to STRANGE's DCS Updater Utility!"
write-host "The script will walk you through whats happening and ask you some questions if it has trouble"
$proceed = read-host "Would you like to continue? (y/n)"
if($proceed -eq "y" -or $proceed -eq "yes"){

# This piece now looks to see if it finds an Eagle Dynamics folder on your C: drive
write-host "Checking to find your ED directory..."
write-host "Give me just a sec..."
$directory = Get-ChildItem "c:\Program Files\" -Recurse | Where-Object {$_.PSIsContainer -eq $true -and $_.Name -match "Eagle Dynamics"}

# If an Eagle Dynamics folder isn't found, prompt you to enter the Drive Letter for your install directory
if($directory -eq $null){
    write-host "Looks like your install directory isn't on C:\"
    $installdrive = Read-Host "Please provide the drive letter for your install (Ex. E)"
        $installdirectory = Get-ChildItem "$installdrive" -Recurse | Where-Object {$_.PSIsContainer -eq $true -and $_.Name -match "Eagle Dynamics"}
        write-host "Cool, I see your install directory is listed at:

        $installdirectory
        "
    }
else{
        $installdirectory = $env:ProgramFiles + "\Eagle Dynamics\DCS World\bin"
        write-host "Cool, I see your install directory is listed at:

        $installdirectory
        "
    }

# Now we need to setup some variables for the updater EXE to tell it what we want to do
    $DCSEXE = "dcs_updater.exe"
    $udpate = 'update'
    $repair = 'repair'
    $openbeta = 'update @openbeta'
    $openalpha = 'update @openalpha'
    $release = 'update @release'
    $cleanup = 'cleanup'

# Next we need some choice options for the variables we just set above
$Title = "Alrighty, lets get this rolling..."
$Info = "Choose and option to proceed!"
 
$options = [System.Management.Automation.Host.ChoiceDescription[]] @("&None", "&Update", "&Repair", "Update:Open&Beta","Update:Open&Alpha","Update:Re&lease", "&Cleanup")
    [int]$defaultchoice = 0
$opt = $host.UI.PromptForChoice($Title , $Info , $Options,$defaultchoice)

# Finally depending on what option you choose, the updater will do whatever you want it to!
switch($opt)
    {
    0 {write-host "Have a good one!" -foregroundcolor Green
       read-host "Press any key to continue" }
    1 {& "$installdirectory\$DCSEXE" $update }
    2 {& "$installdirectory\$DCSEXE" $repair }
    3 {& "$installdirectory\$DCSEXE" $openbeta }
    4 {& "$installdirectory\$DCSEXE" $openalpha }
    5 {& "$installdirectory\$DCSEXE" $release } 
    6 {& "$installdirectory\$DCSEXE" $cleanup }
    }

}

# Lastly, if the user said above that they didn't want to continue, kill the script :)
else{write-host "Thanks for giving me a try!" -foregroundcolor Green
write-host "Press any key to continue..."
$HOST.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown") | OUT-NULL
$HOST.UI.RawUI.Flushinputbuffer()}
