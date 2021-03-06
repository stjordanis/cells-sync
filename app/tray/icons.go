/*
 * Copyright 2019 Abstrium SAS
 *
 *  This file is part of Cells Sync.
 *
 *  Cells Sync is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  Cells Sync is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with Cells Sync.  If not, see <https://www.gnu.org/licenses/>.
 */

package tray

import (
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pydio/systray"

	coloricon "github.com/pydio/cells-sync/app/tray/color/icon"
	coloractive "github.com/pydio/cells-sync/app/tray/color/iconactive"
	coloractive2 "github.com/pydio/cells-sync/app/tray/color/iconactive2"
	colorerror "github.com/pydio/cells-sync/app/tray/color/iconerror"
	colorpause "github.com/pydio/cells-sync/app/tray/color/iconpause"

	darkicon "github.com/pydio/cells-sync/app/tray/dark/icon"
	darkactive "github.com/pydio/cells-sync/app/tray/dark/iconactive"
	darkactive2 "github.com/pydio/cells-sync/app/tray/dark/iconactive2"
	darkerror "github.com/pydio/cells-sync/app/tray/dark/iconerror"
	darkpause "github.com/pydio/cells-sync/app/tray/dark/iconpause"

	lighticon "github.com/pydio/cells-sync/app/tray/light/icon"
	lightactive "github.com/pydio/cells-sync/app/tray/light/iconactive"
	lightactive2 "github.com/pydio/cells-sync/app/tray/light/iconactive2"
	lighterror "github.com/pydio/cells-sync/app/tray/light/iconerror"
	lightpause "github.com/pydio/cells-sync/app/tray/light/iconpause"
)

var (
	iconData        = coloricon.Data
	iconActiveData  = coloractive.Data
	iconActive2Data = coloractive2.Data
	iconErrorData   = colorerror.Data
	iconPauseData   = colorpause.Data
	activeToggler   bool
	crtStatus       string
	status          chan string
)

func init() {
	if runtime.GOOS == "darwin" {
		useLightIcons := false
		cmd := exec.Command("defaults", "read", "-g", "AppleInterfaceStyle")
		if output, err := cmd.Output(); err == nil {
			if strings.Contains(string(output), "Dark") {
				useLightIcons = true
			}
		}
		if useLightIcons {
			iconData = lighticon.Data
			iconActiveData = lightactive.Data
			iconActive2Data = lightactive2.Data
			iconErrorData = lighterror.Data
			iconPauseData = lightpause.Data
		} else {
			iconData = darkicon.Data
			iconActiveData = darkactive.Data
			iconActive2Data = darkactive2.Data
			iconErrorData = darkerror.Data
			iconPauseData = darkpause.Data
		}
	}
	status = make(chan string, 1)
	go func() {
		for {
			select {
			case <-time.After(750 * time.Millisecond):
				if crtStatus != "active" {
					break
				}
				if !activeToggler {
					systray.SetIcon(iconActiveData)
				} else {
					systray.SetIcon(iconActive2Data)
				}
				activeToggler = !activeToggler

			case s := <-status:
				if crtStatus == s {
					break
				}
				var data []byte
				crtStatus = s
				switch s {
				case "active":
					activeToggler = false
					data = iconActiveData
				case "idle":
					data = iconData
				case "error":
					data = iconErrorData
				case "pause":
					data = iconPauseData
				}
				systray.SetIcon(data)
			}
		}
	}()
}

func setIconActive() {
	status <- "active"
}

func setIconIdle() {
	status <- "idle"
}

func setIconError(msg ...string) {
	if len(msg) > 0 && crtStatus != "error" {
		// TODO
		// notify("CellsSync", msg[0])
	}
	status <- "error"
}

func setIconPause() {
	status <- "pause"
}
