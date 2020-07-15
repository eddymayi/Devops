package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
        "os"
)

type HostStatus struct {
	Recordcount string `json:"recordcount"`
	Hoststatus  []struct {
//		attributes struct {
//			ID string `json:"id"`
//		} `json:"@attributes"`
		InstanceID       string `json:"instance_id"`
		HostID           string `json:"host_id"`
		Name             string `json:"name"`
		DisplayName      string `json:"display_name"`
		Address          string `json:"address"`
		Alias            string `json:"alias"`
		StatusUpdateTime string `json:"status_update_time"`
		StatusText       string `json:"status_text"`
		StatusTextLong   struct {
		} `json:"status_text_long"`
		CurrentState string `json:"current_state"`
//		iconImage    struct {
//		} `json:"icon_image"`
//		iconImageAlt struct {
//		} `json:"icon_image_alt"`
		PerformanceData           string `json:"performance_data"`
		ShouldBeScheduled         string `json:"should_be_scheduled"`
		CheckType                 string `json:"check_type"`
		LastStateChange           string `json:"last_state_change"`
		LastHardStateChange       string `json:"last_hard_state_change"`
		LastHardState             string `json:"last_hard_state"`
		LastTimeUp                string `json:"last_time_up"`
		LastTimeDown              string `json:"last_time_down"`
		LastTimeUnreachable       string `json:"last_time_unreachable"`
		LastNotification          string `json:"last_notification"`
		NextNotification          string `json:"next_notification"`
		NoMoreNotifications       string `json:"no_more_notifications"`
		AcknowledgementType       string `json:"acknowledgement_type"`
		CurrentNotificationNumber string `json:"current_notification_number"`
		EventHandlerEnabled       string `json:"event_handler_enabled"`
		ProcessPerformanceData    string `json:"process_performance_data"`
		ObsessOverHost            string `json:"obsess_over_host"`
		ModifiedHostAttributes    string `json:"modified_host_attributes"`
		EventHandler              struct {
		} `json:"event_handler"`
		CheckCommand           string `json:"check_command"`
		NormalCheckInterval    string `json:"normal_check_interval"`
		RetryCheckInterval     string `json:"retry_check_interval"`
		CheckTimeperiodID      string `json:"check_timeperiod_id"`
		HasBeenChecked         string `json:"has_been_checked"`
		CurrentCheckAttempt    string `json:"current_check_attempt"`
		MaxCheckAttempts       string `json:"max_check_attempts"`
		LastCheck              string `json:"last_check"`
		NextCheck              string `json:"next_check"`
		StateType              string `json:"state_type"`
		NotificationsEnabled   string `json:"notifications_enabled"`
		ProblemAcknowledged    string `json:"problem_acknowledged"`
		PassiveChecksEnabled   string `json:"passive_checks_enabled"`
		ActiveChecksEnabled    string `json:"active_checks_enabled"`
		FlapDetectionEnabled   string `json:"flap_detection_enabled"`
		IsFlapping             string `json:"is_flapping"`
		PercentStateChange     string `json:"percent_state_change"`
		Latency                string `json:"latency"`
		ExecutionTime          string `json:"execution_time"`
		ScheduledDowntimeDepth string `json:"scheduled_downtime_depth"`
		Notes                  struct {
		} `json:"notes"`
		NotesURL struct {
		} `json:"notes_url"`
		ActionURL struct {
		} `json:"action_url"`
	} `json:"hoststatus"`
}

func getGoStruct(body []byte) (*HostStatus, error) {
	var s = new(HostStatus)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Json parse error", err)
                os.Exit(7)
	}
	return s, err
}

func getNagiosHosts(hostn string, apikey string) (*HostStatus, error) {
        url := "https://" + hostn + "/nagiosxi/api/v1/objects/hoststatus?apikey=" + apikey
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		fmt.Println("Cannot reach url", err)
                os.Exit(5)
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Cannot read body")
                os.Exit(6)
	}
	s, err := getGoStruct([]byte(body))
	return s, err
}
