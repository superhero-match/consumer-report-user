/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"github.com/superhero-match/consumer-report-user/cmd/consumer/reader"
	"github.com/superhero-match/consumer-report-user/internal/config"
	"github.com/superhero-match/consumer-report-user/internal/health"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	client := health.NewClient(cfg)

	r, err := reader.NewReader(cfg)
	if err != nil {
		_ = client.ShutdownHealthServer()

		panic(err)
	}

	err = r.Read()
	if err != nil {
		_ = client.ShutdownHealthServer()

		panic(err)
	}

	_ = client.ShutdownHealthServer()
}
