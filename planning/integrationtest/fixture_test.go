package integrationtest

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/szabba/assert"
	"gocloud.dev/docstore"
)

func loadFixture(t *testing.T, fixtureName string) {
	f, err := os.Open(fixtureName)
	assert.That(err == nil, t.Fatalf, "cannot load fixture %q data: %s", fixtureName, err)
	defer f.Close()

	colls := make(Fixture)
	dec := json.NewDecoder(f)
	err = dec.Decode(&colls)
	assert.That(err == nil, t.Fatalf, "cannot parse fixture %q data: %s", fixtureName, err)
	colls.Load(t)
}

type Fixture map[string]CollectionFixture

func (colls Fixture) Load(t *testing.T) {
	for url, docs := range colls {
		docs.Fill(t, url)
	}
}

type CollectionFixture []DocumentFixture

func (docs CollectionFixture) Fill(t *testing.T, url string) {
	coll, err := docstore.OpenCollection(context.Background(), url)
	assert.That(err == nil, t.Fatalf, "cannot open collection %q: %s", url, err)

	for _, d := range docs {
		d.Add(t, url, coll)
	}
}

type DocumentFixture map[string]interface{}

func (d DocumentFixture) Add(t *testing.T, url string, coll *docstore.Collection) {
	if rev, ok := d["DocumentRevision"].(float64); ok {
		d["DocumentRevision"] = int64(rev)
	}
	var rawDoc map[string]interface{} = d
	err := coll.Put(context.Background(), rawDoc)
	assert.That(err == nil, t.Fatalf, "cannot fill collection %s: cannot store document %#v: %s", url, d, err)
}
