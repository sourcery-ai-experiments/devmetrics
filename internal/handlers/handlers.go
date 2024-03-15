package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/rybalka1/devmetrics/internal/memstorage"
)

func UpdateMetricsHandle(store memstorage.Storage) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		metricsInfo, found := strings.CutPrefix(r.URL.Path, "/update/")
		if !found {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		pieces := strings.Split(metricsInfo, "/")
		if len(pieces) != 3 {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		mType := pieces[0]
		mName := pieces[1]
		mValue := pieces[2]

		switch mType {
		case "gauge":
			val, err := strconv.ParseFloat(mValue, 64)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
			store.UpdateGauges(mName, val)
		case "counter":
			val, err := strconv.ParseInt(mValue, 10, 64)
			if err != nil {
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
			store.UpdateCounters(mName, val)
		default:
			rw.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println(store)
		rw.WriteHeader(http.StatusOK)
	}
}