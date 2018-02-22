#$language = "VBScript"
#$interface = "1.0"

crt.Screen.Synchronous = Ture

Sub Main    
    Dim logfile, Ctime, Cmonth, Cday, Chour
    Dim slotIp, userName, password, logPath
    Dim needLog, VmIp, VM_PASSWD
    Dim command, HOST_ROOT, VM_ROOT
    Dim CCBIT , ftp_ip ,telnet_ip
    
    Cmonth = Month(Now) 
    Cday   = Day(Now)
    Chour  = Hour(Now)
    
	
	slotIp  = crt.Arguments(0)
	userName = crt.Arguments(1)
	password = crt.Arguments(2)
	
	
    
    logfile = slotIp&"_host"
    Ctime   = Cmonth&"m"&Cday&"d"&Chour&"h.log"   
        
 
    set tab = crt.Session.ConnectInTab("/s " &slotIp)
    tab.Caption = slotIp & "(change-software" & slotIp & ")"
    
	Do
		If tab.Screen.WaitForString("#", 5) = True Then
			Exit Do
			else
				crt.Sleep 1000
				tab.Screen.Send  chr(13)
		End If
	Loop
	Dim plen 
		plen=Len(slotIp)
	if plen <> 2 Then
		Do
			command="ssh "& slotIp &" "
			tab.Screen.Send command & chr(13)
			Dim result
			result = tab.Screen.WaitForStrings("Password: ", "(yes/no)? ","or 'no': ", 10)
			If result = 1 Then
				tab.Screen.Send ""& password &"" & chr(13)
				Exit Do
			ElseIf result = 2 Then
				tab.Screen.Send "yes" & VbCr
				If tab.Screen.WaitForString("Password: ", 5) = True Then
					tab.Screen.Send ""& password &"" & chr(13)
					Exit Do
				End If
			ElseIf result = 3 Then
				tab.Screen.Send "yes"  & VbCr
				If tab.Screen.WaitForString("Password: ", 5) = True Then
					tab.Screen.Send ""& password &"" & chr(13)
					Exit Do
				End If
			End If
		Loop
		tab.Screen.Send "ll" & chr(13)
		tab.Screen.Send "echo hello" & chr(13)

		

	End If
	set tab1 = crt.Session.ConnectInTab("/s " &slotIp)
    tab1.Caption = slotIp & "(change-software2" & slotIp & ")"
	if plen <> 2 Then
		Do
			command="ssh "& slotIp &" "
			tab1.Screen.Send command & chr(13)
			
			result = tab1.Screen.WaitForStrings("Password: ", "(yes/no)? ","or 'no': ", 10)
			If result = 1 Then
				tab1.Screen.Send ""& password &"" & chr(13)
				Exit Do
			ElseIf result = 2 Then
				tab1.Screen.Send "yes" & VbCr
				If tab1.Screen.WaitForString("Password: ", 5) = True Then
					tab1.Screen.Send ""& password &"" & chr(13)
					Exit Do
				End If
			ElseIf result = 3 Then
				tab1.Screen.Send "yes"  & VbCr
				If tab1.Screen.WaitForString("Password: ", 5) = True Then
					tab1.Screen.Send ""& password &"" & chr(13)
					Exit Do
				End If
			End If
		Loop
		tab1.Screen.Send "ll" & chr(13)
		tab1.Screen.Send "echo nihao" & chr(13)
	End If
	crt.Sleep 1000
    'tab.Session.Disconnect
    'crt.Quit
End Sub