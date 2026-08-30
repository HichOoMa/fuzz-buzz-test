package service

import (
	"net/url"
	"reflect"
	"testing"

	"fizz-buzz.com/internal/model"
)

// resetStats clears the package-level stats state so tests don't leak
// into one another (RecordRequest/MostUsedRequest share global state).
func resetStats() {
	statsMu.Lock()
	defer statsMu.Unlock()
	statsCounts = make(map[string]int)
	statsParams = make(map[string]map[string]string)
}

func TestRecordRequest(t *testing.T) {
	t.Run("records params and increments count for repeated requests", func(t *testing.T) {
		resetStats()

		values := url.Values{"int1": {"3"}, "int2": {"5"}}
		RecordRequest(values)
		RecordRequest(values)

		key := values.Encode()

		statsMu.Lock()
		gotCount := statsCounts[key]
		gotParams := statsParams[key]
		statsMu.Unlock()

		if gotCount != 2 {
			t.Errorf("statsCounts[%q] = %d, want 2", key, gotCount)
		}

		wantParams := map[string]string{"int1": "3", "int2": "5"}
		if !reflect.DeepEqual(gotParams, wantParams) {
			t.Errorf("statsParams[%q] = %v, want %v", key, gotParams, wantParams)
		}
	})

	t.Run("distinct query values are tracked separately", func(t *testing.T) {
		resetStats()

		RecordRequest(url.Values{"int1": {"3"}})
		RecordRequest(url.Values{"int1": {"7"}})

		statsMu.Lock()
		defer statsMu.Unlock()

		if len(statsCounts) != 2 {
			t.Errorf("len(statsCounts) = %d, want 2", len(statsCounts))
		}
		if statsCounts[url.Values{"int1": {"3"}}.Encode()] != 1 {
			t.Errorf("expected count 1 for int1=3")
		}
		if statsCounts[url.Values{"int1": {"7"}}.Encode()] != 1 {
			t.Errorf("expected count 1 for int1=7")
		}
	})
}

func TestMostUsedRequest(t *testing.T) {
	t.Run("no requests recorded returns zero value", func(t *testing.T) {
		resetStats()

		got := MostUsedRequest()
		want := model.RequestStats{Params: nil, Count: 0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MostUsedRequest() = %+v, want %+v", got, want)
		}
	})

	t.Run("returns the most frequently recorded request", func(t *testing.T) {
		resetStats()

		least := url.Values{"int1": {"3"}, "int2": {"5"}}
		most := url.Values{"int1": {"7"}, "int2": {"11"}}

		RecordRequest(least)
		RecordRequest(most)
		RecordRequest(most)
		RecordRequest(most)

		got := MostUsedRequest()
		want := model.RequestStats{
			Params: map[string]string{"int1": "7", "int2": "11"},
			Count:  3,
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MostUsedRequest() = %+v, want %+v", got, want)
		}
	})

	t.Run("tie returns one of the top-count requests", func(t *testing.T) {
		resetStats()

		first := url.Values{"int1": {"1"}}
		second := url.Values{"int1": {"2"}}

		RecordRequest(first)
		RecordRequest(second)

		got := MostUsedRequest()
		if got.Count != 1 {
			t.Fatalf("MostUsedRequest().Count = %d, want 1", got.Count)
		}

		firstParams := map[string]string{"int1": "1"}
		secondParams := map[string]string{"int1": "2"}
		if !reflect.DeepEqual(got.Params, firstParams) && !reflect.DeepEqual(got.Params, secondParams) {
			t.Errorf("MostUsedRequest().Params = %v, want %v or %v", got.Params, firstParams, secondParams)
		}
	})
}
