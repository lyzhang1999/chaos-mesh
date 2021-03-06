// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package podchaos_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
	"github.com/chaos-mesh/chaos-mesh/controllers/podchaos"
)

func TestPodChaos(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"PodChaos Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))
	close(done)
}, 60)

var _ = AfterSuite(func() {
})

var _ = Describe("PodChaos", func() {
	Context("PodChaos", func() {
		invalidDuration := "invalid duration"
		meta := metav1.TypeMeta{
			Kind:       "PodChaos",
			APIVersion: "v1",
		}

		r := podchaos.Reconciler{
			Client:        fake.NewFakeClientWithScheme(scheme.Scheme),
			EventRecorder: &record.FakeRecorder{},
			Log:           ctrl.Log.WithName("controllers").WithName("PodChaos"),
		}

		It("PodChaos Reconcile", func() {
			var err error

			_, err = r.Reconcile(ctrl.Request{}, &v1alpha1.PodChaos{
				TypeMeta: meta,
				Spec:     v1alpha1.PodChaosSpec{Scheduler: nil, Duration: &invalidDuration},
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid duration"))

			_, err = r.Reconcile(ctrl.Request{}, &v1alpha1.PodChaos{
				TypeMeta: meta,
				Spec:     v1alpha1.PodChaosSpec{Scheduler: nil, Duration: nil, Action: v1alpha1.PodKillAction},
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unsupported chaos action"))
			_, err = r.Reconcile(ctrl.Request{}, &v1alpha1.PodChaos{
				TypeMeta: meta,
				Spec:     v1alpha1.PodChaosSpec{Scheduler: nil, Duration: nil, Action: v1alpha1.ContainerKillAction},
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("unsupported chaos action"))
			_, err = r.Reconcile(ctrl.Request{}, &v1alpha1.PodChaos{
				TypeMeta: meta,
				Spec:     v1alpha1.PodChaosSpec{Scheduler: nil, Duration: nil, Action: ""},
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid chaos action"))

			_, err = r.Reconcile(ctrl.Request{}, &v1alpha1.PodChaos{
				TypeMeta: meta,
				Spec:     v1alpha1.PodChaosSpec{Scheduler: &v1alpha1.SchedulerSpec{}, Duration: nil, Action: ""},
			})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid chaos action"))
		})
	})
})
