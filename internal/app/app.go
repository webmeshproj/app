/*
Copyright 2023 Avi Zimmerman <avi.zimmerman@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/webmeshproj/app/internal/daemon"
)

const appID = "com.webmeshproj.app"

type App struct {
	fyne.App
	cli daemon.Client
}

func New() *App {
	a := app.NewWithID(appID)
	app := &App{
		App: a,
		cli: daemon.NewClient(),
	}
	app.setup()
	return app
}

func (app *App) setup() {}
