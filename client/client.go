//go:generate mockgen -package client -source=client.go -destination client_mock.go

package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eggsbenjamin/open_service_broker_api/models"
	"github.com/pkg/errors"
)

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type Client interface {
	GetCatalog() (*models.Catalog, error)
}

type client struct {
	host, version string
	httpClient    httpClient
}

func NewClient(httpClient httpClient, host, version string) Client {
	return &client{
		httpClient: httpClient,
		host:       host,
		version:    version,
	}
}

func (c *client) GetCatalog() (*models.Catalog, error) {
	url := fmt.Sprintf("%s/%s/catalog", c.host, c.version)
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error getting catalog")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, errors.Wrap(err, "error parsing response body")
		}
		return nil, errors.Wrapf(err, "error getting catalog: %d %s", res.StatusCode, body)
	}

	catalog := &models.Catalog{}
	return catalog, json.NewDecoder(res.Body).Decode(catalog)
}
