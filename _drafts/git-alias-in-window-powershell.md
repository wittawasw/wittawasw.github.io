---
layout: post
title: git alias in window powershell
---

```sh
. $profile
g config --global push.autoSetupRemote true
```

```sh
# Functions for Git commands
Function gitCommand { git $args }
Function gitCommit { git commit --verbose }
Function gitCheckout { git checkout $args }
Function gitCheckoutBranch { git checkout -b $args }
Function gitRemote { git remote -v }
Function gitRemoteRemove { git remote remove $args }
Function gitRemoteAdd { git remote add $args }
Function gitPull { git pull origin $args }

Function gitPush {
    param (
        [string[]]$args
    )

    if ($args.Count -ne 0 -and $args.Count -ne 1) {
        git push origin $args
    } else {
        if ($args.Count -eq 0) {
            $b = (git rev-parse --abbrev-ref HEAD)
        }
        git push origin $($b = if ($args.Count -eq 0) { $b } else { $args[0] })
    }
}

Function gitPushForce {
    param (
        [string]$arg
    )

    if ($args.Count -ne 1) {
        $b = (git rev-parse --abbrev-ref HEAD)
    }
    git push --force origin $b
}

Function gitFetch { git fetch --all --tags --prune --jobs=10 }
Function gitStatus { git status }
Function gitSwitch { git switch $args }
Function gitResetHard { git reset --hard $args }
Function gitResetHardHead { git reset --hard HEAD }
Function gitBranch { git branch $args }
Function gitBranchDelete { git branch -d $args }
Function gitLogGraph { git log --oneline --graph --decorate }
Function gitLogAllGraph { git log --oneline --all --graph --decorate }
Function gitAdd { git add $args }
Function gitAddAll { git add . }

# Set aliases pointing to functions
Set-Alias -Name g -Value gitCommand
Set-Alias -Name gcc -Value gitCommit
Set-Alias -Name gco -Value gitCheckout
Set-Alias -Name gcb -Value gitCheckoutBranch
Set-Alias -Name grv -Value gitRemote
Set-Alias -Name grrm -Value gitRemoteRemove
Set-Alias -Name gra -Value gitRemoteAdd
Set-Alias -Name ggl -Value gitPull
Set-Alias -Name ggp -Value gitPush
Set-Alias -Name ggf -Value gitPushForce
Set-Alias -Name gst -Value gitStatus
Set-Alias -Name gsw -Value gitSwitch
Set-Alias -Name grh -Value gitResetHard
Set-Alias -Name grhh -Value gitResetHardHead
Set-Alias -Name gb -Value gitBranch
Set-Alias -Name gbd -Value gitBranchDelete
Set-Alias -Name glgg -Value gitLogGraph
Set-Alias -Name glola -Value gitLogAllGraph
Set-Alias -Name gf -Value gitFetch
Set-Alias -Name ga -Value gitAdd
Set-Alias -Name gaa -Value gitAddAll

```
