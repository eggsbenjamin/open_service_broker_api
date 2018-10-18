// +build service

package service_test

import (
	"io/ioutil"
	"net/http"

	"github.com/eggsbenjamin/open_service_broker_api/db"
	"github.com/eggsbenjamin/open_service_broker_api/testutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("System", func() {

	It("responds to GET /v1/catalog with the correct formatted response", func() {
		db, err := db.NewConnection("localhost", "32768", "postgres", "postgres", "service_catalog")
		Expect(err).NotTo(HaveOccurred())
		err = testutils.TeardownDB(db)
		Expect(err).NotTo(HaveOccurred())
		defer testutils.TeardownDB(db)

		SQLFixtures, err := ioutil.ReadFile("./testdata/catalog.sql")
		Expect(err).NotTo(HaveOccurred())

		_, err = db.Exec(string(SQLFixtures))
		Expect(err).NotTo(HaveOccurred())

		JSONFixture, err := ioutil.ReadFile("./testdata/GET_catalog_200.json")
		Expect(err).NotTo(HaveOccurred())

		res, err := http.Get("http://localhost:8080/v1/catalog")
		Expect(err).NotTo(HaveOccurred())
		defer res.Body.Close()

		Expect(res.StatusCode).To(Equal(200))
		body, err := ioutil.ReadAll(res.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(body).To(MatchJSON(JSONFixture))
	})
})
