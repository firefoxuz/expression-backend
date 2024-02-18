package services

import "time"

var activeAgents map[string]time.Time

type Agent struct {
	Id       string `json:"id"`
	LastPing string `json:"last_ping"`
	IsActive bool   `json:"is_active"`
}

func NewAgentCounters() map[string]time.Time {
	activeAgents = make(map[string]time.Time)

	return activeAgents
}

func AddAgent(id string, lastPingTime time.Time) {
	activeAgents[id] = lastPingTime
}

func UpdateAgent(id string, lastPingTime time.Time) {
	activeAgents[id] = lastPingTime
}

func RemoveAgent(id string) {
	delete(activeAgents, id)
}

func GetAgents() map[string]time.Time {
	return activeAgents
}
