package api

import (
	"context"
	"fmt"
	"github.com/squadcast/terraform-provider-squadcast/internal/tf"
	"net/http"
	"time"
)

type Timezone struct {
	Name string `json:"name"`
}

type Repetition struct {
	Frequency int    `json:"frequency"`
	Type      string `json:"type"`
}

type Config struct {
	Timezone          Timezone   `json:"timezone"`
	IsForever         bool       `json:"is_forever"`
	Rotate            bool       `json:"rotate"`
	RotationFrequency int        `json:"rotation_frequency"`
	Repeat            bool       `json:"repeat"`
	Repetition        Repetition `json:"repetition"`
}

type RotationSets struct {
	UserIds  []string `json:"user_ids" tf:"user_ids"`
	SquadIds []string `json:"squad_ids" tf:"squad_ids"`
}

type OnCall struct {
	Name         string          `json:"name" tf:"name"`
	StartTime    time.Time       `json:"start_time" tf:"start_time"`
	EndTime      time.Time       `json:"end_time" tf:"end_time"`
	Config       Config          `json:"config" tf:"config"`
	RotationSets []*RotationSets `json:"rotation_sets" tf:"rotation_sets"`
}

type OnCallEvents struct {
	Id         string    `json:"id"`
	CalendarId string    `json:"calendar_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Name       string    `json:"name"`
	UserIds    []string  `json:"user_ids"`
	SeriesId   string    `json:"series_id"`
	SquadIds   []string  `json:"squad_ids"`
	IsOverride bool      `json:"is_override"`
	ScheduleId string    `json:"schedule_id"`
}

func (o *OnCallEvents) Encode() (tf.M, error) {
	m, err := tf.Encode(o)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (client *Client) GetOnCallEventById(ctx context.Context, scheduleId string, eventId string) (*OnCallEvents, error) {
	url := fmt.Sprintf("%s/schedules/%s/events/%s", client.BaseURLV3, scheduleId, eventId)

	return Request[any, OnCallEvents](http.MethodGet, url, client, ctx, nil)
}

func (client *Client) GetAllOnCallEvents(ctx context.Context, scheduleId string) ([]*OnCallEvents, error) {
	url := fmt.Sprintf("%s/schedules/%s", client.BaseURLV3, scheduleId)

	return RequestSlice[any, OnCallEvents](http.MethodGet, url, client, ctx, nil)
}

func (client *Client) GetOnCallFirstEvent(ctx context.Context, scheduleId string) (*OnCallEvents, error) {
	onCallEvents, err := client.GetAllOnCallEvents(ctx, scheduleId)
	if err != nil {
		return nil, err
	}

	if len(onCallEvents) > 0 {
		return onCallEvents[0], nil
	}

	return nil, fmt.Errorf("there is no on-call events for this schedule")
}

func (client *Client) CreateOnCall(ctx context.Context, req *OnCall, scheduleId string) ([]*OnCallEvents, error) {
	url := fmt.Sprintf("%s/schedules/%s", client.BaseURLV3, scheduleId)

	return RequestSlice[OnCall, OnCallEvents](http.MethodPost, url, client, ctx, req)
}

//func (client *Client) DeleteSchedule(ctx context.Context, id string) (*any, error) {
//	url := fmt.Sprintf("%s/schedules/%s", client.BaseURLV3, id)
//	return Request[any, any](http.MethodDelete, url, client, ctx, nil)
//}
