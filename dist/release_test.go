// Copyright (C) 2020  Matthew "strager" Glazar
// See end of file for extended copyright information.

package main

import "io/fs"
import "os"
import "path/filepath"
import "testing"
import "time"

func AssertStringsEqual(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("expected %#v, but got %#v", expected, actual)
	}
}

func Check(t *testing.T, err error) {
	if err != nil {
		t.Errorf("%#v", err)
	}
}

func TestReadVersionFileData(t *testing.T) {
	t.Run("1.0.0", func(t *testing.T) {
		version := ReadVersionFileData([]byte("1.0.0\n2021-12-13\n"))
		AssertStringsEqual(t, version.VersionNumber, "1.0.0")
		if version.ReleaseDate.Year() != 2021 || version.ReleaseDate.Month() != 12 || version.ReleaseDate.Day() != 13 {
			t.Errorf("expected 2021-12-13, but got %#v", version.ReleaseDate)
		}
	})
}

func TestWriteVersionFile(t *testing.T) {
	t.Run("replace existing version file", func(t *testing.T) {
		dir := t.TempDir()
		Check(t, os.Chdir(dir))
		Check(t, os.WriteFile("version", []byte{}, fs.FileMode(0644)))

		Check(t, WriteVersionFile(VersionFileInfo{
			VersionNumber: "2.0.0",
			ReleaseDate:   time.Date(2022, 2, 8, 16, 56, 37, 0, time.Local),
		}))

		versionData, err := os.ReadFile("version")
		Check(t, err)
		AssertStringsEqual(t, string(versionData), "2.0.0\n2022-02-08\n")
	})
}

func TestDebianChangelogEntry(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		losAngeles, err := time.LoadLocation("America/Los_Angeles")
		Check(t, err)

		entry := DebianChangelogEntry(VersionFileInfo{
			VersionNumber: "2.0.0",
			ReleaseDate:   time.Date(2022, 2, 8, 16, 56, 37, 0, losAngeles),
		})
		AssertStringsEqual(t, entry,
			"quick-lint-js (2.0.0-1) unstable; urgency=medium\n"+
				"\n"+
				"  * New release.\n"+
				"\n"+
				" -- Matthew \"strager\" Glazar <strager.nds@gmail.com>  Tue, 08 Feb 2022 16:56:37 -0800\n"+
				"\n")
	})
}

func TestUpdateDebianChangelog(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		dir := t.TempDir()
		changelogFilePath := filepath.Join(dir, "changelog")
		Check(t, os.WriteFile(changelogFilePath, []byte(
			"quick-lint-js (0.1.0-1) unstable; urgency=medium\n"+
				"\n"+
				"  * Initial release.\n"+
				"\n"+
				" -- Matthew \"strager\" Glazar <strager.nds@gmail.com>  Thu, 14 Jan 2021 19:25:47 -0800\n"), fs.FileMode(0644)))
		losAngeles, err := time.LoadLocation("America/Los_Angeles")
		Check(t, err)

		Check(t, UpdateDebianChangelog(changelogFilePath, VersionFileInfo{
			VersionNumber: "0.2.0",
			ReleaseDate:   time.Date(2021, 4, 5, 19, 28, 8, 0, losAngeles),
		}))

		data, err := os.ReadFile(changelogFilePath)
		Check(t, err)
		AssertStringsEqual(t, string(data),
			"quick-lint-js (0.2.0-1) unstable; urgency=medium\n"+
				"\n"+
				"  * New release.\n"+
				"\n"+
				" -- Matthew \"strager\" Glazar <strager.nds@gmail.com>  Mon, 05 Apr 2021 19:28:08 -0700\n"+
				"\n"+
				"quick-lint-js (0.1.0-1) unstable; urgency=medium\n"+
				"\n"+
				"  * Initial release.\n"+
				"\n"+
				" -- Matthew \"strager\" Glazar <strager.nds@gmail.com>  Thu, 14 Jan 2021 19:25:47 -0800\n")
	})
}

// quick-lint-js finds bugs in JavaScript programs.
// Copyright (C) 2020  Matthew "strager" Glazar
//
// This file is part of quick-lint-js.
//
// quick-lint-js is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// quick-lint-js is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with quick-lint-js.  If not, see <https://www.gnu.org/licenses/>.
