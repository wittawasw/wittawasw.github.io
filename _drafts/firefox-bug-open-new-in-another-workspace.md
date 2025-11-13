---
layout: post
title: Firefox bug open new in another workspace
---

Original bug
https://bugzilla.mozilla.org/show_bug.cgi?id=1423768

Report of bugfix
https://windowsreport.com/firefox-fixes-an-8-year-old-windows-bug-that-broke-virtual-desktop-use/

Created another bug
https://www.reddit.com/r/firefox/comments/1nkjjg1/how_do_i_open_links_on_a_different_virtual/
https://support.mozilla.org/gl/questions/1536551

and was filed
https://bugzilla.mozilla.org/show_bug.cgi?id=1994825

Fix is release but not in release note
https://www.firefox.com/en-US/firefox/145.0/releasenotes/

Solution is to go to about:config and set this variable to false
widget.prefer_windows_on_current_virtual_desktop
