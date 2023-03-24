// Copyright 2021-2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"

	"github.com/vmware-tanzu/tanzu-framework/pinniped-components/common/pkg/pinnipedinfo"
	netutil "github.com/vmware-tanzu/tanzu-framework/tkg/utils/net"
)

const (
	KubePublicNamespace       = "kube-public"
	PinnipedInfoConfigMapName = "pinniped-info"
)

type clusterInfoConfig struct {
	Version string `json:"apiVersion"`
	Data    struct {
		Kubeconfig string `json:"kubeconfig"`
	}
	Kind string `json:"kind"`
}

// GetClusterNameFromKubeconfigAndContext gets name of the cluster from kubeconfig file
// and kube-context
func GetClusterNameFromKubeconfigAndContext(kubeConfigPath, context string) (string, error) {
	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to load kubeconfig file from %q", kubeConfigPath)
	}

	if context == "" {
		context = config.CurrentContext
	}

	for contextName, ctx := range config.Contexts {
		if contextName == context {
			return ctx.Cluster, nil
		}
	}
	return "", errors.Errorf("unable to find cluster name from kubeconfig file: %q", kubeConfigPath)
}

// GetClusterServerFromKubeconfigAndContext gets apiserver URL of the cluster from kubeconfig file
// and kube-context
func GetClusterServerFromKubeconfigAndContext(kubeConfigPath, context string) (string, error) {
	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to load kubeconfig file from %q", kubeConfigPath)
	}

	if context == "" {
		context = config.CurrentContext
	}
	clusterName := ""
	for contextName, ctx := range config.Contexts {
		if contextName == context {
			clusterName = ctx.Cluster
			break
		}
	}
	if clusterName == "" {
		return "", errors.Errorf("unable to find cluster apiserver url from kubeconfig file: %q", kubeConfigPath)
	}
	return config.Clusters[clusterName].Server, nil
}

// GetClusterInfoFromCluster gets the cluster Info by accessing the cluster-info configMap in kube-public namespace
func GetClusterInfoFromCluster(clusterAPIServerURL, configmapName string) (*clientcmdapi.Cluster, error) {
	clusterAPIServerURL = strings.TrimSpace(clusterAPIServerURL)
	if !strings.HasPrefix(clusterAPIServerURL, "https://") && !strings.HasPrefix(clusterAPIServerURL, "http://") {
		clusterAPIServerURL = "https://" + clusterAPIServerURL
	}
	_, err := url.Parse(clusterAPIServerURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse endpoint URL")
	}

	clusterAPIServerURL = strings.TrimRight(clusterAPIServerURL, " /")
	clusterInfoURL := clusterAPIServerURL + fmt.Sprintf("/api/v1/namespaces/%s/configmaps/%s", KubePublicNamespace, configmapName)
	//nolint:noctx
	req, _ := http.NewRequest("GET", clusterInfoURL, http.NoBody)
	// To get the cluster ca certificate first time, we need to use skip verify the server certificate,
	// all the later communications to cluster would be using CA after this call
	clusterClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			// #nosec
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 10,
	}
	response, err := clusterClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get cluster-info from the end-point")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get cluster-info from the end-point")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the response body")
	}

	var clusterIC clusterInfoConfig
	err = json.Unmarshal(responseBody, &clusterIC)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing http response body ")
	}

	config, err := clientcmd.Load([]byte(clusterIC.Data.Kubeconfig))
	if err != nil {
		return nil, errors.Wrap(err, "failed to load the kubeconfig")
	}

	if len(config.Clusters) == 0 {
		return nil, errors.New("failed to get cluster information ")
	}
	// since it is a map with one cluster object, get the first entry
	var cluster *clientcmdapi.Cluster
	for _, cluster = range config.Clusters {
		break
	}

	return cluster, nil
}

// GetPinnipedInfoFromCluster gets the Pinniped Info by accessing the pinniped-info configMap in kube-public namespace
// 'discoveryPort' is used to optionally override the port used for discovery. This may be needed on setups that expose
// discovery information to unauthenticated users on a different port (for instance, to avoid the need to anonymous auth
// on the apiserver). By default, the endpoint from the cluster-info is used.
func GetPinnipedInfoFromCluster(clusterInfo *clientcmdapi.Cluster, discoveryPort *int) (*pinnipedinfo.PinnipedInfo, error) {
	endpoint := strings.TrimRight(clusterInfo.Server, " /")
	var err error
	if discoveryPort != nil {
		endpoint, err = netutil.SetPort(clusterInfo.Server, *discoveryPort)
		if err != nil {
			return nil, errors.Wrap(err, "failed to override discovery port")
		}
	}
	pinnipedInfoURL := endpoint + fmt.Sprintf("/api/v1/namespaces/%s/configmaps/pinniped-info", KubePublicNamespace)
	//nolint:noctx
	req, _ := http.NewRequest("GET", pinnipedInfoURL, http.NoBody)
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(clusterInfo.CertificateAuthorityData)
	clusterClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs:    pool,
				MinVersion: tls.VersionTLS12,
			},
		},
		Timeout: time.Second * 10,
	}

	response, err := clusterClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pinniped-info from the cluster")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get pinniped-info from the cluster. Status code: %+v", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read the response body")
	}

	return pinnipedinfo.ByteArrayToPinnipedInfo(responseBody)
}
