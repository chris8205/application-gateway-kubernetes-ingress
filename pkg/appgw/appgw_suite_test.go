// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

// +build unittest

package appgw

import (
	"flag"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

func TestAppgw(t *testing.T) {
	klog.InitFlags(nil)
	_ = flag.Set("v", "3")
	_ = flag.Lookup("logtostderr").Value.Set("true")

	RegisterFailHandler(Fail)
	RunSpecs(t, "Appgw Suite")
}
