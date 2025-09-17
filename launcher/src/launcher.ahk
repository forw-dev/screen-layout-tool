#SingleInstance Force

; --- global -------------------------------------------------
ConfigDir := A_AppData . "\screen-layout-tool-launcher"
if !FileExist(ConfigDir) {
    DirCreate ConfigDir
}
ConfigFilePath := ConfigDir . "\launcher.ini"

LayoutDirPath := ".\layouts\"

; --- tray ----------------------------------------------------
A_IconTip := "Screen Layout Tool"
if (A_IsCompiled) {
    TraySetIcon(A_ScriptName)
}
UpdateMenu()
UpdateAutoStartupShortcut()

; --- menu --------------------------------------------------
UpdateMenu() {
    Language := IniRead(ConfigFilePath, "Running", "Language", "en")
    IsAutoStartup := IniRead(ConfigFilePath, "Setting", "AutoStartup", "true")
    IsCrossMonitor := IniRead(ConfigFilePath, "Setting", "CrossMonitor", "true")
    CurrLayoutName := IniRead(ConfigFilePath, "Running", "Layout", "Basic")
    ; 语言菜单
    LanguageMenu := Menu()
    LanguageMenu.Add("English", OnLanguageEnClicked)
    if (Language == "en") {
        LanguageMenu.Check("English")
    }
    LanguageMenu.Add("中文", OnLanguageZhClicked)
    if (Language == "zh") {
        LanguageMenu.Check("中文")
    }
    ; 布局菜单
    LayoutMenu := Menu()
    LayoutNames := GetLayoutNames()
    for LayoutName in LayoutNames {
        LayoutMenu.Add(LayoutName, OnLayoutClicked)
        if (CurrLayoutName == LayoutName) {
            LayoutMenu.Check(LayoutName)
        }
    }
    ; 主菜单
    MyMenu := A_TrayMenu
    MyMenu.Delete() ; 删除默认的菜单项
    if (Language == "en") {
        ; MyMenu.Add("Support Us", OnSupportUsClicked)
        ; MyMenu.Add()
        MyMenu.Add("Layout", LayoutMenu)
        MyMenu.Add("Open Layout Folder", OnOpenLayoutFolderClicked)
        MyMenu.Add("Get Layout", OnGetLayoutClicked)
        MyMenu.Add()
        MyMenu.Add("Auto Start", OnAutoStartClicked)
        MyMenu.Add("Cross Monitor", OnCrossMonitorClicked)
        MyMenu.Add("Language", LanguageMenu)
        MyMenu.Add()
        MyMenu.Add("About", OnAboutClicked)
        MyMenu.Add("Help", OnHelpClicked)
        MyMenu.Add()
        MyMenu.Add("Reload", OnReloadClicked)
        MyMenu.Add("Exit", OnExitClicked)
        if (IsAutoStartup = "true") {
            MyMenu.Check("Auto Start")
        } else {
            MyMenu.Uncheck("Auto Start")
        }
        if (IsCrossMonitor = "true") {
            MyMenu.Check("Cross Monitor")
        } else {
            MyMenu.Uncheck("Cross Monitor")
        }
    }
    if (Language == "zh") {
        LanguageMenu.Check("中文")
        ; MyMenu.Add("支持我们", OnSupportUsClicked)
        ; MyMenu.Add()
        MyMenu.Add("布局", LayoutMenu)
        MyMenu.Add("打开布局文件目录", OnOpenLayoutFolderClicked)
        MyMenu.Add("获取布局", OnGetLayoutClicked)
        MyMenu.Add()
        MyMenu.Add("自动启动", OnAutoStartClicked)
        MyMenu.Add("跨显示器", OnCrossMonitorClicked)
        MyMenu.Add("语言", LanguageMenu)
        MyMenu.Add()
        MyMenu.Add("关于", OnAboutClicked)
        MyMenu.Add("帮助", OnHelpClicked)
        MyMenu.Add()
        MyMenu.Add("重新加载", OnReloadClicked)
        MyMenu.Add("退出", OnExitClicked)
        if (IsAutoStartup = "true") {
            MyMenu.Check("自动启动")
        } else {
            MyMenu.Uncheck("自动启动")
        }
        if (IsCrossMonitor = "true") {
            MyMenu.Check("跨显示器")
        } else {
            MyMenu.Uncheck("跨显示器")
        }
    }
}

UpdateAutoStartupShortcut() {
    IsAutoStartup := IniRead(ConfigFilePath, "Setting", "AutoStartup", "true")
    AutoStartupShortcutPath := A_Startup "\" A_ScriptName ".lnk"
    if (IsAutoStartup = "true") {
        if !FileExist(AutoStartupShortcutPath) {
            FileCreateShortcut(A_ScriptFullPath, AutoStartupShortcutPath, A_ScriptDir)
        }
    } else {
        if FileExist(AutoStartupShortcutPath) {
            FileDelete(AutoStartupShortcutPath)
        }
    }
}

; --- menu event -------------------------------------------
OnLayoutClicked(ItemName, *) {
    IniWrite(ItemName, ConfigFilePath, "Running", "Layout")
    UpdateMenu()
}

