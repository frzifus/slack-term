package views

import (
	"github.com/erroneousboat/termui"

	"github.com/theremix/slack-term/components"
	"github.com/theremix/slack-term/config"
	"github.com/theremix/slack-term/service"
)

type View struct {
	Config   *config.Config
	Input    *components.Input
	Chat     *components.Chat
	Channels *components.Channels
	Mode     *components.Mode
	Debug    *components.Debug
}

func CreateView(config *config.Config, svc *service.SlackService) (*View, error) {
	// Create Input component
	input := components.CreateInputComponent()

	// Channels: create the component
	channels := components.CreateChannelsComponent(input.Par.Height)

	// Channels: fill the component
	slackChans, err := svc.GetChannels()
	if err != nil {
		return nil, err
	}

	// Channels: set channels in component
	channels.SetChannels(slackChans)

	// Chat: create the component
	chat := components.CreateChatComponent(input.Par.Height)

	// Chat: fill the component
	msgs, err := svc.GetMessages(
		channels.ChannelItems[channels.SelectedChannel].ID,
		chat.GetMaxItems(),
	)

	if err != nil {
		return nil, err
	}

	// Chat: set messages in component
	chat.SetMessages(msgs)

	chat.SetBorderLabel(
		channels.ChannelItems[channels.SelectedChannel].GetChannelName(),
	)

	// Debug: create the component
	debug := components.CreateDebugComponent(input.Par.Height)

	// Mode: create the component
	mode := components.CreateModeComponent()

	view := &View{
		Config:   config,
		Input:    input,
		Channels: channels,
		Chat:     chat,
		Mode:     mode,
		Debug:    debug,
	}

	return view, nil
}

func (v *View) Refresh() {
	termui.Render(
		v.Input,
		v.Chat,
		v.Channels,
		v.Mode,
	)
}
