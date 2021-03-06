//     Digota <http://digota.com> - eCommerce microservice
//     Copyright (C) 2017  Yaron Sumel <yaron@digota.com>. All Rights Reserved.
//
//     This program is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Affero General Public License as published
//     by the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     This program is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Affero General Public License for more details.
//
//     You should have received a copy of the GNU Affero General Public License
//     along with this program.  If not, see <http://www.gnu.org/licenses/>.

package sdk

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	if _, err := NewClient("notvalidhost", nil); err == nil {
		t.Fatal()
	}
	if _, err := NewClient("notvalidhost", &ClientOpt{
		WithBlock: true,
		Crt:       "none",
		Key:       "none",
	}); err == nil {
		t.Fatal()
	}
	if _, err := NewClient("notvalidhost", &ClientOpt{
		InsecureSkipVerify: false,
		ServerName:         "server.com",
		CaCrt:              "testFiles/notvalid.crt",
		Crt:                "testFiles/client.com.crt",
		Key:                "testFiles/client.com.key",
	}); err == nil {
		t.Fatal()
	}
	if _, err := NewClient("notvalidhost", &ClientOpt{
		InsecureSkipVerify: false,
		ServerName:         "server.com",
		CaCrt:              "testFiles/client.com.key",
		Crt:                "testFiles/client.com.crt",
		Key:                "testFiles/client.com.key",
	}); err == nil {
		t.Fatal()
	}
	if _, err := NewClient("notvalidhost", &ClientOpt{
		WithBlock:          true,
		InsecureSkipVerify: false,
		ServerName:         "server.com",
		CaCrt:              "testFiles/ca.crt",
		Crt:                "testFiles/client.com.crt",
		Key:                "testFiles/client.com.key",
	}); err == nil {
		t.Fatal()
	}
	if _, err := NewClient("notvalidhost", &ClientOpt{
		InsecureSkipVerify: false,
		ServerName:         "server.com",
		CaCrt:              "testFiles/ca.crt",
		Crt:                "testFiles/client.com.crt",
		Key:                "testFiles/client.com.key",
	}); err != nil {
		t.Fatal()
	}
}
