// Copyright (C) 2020  Matthew "strager" Glazar
// See end of file for extended copyright information.

#ifndef QUICK_LINT_JS_MOCK_INOTIFY_H
#define QUICK_LINT_JS_MOCK_INOTIFY_H

#include <quick-lint-js/configuration/change-detecting-filesystem.h>
#include <quick-lint-js/port/have.h>

namespace quick_lint_js {
#if QLJS_HAVE_INOTIFY
class mock_inotify_init_error_guard {
 public:
  explicit mock_inotify_init_error_guard(int error) noexcept
      : old_error_(mock_inotify_force_init_error) {
    mock_inotify_force_init_error = error;
  }

  mock_inotify_init_error_guard(const mock_inotify_init_error_guard&) = delete;
  mock_inotify_init_error_guard& operator=(
      const mock_inotify_init_error_guard&) = delete;

  ~mock_inotify_init_error_guard() {
    mock_inotify_force_init_error = this->old_error_;
  }

 private:
  int old_error_;
};

class mock_inotify_add_watch_error_guard {
 public:
  explicit mock_inotify_add_watch_error_guard(int error) noexcept
      : old_error_(mock_inotify_force_add_watch_error) {
    mock_inotify_force_add_watch_error = error;
  }

  mock_inotify_add_watch_error_guard(
      const mock_inotify_add_watch_error_guard&) = delete;
  mock_inotify_add_watch_error_guard& operator=(
      const mock_inotify_add_watch_error_guard&) = delete;

  ~mock_inotify_add_watch_error_guard() {
    mock_inotify_force_add_watch_error = this->old_error_;
  }

 private:
  int old_error_;
};
#endif
}

#endif

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
