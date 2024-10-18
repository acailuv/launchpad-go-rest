package cron_handler

import "github.com/rs/zerolog/log"

func (h handler) Heartbeat() {
	log.Info().Msg("Heartbeat...")
}
