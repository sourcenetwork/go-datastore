package datastore

import (
	"reflect"
	"testing"
)

func TestFeaturesByName(t *testing.T) {
	feats := FeaturesByName()
	if feats != nil {
		t.Fatalf("expected nil features, got %v", feats)
	}

	feats = FeaturesByName("Batching")
	if len(feats) != 1 ||
		feats[0].Name != "Batching" ||
		feats[0].Interface != (*BatchingFeature)(nil) ||
		feats[0].DatastoreInterface != (*Batching)(nil) {
		t.Fatalf("expected a batching feature, got %v", feats)
	}

	feats = FeaturesByName("Batching", "UnknownFeature")
	if len(feats) != 1 || feats[0].Name != "Batching" {
		t.Fatalf("expected a batching feature, got %v", feats)
	}
}

func TestFeaturesForDatastore(t *testing.T) {
	cases := []struct {
		name             string
		d                Datastore
		expectedFeatures []string
	}{
		{
			name:             "MapDatastore",
			d:                &MapDatastore{},
			expectedFeatures: []string{"Batching"},
		},
		{
			name:             "NullDatastore",
			d:                &NullDatastore{},
			expectedFeatures: []string{"Batching"},
		},
		{
			name:             "LogDatastore",
			d:                &LogDatastore{},
			expectedFeatures: []string{"Batching", "Checked", "GC", "Persistent", "Scrubbed"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			feats := FeaturesForDatastore(c.d)
			if len(feats) != len(c.expectedFeatures) {
				t.Fatalf("expected %d features, got %v", len(c.expectedFeatures), feats)
			}
			expectedFeats := FeaturesByName(c.expectedFeatures...)
			if !reflect.DeepEqual(expectedFeats, feats) {
				t.Fatalf("expected features %v, got %v", c.expectedFeatures, feats)
			}
		})
	}
}
