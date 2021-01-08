package main

import (
	"testing"
)

func TestOptions_findErrors(t *testing.T) {

	//var err error

	//// getting the current namespace is found from a local kube config file
	//err = os.Setenv("KUBECONFIG", filepath.Join("test_data", "test-config"))
	//assert.NoError(t, err)
	//
	//o := Options{}
	//
	//tests := []struct {
	//	name      string
	//	namespace string
	//	want      []string
	//	wantErr   bool
	//}{
	//	{name: "no_error", namespace: "", want: []string{}, wantErr: false},
	//	{name: "certificate_error", namespace: "", want: []string{"A bad thing happened"}, wantErr: false},
	//	{name: "certificate_error", namespace: "cheese", want: []string{"A bad thing happened"}, wantErr: false},
	//	{name: "certificate_error", namespace: "test", want: []string{}, wantErr: false},
	//	{name: "certificate_request_error", namespace: "", want: []string{"Waiting on certificate issuance from order jx/tls-pr-1956-2-gke-tls-jenkinsxlabs-com-s-wbnsn-458730623: \"pending\""}, wantErr: false},
	//	{name: "challenge_error", namespace: "", want: []string{"Waiting for DNS-01 challenge propagation: DNS record for \"pr-1956-2-gke-tls.jenkinsxlabs.com\" not yet propagated"}, wantErr: false},
	//	{name: "clusterissuer_error", namespace: "", want: []string{"A bad thing happened"}, wantErr: false},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//
	//		os.Setenv(envVarTargetNamespace, tt.namespace)
	//
	//		objects := loadDir(t, "cheese", filepath.Join("test_data", tt.name))
	//		o.cmClient = cmFakeClient.NewSimpleClientset(objects...)
	//
	//		got, err := o.findErrors()
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("findErrors() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("findErrors() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

