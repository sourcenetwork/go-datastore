package scoped

import (
	"reflect"
	"testing"

	ds "github.com/ipfs/go-datastore"
)

func TestWithFeatures(t *testing.T) {
	cases := []struct {
		name     string
		dstore   ds.Datastore
		features []ds.Feature

		expectedFeatures []ds.Feature
	}{
		{
			name:             "no features should return a base datastore",
			dstore:           &ds.MapDatastore{},
			features:         nil,
			expectedFeatures: nil,
		},
		{
			name:             "identity case",
			dstore:           &ds.MapDatastore{},
			features:         ds.FeaturesByName("Batching"),
			expectedFeatures: ds.FeaturesByName("Batching"),
		},
		{
			name:             "should scope down correctly",
			dstore:           &ds.LogDatastore{},
			features:         ds.FeaturesByName("Batching"),
			expectedFeatures: ds.FeaturesByName("Batching"),
		},
		{
			name:             "takes intersection of features",
			dstore:           &ds.MapDatastore{},
			features:         ds.FeaturesByName("Batching", "Checked"),
			expectedFeatures: ds.FeaturesByName("Batching"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			new := WithFeatures(c.dstore, c.features)
			newFeats := ds.FeaturesForDatastore(new)
			if len(newFeats) != len(c.expectedFeatures) {
				t.Fatalf("expected %d features, got %v", len(c.expectedFeatures), newFeats)
			}
			if !reflect.DeepEqual(newFeats, c.expectedFeatures) {
				t.Fatalf("expected features %v, got %v", c.expectedFeatures, newFeats)
			}
		})
	}
}
