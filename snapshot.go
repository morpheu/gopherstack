package gopherstack

import (
	"net/url"
)

// List snapshots
func (c CloudstackClient) ListSnapshots(volumeid string, projectid string, account string) (ListSnapshotsResponse, error) {
	var resp ListSnapshotsResponse
	params := url.Values{}

	if volumeid != "" {
		params.Set("volumeid", volumeid)
	}

	if account != "" {
		params.Set("account", account)
	}

	if projectid != "" {
		params.Set("projectid", projectid)
	}

	response, err := NewRequest(c, "listSnapshots", params)
	if err != nil {
		return resp, err
	}

	resp = response.(ListSnapshotsResponse)

	return resp, err
}

func (c CloudstackClient) createSnapshot(volumeid string, projectid string, account string) (CreateSnapshotResponse, error) {
	var resp CreateSnapshotResponse
	params := url.Values{}

	if volumeid != "" {
		params.Set("volumeid", volumeid)
	}

	if account != "" {
		params.Set("account", account)
	}

	if projectid != "" {
		params.Set("projectid", projectid)
	}

	response, err := NewRequest(c, "createSnapshot", params)
	if err != nil {
		return resp, err
	}

	resp = response.(CreateSnapshotResponse)

	return resp, err
}

func (c CloudstackClient) deleteSnapshot(snapshotid string, projectid string, account string) (DeleteSnapshotResponse, error) {
	var resp DeleteSnapshotResponse
	params := url.Values{}

	if account != "" {
		params.Set("account", account)
	}

	if projectid != "" {
		params.Set("projectid", projectid)
	}

	params.Set("id", snapshotid)

	response, err := NewRequest(c, "deleteSnapshot", params)
	if err != nil {
		return resp, err
	}

	resp = response.(DeleteSnapshotResponse)

	return resp, err
}

type Snapshot struct {
	ID           string        `json:"id"`
	Account      string        `json:"account"`
	Created      string        `json:"created"`
	Domain       string        `json:"domain"`
	Domainid     string        `json:"domainid"`
	IntervalType string        `json:"intervaltype"`
	Name         string        `json:"name"`
	Project      string        `json:"project"`
	ProjectId    string        `json:"projectid"`
	SnapshotType string        `json:"snapshottype"`
	State        string        `json:"state"`
	VolumeId     string        `json:"volumeid"`
	VolumeName   string        `json:"volumename"`
	VolumeType   string        `json:"volumetype"`
	Tags         []interface{} `json:"tags"`
}

type ListSnapshotsResponse struct {
	Listsnapshotsresponse struct {
		Count    float64    `json:"count"`
		Snapshot []Snapshot `json:"snapshot"`
	} `json:"listsnapshotsresponse"`
}

type CreateSnapshotResponse struct {
	Createsnapshotresponse struct {
		ID    string `json:"id"`
		Jobid string `json:"jobid"`
	} `json:"createsnapshotresponse"`
}

type DeleteSnapshotResponse struct {
	Deletesnapshotresponse struct {
		DisplayText string `json:"displaytext"`
		Success     string `json:"success"`
	} `json:"deletesnapshotresponse"`
}
