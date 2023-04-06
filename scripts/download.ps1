mkdir msi
$dotnet = New-Object net.webclient
$dotnet.Downloadfile("https://download.visualstudio.microsoft.com/download/pr/c6ad374b-9b66-49ed-a140-588348d0c29a/78084d635f2a4011ccd65dc7fd9e83ce/dotnet-sdk-7.0.202-win-x64.exe"  ,"msi/dotnet.msi")
$go = New-Object net.webclient
$go.Downloadfile("https://go.dev/dl/go1.20.3.windows-amd64.msi"                                                                                                                      ,"msi/go.msi")
$gcc = New-Object net.webclient
$gcc.Downloadfile("https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe"                                                                     ,"msi/gcc.msi")
$gstreamerdev = New-Object net.webclient
$gstreamerdev.Downloadfile("https://gstreamer.freedesktop.org/data/pkg/windows/1.22.1/msvc/gstreamer-1.0-devel-msvc-x86_64-1.22.1.msi"                                               ,"msi/gstreamer-dev.msi")
$gstreamer = New-Object net.webclient
$gstreamer.Downloadfile("https://gstreamer.freedesktop.org/data/pkg/windows/1.22.1/msvc/gstreamer-1.0-msvc-x86_64-1.22.1.msi"                                                        ,"msi/gstreamer.msi")
$git = New-Object net.webclient
$git.Downloadfile("https://github.com/git-for-windows/git/releases/download/v2.40.0.windows.1/Git-2.40.0-64-bit.exe"                                                                 ,"msi/git.msi")

MSIEXEC /i msi/dotnet.msi, /qn                  ,WAIT
MSIEXEC /i msi/go.msi, /qn                      ,WAIT
MSIEXEC /i msi/gcc.msi, /qn                     ,WAIT
MSIEXEC /i msi/gstreamer-dev.msi, /qn           ,WAIT
MSIEXEC /i msi/gstreamer.msi, /qn               ,WAIT
MSIEXEC /i msi/git.msi, /qn                     ,WAIT