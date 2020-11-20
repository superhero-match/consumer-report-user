/*
  Copyright (C) 2019 - 2021 MWSOFT
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
package config

// DB holds the configuration values for the database.
type DB struct {
	Host     string `env:"DB_HOST" yaml:"host" default:"192.168.0.64"`
	Port     int    `env:"DB_PORT" yaml:"port" default:"3306"`
	User     string `env:"DB_USER" yaml:"user" default:"dev"`
	Password string `env:"DB_PASSWORD" yaml:"password" default:"Awesome85**"`
	Name     string `env:"CONSUMER_REPORT_USER_DB_NAME" yaml:"name" default:"report"`
}