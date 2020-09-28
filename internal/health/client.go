/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package health

import (
	"fmt"
	"github.com/superhero-match/consumer-report-user/internal/config"
	"log"
	"net"
)

// Client holds health client related data.
type Client struct {
	HealthServerURL string
	ContentType     string
}

// NewClient return new health client.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		HealthServerURL: fmt.Sprintf("http://%s%s%s", getIPAddress(), cfg.Health.Port, cfg.Health.ShutdownEndpoint),
		ContentType:     cfg.Health.ContentType,
	}
}

// Get preferred outbound ip of this machine.
func getIPAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}