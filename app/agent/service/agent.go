package service

import (
	"china-russia/app/agent/swag/response"
	"china-russia/extends"
	"china-russia/model"
)

type AgentService struct{}

func (AgentService) Info(agent *model.Agent) response.AgentInfo {
	jwtService := extends.JwtUtils{}
	token := jwtService.NewToken(agent.Id, agent.Token)
	return response.AgentInfo{
		Id:         agent.Id,
		Name:       agent.Account,
		Token:      token,
		Status:     agent.Status,
		CreateTime: agent.CreateTime,
		UpdateTime: agent.UpdateTime,
	}
}