OnOpenLayoutFolderClicked(*) {
    Run(A_ScriptDir . LayoutDirPath)
}

OnGetLayoutClicked(*) {
    Language := IniRead(ConfigFilePath, "Running", "Language", "en")
    Run("https://forw.cc/screen-layout-tool/layouts/")
}

OnAutoStartClicked(*) {
    IsAutoStartup := IniRead(ConfigFilePath, "Setting", "AutoStartup", "true")
    IsAutoStartup := (IsAutoStartup = "true") ? "false" : "true"
    IniWrite(IsAutoStartup, ConfigFilePath, "Setting", "AutoStartup")
    UpdateMenu()
    UpdateAutoStartupShortcut()
}

OnCrossMonitorClicked(*) {
    IsCrossMonitor := IniRead(ConfigFilePath, "Setting", "CrossMonitor", "true")
    IsCrossMonitor := (IsCrossMonitor = "true") ? "false" : "true"
    IniWrite(IsCrossMonitor, ConfigFilePath, "Setting", "CrossMonitor")
    UpdateMenu()
}

OnLanguageEnClicked(*) {
    IniWrite("en", ConfigFilePath, "Running", "Language")
    UpdateMenu()
}
OnLanguageZhClicked(*) {
    IniWrite("zh", ConfigFilePath, "Running", "Language")
    UpdateMenu()
}

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;; 关于 ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
OnAboutClicked(*) {
    TrayTip("v0.21.20`nforw.cc", "Screen Layout Tool", 17)
}

OnHelpClicked(*) {
    Language := IniRead(ConfigFilePath, "Running", "Language", "en")
    Run("https://forw.cc/screen-layout-tool/get-help.md")
}

OnReloadClicked(*) {
    Run(A_ScriptFullPath)
    ExitApp()
}

OnExitClicked(*) {
    ExitApp()
}

; --- layout function ------------------
GetLayoutNames() {
    LayoutNames := []
    LayoutFilePattern := LayoutDirPath . "*.json"
    loop files, LayoutFilePattern {
        SplitPath A_LoopFileName, , , , &LayoutName
        LayoutNames.Push(LayoutName)
    }
    return LayoutNames
}

; --- hotkey ---------------------------
#+WheelDown::
#+Right:: {
    ScrollCurrLayout("1")
}
#+WheelUp::
#+Left:: {
    ScrollCurrLayout("-1")
}

#WheelDown:: 
#Right:: {
    ScrollCurrWindow("1")
}
#WheelUp::
#Left:: {
    ScrollCurrWindow("-1")
}

; --- hotkey function -------------------------
ScrollCurrLayout(step) {
    static LastScrollTime := 0
    if (A_TickCount - LastScrollTime < 300)
        return
    CurrLayoutName := IniRead(ConfigFilePath, "Running", "Layout", "Basic")
    LayoutNames := GetLayoutNames()
    CurrentIndex := 0
    for Index, LayoutName in LayoutNames {
        if (LayoutName = CurrLayoutName) {
            CurrentIndex := Index
            break
        }
    }
    n := CurrentIndex + step
    if n > LayoutNames.Length {
        n := 1
    }
    if n <= 0 {
        n := LayoutNames.Length
    }
    LayoutName := LayoutNames[n]
    IniWrite(LayoutName, ConfigFilePath, "Running", "Layout")
    TrayTip()
    TrayTip(LayoutName, "Current Layout", 16)
    UpdateMenu()
}

ScrollCurrWindow(step) {
    static LastScrollTime := 0
    if (A_TickCount - LastScrollTime < 300)
        return
    IsCrossMonitor := IniRead(ConfigFilePath, "Setting", "CrossMonitor", "true")
    CurrLayoutName := IniRead(ConfigFilePath, "Running", "Layout", "Basic")
    LayoutFilePath := LayoutDirPath . CurrLayoutName . ".json"
    LastScrollTime := A_TickCount
    cmd := "controller.exe window scroll"
    cmd := cmd . " -s " . step
    cmd := cmd . " -l `"" . LayoutFilePath . "`""

    if (IsCrossMonitor = "true") {
        cmd := cmd . " -c "
    }
    ExitCode := RunWait(cmd, , "Hide")
    Warning(ExitCode)
}

; --- other -----------------------------------
Warning(Code) {
    if (Code != 0) {
        TrayTip()
    }
    if (Code = 1) {
        TrayTip("Soming wrong", "Warning", 2)
    }
    if (Code = 11) {
        TrayTip("Position result unexpected. You can adjust it manually and try next position", "Warning", 2)
    }
    if (Code = 12) {
        TrayTip("Layout file parse error", "Warning", 2)
    }
    if (Code = 13) {
        TrayTip("Layout file not found", "Warning", 2)
    }
    if (Code = 14) {
        TrayTip("No such a layout", "Warning", 2)
    }
    if (Code = 15) {
        TrayTip("Duplicate positions in layout", "Warning", 2)
    }
    if (Code = 16) {
        TrayTip("Monitor not found", "Warning", 2)
    }
    if (Code = 17) {
        TrayTip("Load layout file failed", "Warning", 2)
    }
}
