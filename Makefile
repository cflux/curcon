# Copyright 2009 The Daniel Polnoff. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=curcon
GOFILES=\
	curcon.go\

include $(GOROOT)/src/Make.pkg
