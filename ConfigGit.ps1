$MyPat = 'q66jjy4i45txj57fdbvjhmxcnlesapfjbhfsldkpvsb3irdffmwa'

Write-Output $MyPat

$B64Pat = [Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes("`:$MyPat"))

Write-Output $B64Pat

git config --global --unset http.extraHeader

git -c http.extraHeader="Authorization: Basic $B64Pat" ls-remote https://dev.azure.com/slb1-swt/wireline-automation/_git/dev-pdp-maxwell-interface
# git -c http.extraHeader="Authorization: Basic $B64Pat" clone https://dev.azure.com/slb1-swt/wireline-automation/_git/dev-pdp-maxwell-interface



# git clone https://slb1-swt@dev.azure.com/slb1-swt/wireline-automation/_git/dev-pdp-maxwell-interface
# git clone https://         dev.azure.com/slb1-swt/wireline-automation/_git/dev-pdp-maxwell-interface


# git ls-remote https://slb1-swt@dev.azure.com/slb1-swt/wireline-automation/_git/dev-pdp-maxwell-interface



[System.Environment]::SetEnvironmentVariable("MY_VARIABLE", "MyValue", "Machine")