// Copyright 2016 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dockermachine

import (
	"github.com/docker/machine/drivers/generic"
	"github.com/docker/machine/drivers/google"
	"github.com/docker/machine/libmachine/drivers/plugin"
	"github.com/docker/machine/libmachine/drivers/plugin/localbinary"
	"github.com/pkg/errors"
)

func init() {
	localbinary.CoreDrivers = append(localbinary.CoreDrivers, "cloudstack")
}

func RunDriver(driverName string) error {
	switch driverName {
	case "generic":
		plugin.RegisterDriver(generic.NewDriver("", ""))
	case "google":
		plugin.RegisterDriver(google.NewDriver("", ""))
	default:
		return errors.Errorf("Unsupported driver: %s\n", driverName)
	}
	return nil
}

func DefaultParamsForDriver(driverName string) map[string]interface{} {
	params := make(map[string]interface{})
	switch driverName {
	case "google":
		params["google-use-internal-ip"] = true
	}
	return params
}
