/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient/fakes"
)

var _ = Describe("Virtual Switch", func() {
	var (
		manager    *VirtualSwitchManager
		fakeClient *fakes.ClientAPI
		cpcid      string
		vswitchId  string
		url        *url.URL
	)

	BeforeEach(func() {
		fakeClient = &fakes.ClientAPI{}
		manager = NewVirtualSwitchManager(fakeClient)

		url, _ = url.Parse("https://127.0.0.1:443")
		cpcid = "cpcid"
		vswitchId = "vswitchId"
	})

	Describe("ListVirtualSwitches", func() {
		var (
			virtualSwitches      []VirtualSwitch
			virtualSwitchesArray VirtualSwitchesArray
			bytes                []byte
		)

		BeforeEach(func() {
			virtualSwitches = []VirtualSwitch{
				{
					URI:  "uri1",
					Name: "name1",
					Type: VIRTUALSWITCH_TYPE_HIPERSOCKET,
				},
				{
					URI:  "uri2",
					Name: "name2",
					Type: VIRTUALSWITCH_TYPE_HIPERSOCKET,
				},
			}

			virtualSwitchesArray = VirtualSwitchesArray{
				virtualSwitches,
			}
			bytes, _ = json.Marshal(virtualSwitchesArray)
		})

		Context("When list virtual switches and returns correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, err := manager.ListVirtualSwitches(cpcid)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(rets[0]).To(Equal(virtualSwitches[0]))
			})
		})

		Context("When list virtual switches and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, errors.New("error"))
				rets, err := manager.ListVirtualSwitches(cpcid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When list virtual switches and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.ListVirtualSwitches(cpcid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When list virtual switches and returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, err := manager.ListVirtualSwitches(cpcid)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})

	Describe("GetVirtualSwitchProperties", func() {
		var (
			virtualSwitch VirtualSwitch
			bytes         []byte
		)

		BeforeEach(func() {
			virtualSwitch = VirtualSwitch{
				URI:  "uri1",
				Name: "name1",
				Type: VIRTUALSWITCH_TYPE_HIPERSOCKET,
			}
			bytes, _ = json.Marshal(virtualSwitch)
		})

		Context("When get virtual switch properties return correctly", func() {
			It("check the results succeed", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, nil)
				rets, err := manager.GetVirtualSwitchProperties(vswitchId)

				Expect(err).To(BeNil())
				Expect(rets).ToNot(BeNil())
				Expect(*rets).To(Equal(virtualSwitch))
			})
		})

		Context("When get virtual switches and returns error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, bytes, errors.New("error"))
				rets, err := manager.GetVirtualSwitchProperties(vswitchId)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When get virtual switches and unmarshal error", func() {
			It("check the error happened", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusOK, []byte("incorrect json bytes"), errors.New("error"))
				rets, err := manager.ListVirtualSwitches(vswitchId)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})

		Context("When get virtual switches and returns incorrect status", func() {
			It("check the results is empty", func() {
				fakeClient.CloneEndpointURLReturns(url)
				fakeClient.ExecuteRequestReturns(http.StatusForbidden, bytes, nil)
				rets, err := manager.ListVirtualSwitches(vswitchId)

				Expect(err).ToNot(BeNil())
				Expect(rets).To(BeNil())
			})
		})
	})
})
