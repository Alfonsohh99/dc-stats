package tasks

import (
	"context"
	"dc-stats/database"
	"dc-stats/model/processed"
	"log"
	"sort"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func ProcessVoiceStats(goBot *discordgo.Session, ctx context.Context, wait *sync.WaitGroup) {
	defer wait.Done()

	for _, guild := range goBot.State.Guilds {

		guildId := guild.ID

		guildObject, err := database.FindDataGuild(ctx, guildId)
		if err != nil {
			log.Println("Cannot find guild to process")
			continue
		}

		scores := []processedModel.UserScore{}
		userData := map[string]processedModel.User{}
		for _, user := range guildObject.Users {
			channelData := []processedModel.ChannelData{}
			var totalScore uint64
			for channelId, value := range user.UserVoiceActivity {
				channelName, exists := guildObject.ChannelNameMap[channelId]
				if !exists {
					channelName = "[" + channelId + "], "
				}
				channelData = append(channelData, processedModel.ChannelData{ChannelName: channelName, Score: value})
				totalScore += value

			}
			if totalScore != 0 {
				scores = append(scores, processedModel.UserScore{Username: guildObject.UserNicknameMap[user.UserID], Score: totalScore})
				sort.SliceStable(channelData, func(i, j int) bool {
					return channelData[i].Score > channelData[j].Score
				})
				userData[user.UserID] = processedModel.User{Score: totalScore, ChannelData: channelData}
			}

		}

		sort.SliceStable(scores, func(i, j int) bool {
			return scores[i].Score > scores[j].Score
		})

		database.SaveOrUpdateProcessedGuildFromVoice(guildId, scores, userData, ctx)

	}

}
